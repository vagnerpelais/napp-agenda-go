package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		stores := main.Group("napp-agenda")
		{
			stores.GET("/stores", controllers.GetStores)
			stores.GET("/stores/:id", controllers.GetStoreByID)
			stores.POST("/stores", controllers.CreateStore)
			stores.PATCH("/stores/:id", controllers.UpdateStore)

			stores.GET("/integrationteams", controllers.GetIntegrationTeams)
			stores.GET("/integrationteams/:id", controllers.GetIntegrationTeamByID)
			stores.POST("/integrationteams", controllers.CreateIntegrationTeam)
			stores.PATCH("/integrationteams/:id", controllers.UpdateIntegrationTeam)

			stores.GET("/users", controllers.GetUsers)
			stores.GET("/users/:id", controllers.GetUserByID)
			stores.POST("/users", controllers.CreateUser)
			stores.PATCH("/users/:id", controllers.UpdateUser)

			stores.GET("/times", controllers.GetTimes)
			stores.GET("/times/:id", controllers.GetTimeByID)
			stores.POST("/times", controllers.CreateTime)
			stores.PATCH("/times/:id", controllers.UpdateTime)

			stores.GET("/tasktypes", controllers.GetTaskTypes)
			stores.GET("/tasktypes/:id", controllers.GetTaskTypeByID)
			stores.POST("/tasktypes", controllers.CreateTaskType)
			stores.PATCH("/tasktypes/:id", controllers.UpdateTaskType)

			stores.GET("/tasks", controllers.GetTasks)
			stores.GET("/tasks/:id", controllers.GetTaskByID)
			stores.POST("/tasks", controllers.CreateTask)
			stores.PATCH("/tasks/:id", controllers.UpdateTask)
		}
	}

	return router
}
