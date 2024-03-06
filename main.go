package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"qgis-repository/config"
	"qgis-repository/xml"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("DEBUG") == "true" {
		log.SetLevel(log.DebugLevel)
		log.Warn("DEBUG MODE ENABLED")
	}

	log.SetFormatter(&log.TextFormatter{TimestampFormat: "02.01.2006 15:04:05", FullTimestamp: true})

	log.Info("QGIS Repository has started.")
	config.LoadConfig()

	// Create new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	// Add a path
	err = watcher.Add(config.Env.PluginDir)
	if err != nil {
		log.Fatal(err)
	}

	// Start listening for plugin changes
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Debug("event:", event)
				if filepath.Ext(event.Name) == ".zip" || event.Name == config.Env.PluginDir {
					log.Info("modified file: ", event.Name)
					// Rebuild the plugin list
					plugins, err := xml.FindPlugins(config.Env.PluginDir)
					if err != nil {
						log.WithError(err).Fatal("Error finding plugins")
					}
					log.Infof("Found %d plugins", len(plugins))
					xml.SetPluginList(plugins)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.WithError(err).Error("Watcher error")
			}
		}
	}()

	// Trigger the initial plugin list
	watcher.Events <- fsnotify.Event{Name: config.Env.PluginDir, Op: fsnotify.Write}

	// Make the plugins available and the interface
	http.Handle("/plugins/", http.StripPrefix("/plugins/", http.FileServer(http.Dir(config.Env.PluginDir))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Endpoint for the plugins overview xml
	http.HandleFunc("/plugins/plugins.xml", func(w http.ResponseWriter, r *http.Request) {
		// Set server url based on the request
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		ctx := context.WithValue(r.Context(), xml.ServerUrlKey{}, fmt.Sprintf("%s://%s", scheme, r.Host))

		// Check the qgis version from the request
		qgisVersion, err := version.NewVersion(r.URL.Query().Get("qgis"))
		if err != nil {
			qgisVersion, _ = version.NewVersion(config.Env.MinimumQgisVersion)
		}

		pluginsXml, err := xml.GenerateXML(ctx, *qgisVersion)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		io.WriteString(w, pluginsXml)
	})
	// Redirect from shortcut to correct endpoint
	http.HandleFunc("/plugins", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/plugins/plugins.xml", http.StatusSeeOther)
	})

	log.Infof("Connect to http://localhost:%d/plugins/plugins.xml for QGIS Repository", config.Env.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.FormatUint(uint64(config.Env.Port), 10), nil))
}
