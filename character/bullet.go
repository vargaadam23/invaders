package character

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/gameobject"
)

const speed = 5
const size = 10.00

var colorb color.RGBA = rl.Red

type Point struct {
	X int32
	Y int32
}

type Bullet struct {
	Position  rl.Rectangle
	IsSpawned bool
	Speed     int32
	Size      int32
	Color     color.RGBA
	Direction int32
	Type      gameobject.ObjectType
}

func (bullet Bullet) GetPosition() rl.Rectangle {
	return bullet.Position
}

func (bullet Bullet) GetObjectType() gameobject.ObjectType {
	return gameobject.PLAYER
}

func (bullet *Bullet) DrawObject() {
	bullet.RenderBullet()
}

func InitBullet(X, Y float32, Direction int32, Issuer gameobject.ObjectType) *Bullet {
	var obtype gameobject.ObjectType

	if Issuer == gameobject.ENEMY {
		obtype = gameobject.ENEMY_BULLET
	} else {
		obtype = gameobject.PLAYER_BULLET
	}

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
		Direction: Direction,
		Type:      obtype,
	}
}

func (bullet *Bullet) advanceBullet() {
	bullet.Position.Y -= float32(bullet.Direction * bullet.Speed)
}

func (bullet *Bullet) RenderBullet() bool {
	if bullet.IsSpawned {
		rl.DrawRectangleRec(bullet.Position, bullet.Color)
		bullet.advanceBullet()
	}

	if bullet.Position.Y < 0 || bullet.Position.Y > 800 {
		bullet.IsSpawned = false
		return false
	}

	return true
}

func (bullet *Bullet) GetHandlerByType(obType gameobject.ObjectType) gameobject.CollisionHandlerFunction {
	switch obType {
	case gameobject.ENEMY:
		if bullet.GetObjectType() == gameobject.ENEMY_BULLET {
			return func() {
				fmt.Println("Player hit enemy with bullet!")
			}
		} else {
			return func() {
				fmt.Println("Enemy hit player with bullet!")
			}
		}

	}

	return func() {}
}
