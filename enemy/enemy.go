package enemy

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/character"
	"github.com/vargaadam23/invaders/util"
)

type Enemy struct {
	character.Character
	BulletTimer *util.Timer
}

func (enemy *Enemy) DrawCharacter() {
	rl.DrawRectangle(int32(enemy.Hitbox.X), int32(enemy.Hitbox.Y), int32(enemy.Hitbox.Width), int32(enemy.Hitbox.Height), enemy.Color)
}

func (enemy *Enemy) SpawnEnemyBullet() {
	if enemy.IsAlive && enemy.BulletTimer.TimerDone() {
		enemy.Bullets.Append(character.NewBullet(int32(enemy.Hitbox.X), int32(enemy.Hitbox.Y), -1))
		enemy.BulletTimer.StartTimer(2)
	}
}

func (enemy *Enemy) RenderCharacterBullets() {
	enemy.SpawnEnemyBullet()

	current := enemy.Bullets.Head

	for current.Next != nil {

		if !current.Value.RenderBullet() {
			current = current.Unlink(enemy.Bullets)
		} else {
			current = current.Next
		}
	}
}

func InitCharacter(windowWidth, windowHeight int32) *Enemy {
	timer := &util.Timer{}
	timer.StartTimer(2)

	return &Enemy{
		Character: character.Character{
			Hitbox: rl.Rectangle{
				X:      float32(windowWidth)/2 - 50,
				Y:      200,
				Width:  50,
				Height: 50,
			},
			Color:   rl.Green,
			Speed:   10,
			Bullets: character.InitLinkedList(),
			IsAlive: true,
		},
		BulletTimer: timer,
	}
}
