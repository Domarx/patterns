package main

import (
	"github.com/domarx/patterns/structural/bridge/device"
)

func main() {
	remote := device.NewRemote(device.NewTV())
	remote.TogglePower()
	remote.IncreaseVolume()
	remote.DecreaseVolume()
	remote.TogglePower()

	remote.ChangeDevice(device.NewRadio())
	remote.TogglePower()
	remote.IncreaseVolume()
	remote.DecreaseVolume()
	remote.TogglePower()

	modernRemote := device.NewModernRemote(device.NewRadio())
	modernRemote.TogglePower()
	modernRemote.IncreaseVolume()
	modernRemote.IncreaseVolume()
	modernRemote.DecreaseVolume()
	modernRemote.Mute()
	modernRemote.UnMute()
	modernRemote.TogglePower()

	modernRemote.ChangeDevice(device.NewTV())
	modernRemote.TogglePower()
	modernRemote.IncreaseVolume()
	modernRemote.IncreaseVolume()
	modernRemote.Mute()
	modernRemote.UnMute()
	modernRemote.DecreaseVolume()
	modernRemote.TogglePower()
}
