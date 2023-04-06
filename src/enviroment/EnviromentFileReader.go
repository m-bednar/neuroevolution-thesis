package enviroment

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func ReadEnviromentFile(filename string) []TileType {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	return ReadTiles(reader)
}

func ReadTiles(reader *bufio.Reader) []TileType {
	tiles := make([]TileType, 0)
	for {
		c, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Ignore non-digits
		if !unicode.IsDigit(c) {
			continue
		}

		// Convert ascii value to tile type
		value, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}

		tile := TileType(value)
		tiles = append(tiles, tile)
	}
	return tiles
}
