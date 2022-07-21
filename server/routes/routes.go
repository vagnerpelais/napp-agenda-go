package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		api := main.Group("napp-agenda")
		{
			api.GET("/stores", controllers.GetStores)
			api.GET("/stores/:id", controllers.GetStoreByID)
			api.POST("/stores", controllers.CreateStore)
			api.PATCH("/stores/:id", controllers.UpdateStore)
			api.DELETE("/stores/:id", controllers.DeleteStore)

			api.GET("/integrationteams", controllers.GetIntegrationTeams)
			api.GET("/integrationteams/:id", controllers.GetIntegrationTeamByID)
			api.POST("/integrationteams", controllers.CreateIntegrationTeam)
			api.PATCH("/integrationteams/:id", controllers.UpdateIntegrationTeam)
			api.DELETE("/integrationteams/:id", controllers.DeleteIntegrationTeam)

			api.GET("/users", controllers.GetUsers)
			api.GET("/users/:id", controllers.GetUserByID)
			api.POST("/users", controllers.CreateUser)
			api.PATCH("/users/:id", controllers.UpdateUser)
			api.DELETE("/users/:id", controllers.DeleteUser)

			api.GET("/times", controllers.GetTimes)
			api.GET("/times/:id", controllers.GetTimeByID)
			api.POST("/times", controllers.CreateTime)
			api.PATCH("/times/:id", controllers.UpdateTime)
			api.DELETE("/times/:id", controllers.DeleteTime)

			api.GET("/tasktypes", controllers.GetTaskTypes)
			api.GET("/tasktypes/:id", controllers.GetTaskTypeByID)
			api.POST("/tasktypes", controllers.CreateTaskType)
			api.PATCH("/tasktypes/:id", controllers.UpdateTaskType)
			api.DELETE("/tasktypes/:id", controllers.DeleteTaskType)

			api.GET("/tasks", controllers.GetTasks)
			api.GET("/tasks/:id", controllers.GetTaskByID)
			api.POST("/tasks", controllers.CreateTask)
			api.PATCH("/tasks/:id", controllers.UpdateTask)
			api.DELETE("/tasks/:id", controllers.DeleteTaskType)

			api.GET("/events", controllers.GetEvents)
			api.GET("/events/:id", controllers.GetEventByID)
			api.POST("/events", controllers.CreateEvent)
			api.PATCH("/events/:id", controllers.UpdateEvent)
			api.DELETE("/events/:id", controllers.DeleteEvent)
		}
	}

	return router
}
