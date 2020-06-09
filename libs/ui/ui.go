package ui

import "gioui.org/app"

func Windows() {
	go func() {
		w := app.NewWindow()
		for range w.Events() {

		}
	}()
	app.Main()
}
