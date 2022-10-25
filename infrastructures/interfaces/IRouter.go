package interfaces

import "github.com/go-chi/chi"

// type IRouter interface {
// 	InitRouter()
// }

type IChiRouter interface {
	InitRouter(func(*chi.Mux)) *chi.Mux
}
