package device

func NewRemote(device Device) *Remote {
	return &Remote{device: device}
}

type Remote struct {
	device Device
}

func (r *Remote) ChangeDevice(device Device) { r.device = device }

func (r *Remote) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *Remote) IncreaseVolume() { r.device.SetVolume(r.device.Volume() + 5) }

func (r *Remote) DecreaseVolume() { r.device.SetVolume(r.device.Volume() - 5) }

func NewModernRemote(device Device) *ModernRemote {
	return &ModernRemote{
		Remote:           Remote{device: device},
		volumeBeforeMute: 0,
	}
}

type ModernRemote struct {
	Remote
	volumeBeforeMute float64
}

func (m *ModernRemote) Mute() {
	m.volumeBeforeMute = m.device.Volume()
	m.device.SetVolume(0)
}

func (m *ModernRemote) UnMute() { m.device.SetVolume(m.volumeBeforeMute) }
