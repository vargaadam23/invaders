package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	pl "github.com/vargaadam23/invaders/player"
)

const windowHeight = 800
const windowWidth = 1000

type BulletArray []*pl.Bullet

func main() {
	rl.InitWindow(windowWidth, windowHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	player := pl.InitPlayer(windowWidth, windowHeight)

	var bulletArray BulletArray = BulletArray{}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		player.HandleMovement()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyPressed(rl.KeySpace) {
			bulletArray = append(bulletArray, pl.NewBullet(int32(player.Hitbox.X), int32(player.Hitbox.Y)))
			bulletArray = append(bulletArray, pl.NewBullet(int32(player.Hitbox.X+player.Hitbox.Width), int32(player.Hitbox.Y)))
			fmt.Print(len(bulletArray))
		}

		for index, bullet := range bulletArray {
			if !bullet.RenderBullet() {
				bulletArray = RemoveIndex(bulletArray, index)
			}
		}

		AddWalls(10)

		player.DrawPlayer()

		rl.EndDrawing()
	}
}

func AddWalls(wallWidth int32) {
	rl.DrawRectangle(0, 0, windowWidth, wallWidth, rl.Black)
	rl.DrawRectangle(0, 0, wallWidth, windowHeight, rl.Black)
	rl.DrawRectangle(0, windowHeight-wallWidth, windowWidth, wallWidth, rl.Black)
	rl.DrawRectangle(windowWidth-wallWidth, 0, wallWidth, windowHeight, rl.Black)
}

func RemoveIndex(s BulletArray, index int) BulletArray {
	return append(s[:index], s[index+1:]...)
}
