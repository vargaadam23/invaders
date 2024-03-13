package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/gamecontext"
)

const windowHeight = 800
const windowWidth = 1000

func main() {
	rl.InitWindow(windowWidth, windowHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	context := gamecontext.InitGameContext()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		context.Player.HandleMovement()

		rl.BeginDrawing()

		context.DrawGameElements()

		rl.DrawFPS(100, 100)
		rl.EndDrawing()
	}
}
