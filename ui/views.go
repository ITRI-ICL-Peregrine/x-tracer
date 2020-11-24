package ui

import (
	"fmt"
	"strings"
	"time"
	"github.com/jroimartin/gocui"
	"github.com/willf/pad"
)

// View: Overlay
func viewOverlay(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("overlay", 0, 0, lMaxX, lMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = false
	}

	return nil
}

// View: Title bar
func viewTitle(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("title", -1, -1, lMaxX, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = true
		v.BgColor = gocui.ColorDefault | gocui.AttrReverse
		v.FgColor = gocui.ColorDefault | gocui.AttrReverse

		// Content
		fmt.Fprintln(v, versionTitle(lMaxX))

	}

	return nil
}

//View: Information

func viewInfo(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("info", 9, lMaxY-3, lMaxX, lMaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		//settings

		v.Frame = true
		v.BgColor = gocui.ColorDefault | gocui.AttrReverse
		v.FgColor = gocui.ColorDefault | gocui.AttrReverse
		fmt.Fprintln(v, strings.Repeat("─", lMaxX))

		//content
		fmt.Fprintln(v, textPadCenter("hello", lMaxX))
		g.SetCurrentView(v.Name())
	}

	return nil
}

func viewLogs(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("logs", 2, 2, lMaxX-4, lMaxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = " Logs "
		v.Autoscroll = false

		v.SetCursor(1, 3)
	}
	return nil
}

func viewTcpLogs(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("tcplogs", 1, 1, lMaxX/2, lMaxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = " Tcp Logs "
		v.Autoscroll = true
		v.Wrap = true

		v.SetCursor(1, 3)

	}

	return nil
}

func viewTcpLifeLogs(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("tcplife", lMaxX/2, 1, lMaxX, lMaxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = " TcpLife "
		v.Autoscroll = true
		v.Wrap = true

		v.SetCursor(1, 3)

	}

	return nil
}

func viewExecSnoopLogs(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("execsnoop", 1, lMaxY/2, lMaxX/2, lMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = " ExecSnoop "
		v.Autoscroll = true
		v.Wrap = true

		v.SetCursor(1, 3)

	}

	return nil
}

func viewCacheStatLogs(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("cachestat", lMaxX/2, lMaxY/2, lMaxX, lMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = " CacheStats "
		v.Autoscroll = true
		v.Wrap = true

		v.SetCursor(1, 3)

	}

	return nil
}

// View: Namespace
func viewNamespaces(g *gocui.Gui, lMaxX int, lMaxY int) error {
	w := lMaxX / 2
	h := lMaxY / 4
	minX := (lMaxX / 2) - (w / 2)
	minY := (lMaxY / 2) - (h / 2)
	maxX := minX + w
	maxY := minY + h

	// Main view
	if v, err := g.SetView("namespaces", minX, minY, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Configure view
		v.Title = " Namespaces "
		v.Frame = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		// Display list
		go viewNamespacesRefreshList(g)

	}
	return nil
}

// Actualize list in namespaces view
func viewNamespacesRefreshList(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("namespaces")
		if err != nil {
			return err
		}

		namespaces, err := getNamespaces()
		if err != nil {
			displayError(g, err)
			return nil
		}
		hideError(g)

		var ns []string

		v.Clear()

		if len(namespaces.Items) > 0 {
			for _, namespace := range namespaces.Items {
				fmt.Fprintln(v, namespace.GetName())
				ns = append(ns, namespace.GetName())
			}
		} else {

		}

		setViewCursorToLine(g, v, ns, NAMESPACE)

		return nil
	})
}

// View: Pods
func viewPods(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("pods", -1, 5, lMaxX, lMaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = true
		v.Title = " Pods "
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.SetCursor(0, 2)

		// Set as current view
		g.SetCurrentView(v.Name())

		// Content
		go viewPodsShowWithAutoRefresh(g)
	}

	return nil
}

// Auto refresh view pods
func viewPodsShowWithAutoRefresh(g *gocui.Gui) {
	c := getConfig()
	t := time.NewTicker(time.Duration(c.frequency) * time.Second)
	go viewPodsRefreshList(g)
	for {
		select {
		case <-t.C:
			go viewPodsRefreshList(g)
		}
	}
}

