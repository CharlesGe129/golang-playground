package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	add()
	low()
	event()
}

func add() {
	robotgo.EventHook(hook.KeyDown, []string{"z", "1"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		color := robotgo.GetPixelColor(x, y)
		fmt.Printf(`{ X: %d, Y: %d, Color: "%s" },`+"\n", x, y, color)
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func low() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		fmt.Println("hook: ", ev)
	}
}

func event() {
	if robotgo.AddEvents("z", "1") {
		fmt.Println("you press... ", "k")
	}
}
