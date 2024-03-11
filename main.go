package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/enemy"
	pl "github.com/vargaadam23/invaders/player"
)

const windowHeight = 800
const windowWidth = 1000

func main() {
	rl.InitWindow(windowWidth, windowHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	player := pl.InitCharacter(windowWidth, windowHeight)
	enemy := enemy.InitCharacter(windowWidth, windowHeight)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		player.HandleMovement()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		AddWalls(10)

		player.DrawCharacter()

		player.RenderCharacterBullets()

		enemy.DrawCharacter()

		enemy.RenderCharacterBullets()

		rl.EndDrawing()
	}
}

func AddWalls(wallWidth int32) {
	rl.DrawRectangle(0, 0, windowWidth, wallWidth, rl.Black)
	rl.DrawRectangle(0, 0, wallWidth, windowHeight, rl.Black)
	rl.DrawRectangle(0, windowHeight-wallWidth, windowWidth, wallWidth, rl.Black)
	rl.DrawRectangle(windowWidth-wallWidth, 0, wallWidth, windowHeight, rl.Black)
}
