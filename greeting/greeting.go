package greeting

import (
	"fmt"

	"github.com/raye17/subDemo/utils"
)

func Hello(name string) string {
	fmt.Println("this is subDemo")
	fmt.Println("hhhhhhh")
	fmt.Println("wwwwwwwwwwwwww")
	return "Hello, World!,my name is " + name
}
func Hello2(name string) {
	utils.Say()
	fmt.Println("this is subDemo2")
}
