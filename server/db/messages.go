package db

import (
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

// Retrieves all messages from the database and returns a Messages slice.
func GetAllMessages() ([]Message, error) {
	// Initialize database connection.
	db, err := InitializeDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// An messages slice to hold data from returned rows.
	var messages []Message = make([]Message, 0)

	// Retrieve message data from database.
	rows, err := db.Query("SELECT id, username, content, likes, ys, created_at FROM messages ORDER BY created_at DESC")
	if err != nil {
		return nil, err
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
func GetMessageById(id int) (*Message, error) {
	// Initialize database connection.
	db, err := InitializeDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, username, content, likes, ys, created_at FROM messages WHERE id = ?", id)

	var message Message
	if err := row.Scan(&message.ID, &message.Username, &message.Content, &message.Likes, &message.Ys, &message.CreatedAt); err != nil {
		return nil, fmt.Errorf("getMessageById %v", err)
	}

	return &message, nil
}

// Add a message to the database and return the message struct
func AddMessage(message *Message) (*Message, error) {
	// Initialize database connection.
	db, err := InitializeDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO messages (username, content, likes, ys) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(&message.Username, &message.Content, &message.Likes, &message.Likes)
	if err != nil {
		return nil, err
	}
	// Return success response
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	updatedMessage, err := GetMessageById(int(id))
	if err != nil {
		return nil, err
	}

	// Grab the message from the database
	return updatedMessage, nil
}

// Update a message its likes, ys, or content
func UpdateMessage(message *Message) (*Message, error) {
	// Initialize database connection.
	db, err := InitializeDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE messages SET likes = ?, ys = ?, content = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(&message.Likes, &message.Ys, &message.Content, &message.ID)
	if err != nil {
		return nil, err
	}
	// Return success response
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	message.ID = id
	return message, nil
}
