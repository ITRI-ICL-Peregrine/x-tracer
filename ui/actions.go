package ui

import (
	"github.com/Sheenam3/x-tracer-gocui/events"
	"github.com/jroimartin/gocui"
	"time"
)


var NAMESPACES_DISPLAYED bool = false

// Global action: Quit

func actionGlobalQuit(g *gocui.Gui, v *gocui.View) error {

	if err := deletePod("x-agent"); err != nil {
		return err
	}

	return gocui.ErrQuit
}


// View namespaces: Toggle display
func actionGlobalToggleViewNamespaces(g *gocui.Gui, v *gocui.View) error {
	vn := "namespaces"

	if !NAMESPACES_DISPLAYED {
		g.SetViewOnTop(vn)
		g.SetCurrentView(vn)
		changeStatusContext(g, "SE")
	} else {
		g.SetViewOnBottom(vn)
		g.SetCurrentView("pods")
		changeStatusContext(g, "D")
	}

	NAMESPACES_DISPLAYED = !NAMESPACES_DISPLAYED

	return nil
}

// View pods: Up
func actionViewPodsUp(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorUp(g, v, 2)
	return nil
}

// View pods: Down
func actionViewPodsDown(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorDown(g, v, false)
	return nil
}

//Display Probe Tools after Pod select
func actionViewPodsSelect(g *gocui.Gui, v *gocui.View) error {
	line, err := getViewLine(g, v)
	if err != nil {
		return err
	}
	LOG_MOD = "pod"
	errr := showSelectProbe(g)

	changeStatusContext(g, "SL")
	displayConfirmation(g, line+" Pod selected")
	return errr

}

// View pods: Delete
func actionViewPodsDelete(g *gocui.Gui, v *gocui.View) error {
	p, err := getSelectedPod(g)
	if err != nil {
		return err
	}

	if err := deletePod(p); err != nil {
		return err
	}

	displayConfirmation(g, p+" pod deleted")

	go viewPodsRefreshList(g)

	return nil
}


// View pod logs: Up
func actionViewPodsLogsUp(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorUp(g, v, 0)
	events.PublishEvent("logs:refresh", events.EmptyMessage{})
	return nil
}

// View pod logs: Down
func actionViewPodsLogsDown(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorDown(g, v, false)
	events.PublishEvent("logs:refresh", events.EmptyMessage{})
	return nil
}

// View logs: Hide
func actionViewPodsLogsHide(g *gocui.Gui, v *gocui.View) error {
	g.SetViewOnBottom("logs")
	g.SetViewOnBottom("logs-containers")
	g.SetCurrentView("pods")
	v.Clear()
	changeStatusContext(g, "D")
	return nil
}

// View namespaces: Up
func actionViewNamespacesUp(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorUp(g, v, 0)
	return nil
}

// View namespaces: Down
func actionViewNamespacesDown(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorDown(g, v, false)
	return nil
}

// Namespace: Choose
func actionViewNamespacesSelect(g *gocui.Gui, v *gocui.View) error {
	line, err := getViewLine(g, v)
	NAMESPACE = line
	go viewPodsRefreshList(g)
	actionGlobalToggleViewNamespaces(g, v)
	displayConfirmation(g, line+" namespace selected")
	return err
}

// Probes:  Choose
func actionViewProbesSelect(g *gocui.Gui, v *gocui.View) error {
	line, err := getViewLine(g, v)
	LOG_MOD = "probe"

	G, p, lv := showViewPodsLogs(g)
	displayConfirmation(g, line+" probe selected")
	startAgent(G, p, lv, line)
	G.SetViewOnTop("logs")
	G.SetCurrentView("logs")
	return err
}

func actionViewProbesList(g *gocui.Gui, v *gocui.View) error {
	if err := deletePod("x-agent"); err != nil {
		return err
	}
	LOG_MOD = "pod"
	errr := showSelectProbe(g)
	changeStatusContext(g, "SL")
	time.Sleep(10 * time.Second)
	return errr
}
