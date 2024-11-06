package nest

import (
	"net/http"
)

type router struct {
	mux *http.ServeMux
}

func (r *router) defaultHandler(h *handler) {
	r.mux.HandleFunc(h.path,
		func(w http.ResponseWriter, r *http.Request) {
			c := newCtx(w, r)
			defer c.reset()

			i := 0
			canActivate := true
			c.Next = func() {
				if i < len(h.midWares) {
					mw := h.midWares[i]
					i++
					mw(c)
					canActivate = false
					if i <= len(h.midWares) {
						c.Next()
					}
				} else if h.endPoint != nil && canActivate {
					c.Next = func() {}
					err := (*h.endPoint)(c)
					if err != nil {
						http.Error(c.Res(), err.Error(), http.StatusInternalServerError)
					}
				}
			}
			c.Next()
		})
}

func newRouter() *router {
	return &router{
		mux: http.NewServeMux(),
	}
}
