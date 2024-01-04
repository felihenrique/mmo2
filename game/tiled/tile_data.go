package tiled

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type TileData []int

func (s *TileData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	lines := strings.Split(content, "\n")
	*s = make(TileData, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		tiles := strings.Split(line, ",")
		for _, tile := range tiles {
			if tile == "" {
				continue
			}
			val, err := strconv.Atoi(tile)
			if err != nil {
				panic(err)
			}
			*s = append(*s, val)
		}
	}
	return nil
}
