package main

import (
	"fmt"
	"time"

	"github.com/intelsdilabs/gomit"
	"github.com/intelsdilabs/gomit/control"
)

func HandleStuff(e gomit.Event) {

	switch x := e.SourceEvent.(type) {
	case control.ZEvent:
		fmt.Printf("Z EVENT => name:%s & address:%s\n", x.Name, x.Address)
	case control.YEvent:
		fmt.Printf("Y EVENT => name:%s & address:%s\n", x.Name, x.Address)
	}
}

func main() {
	pluginCtrl := control.NewControl()

	fmt.Println(pluginCtrl)

	pluginCtrl.EventEmitter.AddHandler(HandleStuff)

	pluginCtrl.SendZEvent()
	pluginCtrl.SendYEvent()
	pluginCtrl.SendXEvent()

	time.Sleep(time.Second * 1)
}
