package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	// 1、插入单条数据
	//user := User{
	//	Name: "Alice",
	//	Age:  21,
	//}
	//result1 := db.Create(&user)
	//
	// 2、插入多条数据
	//users := []*User{
	//	{Name: "Bob", Age: 22},
	//	{Name: "Cindy", Age: 23},
	//}
	//result2 := db.Create(users)

	//fmt.Println(result1.RowsAffected)
	//fmt.Println(result2.RowsAffected)

	// 3、用指定的字段去插入数据
	//user3 := User{
	//	Name: "sunl",
	//	Age:  27,
	//}
	//db.Select("Name", "Age", "CreatedAt").Create(&user3)

	// 4、省略指定字段去插入数据
	//user4 := User{
	//	Name: "test",
	//	Age:  33,
	//}
	//db.Omit("Name").Create(&user4)

	// 5、批量插入
	var users = []User{{Name: "test1"}, {Name: "test2"}, {Name: "test3"}}
	db.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID)
	}
}
