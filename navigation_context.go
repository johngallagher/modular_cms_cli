package main

import "strings"

// Model definition
type NavigationContext struct {
	Path []string
}

func (nc *NavigationContext) Breadcrumb() string {
	return strings.Join(nc.Path, " > ")
}

func (nc *NavigationContext) Push(level string) {
	nc.Path = append(nc.Path, level)
}

func (nc *NavigationContext) Pop() {
	if len(nc.Path) > 1 {
		nc.Path = nc.Path[:len(nc.Path)-1]
	}
}
