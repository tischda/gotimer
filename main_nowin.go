// -build windows

package main

import "github.com/tischda/timer/registry"

func init() {
	t = timer{registry: registry.NewMockRegistry()}
}
