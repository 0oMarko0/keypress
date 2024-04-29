package main

import "github.com/micmonay/keybd_event"

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
