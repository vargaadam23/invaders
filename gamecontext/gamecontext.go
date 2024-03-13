package gamecontext

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/character"
	"github.com/vargaadam23/invaders/enemy"
	"github.com/vargaadam23/invaders/player"
)

const wallWidth = 10

// TODO: Implement game context, will hold wall information and other params needed for player and enemies, collidable objects
type GameContext struct {
	Player  *player.Player
	Enemies []*enemy.Enemy
	Window  *GameWindow
}

type GameWindow struct {
	Pos        rl.Rectangle
	wallWidth  int32
	wallColor  rl.Color
	LeftWall   *rl.Rectangle
	RightWall  *rl.Rectangle
	TopWall    *rl.Rectangle
	BottomWall *rl.Rectangle
}

type Wall struct {
	rl.Rectangle
}

func (wall *Wall) GetHitbox() rl.Rectangle {
	return wall.Rectangle
}

func (window *GameWindow) DrawWindow() {
	rl.ClearBackground(rl.RayWhite)

	rl.DrawRectangleRec(*window.LeftWall, window.wallColor)
	rl.DrawRectangleRec(*window.RightWall, window.wallColor)
	rl.DrawRectangleRec(*window.TopWall, window.wallColor)
	rl.DrawRectangleRec(*window.BottomWall, window.wallColor)
}

func InitGameContext() *GameContext {
	window := &GameWindow{
		wallWidth:  wallWidth,
		wallColor:  rl.Black,
		LeftWall:   nil,
		TopWall:    nil,
		RightWall:  nil,
		BottomWall: nil,
		Pos: rl.Rectangle{
			Width:  1000,
			Height: 800,
			X:      0,
			Y:      0,
		},
	}

	window.LeftWall = &rl.Rectangle{
		Width:  float32(window.wallWidth),
		Height: window.Pos.Height,
		X:      0,
		Y:      0,
	}

	window.TopWall = &rl.Rectangle{
		Width:  float32(window.Pos.Width),
		Height: wallWidth,
		X:      0,
		Y:      0,
	}

	window.RightWall = &rl.Rectangle{
		Width:  window.Pos.Width,
		Height: wallWidth,
		X:      0,
		Y:      window.Pos.Height - wallWidth,
	}

	window.BottomWall = &rl.Rectangle{
		Width:  float32(window.wallWidth),
		Height: window.Pos.Height,
		X:      window.Pos.Width - wallWidth,
		Y:      0,
	}

	player := player.InitCharacter(int32(window.Pos.Width), int32(window.Pos.Height))

	enemyCollidables := []character.Collidable{}
	enemyCollidables = append(enemyCollidables, player)

	enemies := []*enemy.Enemy{}
	enemies = append(enemies, enemy.InitCharacter(int32(window.Pos.Width), int32(window.Pos.Height), enemyCollidables))

	player.Collidables = append(player.Collidables, enemies[0])

	return &GameContext{
		Player:  player,
		Enemies: enemies,
		Window:  window,
	}
}

func (context *GameContext) DrawGameElements() {
	context.Window.DrawWindow()

	context.Player.DrawCharacter()

	for _, enemy := range context.Enemies {
		enemy.DrawCharacter()
	}
}
