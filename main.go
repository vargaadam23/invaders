package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/enemy"
	"github.com/vargaadam23/invaders/gamecontext"
	"github.com/vargaadam23/invaders/player"
)

const windowHeight = 800
const windowWidth = 1000

func main() {
	rl.InitWindow(windowWidth, windowHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	context := gamecontext.GetInstance()
	player := player.InitCharacter(int32(context.Window.Width), int32(context.Window.Height))

	obenemy := enemy.InitCharacter(int32(context.Window.Width), int32(context.Window.Height))

	context.GameObjects.Append(player)
	context.GameObjects.Append(obenemy)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		context.DrawGameElements()

		rl.DrawFPS(100, 100)
		rl.EndDrawing()
	}
}
