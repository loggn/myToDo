// package pkg

// import (
// 	"fmt"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// type User struct {
// 	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
// 	Name     string `gorm:"default:CONCAT('User_', TO_CHAR(NOW(), 'YYYYMMDD_HH24MISS'))" json:"name"`
// 	Account  string `json:"account"`
// 	Password string `json:"password"`
// 	Todos    []Todo `gorm:"foreignKey:UserID"`
// }

// type Todo struct {
// 	ID            uint      `gorm:"primaryKey" json:"id"`          // 任务唯一标识
// 	UserID        string    `gorm:"index;not null" json:"user_id"` // 关联用户的ID
// 	Text          string    `json:"text"`                          // 任务内容
// 	IsFinish      bool      `json:"isFinish"`                      // 任务是否完成
// 	IsImportant   bool      `json:"isImportant"`                   // 任务是否重要
// }

// // 自动生成 UUID 和 Name（数据库默认值）
// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	// ID 和 Name 字段由数据库自动填充，不需要在 GORM 中手动设置
// 	return nil
// }

// func DatabaseInit() {
// 	var err error
// 	DB, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("连接数据库失败：", err)
// 		return
// 	}

// 	err = DB.AutoMigrate(&User{}, &Todo{})
// 	if err != nil {
// 		fmt.Println("数据库迁移失败：", err)
// 		return
// 	}

//		fmt.Println("数据库初始化完成！")
//	}
package pkg

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID       string `gorm:"primaryKey;type:TEXT" json:"id"`
	Name     string `json:"name"` 
	Account  string `json:"account"`
	Password string `json:"password"`
	Todos    []Todo `gorm:"foreignKey:UserID"`
}

type Todo struct {
	ID          uint   `gorm:"primaryKey" json:"id"`          // 任务唯一标识
	UserID      string `gorm:"index;not null;constraint:OnDelete:CASCADE" json:"user_id"`  // 关联用户的ID
	Text        string `json:"text"`                          // 任务内容
	IsFinish    bool   `json:"isFinish"`                      // 任务是否完成
	IsImportant bool   `json:"isImportant"`                   // 任务是否重要
}

// BeforeCreate 钩子：自动填充 ID 和 Name 字段
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 自动生成 UUID
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	// 自动生成 Name 字段（例如：User_20250220_201234）
	if u.Name == "" {
		u.Name = fmt.Sprintf("User_%s", time.Now().Format("20060102_150405"))
	}
	return nil
}

func DatabaseInit() {
	var err error
	DB, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}

	// 自动迁移数据库结构
	err = DB.AutoMigrate(&User{}, &Todo{})
	if err != nil {
		fmt.Println("数据库迁移失败：", err)
		return
	}

	fmt.Println("数据库初始化完成！")
}
