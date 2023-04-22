package middlewares

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	appErrors "ms-gateway/internal/app-errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var appErr *appErrors.AppError
		err := h(w, r, params)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if errors.As(err, &appErr) {
				if errors.Is(err, appErrors.ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(appErrors.ErrNotFound.Marshal())
					return
				}

				err = err.(*appErrors.AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
				return
			}

			w.WriteHeader(http.StatusTeapot)
			w.Write(appErrors.SystemError(err).Marshal())
		}
	}
}
