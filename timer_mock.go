package main

type MockTimer struct {
	name        string
	startCalled bool
	stopCalled  bool
	readCalled  bool
	clearCalled bool
	listCalled  bool
	execCalled  bool
}

func (t *MockTimer) start(name string) {
	t.startCalled = true
}

func (t *MockTimer) stop(name string) {
	t.stopCalled = true
}

func (t *MockTimer) read(name string) {
	t.readCalled = true
}

func (t *MockTimer) clear(name string) {
	t.clearCalled = true
}

func (t *MockTimer) list(name string) {
	t.listCalled = true
}

func (t *MockTimer) exec(process string) {
	t.execCalled = true
}
