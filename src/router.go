package main

// Handler ...
type Handler func(*BotCommand)

// Router ...
type Router struct {
	rules map[string]Handler
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

// Handle ...
func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

// FindHandler ...
func (r *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found
}
