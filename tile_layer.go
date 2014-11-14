package game_manager

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"io"
)

type TileLayer struct {
	XMLName    xml.Name `xml:"layer"`
	name       string   `xml:"name"`
	width      int      `xml:"width"`
	height     int      `xml:"height"`
	tileIDs    []int
	numColumns int
	numRows    int
	position   Vector2D
	velocity   Vector2D
	tilesets   []Tileset
	data       [][]int
}

//*****************************************************************************
// decompressData
//*****************************************************************************
func (tl *TileLayer) decompressData(string_data string) {
	decoded_data, _ := encoding.base64.StdEncoding.DecodeString(string_data)

	ids := bytes.NewBuffer(nil)
	b := bytes.NewReader(decoded_data)

	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(ids, r)

	r.Close()

	for rows := 0; rows < tl.height; rows++ {
		for cols := 0; cols < tl.width; cols++ {
			final_data[rows][cols] = ids[rows*tl.width+cols]
		}
	}
}

//*****************************************************************************
// SetTileIDs
//*****************************************************************************
func (tl *TileLayer) SetTileIDs(data []ints) {
	tl.tileIDs = data
}

//*****************************************************************************
// SetTileSize
//*****************************************************************************
func (tl *TileLayer) setTileSize(tileSize int) {
	tl.tileSize = tileSize
}

//*****************************************************************************
// TilesetByID
//*****************************************************************************
func (tl *TileLayer) TilesetByID(tileID int) {
}

//*****************************************************************************
// render
//*****************************************************************************
func (tl *TileLayer) render() {
}

//*****************************************************************************
// update
//*****************************************************************************
func (tl *TileLayer) update() {
}
