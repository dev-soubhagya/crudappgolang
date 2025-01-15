package models

import (
	"crudappgolang/config"
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/xuri/excelize/v2"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ProcessExcel(xlsx *excelize.File) {
	rows, _ := xlsx.GetRows("Sheet1")
	for i, row := range rows {
		if i == 0 || len(row) < 2 {
			continue
		}
		_, err := config.DB.Exec("INSERT IGNORE INTO users (name, email) VALUES (?, ?)", row[0], row[1])
		if err != nil {
			log.Printf("Failed to insert row %d: %v", i, err)
		}
	}
}

func GetCachedData() ([]User, error) {
	conn := config.RedisPool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", "imported_data"))
	if err != nil {
		return nil, err
	}

	var users []User
	json.Unmarshal([]byte(data), &users)
	return users, nil
}

func FetchAllUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT name, email FROM users") //we can do pagination for large data sets
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Name, &user.Email)
		users = append(users, user)
	}

	cacheData, _ := json.Marshal(users)
	conn := config.RedisPool.Get()
	conn.Do("SETEX", "imported_data", 300, cacheData)

	return users, nil
}

func UpdateUser(id string, user User) error {
	_, err := config.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err == nil {
		conn := config.RedisPool.Get()
		defer conn.Close()
		conn.Do("DEL", "imported_data")
	}
	return err
}
