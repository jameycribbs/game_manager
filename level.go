package game_manager

import (
	"encoding/xml"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"os"
)

type Level struct {
	XMLName  xml.Name   `xml:"map"`
	tilesets []*Tileset `xml:"tileset"`
	layers   []*Layer   `xml:"layer"`
}

//*****************************************************************************
// ParseLevel
//*****************************************************************************
func ParseLevel(levelFile string) *Level {
	var level *Level

	xmlFile, err := os.Open(levelFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(xmlData, &level)

	return level
}

//*****************************************************************************
// GetLayers
//*****************************************************************************
func (level *Level) GetLayers() []*Layer {
	return level.layers
}

//*****************************************************************************
// GetTilesets
//*****************************************************************************
func (level *Level) GetTilesets() []*Tileset {
	return level.tilesets
}

//*****************************************************************************
// update
//*****************************************************************************
func (level *Level) update(game *Game) {
	for _, layer := range level.layers {
		layer.update()
	}
}

//*****************************************************************************
// render
//*****************************************************************************
func (level *Level) render(game *Game) {
	for _, layer := range level.layers {
		layer.render()
	}
}
