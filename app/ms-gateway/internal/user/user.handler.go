package user

import (
	"ms-gateway/internal/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
}

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

func New() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetAll)
	router.GET(userURL, h.GetById)
	router.POST(usersURL, h.Create)
	router.PUT(userURL, h.UpdateOrCreate)
	router.PATCH(userURL, h.Update)
	router.DELETE(userURL, h.Delete)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetUsers"))
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetUserById"))
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("CreateUser"))
}

func (h *handler) UpdateOrCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("UpdateOrCreate"))
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update"))
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete"))
}
