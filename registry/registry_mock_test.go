// -build windows

package registry

var registry Registry

func init() {
	registry = NewMockRegistry()
}
