package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/domarx/patterns/creational/singleton"
)

func main() {
	reader := bufio.NewReader(os.Stdout)

	fmt.Println("Enter your name to be greeted")
	name, _ := reader.ReadString('\n')

	var greeter singleton.Greeter
	for i := 0; i < 10; i++ {
		go func() {
			greeter = singleton.Create()
			greeter.Greet(name)
		}()
	}

	fmt.Println("press any key to exit")
	_, _ = reader.ReadByte()
}
