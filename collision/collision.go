package collision

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/gameobject"
)

func CheckCollision(objOne, objTwo gameobject.GameObject) {
	if rl.CheckCollisionRecs(objOne.GetPosition(), objTwo.GetPosition()) {
		objOne.GetHandlerByType(objTwo.GetObjectType())()
		objTwo.GetHandlerByType(objOne.GetObjectType())()
	}
}
