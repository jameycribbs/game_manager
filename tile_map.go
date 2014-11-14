package game_manager

import (
	"encoding/xml"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"os"
)

type TileMap struct {
	XMLName  xml.Name `xml:"map"`
	tileSize int
	width    int
	height   int
}
