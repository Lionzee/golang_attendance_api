package controller

import (
	"DailyActivity/helper"
	"DailyActivity/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

type AttendanceController interface {
	CheckIn(ctx *gin.Context)
}

type attendanceController struct {
	attendanceService service.AttendanceService
	jwtService        service.JWTService
}

func NewAttendanceController(attendanceService service.AttendanceService, jwtService service.JWTService) AttendanceController {
	return &attendanceController{
		attendanceService: attendanceService,
		jwtService:        jwtService,
	}
}

func (c *attendanceController) CheckIn(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	if !c.attendanceService.IsDuplicate(userID) {
		response := helper.BuildErrorResponse("Failed to process request", "Already Checked-in", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		res, err := c.attendanceService.CheckIn(userID)
		if err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
			return
		}

		response := helper.BuildResponse(true, "Check-in success!", res)
		ctx.JSON(http.StatusCreated, response)
	}
}
