package infrastructures

import (
	"sync"

	"github.com/go-chi/chi"

	"listing-service/infrastructures/interfaces"
)

type Router struct{}

func (router *Router) InitRouter(routerFunc func(*chi.Mux)) *chi.Mux {
	r := chi.NewRouter()
	routerFunc(r)
	return r
}

var (
	m          *Router
	routerOnce sync.Once
)

func ChiRouter() interfaces.IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &Router{}
		})
	}
	return m
}
