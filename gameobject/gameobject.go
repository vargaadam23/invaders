package gameobject

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObjectType int

const (
	PLAYER ObjectType = iota
	ENEMY
	MAP_ELEMENT
	PLAYER_BULLET
	ENEMY_BULLET
	WALL
)

type CollisionHandlerFunction func()

type HandlerMap map[ObjectType]CollisionHandlerFunction

type GameObject interface {
	GetPosition() rl.Rectangle
	GetObjectType() ObjectType
	DrawObject()
	GetHandlerByType(obType ObjectType) CollisionHandlerFunction
}
