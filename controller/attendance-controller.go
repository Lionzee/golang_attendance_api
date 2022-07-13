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
	CheckOut(ctx *gin.Context)
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
		if !c.attendanceService.IsCheckedOut(userID) {
			response := helper.BuildErrorResponse("Failed to process request", "You already Checked-out", helper.EmptyObj{})
			ctx.JSON(http.StatusConflict, response)
		}

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

func (c *attendanceController) CheckOut(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	if c.attendanceService.IsCheckedOut(userID) {
		response := helper.BuildErrorResponse("Failed to process request", "You are already Checked-out today", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	}

	if !c.attendanceService.IsCheckedIn(userID) {
		response := helper.BuildErrorResponse("Failed to process request", "You are not yet Checked-in", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		res, err := c.attendanceService.CheckOut(userID)
		if err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
			return
		}

		response := helper.BuildResponse(true, "Check-out success!", res)
		ctx.JSON(http.StatusCreated, response)
	}

}
