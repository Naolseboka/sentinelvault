package httpx

import "net/http"

// Router is a minimal chi-compatible mux using Go 1.22 patterns.
type Router struct {
	mux *http.ServeMux
}

// NewRouter creates a new router.
func NewRouter() *Router {
	return &Router{mux: http.NewServeMux()}
}

// Handle registers a handler for method + pattern.
func (r *Router) Handle(method, pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(method+" "+pattern, handler)
}

// GET registers a GET handler.
func (r *Router) GET(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodGet, pattern, handler)
}

// POST registers a POST handler.
func (r *Router) POST(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPost, pattern, handler)
}

// PUT registers a PUT handler.
func (r *Router) PUT(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPut, pattern, handler)
}

// DELETE registers a DELETE handler.
func (r *Router) DELETE(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodDelete, pattern, handler)
}

// Mount mounts a sub-handler at prefix.
func (r *Router) Mount(prefix string, handler http.Handler) {
	r.mux.Handle(prefix+"/", http.StripPrefix(prefix, handler))
}

// Use wraps the router handler with middleware (outermost first).
func (r *Router) Use(middleware ...func(http.Handler) http.Handler) http.Handler {
	var h http.Handler = r.mux
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

// Handler returns the underlying handler with optional middleware.
func (r *Router) Handler(middleware ...func(http.Handler) http.Handler) http.Handler {
	return r.Use(middleware...)
}
