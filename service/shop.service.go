package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"restAPI/lib"
	"restAPI/model"
)

func GetProducts() ([]model.Product, int, error) {

	resp, err := http.Get(lib.BaseURL + "/products")
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("failed to get products")
	}

	var products []model.Product

	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return products, http.StatusOK, nil
}

func GetSingleProduct(id int) (model.Product, int, error) {

	resp, err := http.Get(
		lib.BaseURL + "/products/" + strconv.Itoa(id),
	)

	if err != nil {
		return model.Product{}, http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Product{}, resp.StatusCode, fmt.Errorf("product not found")
	}

	var product model.Product

	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return model.Product{}, http.StatusInternalServerError, err
	}

	return product, http.StatusOK, nil
}