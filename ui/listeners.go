package ui

import (
	"fmt"
	"github.com/ITRI-ICL-Peregrine/x-tracer/events"
	"github.com/ITRI-ICL-Peregrine/x-tracer/pkg"
	"github.com/jroimartin/gocui"
)

func refreshIntegratedLogs(e events.Event) {

	if e, ok := e.(events.EmptyMessage); ok {

		g.Update(func(g *gocui.Gui) error {

			pn := e.Pn
			if pn == "tcplife" {
				viewtl, err := g.View("tcplife")
				if err != nil {
					return err
				}
				viewtl.Clear()

				_, _ = fmt.Fprint(viewtl, pkg.GetActiveLogs(pn))

				g.SetViewOnTop("tcplife")

				viewtl.Autoscroll = true

				return nil
			} else if pn == "cachestat" {
				viewcs, err := g.View("cachestat")
				if err != nil {
					return err
				}
				viewcs.Clear()

				_, _ = fmt.Fprint(viewcs, pkg.GetActiveLogs(pn))

				g.SetViewOnTop("cachestat")
				viewcs.Autoscroll = true

				return nil
			} else if pn == "execsnoop" {
				viewes, err := g.View("execsnoop")
				if err != nil {
					return err
				}
				viewes.Clear()

				_, _ = fmt.Fprint(viewes, pkg.GetActiveLogs(pn))

				g.SetViewOnTop("execsnoop")


				viewes.Autoscroll = true

				return nil
			} else {
				view, err := g.View("tcplogs")
				if err != nil {
					return err
				}
				view.Clear()

				_, _ = fmt.Fprint(view, pkg.GetActiveLogs(pn))
				g.SetViewOnTop("tcplogs")

				view.Autoscroll = true

				return nil
			}

		})
	}
}

func refreshSingleLogs(e events.Event) {

	if e, ok := e.(events.EmptyMessage); ok {

		g.Update(func(g *gocui.Gui) error {

			pn := e.Pn
			view, err := g.View("logs")
			if err != nil {
				return err
			}
			view.Clear()

			_, _ = fmt.Fprint(view, pkg.GetActiveLogs(pn))

			g.SetViewOnTop("logs")
			g.SetCurrentView("logs")

			view.Autoscroll = true

			return nil
		})
	}

}

func SubscribeListeners() {
	events.Subscribe(refreshIntegratedLogs, "logs:refreshinteg")
	events.Subscribe(refreshSingleLogs, "logs:refreshsingle")
}
