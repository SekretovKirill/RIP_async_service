package main

import (
    "async/internal/pkg/api"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Обработчик POST-запроса для set_status
    r.POST("/check", api.SetStatusHandler)

    r.Run(":8080") // Замените на ваш порт
}
