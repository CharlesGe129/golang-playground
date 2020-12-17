package main

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/rs/zerolog/log"
)

type Point struct {
	X     int
	Y     int
	Color string
}

type Hero struct {
	NamePointList   []Point
	AttrPointList   []Point
	NumberPointList []Point
}

func NewHero(namePointList, attrPointList, numberPointList []Point) *Hero {
	return &Hero{
		NamePointList:   namePointList,
		AttrPointList:   attrPointList,
		NumberPointList: numberPointList,
	}
}

func (h *Hero) Start() {
	time.Sleep(time.Second * 3)
	for {
		h.Refresh()
		if h.Check() {
			break
		}
	}
}

func (h *Hero) LeftClick(x, y int, interval time.Duration) {
	robotgo.MoveMouse(x, y)
	time.Sleep(interval)
	robotgo.MouseClick("left", true)
	time.Sleep(interval)
}

func (h *Hero) Refresh() {
	refreshPoint := Point{
		X: 757,
		Y: 477,
	}
	confirmPoint := Point{
		X: 552,
		Y: 471,
	}
	// 刷新
	h.LeftClick(refreshPoint.X, refreshPoint.Y, time.Millisecond*250)

	// 确认
	h.LeftClick(confirmPoint.X, confirmPoint.Y, time.Millisecond*250)

	// 刷新
	h.LeftClick(refreshPoint.X, refreshPoint.Y, time.Millisecond*250)
}

func (h *Hero) Check() bool {
	// name
	for _, point := range h.NamePointList {
		if robotgo.GetPixelColor(point.X, point.Y) != point.Color {
			return false
		}
	}
	log.Info().Msg("name matched")
	// attr
	for _, point := range h.AttrPointList {
		if robotgo.GetPixelColor(point.X, point.Y) != point.Color {
			return false
		}
	}
	log.Info().Msg("attr matched")
	// number
	for _, point := range h.NumberPointList {
		if robotgo.GetPixelColor(point.X, point.Y) != point.Color {
			return false
		}
	}
	log.Info().Msg("number matched")
	return true
}

func main() {
	// 夏+暴击伤害
	name := []Point{
		{X: 702, Y: 328, Color: "a0a145"},
		{X: 702, Y: 329, Color: "d3d751"},
		{X: 702, Y: 332, Color: "534f33"},
	}
	attr := []Point{

	}
	numebr := []Point{

	}
	NewHero(name, attr, numebr).Start()
}
