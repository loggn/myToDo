package pkg

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)


var rdb *redis.Client
var ctx = context.Background()


// 初始化 Redis 客户端
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		// Password: "",               // 没有密码的话，填空
		// DB:       0,                // 默认数据库 0
	})

	// 检查 Redis 连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}

}

// // 验证用户输入的验证码
// func VerifyCaptcha(phone, inputCaptcha string) (bool, error) {
// 	storedCaptcha, err := rdb.Get(ctx, phone).Result()
// 	if err != nil {
// 		return false, err
// 	}
// 	return storedCaptcha == inputCaptcha, nil
// }

// 生成 6 位数字的验证码
func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999-100000) + 100000 // 生成6位数验证码
	return strconv.Itoa(code)
}

// 发送短信的函数，假设使用阿里云短信服务
// func sendCaptchaToPhone(phone string, captcha string) error {
// 	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<your-access-key-id>", "<your-access-key-secret>")
// 	if err != nil {
// 		return err
// 	}

// 	request := dysmsapi.CreateSendSmsRequest()
// 	request.Scheme = "https"
// 	request.PhoneNumbers = phone
// 	request.SignName = "<your-sign-name>"
// 	request.TemplateCode = "<your-template-code>"
// 	request.TemplateParam = fmt.Sprintf("{\"code\":\"%s\"}", captcha)

// 	_, err = client.SendSms(request)
// 	return err
// }

// 获取验证码接口
func GetVerificationCode(c *gin.Context) {
	// 获取用户手机号
	phone := c.DefaultQuery("phone", "")
	if phone == "" {
		c.JSON(400, gin.H{"error": "手机号不能为空"})
		return
	}

	// 生成验证码
	code := generateVerificationCode()

	// 将验证码存储到 Redis，设置有效期为 5 分钟
	err := rdb.Set(ctx, phone, code, 5*time.Minute).Err()
	if err != nil {
		log.Printf("Redis set error: %v", err)
		c.JSON(500, gin.H{"error": "验证码存储失败"})
		return
	}

	// 此处模拟发送验证码，实际上你可以调用短信服务（如阿里云短信服务）发送短信
	log.Printf("发送验证码: %s 到手机号: %s", code, phone)

	// 返回成功信息
	c.JSON(200, gin.H{"message": "验证码已发送", "phone": phone})
}

// 校验验证码接口
func VerifyCode(c *gin.Context) {
	// 获取手机号和验证码
	var json struct {
		Phone    string `json:"phone"`
		Code     string `json:"code"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "请求参数错误"})
		return
	}

	// 从 Redis 获取验证码
	storedCode, err := rdb.Get(ctx, json.Phone).Result()
	if err == redis.Nil {
		c.JSON(400, gin.H{"error": "验证码已过期或不存在"})
		return
	}
	if err != nil {
		log.Printf("Redis get error: %v", err)
		c.JSON(500, gin.H{"error": "获取验证码失败"})
		return
	}

	// 校验验证码
	if storedCode == json.Code {
		c.JSON(200, gin.H{"message": "验证码校验成功"})
	} else {
		c.JSON(400, gin.H{"error": "验证码错误"})
	}
}
