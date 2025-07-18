package internal

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// fmt.Println("hook start...")
	// evChan := hook.Start()
	// defer hook.End()

	// for ev := range evChan {
	// 	fmt.Println("hook: ", ev)
	// 	if ev.Keychar == 'q' {
	// 		break
	// 	}
	// }
	fmt.Println(KeyByListen(Params{}, Config{}))
}
