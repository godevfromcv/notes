package models

import (
	"backend/database"
)

type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

func AddNote(note Note) (int64, error) {
	db, _ := database.InitDB()
	defer db.Close()

	res, err := db.Exec("INSERT INTO notes(content, user_id) VALUES (?, ?)", note.Content, note.UserID)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func GetNotesByUserID(userID int) ([]Note, error) {
	db, _ := database.InitDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, content FROM notes WHERE user_id=?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Content)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
