package systray

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/adrielparedes/kogito-local-server/pkg/images"
	"github.com/adrielparedes/kogito-local-server/pkg/server"
	"github.com/getlantern/systray"
)

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
	return proxy.CheckStatus()
}

func onReady() {

	systray.SetTemplateIcon(images.Data, images.Data)
	systray.SetTooltip(NAME)

	openModeler := systray.AddMenuItem(BUSINESS_MODELER, BUSINESS_MODELER)

	statusItem := systray.AddMenuItem(SERVER_STATUS, SERVER_STATUS)

	systray.AddMenuItem(OTHER_KOGITO_SERVICES, OTHER_KOGITO_SERVICES)

	systray.AddSeparator()

	toggleItem := systray.AddMenuItem(START, START)
	restartItem := systray.AddMenuItem(RESTART, RESTART)

	systray.AddSeparator()
	quitItem := systray.AddMenuItem(QUIT, QUIT)

	go CheckAndStatus(toggleItem, statusItem)

	for {
		select {
		case <-openModeler.ClickedCh:
			openBrowser(MODELER_LINK)
		case <-restartItem.ClickedCh:
			if CheckStatus() {
				Stop()
				Start()
			} else {
				Start()
			}
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

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
