package main

import (
	"TodoServer/pkg"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	pkg.DatabaseInit()
	// pkg.InitRedis()

	fmt.Println("doto服务器启动中...")
	r := gin.Default()

	// CORS 配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// // 获取验证码(这里不用检查手机号是否被注册，但是应该限制次数和时间)
	// r.GET("/grtCaptcha", pkg.GetVerificationCode)
	// // 校验验证码接口
	// r.POST("/verify_code", pkg.VerifyCode)

	// // 用户注册
	// r.POST("/register", func(c *gin.Context) {
	// 	var user pkg.User
	// 	if err := c.ShouldBindJSON(&user); err != nil {
	// 		c.JSON(400, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	fmt.Println("手机号:", user.PhoneNumber)
	//   // 检查手机号是否已存在
	// 	var existingUser pkg.User
	//   if err := pkg.DB.Where("phone = ?", user.PhoneNumber).First(&existingUser).Error; err == nil {
	//       c.JSON(400, gin.H{"error": "手机号已被注册"})
	//       return
	//   }

	//   // 插入新用户
	//   if err := pkg.DB.Create(&user).Error; err != nil {
	//       c.JSON(500, gin.H{"error": "用户注册失败"})
	//       return
	//   }

	// 	token, err := pkg.GenToken(user.ID, user.Name)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{"error": "生成 Token 失败"})
	// 		return
	// 	}

	// 	c.JSON(200, gin.H{"message": "注册成功", "user": user, "token": token})
	// })

	// 用户注册
	r.POST("/API/register", func(c *gin.Context) {
		var user pkg.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if user.Account == "" {
			c.JSON(400, gin.H{"error": "账号不得为空"})
		}

		if user.Password == "" {
			c.JSON(400, gin.H{"error": "密码不得为空"})
		}

		var existingUser pkg.User
		if err := pkg.DB.Where("account = ?", user.Account).First(&existingUser).Error; err == nil {
			c.JSON(400, gin.H{"error": "账号已存在"})
			return
		}

		if err := pkg.DB.Create(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": "用户注册失败"})
			return
		}

		c.JSON(200, gin.H{"message": "注册成功", "user": user})
	})

	// 用户登录
	r.POST("/API/login", func(c *gin.Context) {
		var user pkg.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if user.Account == "" {
			c.JSON(400, gin.H{"error": "账号不得为空"})
		}

		if user.Password == "" {
			c.JSON(400, gin.H{"error": "密码不得为空"})
		}

		var existingUser pkg.User
		if err := pkg.DB.Where("account = ?", user.Account).First(&existingUser).Error; err != nil {
			c.JSON(401, gin.H{"error": "账号错误"})
			return
		}

		if existingUser.Password != user.Password {
			c.JSON(401, gin.H{"error": "密码错误"})
			return
		}

		token, err := pkg.GenToken(user.ID, user.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "生成 Token 失败"})
			return
		}

		var User pkg.User
		if err := pkg.DB.Where("account = ?", user.Account).First(&User).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "登录成功", "token": token, "userName": User.Name, "userId": User.ID})
	})

	// 用户task数据存储
	r.POST("/API/addTask", func(c *gin.Context) {
		var task pkg.Todo
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(task.Text)

		task.IsFinish = false
		task.IsImportant = false

		if err := pkg.DB.Create(&task).Error; err != nil {
			c.JSON(500, gin.H{"error": "事件存储失败"})
			return
		}

		c.JSON(200, gin.H{"message": "事件存储成功"})
	})

	r.POST("/API/getTask", func(c *gin.Context) {
		var user pkg.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var User pkg.User
		if err := pkg.DB.Where("id = ?", user.ID).Preload("Todos").First(&User).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "事件查询成功", "todos": User.Todos})
	})

	r.POST("/API/changeUserName", func(c *gin.Context) {
		var user pkg.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var User pkg.User
		if err := pkg.DB.Where("id = ?", user.ID).First(&User).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := pkg.DB.Model(&User).Update("name", user.Name).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var newUser pkg.User
		if err := pkg.DB.Where("id = ?", user.ID).First(&newUser).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "账号修改成功", "name": newUser.Name})
	})

	r.POST("/API/changeTaskText", func(c *gin.Context) {
		var todo pkg.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var Todo pkg.Todo
		if err := pkg.DB.Where("id = ?", todo.ID).First(&Todo).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := pkg.DB.Model(&Todo).Update("text", todo.Text).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var newTodo pkg.Todo
		if err := pkg.DB.Where("id = ?", todo.ID).First(&newTodo).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "文本修改成功", "name": newTodo.Text})

	})

	r.POST("/API/changeTaskIsFinish", func(c *gin.Context) {
		var todo pkg.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var Todo pkg.Todo
		if err := pkg.DB.Where("id = ?", todo.ID).First(&Todo).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := pkg.DB.Model(&Todo).Update("IsFinish", todo.IsFinish).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var newTodo pkg.Todo
		if err := pkg.DB.Where("id = ?", todo.ID).First(&newTodo).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "文本修改成功", "name": newTodo.IsFinish})
	})

	r.POST("/API/changeTaskIsImportant", func(c *gin.Context) {
		var todo pkg.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var Todo pkg.Todo
		if err := pkg.DB.Where("id = ?", todo.ID).First(&Todo).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := pkg.DB.Model(&Todo).Update("IsImportant", todo.IsImportant).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var newTodo pkg.Todo
		if err := pkg.DB.Where("id = ?", todo.ID).First(&newTodo).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "文本修改成功", "name": newTodo.IsImportant})
	})

	r.POST("/API/deleteTask", func(c *gin.Context) {
		var todo pkg.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var Todo pkg.Todo
		if err := pkg.DB.Delete(&Todo, todo.ID).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "任务删除成功"})
	})

	r.POST("/API/deleteUser", func(c *gin.Context) {
		var user pkg.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if user.ID == "" {
			c.JSON(400, gin.H{"error": "ID 不能为空"})
			return
		}

		if err := pkg.DB.Where("id = ?", user.ID).Delete(&pkg.User{}).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "用户删除成功"})
	})

	r.Run(":8087")

}
