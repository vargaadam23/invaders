package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/enemy"
	"github.com/vargaadam23/invaders/gamecontext"
	"github.com/vargaadam23/invaders/mapobjects"
	"github.com/vargaadam23/invaders/player"
)

const windowHeight = 800
const windowWidth = 1000

func main() {
	rl.InitWindow(windowWidth, windowHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	window := rl.Rectangle{X: 0, Y: 0, Height: windowHeight, Width: windowWidth}

	context := gamecontext.GetInstance()
	player := player.InitCharacter(int32(context.Window.Width), int32(context.Window.Height))

	obenemy := enemy.InitCharacter(int32(context.Window.Width), int32(context.Window.Height))
	wallT := mapobjects.InitWall(10, "TOP", window)
	wallB := mapobjects.InitWall(10, "BOTTOM", window)
	wallL := mapobjects.InitWall(10, "LEFT", window)
	wallR := mapobjects.InitWall(10, "RIGHT", window)

	context.GameObjects.Append(player)
	context.GameObjects.Append(obenemy)

	context.GameObjects.Append(wallT)
	context.GameObjects.Append(wallB)
	context.GameObjects.Append(wallL)
	context.GameObjects.Append(wallR)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		context.DrawGameElements()

		rl.DrawFPS(100, 100)
		rl.EndDrawing()
	}
}
