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

func (player *Player) DrawCharacter() {
	rl.DrawRectangle(int32(player.Hitbox.X), int32(player.Hitbox.Y), int32(player.Hitbox.Width), int32(player.Hitbox.Height), player.Color)
}

func (player *Player) RenderCharacterBullets() {
	if rl.IsKeyPressed(rl.KeySpace) {
		player.Bullets.Append(character.NewBullet(int32(player.Hitbox.X), int32(player.Hitbox.Y), 1))
		player.Bullets.Append(character.NewBullet(int32(player.Hitbox.X+player.Hitbox.Width), int32(player.Hitbox.Y), 1))
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
