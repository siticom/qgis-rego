package xml

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"qgis-repository/config"
	"qgis-repository/model"

	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
)

// Necessary to embed the web interface
const xmlHeader = `<?xml version = "1.0" encoding = "UTF-8"?><?xml-stylesheet type="text/xsl" href="/static/plugins.xsl" ?>`

var plugin_list model.PluginList

// Generates the XML output for the plugins found in the plugin directory
func GenerateXML(queryQgisVersion version.Version) (string, error) {
	var filtered_plugin_list model.PluginList
	for _, plugin := range plugin_list {
		// Parse version information or set default values
		pluginMinQgisVersion, err := version.NewVersion(plugin.QgisMinimumVersion)
		if err != nil {
			pluginMinQgisVersion, _ = version.NewVersion("2.14")
		}
		pluginMaxQgisVersion, err := version.NewVersion(plugin.QgisMaximumVersion)
		if err != nil {
			plugin.QgisMaximumVersion = fmt.Sprintf("%d.98", pluginMinQgisVersion.Segments()[0])
			pluginMaxQgisVersion, _ = version.NewVersion(plugin.QgisMaximumVersion)
		}

		// Compare queried version against plugin version, only add if plugin supports version
		if queryQgisVersion.GreaterThanOrEqual(pluginMinQgisVersion) && queryQgisVersion.LessThanOrEqual(pluginMaxQgisVersion) {
			filtered_plugin_list = append(filtered_plugin_list, plugin)
		}
	}

	// Convert plugin list into XML
	output, err := xml.MarshalIndent(filtered_plugin_list, "", "   ")
	if err != nil {
		log.WithError(err).Fatal("could not generate xml output")
		return "", err
	}
	log.Debug(string(output))
	xml := []byte(xmlHeader + string(output))
	return string(xml), nil
}

func SetPluginList(plugins model.PluginList) {
	plugin_list = plugins
}

func FindPlugins(dir string) (model.PluginList, error) {
	// Get all files from directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.WithError(err).Fatal("could not read plugin directory")
		return nil, err
	}

	// Index all files in the directory and transform them into Plugins
	var plugin_list model.PluginList = make([]model.Plugin, 0, len(files))
	for _, file := range files {
		// We are only interested in ZIP files
		if filepath.Ext(file.Name()) != ".zip" {
			continue
		}
		// Open the ZIP file
		archive, err := zip.OpenReader(filepath.Join(dir, file.Name()))
		if err != nil {
			log.WithError(err).Error("could not open archive: " + file.Name())
			continue
		}
		defer archive.Close()

		// Search for metadata.txt in archive
		for _, f := range archive.File {
			// We are only interested in the metadata file
			if f.FileInfo().Name() != "metadata.txt" {
				continue
			}

			metadataFile, err := f.Open()
			if err != nil {
				log.WithError(err).Error("could not open metadata file")
				break
			}
			defer metadataFile.Close()

			// Get all metadata from the file
			plugin := model.ParseMetadataToModel(metadataFile)
			plugin.SetFileNameAndDownload(file.Name(), config.ServerURL)
			if config.Env.ShowIcons {
				plugin.SetImageBase64(archive)
			}
			plugin.GenerateUniqueId()

			// Add the plugin to the list
			plugin_list = append(plugin_list, plugin)
			break
		}
	}
	return plugin_list, nil
}
