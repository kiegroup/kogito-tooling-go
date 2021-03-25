package kogito

func Systray(port int) {

	proxy := &Proxy{}
	proxy.Port = port
	proxy.view = &KogitoSystray{}
	proxy.view.controller = proxy
	proxy.view.Run()
}
