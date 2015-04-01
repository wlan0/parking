package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	names := []string{"alena", "sid", "will", "shengSr", "shengJr", "oleg", "kirill", "sangeetha", "denise"}

	seed := time.Now().Unix()
	src := rand.NewSource(seed)

	namemap := map[int]int{}

	for len(namemap) < 6 {
		rander := rand.New(src)
		randNum := rander.Uint32()
		namemap[int(randNum)%9] = 1
	}

	for i, _ := range namemap {
		n := time.Now().Unix() % int64(3500)
		duration := time.Duration(n) * time.Millisecond
		time.Sleep(duration)
		fmt.Println(names[i])
	}
}