// Actualize list in pods view
func viewPodsRefreshList(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		lMaxX, _ := g.Size()
		v, err := g.View("pods")
		if err != nil {
			return err
		}

		pods, err := getPods()
		if err != nil {
			displayError(g, err)
			return nil
		}
		hideError(g)

		v.Clear()

		// Content: Add column
		viewPodsAddLine(v, lMaxX, "NAME", "READY", "STATUS", "RESTARTS", "AGE")
		fmt.Fprintln(v, strings.Repeat("─", lMaxX))

		if len(pods.Items) > 0 {
			for _, pod := range pods.Items {
				n := pod.GetName()
				//c := "?" // TODO CPU + Memory #20
				//m := "?" // TODO CPU + Memory #20
				s := columnHelperStatus(pod.Status)
				r := columnHelperRestarts(pod.Status.ContainerStatuses)
				a := columnHelperAge(pod.CreationTimestamp)
				cr := columnHelperReady(pod.Status.ContainerStatuses)
				viewPodsAddLine(v, lMaxX, n, cr, s, r, a)
			}

			// Reset cursor when empty line
			if l, err := getViewLine(g, v); err != nil || l == "" {
				v.SetCursor(0, 2)
			}
		} else {
			v.SetCursor(0, 2)
		}

		return nil
	})
}

// View: Status bar
func viewStatusBar(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("status", -1, lMaxY-2, lMaxX, lMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Frame = false
		v.BgColor = gocui.ColorBlack
		v.FgColor = gocui.ColorWhite

		// Content
		changeStatusContext(g, "D")
	}

	return nil
}

// Change status context
func changeStatusContext(g *gocui.Gui, c string) error {
	lMaxX, _ := g.Size()
	v, err := g.View("status")
	if err != nil {
		return err
	}

	v.Clear()

	i := lMaxX + 4
	b := ""

	switch c {
	case "D":
		i = 150 + i
		b = b + frameText("↑") + " Up   "
		b = b + frameText("↓") + " Down   "
		b = b + frameText("D") + " Delete   "
		b = b + frameText("L") + " Show Logs   "
		b = b + frameText("CTRL+N") + " Namespace   "
	case "SE":
		i = i + 100
		b = b + frameText("↑") + " Up   "
		b = b + frameText("↓") + " Down   "
		b = b + frameText("Enter") + " Select   "
	case "SL":
		i = i + 100
		b = b + frameText("↑") + " Up   "
		b = b + frameText("↓") + " Down   "
		b = b + frameText("L") + " Hide Logs   "
	}
	b = b + frameText("CTRL+C") + " Exit"

	fmt.Fprintln(v, pad.Left(b, i, " "))

	return nil
}

func viewPodsAddLine(v *gocui.View, maxX int, name, ready, status, restarts, age string) {
	wN := maxX - 34 // 54 // TODO CPU + Memory #20
	if wN < 45 {
		wN = 45
	}
	line := pad.Right(name, wN, " ") +
		//pad.Right(cpu, 10, " ") + // TODO CPU + Memory #20
		//pad.Right(memory, 10, " ") + // TODO CPU + Memory #20
		pad.Right(ready, 10, " ") +
		pad.Right(status, 10, " ") +
		pad.Right(restarts, 10, " ") +
		pad.Right(age, 4, " ")
	fmt.Fprintln(v, line)
}

func viewProbes(g *gocui.Gui, lMaxX int, lMaxY int) error {
	w := lMaxX / 2
	h := lMaxY / 4
	minX := (lMaxX / 2) - (w / 2)
	minY := (lMaxY / 2) - (h / 2)
	maxX := minX + w
	maxY := minY + h
	// Main view
	if v, err := g.SetView("probes", minX, minY, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Configure view
		v.Title = " Select Probes "
		v.Frame = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		viewProbeNames(g)

	}
	return nil
}

func viewProbeNames(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {

		v, err := g.View("probes")
		if err != nil {
			return err
		}
		probes := getProbeNames()
		v.Clear()

		if len(probes) >= 0 {
			for i := range probes {
				fmt.Fprintln(v, probes[i])
			}
		} else {

		}

		setViewCursorToLine(g, v, probes, "tcptracer")

		return nil

	})

}

func getProbeNames() []string {

	pn := []string{"tcptracer", "tcpconnect", "tcpaccept", "tcplife", "execsnoop", "biosnoop", "cachestat", "All TCP Probes", "All Probes"}
	return pn

}

