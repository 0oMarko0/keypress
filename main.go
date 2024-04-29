package main

import "github.com/micmonay/keybd_event"

// This simulate a key press CTRL + '
func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	kb.SetKeys(keybd_event.VK_APOSTROPHE)
	kb.HasCTRL(true)

	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}
