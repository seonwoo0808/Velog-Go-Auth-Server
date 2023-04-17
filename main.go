// pentag.kr/BuildinAuthVelog/main.go

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"pentag.kr/BuildinAuthVelog/database" // 추가
)

func main() {
	app := fiber.New()

	database.ConnectDB()    // 추가
	database.ConnectRedis() // 추가

	log.Fatal(app.Listen(":" + "3000"))
}
