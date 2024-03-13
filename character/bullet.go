package character

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const speed = 5
const size = 10.00

var colorb color.RGBA = rl.Red

type Point struct {
	X int32
	Y int32
}

type Bullet struct {
	Position        rl.Rectangle
	IsSpawned       bool
	Speed           int32
	Size            int32
	Color           color.RGBA
	Direction       int32
	CollisionEvents []CollisionEvent
}

type Collidable interface {
	GetHitbox() rl.Rectangle
	GetType() string
}

type CollisionEvent struct {
	Collidable Collidable
	Callback   func()
}

func NewBullet(X, Y float32, Direction int32, events []CollisionEvent) *Bullet {
	return &Bullet{
		Position: rl.Rectangle{
			X:      X,
			Y:      Y,
			Width:  size,
			Height: size,
		},
		IsSpawned: true,
		Speed:     speed,
		Size:      size,
		Color:     colorb,
		// -1 =>  DOWN | 1 => UP
		Direction:       Direction,
		CollisionEvents: events,
	}
}

func (bullet *Bullet) advanceBullet() {
	bullet.Position.Y -= float32(bullet.Direction * bullet.Speed)
}

func (bullet *Bullet) RenderBullet() bool {
	if bullet.IsSpawned {
		for _, event := range bullet.CollisionEvents {
			if rl.CheckCollisionRecs(event.Collidable.GetHitbox(), bullet.Position) {
				event.Callback()
				bullet.IsSpawned = false
				return false
			}
		}

		rl.DrawRectangleRec(bullet.Position, bullet.Color)
		bullet.advanceBullet()
	}

	if bullet.Position.Y < 0 || bullet.Position.Y > 800 {
		bullet.IsSpawned = false
		return false
	}

	return true
}
