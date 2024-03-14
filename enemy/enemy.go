package enemy

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/vargaadam23/invaders/character"
	"github.com/vargaadam23/invaders/gamecontext"
	"github.com/vargaadam23/invaders/gameobject"
	"github.com/vargaadam23/invaders/util"
)

type Enemy struct {
	Hitbox      rl.Rectangle
	IsAlive     bool
	BulletTimer *util.Timer
	Color       rl.Color
	Speed       int
}

func (enemy Enemy) GetPosition() rl.Rectangle {
	return enemy.Hitbox
}

func (enemy Enemy) GetObjectType() gameobject.ObjectType {
	return gameobject.PLAYER
}

func (enemy *Enemy) DrawObject() {
	rl.DrawRectangle(int32(enemy.Hitbox.X), int32(enemy.Hitbox.Y), int32(enemy.Hitbox.Width), int32(enemy.Hitbox.Height), enemy.Color)
	enemy.SpawnEnemyBullet()
}

func (enemy *Enemy) GetHandlerByType(obType gameobject.ObjectType) gameobject.CollisionHandlerFunction {
	switch obType {
	case gameobject.PLAYER_BULLET:
		return func() {
			fmt.Println("Enemy collision with player bullet!")
		}
	case gameobject.PLAYER:
		return func() {
			fmt.Println("Enemy collision with player!")
		}
	case gameobject.ENEMY:
		return func() {
			fmt.Println("Enemy collision with enemy!")
		}
	case gameobject.ENEMY_BULLET:
		return func() {
			fmt.Println("Enemy collision with enemy bullet!")
		}
	}

	return func() {}
}

func (enemy *Enemy) SpawnEnemyBullet() {
	if enemy.IsAlive && enemy.BulletTimer.TimerDone() {
		gamecontext.GetInstance().GameObjects.Append(character.InitBullet(enemy.Hitbox.X, enemy.Hitbox.Y, -1, gameobject.ENEMY))
		enemy.BulletTimer.StartTimer(2)
	}
}

func InitCharacter(windowWidth, windowHeight int32) *Enemy {
	timer := &util.Timer{}
	timer.StartTimer(2)

	return &Enemy{
		Hitbox: rl.Rectangle{
			X:      float32(windowWidth)/2 - 50,
			Y:      200,
			Width:  50,
			Height: 50,
		},
		Color:       rl.Green,
		Speed:       10,
		IsAlive:     true,
		BulletTimer: timer,
	}
}
