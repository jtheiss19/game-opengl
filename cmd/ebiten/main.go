package main

import (
	"fmt"
	"game/internal/ecs"
	"game/internal/ecs/objects"
	"game/internal/ecs/systems"
	_ "image/png"
	"log"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
)

const (
	screenWidth  = 500
	screenHeight = 500

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8

	windowName = "Test Game"
)

const (
	fpsCap = 60
)

type Game struct {
	gameWorld *ecs.World
}

func (g *Game) Update() error {
	g.gameWorld.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
	)

	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	logrus.Info("starting up...")

	logrus.Info("creating base systems...")
	world := ecs.NewWorld()
	newRenderer := systems.NewEbitenSpriteRenderer()
	world.AddSystem(newRenderer)

	// newPlayerController := systems.NewPlayerController()
	// world.AddSystem(newPlayerController)

	logrus.Info("creating base objects...")
	objects.NewSprite(world, mgl32.Vec3{0, 0, 8})
	objects.NewSprite(world, mgl32.Vec3{2, 0, 8})
	objects.NewPlayerCamera(world)

	logrus.Info("Configuring Window...")
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle(windowName)

	newGame := &Game{
		gameWorld: world,
	}
	logrus.Info("Running game...")

	err := ebiten.RunGame(newGame)
	if err != nil {
		log.Fatal(err)
	}
}
