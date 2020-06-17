package main

import (
	"github.com/ClashrAuto/alien/libs/clash"
	"github.com/ClashrAuto/alien/libs/window"
	"github.com/andlabs/ui"
)

func main() {
	clash.Clash()
	ui.Main(window.SetupUI)
}
