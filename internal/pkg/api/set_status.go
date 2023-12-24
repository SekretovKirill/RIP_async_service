package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

// MyData структура, представляющая данные для отправки на Django-сервер
type MyData struct {
	EmployeeID string `json:"employee_id"`
	RequestID  string `json:"request_id"`
	Security   bool   `json:"security_value"`
}

// Функция для отправки данных в Django-сервер
func sendDataToDjangoAsync(employeeID, requestID string) {
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(6) + 5
	time.Sleep(time.Duration(delay) * time.Second)
	securityValue := rand.Intn(2) == 0

	// Формирование данных для отправки
	data := MyData{
		EmployeeID: employeeID,
		RequestID:  requestID,
		Security:   securityValue,
	}

	// Сериализация структуры в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка сериализации данных:", err)
		return
	}

	// Создание PUT-запроса
	url := "http://192.168.1.132:8000/put-security/"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return
	}

	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("PUT-запрос успешно обработан. Security: %v\n", securityValue)
	} else {
		fmt.Println("Не удалось обработать PUT-запрос")
	}
}

// SetStatusHandler обработчик POST-запроса для set_status
func SetStatusHandler(c *gin.Context) {
	employeeID := c.PostForm("employee_id")
	requestID := c.PostForm("request_id")

	// Запуск горутины для отправки данных в фоновом режиме
	go sendDataToDjangoAsync(employeeID, requestID)

	c.JSON(http.StatusOK, gin.H{"message": "Инициировано обновление статуса"})
}
