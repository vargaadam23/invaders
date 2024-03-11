package character

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Rectangle hitbox

type Character struct {
	Hitbox  rl.Rectangle
	Color   rl.Color
	Speed   int32
	Bullets *LinkedList
	IsAlive bool
}

type RenderableCharacter interface {
	DrawCharacter()
	RenderCharacterBullets()
	InitCharacter() *Character
}
