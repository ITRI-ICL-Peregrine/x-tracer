package ui

import (
	"fmt"
	"github.com/willf/pad"
)

var HELP = `
x-tracer monitors the Pods on your Kubernetes cluster.
Find more information at https://github.com/sheenam3/x-tracer-gocui
The following options can be passed to any command:
  --help        for more information about x-tracer!
  --version     version
  --frequency   refreshing frequency in seconds (default: 5)
  --kubeconfig  absolute path to the kubeconfig file
`

const APP = "x-tracer"
const AUTHOR = "@sheenam3"

// Get full banner of version
func versionFull() string {
	return fmt.Sprintf("%s %s - By %s", APP, version, AUTHOR)
}

// Get only banner (used in title bar view)
func versionBanner() string {
	return fmt.Sprintf(" %s %s", APP, version)
}

// Get only author (used in title bar view)
func versionAuthor() string {
	return fmt.Sprintf("By %s  ", AUTHOR)
}

// Prepare version title (used in title bar view)
func versionTitle(width int) string {
	return "⣿" + versionBanner() + pad.Left(versionAuthor(), width-len(versionBanner()), " ")
}
