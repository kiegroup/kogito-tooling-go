package systray

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adrielparedes/kogito-local-server/pkg/images"
	"github.com/adrielparedes/kogito-local-server/pkg/server"
	"github.com/getlantern/systray"
)

const NAME = "Kogito"
const URL = "http://127.0.0.1:8000"
const SERVER_STATUS = "Server Status"
const SERVER_STATUS_ON = SERVER_STATUS + ": ON"
const SERVER_STATUS_OFF = SERVER_STATUS + ": OFF"
const START = "Start"
const STOP = "Stop"
const QUIT = "Quit"

var proxy *server.Proxy = &server.Proxy{}

func Systray() {
	onExit := func() {
	}

	systray.Run(onReady, onExit)

}

func ChangeStatus(started bool, item *systray.MenuItem) {
	if started {
		item.SetTitle(SERVER_STATUS_ON)
	} else {
		item.SetTitle(SERVER_STATUS_OFF)
	}
}

func ChangeName(started bool, item *systray.MenuItem) {
	if started {
		item.SetTitle(STOP)
	} else {
		item.SetTitle(START)
	}
}

func Start() {
	fmt.Println("Executing Start command")
	proxy.Start()
}

func Stop() {
	fmt.Println("Executing Stop command")
	proxy.Stop()
}

func CheckAndStatus(toggleItem *systray.MenuItem, statusItem *systray.MenuItem) {
	for true {
		started := CheckStatus()
		fmt.Printf("Checking status: %v\n", started)
		ChangeName(started, toggleItem)
		ChangeStatus(started, statusItem)
		time.Sleep(time.Second * 5)
	}
}

func CheckStatus() bool {
	started := false

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(resp.StatusCode) + resp.Status)
		if resp.StatusCode == 200 {
			started = true
		}
	}

	return started
}

func onReady() {

	systray.SetTemplateIcon(images.Data, images.Data)
	systray.SetTooltip(NAME)

	statusItem := systray.AddMenuItem(SERVER_STATUS_OFF, SERVER_STATUS)

	systray.AddSeparator()

	toggleItem := systray.AddMenuItem(START, START)

	systray.AddSeparator()
	quitItem := systray.AddMenuItem(QUIT, QUIT)

	go CheckAndStatus(toggleItem, statusItem)

	for {
		select {
		case <-toggleItem.ClickedCh:
			if CheckStatus() {
				Stop()
			} else {
				Start()
			}
		case <-quitItem.ClickedCh:
			Stop()
			systray.Quit()
			return
		}
	}
}
