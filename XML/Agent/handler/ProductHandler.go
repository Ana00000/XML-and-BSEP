package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Agent/dto"
	"github.com/xml/XML-and-BSEP/XML/Agent/model"
	"github.com/xml/XML-and-BSEP/XML/Agent/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
)

type ProductHandler struct {
	ProductService * service.ProductService
	Validator * validator.Validate
}

func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO dto.ProductDTO
	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&productDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productId := uuid.New()
	product := model.Product{
		ID:          productId,
		PicturePath: productDTO.PicturePath,
		Amount:      productDTO.Amount,
		Price:       productDTO.Price,
		AgentUserID: productDTO.AgentUserID,
	}


	if err := handler.ProductService.CreateProduct(&product); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

