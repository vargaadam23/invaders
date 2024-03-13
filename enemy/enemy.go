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
	enemy.RenderCharacterBullets()
}

func getCollisionEventArray() {

}

func (enemy *Enemy) SpawnEnemyBullet() {
	if enemy.IsAlive && enemy.BulletTimer.TimerDone() {
		enemy.Bullets.Append(character.NewBullet(enemy.Hitbox.X, enemy.Hitbox.Y, -1, []character.CollisionEvent{}))
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

func (character *Enemy) GetHitbox() rl.Rectangle {
	return character.Hitbox
}

func (character *Enemy) GetType() string {
	return "ENEMY"
}

func InitCharacter(windowWidth, windowHeight int32, collidables []character.Collidable) *Enemy {
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
			Color:       rl.Green,
			Speed:       10,
			Bullets:     character.InitLinkedList(),
			IsAlive:     true,
			Collidables: collidables,
		},
		BulletTimer: timer,
	}
}
