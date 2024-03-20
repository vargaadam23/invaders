package collision

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vargaadam23/invaders/gameobject"
)

func CheckCollision(objOne, objTwo gameobject.GameObject) {
	if !(objOne.GetObjectType() == gameobject.WALL && objTwo.GetObjectType() == gameobject.WALL) && rl.CheckCollisionRecs(objOne.GetPosition(), objTwo.GetPosition()) {
		fmt.Print(objOne.GetObjectTypeString() + "-" + objTwo.GetObjectTypeString() + "\n")
		objOne.GetHandlerByType(objTwo.GetObjectType())()
		objTwo.GetHandlerByType(objOne.GetObjectType())()
	}
}
