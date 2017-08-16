package memo

type memo struct {
	f func_t
	cache map[string]result
}

type func_t func(string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New (f func_t) *memo {
	return &memo{ f:f, cache:make(map[string]result) }
}

func (m *memo) Get(key string) (interface{}, error) {
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	return res.value, res.err
}
