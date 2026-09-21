package route

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"restAPI/model"
	"restAPI/service"
)

// GetAllFromShop godoc
//
//	@Summary		Get all products
//	@Description	Get all products from Fake Store API
//	@Tags			Shop
//	@Produce		json
//	@Success		200	{object}	model.APIResponse
//	@Failure		500	{object}	model.APIResponse
//	@Failure		502	{object}	model.APIResponse
//	@Router			/shop [get]
func GetAllFromShop(w http.ResponseWriter, r *http.Request) {

	products, status, err := service.GetProducts()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		response := model.APIResponse{
			Status:   status,
			DateTime: time.Now(),
			Payload:  map[string]string{"error": err.Error()},
		}

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.APIResponse{
		Status:   status,
		DateTime: time.Now(),
		Payload:  products,
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// GetSingleProduct godoc
//
// @Summary      Get single product
// @Description  Get a single product from Fake Store API
// @Tags         Shop
// @Produce      json
// @Param        id  path  int  true  "Product ID"
// @Success      200  {object}  model.APIResponse
// @Failure      400  {object}  model.APIResponse
// @Failure      404  {object}  model.APIResponse
// @Failure      500  {object}  model.APIResponse
// @Router       /shop/{id} [get]
func GetSingleProduct(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		response := model.APIResponse{
			Status:   http.StatusBadRequest,
			DateTime: time.Now(),
			Payload:  map[string]string{"error": "Invalid product ID"},
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	product, status, err := service.GetSingleProduct(id)

	if err != nil {
		response := model.APIResponse{
			Status:   status,
			DateTime: time.Now(),
			Payload:  map[string]string{"error": err.Error()},
		}

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.APIResponse{
		Status:   status,
		DateTime: time.Now(),
		Payload:  product,
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}