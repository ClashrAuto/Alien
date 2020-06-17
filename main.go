package main

import (
	"github.com/ClashrAuto/alien/libs/clash"
	"github.com/andlabs/ui"
)

func main() {
	clash.Clash()
	ui.Main(window.setupUI)
}
