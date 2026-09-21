package route

import (
	"encoding/json"
	"net/http"
	"time"

	"restAPI/lib"
	"restAPI/model"
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

	resp, err := http.Get(lib.BaseURL + "/products")
	if err != nil {
		response := model.APIResponse{
			Status:   http.StatusInternalServerError,
			DateTime: time.Now(),
			Payload:  map[string]string{"error": "Failed to fetch products"},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		response := model.APIResponse{
			Status:   http.StatusBadGateway,
			DateTime: time.Now(),
			Payload:  map[string]string{"error": "External API returned an error"},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(response)
		return
	}

	var products []model.Product

	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		response := model.APIResponse{
			Status:   http.StatusInternalServerError,
			DateTime: time.Now(),
			Payload:  map[string]string{"error": "Failed to decode response"},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.APIResponse{
		Status:   http.StatusOK,
		DateTime: time.Now(),
		Payload:  products,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}