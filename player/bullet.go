package player

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const speed = 5
const size = 5.00

var colorb color.RGBA = rl.Red

type Point struct {
	X int32
	Y int32
}

type Bullet struct {
	Position  Point
	IsSpawned bool
	Speed     int32
	Size      int32
	Color     color.RGBA
}

func NewBullet(X, Y int32) *Bullet {
	return &Bullet{
		Position: Point{
			X: X,
			Y: Y,
		},
		IsSpawned: true,
		Speed:     speed,
		Size:      size,
		Color:     colorb,
	}
}

func (bullet *Bullet) advanceBullet() {
	bullet.Position.Y -= bullet.Speed
}

func (bullet *Bullet) RenderBullet() bool {
	if bullet.IsSpawned {
		rl.DrawRectangle(bullet.Position.X, bullet.Position.Y, bullet.Size, bullet.Size, bullet.Color)
		bullet.advanceBullet()
	}

	if bullet.Position.Y < 0 {
		bullet.IsSpawned = false
		return false
	}

	return true
}
