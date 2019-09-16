package gadget

import "fmt"

type TapeRecorder struct {
	Microphones int
}

func (tr TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (tr TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (tr TapeRecorder) Stop() {
	fmt.Println("Stopped")
}
