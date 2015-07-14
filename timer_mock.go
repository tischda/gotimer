package main

type mockTimer struct {
	name        string
	startCalled bool
	stopCalled  bool
	readCalled  bool
	clearCalled bool
	listCalled  bool
	execCalled  bool
}

func (t *mockTimer) start(name string) {
	t.startCalled = true
}

func (t *mockTimer) stop(name string) {
	t.stopCalled = true
}

func (t *mockTimer) read(name string) {
	t.readCalled = true
}

func (t *mockTimer) clear(name string) {
	t.clearCalled = true
}

func (t *mockTimer) list(name string) {
	t.listCalled = true
}

func (t *mockTimer) exec(process string) {
	t.execCalled = true
}
