package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (tp *TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (tp *TapePlayer) Stop() {
	fmt.Println("Stopped")
}
