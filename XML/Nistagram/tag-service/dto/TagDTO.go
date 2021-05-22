package dto

type TagDTO struct {
	Name string `json:"name" validate:"required"`
	TagType string `json:"tag_type" validate:"oneof=USER_TAG HASH_TAG"`
}
