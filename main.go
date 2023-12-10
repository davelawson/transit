package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionCancel input.Action = iota
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	logicFrameCtr   int
	graphicFrameCtr int
	inputSystem     input.System
	inputHandler    *input.Handler
}

func NewGame() *Game {
	g := &Game{
		logicFrameCtr:   0,
		graphicFrameCtr: 0,
	}

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	keymap := input.Keymap{
		ActionCancel: {input.KeyEscape},
	}

	g.inputHandler = g.inputSystem.NewHandler(0, keymap)

	return g
}

func (g *Game) Update() error {
	g.inputSystem.Update()
	g.logicFrameCtr++

	if g.inputHandler.ActionIsJustPressed(ActionCancel) {
		fmt.Println("Cancel pressed")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.graphicFrameCtr++
	var frameMessage = fmt.Sprintf("Logic Frame: %d\nGraphic Frame: %d", g.logicFrameCtr, g.graphicFrameCtr)
	ebitenutil.DebugPrint(screen, frameMessage)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
