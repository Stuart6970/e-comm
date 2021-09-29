package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Stuart6970/e-comm-api/api/models"
	"github.com/Stuart6970/e-comm-api/api/responses"
	"github.com/Stuart6970/e-comm-api/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateCatalogItem(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	catalogItem := models.CatalogItem{}
	err = json.Unmarshal(body, &catalogItem)
	fmt.Println(catalogItem.AvailableStock)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = catalogItem.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	catalogItemCteated, err := catalogItem.SaveCatalogItem(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, catalogItemCteated.ID))
	responses.JSON(w, http.StatusCreated, catalogItemCteated)

}

func (server *Server) GetCatalogItems(w http.ResponseWriter, r *http.Request) {

	catalogItem := models.CatalogItem{}

	catalogItems, err := catalogItem.FindAllCatalogItems(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, catalogItems)
}

func (server *Server) GetCatalogItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	catalogItem := models.CatalogItem{}

	catalogItemReceived, err := catalogItem.FindCatalogItemById(server.DB, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, catalogItemReceived)
}

func (server *Server) UpdateCatalogItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Is a valid catalog item id provided to us?
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the catalog item exist
	catalogItem := models.CatalogItem{}
	err = server.DB.Debug().Model(models.CatalogItem{}).Where("id = ?", cid).Take(&catalogItem).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("catalog item not found"))
		return
	}

	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	catalogItemUpdate := models.CatalogItem{}
	err = json.Unmarshal(body, &catalogItemUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Perform validation
	err = catalogItemUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	catalogItemUpdate.ID = catalogItem.ID

	catalogItenUpdated, err := catalogItemUpdate.UpdateCatalogItem(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, catalogItenUpdated)
}

func (server *Server) DeleteCatalogItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Is a valid catalog item id provided to us?
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the catalog item exist
	catalogItem := models.CatalogItem{}
	err = server.DB.Debug().Model(models.CatalogItem{}).Where("id = ?", cid).Take(&catalogItem).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("catalog item not found"))
		return
	}

	_, err = catalogItem.DeleteCatalogItem(server.DB, cid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", cid))
	responses.JSON(w, http.StatusNoContent, "")
}
