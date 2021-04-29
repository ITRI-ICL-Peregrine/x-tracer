package ui

import (
	"fmt"
	"github.com/ITRI-ICL-Peregrine/x-tracer/database"
	"github.com/jroimartin/gocui"
)

func RefreshLogs(pn string, ps int64) {

	if ps == 4 {
		g.Update(func(g *gocui.Gui) error {

			if pn == "tcplife" {
				viewtl, err := g.View("tcplife")
				if err != nil {
					return err
				}
				viewtl.Clear()

				_, _ = fmt.Fprint(viewtl, database.GetActiveLogs(pn))

				g.SetViewOnTop("tcplife")

				viewtl.Autoscroll = true

				return nil
			} else if pn == "cachestat" {
				viewcs, err := g.View("cachestat")
				if err != nil {
					return err
				}
				viewcs.Clear()

				_, _ = fmt.Fprint(viewcs, database.GetActiveLogs(pn))

				g.SetViewOnTop("cachestat")
				viewcs.Autoscroll = true

				return nil
			} else if pn == "execsnoop" || pn == "biosnoop" {
				viewes, err := g.View("execsnoop")
				if err != nil {
					return err
				}
				viewes.Clear()

				_, _ = fmt.Fprint(viewes, database.GetActiveLogs(pn))

				g.SetViewOnTop("execsnoop")

				viewes.Autoscroll = true

				return nil
			} else {
				view, err := g.View("tcplogs")
				if err != nil {
					return err
				}
				view.Clear()

				_, _ = fmt.Fprint(view, database.GetActiveLogs(pn))
				g.SetViewOnTop("tcplogs")

				view.Autoscroll = true

				return nil
			}

		})
	}else if ps == 1{

		 g.Update(func(g *gocui.Gui) error {

                        view, err := g.View("logs")
                        if err != nil {
                                return err
                        }
                        view.Clear()

                        _, _ = fmt.Fprint(view, database.GetActiveLogs(pn))

                        g.SetViewOnTop("logs")
                        g.SetCurrentView("logs")

                        view.Autoscroll = true

                        return nil
                })


	} else {


		g.Update(func(g *gocui.Gui) error {


			if pn == "tcptracer" {
				viewtt, err := g.View("halfscreen")
				if err != nil {
					return err
				}
				viewtt.Clear()

				_, _ = fmt.Fprint(viewtt, database.GetActiveLogs(pn))

				g.SetViewOnTop("halfscreen")
				g.SetCurrentView("halfscreen")

				viewtt.Autoscroll = true

				return nil
			} else if pn == "tcpconnect" {
				viewtc, err := g.View("tcplife")
				if err != nil {
					return err
				}
				viewtc.Clear()

				_, _ = fmt.Fprint(viewtc, database.GetActiveLogs(pn))

				g.SetViewOnTop("tcplife")
				g.SetCurrentView("tcplife")

				viewtc.Autoscroll = true

				return nil
			} else {
				viewtl, err := g.View("tcplogs")
				if err != nil {
					return err
				}
				viewtl.Clear()

				_, _ = fmt.Fprint(viewtl, database.GetActiveLogs(pn))

				g.SetViewOnTop("tcplogs")
				g.SetCurrentView("tcplogs")

				viewtl.Autoscroll = true

				return nil
			}
		})




	}
}
/*
func RefreshSingleLogs(pn string) {

		g.Update(func(g *gocui.Gui) error {

			view, err := g.View("logs")
			if err != nil {
				return err
			}
			view.Clear()

			_, _ = fmt.Fprint(view, database.GetActiveLogs(pn))

			g.SetViewOnTop("logs")
			g.SetCurrentView("logs")

			view.Autoscroll = true

			return nil
		})


}

func RefreshTcpLogs(pn string) {

		g.Update(func(g *gocui.Gui) error {


			if pn == "tcptracer" {
				viewtt, err := g.View("halfscreen")
				if err != nil {
					return err
				}
				viewtt.Clear()

				_, _ = fmt.Fprint(viewtt, database.GetActiveLogs(pn))

				g.SetViewOnTop("halfscreen")
				g.SetCurrentView("halfscreen")

				viewtt.Autoscroll = true

				return nil
			} else if pn == "tcpconnect" {
				viewtc, err := g.View("tcplife")
				if err != nil {
					return err
				}
				viewtc.Clear()

				_, _ = fmt.Fprint(viewtc, database.GetActiveLogs(pn))

				g.SetViewOnTop("tcplife")
				g.SetCurrentView("tcplife")

				viewtc.Autoscroll = true

				return nil
			} else {
				viewtl, err := g.View("tcplogs")
				if err != nil {
					return err
				}
				viewtl.Clear()

				_, _ = fmt.Fprint(viewtl, database.GetActiveLogs(pn))

				g.SetViewOnTop("tcplogs")
				g.SetCurrentView("tcplogs")

				viewtl.Autoscroll = true

				return nil
			}
		})


}
*/
