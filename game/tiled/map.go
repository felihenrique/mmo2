package tiled

import (
	"encoding/xml"
	"os"
	"strings"
)

type Layer struct {
	Id     string   `xml:"id,attr"`
	Name   string   `xml:"name,attr"`
	Width  int      `xml:"width,attr"`
	Height int      `xml:"height,attr"`
	Data   TileData `xml:"data"`
}

type TilesetImage struct {
	Source string `xml:"source,attr"`
}

type Tileset struct {
	Image TilesetImage `xml:"image"`
}

type Map struct {
	Width      int     `xml:"width,attr"`
	Height     int     `xml:"height,attr"`
	TileWidth  int     `xml:"tilewidth,attr"`
	TileHeight int     `xml:"tileheight,attr"`
	Layer      []Layer `xml:"layer"`
	Tileset    Tileset `xml:"tileset"`
}

func Load(path string) (*Map, error) {
	mapFile, err := os.ReadFile("assets/maps/main.tmx")
	if err != nil {
		return nil, err
	}
	var mapData Map
	reader := strings.NewReader(string(mapFile))
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(&mapData)
	if err != nil {
		return nil, err
	}
	return &mapData, nil
}
