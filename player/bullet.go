package player

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const speed = 5
const size = 10.00

var colorb color.RGBA = rl.Red

type Position struct {
	X int32
	Y int32
}

type Bullet struct {
	Position  Position
	IsSpawned bool
	Speed     int32
	Size      float32
	Color     color.RGBA
}

func NewBullet(X, Y int32) *Bullet {
	return &Bullet{
		Position: Position{
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
		rl.DrawCircle(bullet.Position.X, bullet.Position.Y, bullet.Size, bullet.Color)
		bullet.advanceBullet()
	}

	if bullet.Position.Y < 0 {
		bullet.IsSpawned = false
		return false
	}

	return true
}
