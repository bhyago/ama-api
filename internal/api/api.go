package api

import (
	"net/http"

	"github.com/bhyago/ama-api/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
	// upgrader    websocket.Upgrader
	// subscribers map[string]map[*websocket.Conn]context.CancelFunc
	// mu          *sync.Mutex
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
		// upgrader:    websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		// subscribers: make(map[string]map[*websocket.Conn]context.CancelFunc),
		// mu:          &sync.Mutex{},
	}

	r := chi.NewRouter()

	a.r = r
	return a
}
