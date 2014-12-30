package control

import (
	"github.com/intelsdilabs/gomit"
)

type pluginController struct {
	EventEmitter *gomit.Emitter
}

func NewControl() *pluginController {
	p := new(pluginController)
	p.EventEmitter = new(gomit.Emitter)

	return p
}

type ZEvent struct {
	Name, Address string
}

type YEvent struct {
	Name, Address string
}

func (p *pluginController) SendZEvent() {
	z := ZEvent{"nick", "home"}
	// metadata parameters
	m := make(map[string]string)
	m["debug"] = "true"

	event := gomit.NewEvent(z, m)

	p.EventEmitter.Fire(event)
}

func (p *pluginController) SendYEvent() {
	z := YEvent{"daniel", "coffehause"}
	// metadata parameters
	m := make(map[string]string)
	m["debug"] = "true"

	event := gomit.NewEvent(z, m)

	p.EventEmitter.Fire(event)
}

func (p *pluginController) SendXEvent() {
	z := ZEvent{"justin", "pub"}
	// metadata parameters
	m := make(map[string]string)
	m["debug"] = "true"

	event := gomit.NewEvent(z, m)

	p.EventEmitter.Fire(event)
}
