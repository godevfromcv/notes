package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Пример того, как мы можем получить userID. Это зависит от того, как вы храните и передаете userID.
	userID := r.Header.Get("UserID")

	// Конвертируем строковое значение userID в int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Получаем заметки по userID
	notes, err := models.GetNotesByUserID(userIDInt)
	if err != nil {
		http.Error(w, "Failed to retrieve notes", http.StatusInternalServerError)
		return
	}

	// Отправляем заметки обратно клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Добавьте логику для добавления новой заметки
}

// ... другие обработчики, если они есть
