package kogito

func Systray(port int, runner []byte) {
	proxy := NewProxy(port, runner)
	proxy.view = &KogitoSystray{}
	proxy.view.controller = proxy
	proxy.view.Run()
}
