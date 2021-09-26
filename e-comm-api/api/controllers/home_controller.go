package controllers

import (
	"net/http"

	"github.com/Stuart6970/e-comm-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
