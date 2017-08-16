package memo

import "sync"

type memo struct {
	f func_t
	cache map[string]entry
}

type entry struct {
	res result
	ready chan struct{}
	mu sync.Mutex
}

type func_t func(string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New (f func_t) *memo {
	return &memo{ f:f, cache:make(map[string]entry) }
}

func (m *memo) Get(key string) (interface{}, error) {
	mu.Lock()
	e, ok := m.cache[key]
	if !ok {
		e.ready = make(<- chan struct{})
		m.cache[key] = e
		mu.Unlock()
		e.res.value, e.res.err = m.f(key)
		close(e.ready)
	}
	else {
		mu.Unlock()
		<- e.ready
	}
	return e.res.value, e.res.err
}
