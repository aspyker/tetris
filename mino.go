package main

import (
	"math/rand"
	"strings"
	"time"
)

const (
	defaultMinoX, defaultMinoY = 3, -1
	minoWidth, minoHeight      = 4, 4
)

var (
	blocks = []string{
		`
		  ....
			.gg.
			gg..
			....
		`, `
		  ....
			.rr.
			..rr
			....
		`, `
		  ....
			.YY.
			.YY.
			....
		`, `
		  ....
			....
			CCCC
			....
		`, `
		  ....
			.M..
			MMM.
			....
		`, `
		  ....
			.b..
			.bbb
			....
		`, `
		  ....
			..m.
			mmm.
			....
		`,
	}
)

var (
	currentMino *Mino
	nextMino    *Mino
)

type Mino struct {
	block string
	x     int
	y     int
}

func NewMino() *Mino {
	rand.Seed(time.Now().UnixNano())
	block := blocks[rand.Intn(len(blocks))]
	block = strings.Replace(block, "\t", "", -1)
	block = strings.Replace(block, " ", "", -1)
	block = strings.Trim(block, "\n")
	return &Mino{block: block}
}

func (m *Mino) cell(x, y int) rune {
	return rune(m.block[x+(minoWidth+1)*y])
}

func (m *Mino) setCell(x, y int, cell rune) {
	buf := []rune(m.block)
	buf[x+(minoWidth+1)*y] = cell
	m.block = string(buf)
}

func (m *Mino) moveDown() {
	m.y++
}

func (m *Mino) moveLeft() {
	dstMino := *m
	dstMino.x--
	if dstMino.isOnBoard() {
		m.x--
	}
}

func (m *Mino) moveRight() {
	dstMino := *m
	dstMino.x++
	if dstMino.isOnBoard() {
		m.x++
	}
}

func (m *Mino) applyGravity() {
	m.moveDown()
}

func (m *Mino) rotateRight() {
	oldMino := *m

	for j := 0; j < minoHeight; j++ {
		for i := 0; i < minoWidth; i++ {
			m.setCell(minoHeight-j-1, i, oldMino.cell(i, j))
		}
	}
}

func (m *Mino) rotateLeft() {
	oldMino := *m

	for j := 0; j < minoHeight; j++ {
		for i := 0; i < minoWidth; i++ {
			m.setCell(j, minoWidth-i-1, oldMino.cell(i, j))
		}
	}
}

func (m *Mino) isOnBoard() bool {
	for _, position := range m.positions() {
		position.isOnBoard()
		if !position.isOnBoard() {
			return false
		}
	}
	return true
}

func (m *Mino) positions() []*Position {
	positions := []*Position{}
	for i := 0; i < minoWidth; i++ {
		for j := 0; j < minoHeight; j++ {
			if m.cell(i, j) != '.' {
				positions = append(positions, NewPosition(m.x+i, m.y+j))
			}
		}
	}
	return positions
}
