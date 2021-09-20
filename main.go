package main

import (
	"fmt"

	"github.com/dsp-testing/dependency-graph-snek-test/pkg/youneedthis"
)

func main() {
	msg := youneedthis.NewMessage("A Hubber", "Hello DSP! We're all depending on you!")
	fmt.Println(msg)
}
