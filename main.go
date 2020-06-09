package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ClashrAuto/Clashr/config"
	C "github.com/ClashrAuto/Clashr/constant"
	"github.com/ClashrAuto/Clashr/hub"

	log "github.com/sirupsen/logrus"

	"gioui.org/ui/app"

	"github.com/ireoo/alien/libs/clash"
	"github.com/ireoo/alien/libs/ui"
)

func main() {

	go func() {
		Clash()
		Windows()
	}()

	// sigCh := make(chan os.Signal, 1)
	// signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	// <-sigCh
	app.Main()
}
