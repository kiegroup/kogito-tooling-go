package kogito

func Systray(port int) {

	proxy := NewProxy(port)
	proxy.view = &KogitoSystray{}
	proxy.view.controller = proxy
	proxy.view.Run()
}
