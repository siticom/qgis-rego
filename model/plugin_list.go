package model

import "encoding/xml"

// plugins is a list of plugins each having its attributes in a map
type PluginList []Plugin

// Method to transform the the list of plugins into a XML (implement XML interface)
func (p PluginList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "plugins"}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, plugin := range p {
		if err := e.Encode(plugin); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}
