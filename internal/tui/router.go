package tui

type Router struct {
	active string
}

func (r *Router) Switch(name string) {
	r.active = name
}
