package db

import (
	"database/sql"
	"fmt"
)

type Message struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Likes     int64  `json:"likes"`
	Ys        int64  `json:"ys"`
	CreatedAt string `json:"created_at"`
}

type Messages []Message

// Retrieves all messages from the database and returns a Messages slice.
func GetAllMessages(db *sql.DB) (Messages, error) {
	// An messages slice to hold data from returned rows.
	var messages Messages = make([]Message, 0)

	// Retrieve message data from database.
	rows, err := db.Query("SELECT * FROM messages ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("getAllMessages %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var message Message

		if err := rows.Scan(&message.ID, &message.Username, &message.Content, &message.Likes, &message.Ys, &message.CreatedAt); err != nil {
			return nil, fmt.Errorf("getAllMessages %v", err)
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllMessages %v", err)
	}

	return messages, nil
}

// Retrieves a message by id from the database and returns a Message struct.
func GetMessageById(db *sql.DB, id int) (Message, error) {
	var message Message

	row := db.QueryRow("SELECT * FROM messages WHERE id = ?", id)

	if err := row.Scan(&message.ID, &message.Username, &message.Content, &message.Likes, &message.Ys, &message.CreatedAt); err != nil {
		return message, fmt.Errorf("GetMessageById %v", err)
	}

	return message, nil
}

// Add a message to the database and return the message struct
func AddMessage(db *sql.DB, message Message) (Message, error) {
	stmt, err := db.Prepare("INSERT INTO messages (username, content, likes, ys) VALUES (?, ?, ?, ?)")
	if err != nil {
		return message, err
	}
	res, err := stmt.Exec(message.Username, message.Content, message.Likes, message.Likes)
	if err != nil {
		return message, err
	}
	// Return success response
	id, err := res.LastInsertId()
	if err != nil {
		return message, err
	}
	message.ID = id
	return message, nil
}

// Update a message its likes, ys, or content
func UpdateMessage(db *sql.DB, message Message) (Message, error) {
	stmt, err := db.Prepare("UPDATE messages SET likes = ?, ys = ?, content = ? WHERE id = ?")
	if err != nil {
		return message, err
	}
	res, err := stmt.Exec(message.Likes, message.Ys, message.Content, message.ID)
	if err != nil {
		return message, err
	}
	// Return success response
	id, err := res.LastInsertId()
	if err != nil {
		return message, err
	}
	message.ID = id
	return message, nil
}
