package random

import (
	"fmt"
	"math/rand"
	"time"
)

func GetCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}
