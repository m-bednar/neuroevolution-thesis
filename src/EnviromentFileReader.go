package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode"
)

func ReadEnviromentFile(filename string) []TileType {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var reader = bufio.NewReader(file)
	return ReadTiles(reader)
}

func ReadTiles(reader *bufio.Reader) []TileType {
	var tiles = make([]TileType, 0)
	for {
		var b, err = reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Ignore non-digits
		if !unicode.IsDigit(rune(b)) {
			continue
		}

		// Convert ascii value to tile type
		var value = b - 48
		var tile = TileType(value)

		tiles = append(tiles, tile)
	}
	return tiles
}
