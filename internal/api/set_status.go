package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MyData структура, представляющая данные для отправки на Django-сервер
type MyData struct {
	EmployeeID  string `json:"employee_id"`
	RequestID   string `json:"request_id"`
	Security    string `json:"security_value"`
}

// Функция для отправки данных в Django-сервер
func sendDataToDjango(employeeID, requestID, securityValue string) {
	// Формирование данных для отправки
	data := MyData{
		EmployeeID: employeeID,
		RequestID:  requestID,
		Security:   securityValue,
	}

	// Сериализация структуры в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка при сериализации данных:", err)
		return
	}

	// Создание PUT-запроса
	url := "http://your-django-server/put-security/"  // Замените на ваш реальный URL
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}

	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode == http.StatusOK {
		fmt.Println("PUT запрос успешно обработан")
	} else {
		fmt.Println("Не удалось обработать PUT запрос")
	}
}

// Функция для "расчётов" randomStatus
func randomStatus() string {
	time.Sleep(5 * time.Second) // Задержка на 5 секунд

	// Просто для примера, можно заменить на ваш логический расчет
	if time.Now().Unix()%2 == 0 {
		return "true"
	} else {
		return "false"
	}
}

func SetStatusHandler(c *gin.Context) {
	// Получение значения "employee_id" и "request_id" из запроса
	employeeID := c.PostForm("employee_id")
	requestID := c.PostForm("request_id")

	// Выполнение расчётов с randomStatus
	securityValue := randomStatus()

	// Запуск горутины для отправки данных в Django-сервер
	go sendDataToDjango(employeeID, requestID, securityValue)

	c.JSON(http.StatusOK, gin.H{"message": "Инициировано обновление статуса"})
}
