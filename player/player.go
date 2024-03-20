package player

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/vargaadam23/invaders/character"
	"github.com/vargaadam23/invaders/gamecontext"

	"github.com/vargaadam23/invaders/gameobject"
)

// Rectangle hitbox

type Player struct {
	uuid             string
	Hitbox           *rl.Rectangle
	Speed            int
	markedForDestroy bool
}

func (player Player) GetPosition() rl.Rectangle {
	return *player.Hitbox
}

func (player Player) GetObjectType() gameobject.ObjectType {
	return gameobject.PLAYER
}

func (player *Player) DrawObject() {
	player.HandleMovement()

	player.RenderCharacterBullets()
	rl.DrawRectangleRec(*player.Hitbox, rl.SkyBlue)
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

func (player *Player) fireBullet() {
	bullette := character.InitBullet(player.Hitbox.X, player.Hitbox.Y, 1, gameobject.PLAYER)
	gamecontext.GetInstance().GameObjects.Append(bullette)
	fmt.Print(bullette.Type)
}

func (player *Player) RenderCharacterBullets() {
	if rl.IsKeyPressed(rl.KeySpace) {
		player.fireBullet()
	}
}

func InitCharacter(windowWidth, windowHeight int32) *Player {
	return &Player{
		uuid: uuid.NewString(),
		Hitbox: &rl.Rectangle{
			X:      float32(windowWidth)/2 - 50,
			Y:      float32(windowHeight) - 200,
			Width:  200,
			Height: 100,
		},
		Speed:            10,
		markedForDestroy: false,
	}
}

func (player *Player) GetHandlerByType(obType gameobject.ObjectType) gameobject.CollisionHandlerFunction {
	switch obType {
	case gameobject.PLAYER_BULLET:
		return func() {

		}
	}

	return func() {}
}

func (player Player) GetObjectTypeString() string {
	return "PLAYER"
}

func (player Player) GetUuid() string {
	return player.uuid
}

func (player Player) GetMarkedForDestroy() bool {
	return player.markedForDestroy
}

func (player *Player) SetMarkedForDestroy() {
	player.markedForDestroy = true
}
