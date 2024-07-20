package dto

type CreatePropertyRequest struct {
	Name        string `json:"name" binding:"required,alpha,min=3,max=15"`
	Icon        string `json:"icon" binding:"required,min=1,max=1000"`
	CategoryId  int    `json:"category_id" binding:"required"`
	Description string `json:"description" binding:"required,min=10,max=1000"`
	DataType    string `json:"dataType" binding:"required,max=15"`
	Unit        string `json:"unit" binding:"required,max=15"`
}

type UpdatePropertyRequest struct {
	Name        string `json:"name,omitempty" binding:"alpha,min=3,max=15"`
	Icon        string `json:"icon,omitempty" binding:"max=1000"`
	CategoryId  int    `json:"category_id,omitempty" `
	Description string `json:"description,omitempty" binding:"min=10,max=1000"`
	DataType    string `json:"dataType,omitempty" binding:"max=15"`
	Unit        string `json:"unit,omitempty" binding:"max=15"`
}
type GetPropertyResponse struct {
	Id          int                         `json:"id"`
	Name        string                      `json:"name"`
	Icon        string                      `json:"icon" `
	CategoryId  int                         `json:"category_id"`
	Category    GetPropertyCategoryResponse `json:"category,omitempty"`
	Description string                      `json:"description"`
	DataType    string                      `json:"dataType" `
	Unit        string                      `json:"unit" `
}
type CreatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
	Icon string `json:"icon" binding:"min=1,max=1000"`
}
type UpdatePropertyCategoryRequest struct {
	Name string `json:"name,omitempty" binding:"min=3,alpha,max=15"`
	Icon string `json:"icon,omitempty" binding:"min=1,max=1000"`
}
type GetPropertyCategoryResponse struct {
	Id         int                   `json:"id"`
	Name       string                `json:"name"`
	Icon       string                `json:"icon" `
	Properties []GetPropertyResponse `json:"properties,omitempty"`
}
