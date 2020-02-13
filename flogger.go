package main

import (
	"fmt"

	"github.com/BurntSushi/xgb"
)

func main() {
	fmt.Println("You're running KeyFlogger, press ESC to exit")

	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		ev, _ := X.WaitForEvent()
		fmt.Println(ev)
	}

	// fmt.Println("KeyFlogger out!")

}
