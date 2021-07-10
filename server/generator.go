package main

import (
	"math/rand"
	"time"
)

type Generator struct {
	asciiSlice []uint8
	isLong     uint8
	genLink    string
}

var (
	// диапазоны значений между элементами слайсов соответсвуют значениям байтов в таблице ascii

	// LatinaBig латинский алфавит в большом регистре
	LatinaBig = []int{65, 90}
	// LatinaSmall лфтинский алфавит в маленьком регистре
	LatinaSmall = []int{97, 122}
	// Numbers цифры
	Numbers = []int{48, 57}
)

func (g *Generator) SetGenLink(genLink string) *Generator {
	g.genLink = genLink
	return g
}

// SetAsciiSlice создает слайс с допустимыми символами из таблицы ascii
func (g *Generator) SetAsciiSlice(ranges ...[]int) *Generator {

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			g.asciiSlice = append(g.asciiSlice, uint8(i))
		}
	}
	return g
}

// SetIsLong определяет длинну генегируемой ссылки
func (g *Generator) SetIsLong(isLong uint8) *Generator {
	g.isLong = isLong
	return g
}

// GenerateShortLink генерирует короткую ссылку
func (g *Generator) GenerateShortLink() {
	var genByte []uint8
	for len(genByte) < int(g.isLong){
		//rand.Seed(int64(len(g.asciiSlice)))
		rand.Seed(time.Now().UnixNano())
		//rand.Seed()
		randIndex := rand.Intn(len(g.asciiSlice))
		genByte = append(genByte, g.asciiSlice[randIndex])
		g.SetGenLink(string(genByte))
	}
}

