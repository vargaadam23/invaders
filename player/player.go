package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/character"
)

// Rectangle hitbox

type Player struct {
	character.Character
}

// TODO: Add wall values
func (player *Player) HandleMovement() {
	if rl.IsKeyDown(rl.KeyLeft) && player.Hitbox.X > 10 {
		player.Hitbox.X = player.Hitbox.X - float32(player.Speed)
	}

	if rl.IsKeyDown(rl.KeyRight) && player.Hitbox.X < 1000-10-player.Hitbox.Width {
		player.Hitbox.X = player.Hitbox.X + float32(player.Speed)
	}

	if rl.IsKeyDown(rl.KeyDown) && player.Hitbox.Y < 800-10-player.Hitbox.Height {
		player.Hitbox.Y = player.Hitbox.Y + float32(player.Speed)
	}

	if rl.IsKeyDown(rl.KeyUp) && player.Hitbox.Y > 10 {
		player.Hitbox.Y = player.Hitbox.Y - float32(player.Speed)
	}
}

func (character *Player) GetHitbox() rl.Rectangle {
	return character.Hitbox
}

func (character *Player) GetType() string {
	return "PLAYER"
}

func (player *Player) DrawCharacter() {
	rl.DrawRectangle(int32(player.Hitbox.X), int32(player.Hitbox.Y), int32(player.Hitbox.Width), int32(player.Hitbox.Height), player.Color)

	player.RenderCharacterBullets()
}

func (player *Player) fireBullet() {
	player.Bullets.Append(character.NewBullet(player.Hitbox.X, player.Hitbox.Y, 1, []character.CollisionEvent{}))
}

func (player *Player) RenderCharacterBullets() {
	if rl.IsKeyPressed(rl.KeySpace) {
		player.fireBullet()
	}

	current := player.Bullets.Head

	for current.Next != nil {

		if !current.Value.RenderBullet() {
			current = current.Unlink(player.Bullets)
		} else {
			current = current.Next
		}
	}
}

func InitCharacter(windowWidth, windowHeight int32) *Player {
	return &Player{
		Character: character.Character{
			Hitbox: rl.Rectangle{
				X:      float32(windowWidth)/2 - 50,
				Y:      float32(windowHeight) - 200,
				Width:  200,
				Height: 100,
			},
			Color:   rl.Blue,
			Speed:   10,
			Bullets: character.InitLinkedList(),
			IsAlive: true,
		},
	}
}
