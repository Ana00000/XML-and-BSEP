package dto

type ProductDTO struct {
	PicturePath string `json:"picturePath"`
	Amount int `json:"amount"`
	Price float32 `json:"price"`
}
