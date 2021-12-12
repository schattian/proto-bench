package protobench

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var testSize int

func init() {
	rand.Seed(time.Now().UnixNano())
	setSize()
}

func randomI(n int) []int {
	randomI := make([]int, n)
	for i := 0; i < len(randomI); i++ {
		randomI[i] = i
	}
	rand.Shuffle(len(randomI), func(i, j int) {
		randomI[i], randomI[j] = randomI[j], randomI[i]
	})
	return randomI
}

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func setSize() {
	sz, _ := strconv.Atoi(os.Getenv("SIZE"))
	if sz == 0 {
		sz = 1
	}
	testSize = sz
}
