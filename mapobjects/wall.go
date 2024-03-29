package mapobjects

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/vargaadam23/invaders/gameobject"
)

type Wall struct {
	uuid      string
	Pos       *rl.Rectangle
	WallWidth int
	WallColor rl.Color
}

func (wall Wall) GetPosition() rl.Rectangle {
	return *wall.Pos
}

func (wall Wall) GetObjectType() gameobject.ObjectType {
	return gameobject.WALL
}

func (wall Wall) DrawObject() {
	rl.DrawRectangleRec(*wall.Pos, wall.WallColor)
}

func InitWall(wallWidth int, position string, window rl.Rectangle) *Wall {
	var wall rl.Rectangle

	switch position {
	case "TOP":
		wall = rl.Rectangle{
			Width:  float32(window.Width),
			Height: float32(wallWidth),
			X:      0,
			Y:      0,
		}

		break
	case "BOTTOM":
		wall = rl.Rectangle{
			Width:  float32(wallWidth),
			Height: window.Height,
			X:      window.Width - float32(wallWidth),
			Y:      0,
		}

		break
	case "LEFT":
		wall = rl.Rectangle{
			Width:  float32(wallWidth),
			Height: window.Height,
			X:      0,
			Y:      0,
		}

		break
	case "RIGHT":
		wall = rl.Rectangle{
			Width:  window.Width,
			Height: float32(wallWidth),
			X:      0,
			Y:      window.Height - float32(wallWidth),
		}

		break
	}

	return &Wall{
		uuid:      uuid.NewString(),
		Pos:       &wall,
		WallWidth: wallWidth,
		WallColor: rl.Black,
	}
}

func (wall Wall) GetHandlerByType(obType gameobject.ObjectType) gameobject.CollisionHandlerFunction {
	switch obType {
	case gameobject.PLAYER_BULLET:
		return func() {
			fmt.Println("Player bullet hit wall!")
		}
	case gameobject.ENEMY_BULLET:
		return func() {
			fmt.Println("Enemy bullet hit wall!")
		}
	}

	return func() {}
}

func (wall Wall) GetObjectTypeString() string {
	return "WALL"
}

func (wall Wall) GetUuid() string {
	return wall.uuid
}

func (wall Wall) GetMarkedForDestroy() bool {
	return false
}

func (wall *Wall) SetMarkedForDestroy() {

}
