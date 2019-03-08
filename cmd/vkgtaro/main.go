package main

import (
	"flag"
	"os"

	"github.com/vkg/taro/vkgtaro"
)

var (
	l = flag.Bool("little", false, "Little")
	s = flag.Bool("slow", false, "Slow")
)

func main() {
	flag.Parse()
	vkgtaro := vkgtaro.New(os.Stdout)
	if *l {
		vkgtaro.MoveLittle()
		return
	}
	if *s {
		vkgtaro.MoveSlow()
		return
	}
	vkgtaro.Move()
}
