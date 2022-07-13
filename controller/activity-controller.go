package controller

import (
	"DailyActivity/dto"
	"DailyActivity/helper"
	"DailyActivity/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

type ActivityController interface {
	All(ctx *gin.Context)
	CreateActivity(ctx *gin.Context)
}

type activityController struct {
	activityService   service.ActivityService
	attendanceService service.AttendanceService
	jwtService        service.JWTService
}

func NewActivityController(activityService service.ActivityService, attendanceService service.AttendanceService, jwtService service.JWTService) ActivityController {
	return &activityController{
		activityService: activityService,
		jwtService:      jwtService,
	}
}

func (c *activityController) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	acts, err := c.activityService.All(userID)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", acts)
	ctx.JSON(http.StatusOK, response)
}

func (c *activityController) CreateActivity(ctx *gin.Context) {
	var createActivityReq dto.ActivityCreate
	err := ctx.ShouldBind(&createActivityReq)

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	if !c.activityService.IsCheckedIn(userID) {
		response := helper.BuildErrorResponse("Failed to process request", "You are not yet Checked-in", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	}

	if c.activityService.IsCheckedOut(userID) {
		response := helper.BuildErrorResponse("Failed to process request", "You are already Checked-out today", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	}

	res, err := c.activityService.CreateActivity(createActivityReq, userID)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}
