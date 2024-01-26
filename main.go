package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/player"
)

type BulletArray []*player.Bullet

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	x := 10
	y := -10

	var bulletArray BulletArray = BulletArray{}

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyLeft) {
			x = x - 10
		}

		if rl.IsKeyDown(rl.KeyRight) {
			x = x + 10
		}

		if rl.IsKeyDown(rl.KeyDown) {
			y = y + 10
		}

		if rl.IsKeyDown(rl.KeyUp) {
			y = y - 10
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyPressed(rl.KeySpace) {
			bulletArray = append(bulletArray, player.NewBullet(int32(x), int32(y)))
			fmt.Print(len(bulletArray))
		}

		for index, bullet := range bulletArray {
			if !bullet.RenderBullet() {
				bulletArray = RemoveIndex(bulletArray, index)
			}
		}

		rl.DrawCircle(int32(x), int32(y), 100.00, color.RGBA{
			100, 25, 25, 25,
		})

		rl.EndDrawing()
	}
}

func RemoveIndex(s BulletArray, index int) BulletArray {
	return append(s[:index], s[index+1:]...)
}
