package runtime

import (
	"github.com/wailsapp/wails/v2/internal/servicebus"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

// Runtime is a means for the user to interact with the application at runtime
type Runtime struct {
	Browser     Browser
	Events      Events
	Window      Window
	Dialog      Dialog
	System      System
	Menu        Menu
	ContextMenu ContextMenus
	Tray        Tray
	Store       *StoreProvider
	Log         Log
	bus         *servicebus.ServiceBus
}

// New creates a new runtime
func New(serviceBus *servicebus.ServiceBus, menu *menu.Menu, trayMenu *menu.TrayOptions, contextMenus *menu.ContextMenus) *Runtime {
	result := &Runtime{
		Browser:     newBrowser(),
		Events:      newEvents(serviceBus),
		Window:      newWindow(serviceBus),
		Dialog:      newDialog(serviceBus),
		System:      newSystem(serviceBus),
		Menu:        newMenu(serviceBus, menu),
		Tray:        newTray(serviceBus, trayMenu),
		ContextMenu: newContextMenus(serviceBus, contextMenus),
		Log:         newLog(serviceBus),
		bus:         serviceBus,
	}
	result.Store = newStore(result)
	return result
}

// Quit the application
func (r *Runtime) Quit() {
	r.bus.Publish("quit", "runtime.Quit()")
}