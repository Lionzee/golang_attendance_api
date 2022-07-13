package main

import (
	"DailyActivity/config"
	"DailyActivity/controller"
	"DailyActivity/middleware"
	"DailyActivity/repository"
	"DailyActivity/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                   *gorm.DB                        = config.SetupDatabaseConnection()
	userRepository       repository.UserRepository       = repository.NewUserRepository(db)
	attendanceRepository repository.AttendanceRepository = repository.NewAttendanceRepo(db)
	activityRepository   repository.ActivityRepository   = repository.NewActivityRepository(db)
	jwtService           service.JWTService              = service.NewJWTService()
	userService          service.UserService             = service.NewUserService(userRepository)
	authService          service.AuthService             = service.NewAuthService(userRepository)
	activityService      service.ActivityService         = service.NewActivityService(activityRepository)
	attendanceService    service.AttendanceService       = service.NewAttendanceService(attendanceRepository)
	activityController   controller.ActivityController   = controller.NewActivityController(activityService, attendanceService, jwtService)
	authController       controller.AuthController       = controller.NewAuthController(authService, jwtService)
	userController       controller.UserController       = controller.NewUserController(userService, jwtService)
	attendanceController controller.AttendanceController = controller.NewAttendanceController(attendanceService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
	}

	attendanceRoutes := r.Group("api/attendance", middleware.AuthorizeJWT(jwtService))
	{
		attendanceRoutes.POST("/checkin", attendanceController.CheckIn)
		attendanceRoutes.POST("/checkout", attendanceController.CheckOut)
	}

	activityRoutes := r.Group("api/activity", middleware.AuthorizeJWT(jwtService))
	{
		activityRoutes.GET("/all", activityController.All)
		activityRoutes.POST("/", activityController.CreateActivity)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
