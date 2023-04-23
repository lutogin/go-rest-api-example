package middlewares

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	appErrors "ms-users/internal/app-errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error

func ErrorHandlingMiddleware(next appHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Call the next handler in the chain and capture any error
		err := next(w, r, p)

		// If an error is returned, handle it accordingly
		var appErr *appErrors.AppError
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
			w.Write([]byte(err.Error()))
		}
	}
}
