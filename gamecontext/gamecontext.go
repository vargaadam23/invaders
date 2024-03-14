package gamecontext

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/collision"
	"github.com/vargaadam23/invaders/gameobject"
)

const wallWidth = 10

var lock = &sync.Mutex{}
var gameContextInstance *GameContext

// TODO: Implement game context, will hold wall information and other params needed for player and enemies, collidable objects
type GameContext struct {
	GameObjects *gameobject.LinkedList
	Window      *GameWindow
}

type GameWindow struct {
	Width  float32
	Height float32
}

func initGameContext() *GameContext {
	window := &GameWindow{
		Width:  1000,
		Height: 800,
	}

	gameObjects := gameobject.InitLinkedList()

	return &GameContext{
		GameObjects: gameObjects,
		Window:      window,
	}
}

// Singleton, can get game context instance everywhere
func GetInstance() *GameContext {
	if gameContextInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if gameContextInstance == nil {
			gameContextInstance = initGameContext()
		}
	}

	return gameContextInstance
}

func (context *GameContext) DrawGameElements() {
	rl.ClearBackground(rl.RayWhite)

	current := context.GameObjects.Head
	collidable := context.GameObjects.Head

	for current != nil {
		current.Value.DrawObject()

		collidable = current.Next

		for collidable != nil {
			collision.CheckCollision(current.Value, collidable.Value)
			collidable = collidable.Next
		}

		current = current.Next
	}
}
