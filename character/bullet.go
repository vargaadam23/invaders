package character

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"

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
	uuid             string
	Position         *rl.Rectangle
	IsSpawned        bool
	Speed            int32
	Size             int32
	Color            color.RGBA
	Direction        int32
	Type             gameobject.ObjectType
	markedForDestroy bool
}

func (bullet Bullet) GetPosition() rl.Rectangle {
	return *bullet.Position
}

func (bullet Bullet) GetObjectType() gameobject.ObjectType {
	return bullet.Type
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
		uuid: uuid.NewString(),
		Position: &rl.Rectangle{
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
		Direction:        Direction,
		Type:             obtype,
		markedForDestroy: false,
	}
}

func (bullet *Bullet) advanceBullet() {
	bullet.Position.Y -= float32(bullet.Direction * bullet.Speed)
}

func (bullet *Bullet) RenderBullet() bool {
	if bullet.IsSpawned {
		rl.DrawRectangleRec(*bullet.Position, bullet.Color)
		bullet.advanceBullet()
	}

	if bullet.Position.Y < 0 || bullet.Position.Y > 800 {
		bullet.IsSpawned = false
		return false
	}

	return true
}

func (bullet *Bullet) GetHandlerByType(obType gameobject.ObjectType) gameobject.CollisionHandlerFunction {
	if bullet.GetObjectType() == gameobject.ENEMY_BULLET {
		switch obType {
		case gameobject.ENEMY:
			return func() {

			}
		case gameobject.PLAYER:
			return func() {
				bullet.SetMarkedForDestroy()
				fmt.Println("Enemy bullet hit player!")
			}
		case gameobject.WALL:
			return func() {
				bullet.SetMarkedForDestroy()
				fmt.Println("Bullet hit wall, will be destoryed in the next frame!")
			}
		}
	} else {
		switch obType {
		case gameobject.ENEMY:
			return func() {
				bullet.SetMarkedForDestroy()
				fmt.Println("PLayer bullet hit player!")
			}
		case gameobject.PLAYER:
			return func() {

			}
		case gameobject.WALL:
			return func() {
				bullet.SetMarkedForDestroy()
				fmt.Println("Bullet hit wall, will be destoryed in the next frame!")
			}
		}
	}

	return func() {}
}

func (bullet Bullet) GetObjectTypeString() string {
	if bullet.GetObjectType() == gameobject.ENEMY_BULLET {
		return "ENEMY_BULLET"
	} else {
		return "PLAYER_BULLET"
	}
}

func (bullet Bullet) GetUuid() string {
	return bullet.uuid
}

func (bullet Bullet) GetMarkedForDestroy() bool {
	return bullet.markedForDestroy
}

func (bullet *Bullet) SetMarkedForDestroy() {
	bullet.markedForDestroy = true
}
