package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"time"
)

const botToken = "7772983471:AAGNf5n6eMiqL2IENFGYCXYssDnXH7ytTVQ"
const apiURL = "https://api.telegram.org/bot" + botToken
const dbConnStr = "user=postgres password=123 dbname=telegram_bot1 sslmode=disable"

type Update struct {
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			Text string `json:"text"`
			Chat struct {
				ID int `json:"id"`
			} `json:"chat"`
		} `json:"message"`
	} `json:"result"`
}

func getfactFromDB(db *sql.DB) (string, error) {
	var fact string
	err := db.QueryRow("SELECT text FROM facts ORDER BY RANDOM() LIMIT 1").Scan(&fact)
	if err != nil {
		return "Error getting fact", err
	}
	return fact, nil
}

func getUpdates(offset int) (Update, error) {
	resp, err := http.Get(fmt.Sprintf("%s/getUpdates?offset=%d", apiURL, offset))
	if err != nil {
		return Update{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Update{}, err
	}
	var update Update
	err = json.Unmarshal(body, &update)
	if err != nil {
		return Update{}, err
	}
	return update, nil
}

func sendMessage(chatID int, text string) error {
	resp, err := http.Get(fmt.Sprintf("%s/sendMessage?chat_id=%d&text=%s", apiURL, chatID, text))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	offset := 0

	for {
		updates, err := getUpdates(offset)
		if err != nil {
			fmt.Println("Update request error:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		for _, result := range updates.Result {
			offset = result.UpdateID + 1

			if result.Message.Text == "/fact" {
				fact, err := getfactFromDB(db)
				if err != nil {
					fmt.Println("Error getting fact from DB:", err)
					continue
				}
				err = sendMessage(result.Message.Chat.ID, fact)
				if err != nil {
					fmt.Println("Error sending message:", err)
				} else {
					fmt.Println("Fact sent!")
				}
			}

		}
		time.Sleep(2 * time.Second)
	}
}
