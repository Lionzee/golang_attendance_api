package dto

type ActivityCreate struct {
	Description string `form:"description" binding:"required"`
}
