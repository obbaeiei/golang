package golang

import (
	"math/rand"
	"time"
)

var problem int

func scrum() {
	rand.Seed(time.Now().Unix())
	problem = rand.Int()
}
