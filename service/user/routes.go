package user

import (
	"net/http"

	"github.com/nadeem-baig/MHPS-backend/config"
	. "github.com/nadeem-baig/MHPS-backend/middleware"
)

// NewHandler initializes a new router with defined routes for user service.
func NewHandler(h *config.Handler) http.Handler {
	mux := http.NewServeMux()
	store := NewStore(h.DB)

	mux.Handle("/", LoggingMiddleware(MethodHandler(http.MethodGet, HomeHandler(h))))
	mux.Handle("/login", LoggingMiddleware(MethodHandler(http.MethodPost, LoginHandler(h, store))))
	mux.Handle("/register", LoggingMiddleware(MethodHandler(http.MethodPost, RegisterHandler(h, store))))
	mux.Handle("/register/member", LoggingMiddleware(MethodHandler(http.MethodPost, RegisterMemberHandler(h, store))))
	return mux
}
