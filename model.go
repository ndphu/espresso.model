package model

import (
	lirc "github.com/ndphu/lirc"
)

type IREvent struct {
	DeviceId string
	Event    lirc.Event
}

type GPIOCommand struct {
	Pin   int
	State string
}
