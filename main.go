package main

import (
	"log"
	"os/exec"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Stand Up !!!! ")
	c := 2900

	standup := widget.NewLabel("Press the button to countdown")

	go func() {
		for {
			for {

				if c <= 0 {
					break
				}
				t := strconv.Itoa(c)
				standup.SetText("counting:" + t)

				c -= 1
				time.Sleep(1 * time.Second)
			}

			standup.SetText("Time to stand up !!!!")
			cmd := exec.Command("afplay", "/System/Library/Sounds/Ping.aiff")
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(100 * time.Second)
			cmd = exec.Command("afplay", "/System/Library/Sounds/Glass.aiff")
			err = cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			c = 10
		}
	}()

	w.SetContent(container.New(layout.NewVBoxLayout(), standup))

	w.ShowAndRun()

}
