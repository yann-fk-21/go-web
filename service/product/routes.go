package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yann-fk-21/todo-platform/types"
	"github.com/yann-fk-21/todo-platform/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes (router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts)

}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {

	ps, err :=h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	    return
	}

	utils.WriteJson(w, http.StatusAccepted, ps)

}