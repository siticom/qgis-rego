package model

import (
	"archive/zip"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"strings"

	"gopkg.in/ini.v1"

	log "github.com/sirupsen/logrus"
)

// Model according to https://docs.qgis.org/3.22/en/docs/pyqgis_developer_cookbook/plugins/plugins.html#plugin-metadata
// For all cdata field we have to work with embedded as "name,cdata" is not allowed only ",cdata"
type Plugin struct {
	XMLName            xml.Name `xml:"pyqgis_plugin"`
	ID                 int      `ini:"-" xml:"id,attr"`
	Name               string   `ini:"name" xml:"name,attr"`
	QgisMinimumVersion string   `ini:"qgisMinimumVersion" xml:"qgis_minimum_version"`
	QgisMaximumVersion string   `ini:"qgisMaximumVersion" xml:"qgis_maximum_version"`
	Description        *struct {
		Value string `ini:"description" xml:",cdata"`
	} `ini:"general" xml:"description"`
	Version string `ini:"version" xml:"version,attr"`
	Author  *struct {
		Value string `ini:"author" xml:",cdata"`
	} `ini:"general" xml:"author"`
	About *struct {
		Value string `ini:"about" xml:",cdata"`
	} `ini:"general" xml:"about"`
	Email        string  `ini:"email" xml:"email"`
	Changelog    *string `ini:"changelog" xml:"-"`
	Experimental *bool   `ini:"experimental" xml:"experimental"`
	Deprecated   *bool   `ini:"deprecated" xml:"deprecated"`
	Homepage     *struct {
		Value string `ini:"homepage" xml:",cdata"`
	} `ini:"general" xml:"homepage"`
	Repository struct {
		Value string `ini:"repository" xml:",cdata"`
	} `ini:"general" xml:"repository"`
	Tracker *struct {
		Value string `ini:"tracker" xml:",cdata"`
	} `ini:"general" xml:"tracker"`
	Tags *struct {
		Value string `ini:"tags" xml:",cdata"`
	} `ini:"general" xml:"tags"`
	Icon                  *string `ini:"icon" xml:"icon"`
	IconBase64            string  `ini:"-" xml:"icon_base64"`
	Category              *string `ini:"category" xml:"category"`
	PluginDependencies    *string `ini:"plugin_dependencies" xml:"external_dependencies"`
	Server                *bool   `ini:"server" xml:"server"`
	HasProcessingProvider *bool   `ini:"hasProcessingProvider" xml:"-"`

	// Special attributes for the XML
	FileName    string `ini:"-" xml:"file_name"`
	DownloadURL string `ini:"-" xml:"download_url"`
}

func ParseMetadataToModel(file io.ReadCloser) (plugin Plugin) {
	cfg, err := ini.LoadSources(ini.LoadOptions{AllowPythonMultilineValues: true}, file)
	if err != nil {
		log.WithError(err).Fatal("could not read metadata.txt")
		os.Exit(1)
	}
	cfg.Section("general").MapTo(&plugin)

	return plugin
}

func (p *Plugin) GenerateUniqueId() {
	// Generate the unique Plugin ID based on the name
	key := strings.Join([]string{p.Name, p.FileName, p.Version}, ";")
	p.ID = int(crc32.ChecksumIEEE([]byte(key)))
}

func (p *Plugin) SetFileNameAndDownload(file_name string, server_url string) {
	p.FileName = file_name
	p.DownloadURL = fmt.Sprintf("%s/plugins/%s", server_url, file_name)
}

func (p *Plugin) SetImageBase64(zip_reader *zip.ReadCloser) {
	// Encode image as base64 to display it on the webpage
	if p.Icon != nil {
		// Unfortunatly we have to search for the icon file in the archive :(
		for _, f := range zip_reader.File {
			if f.FileInfo().Name() == *p.Icon {
				log.Debug("Found icon file: " + f.Name)
				imgFile, err := f.Open()
				if err != nil {
					log.WithError(err).Error("could not open icon file")
					break
				}
				defer imgFile.Close()
				// Read the image file's content
				imgBytes, err := io.ReadAll(imgFile)
				if err != nil {
					fmt.Println("Error reading image file inside ZIP:", err)
					os.Exit(1)
				}
				p.IconBase64 = base64.StdEncoding.EncodeToString(imgBytes)
				return
			}
		}
	}
}
