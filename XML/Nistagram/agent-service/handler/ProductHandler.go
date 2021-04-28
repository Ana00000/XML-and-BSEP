package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type ProductHandler struct {
	Service * service.ProductService
}

func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product := model.Product{
		ID:          uuid.UUID{},
		PicturePath: productDTO.PicturePath,
		Amount:      productDTO.Amount,
		Price:       productDTO.Price,
	}

	err = handler.Service.CreateProduct(&product)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

