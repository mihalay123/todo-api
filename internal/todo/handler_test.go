package todo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	body := []byte(`{"title":"Test task","done":false}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()

	TodoHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestGetTodos(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()

	TodoHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestUpdateTodo(t *testing.T) {
	body := []byte(`{"title":"Update me","done":false}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()
	TodoHandler(w, req)

	var created Todo
	json.NewDecoder(w.Result().Body).Decode(&created)

	updated := Todo{Title: "Updated", Done: true}
	updatedBody, _ := json.Marshal(updated)
	updateReq := httptest.NewRequest(http.MethodPut, "/todos/"+strconv.Itoa(created.ID), bytes.NewReader(updatedBody))
	updateW := httptest.NewRecorder()
	TodoItemHandler(updateW, updateReq)

	if updateW.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 on update, got %d", updateW.Result().StatusCode)
	}
}

func TestDeleteTodo(t *testing.T) {
	body := []byte(`{"title":"Delete me","done":false}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()
	TodoHandler(w, req)

	var created Todo
	json.NewDecoder(w.Result().Body).Decode(&created)

	deleteReq := httptest.NewRequest(http.MethodDelete, "/todos/"+strconv.Itoa(created.ID), nil)
	deleteW := httptest.NewRecorder()
	TodoItemHandler(deleteW, deleteReq)

	if deleteW.Result().StatusCode != http.StatusNoContent {
		t.Errorf("expected 204 on delete, got %d", deleteW.Result().StatusCode)
	}
}

func TestUpdateTodo_InvalidID(t *testing.T) {
	body := []byte(`{"title":"New","done":true}`)
	req := httptest.NewRequest(http.MethodPut, "/todos/abc", bytes.NewReader(body))
	w := httptest.NewRecorder()

	TodoItemHandler(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 for invalid ID, got %d", w.Result().StatusCode)
	}
}

func TestDeleteTodo_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/todos/999", nil)
	w := httptest.NewRecorder()

	TodoItemHandler(w, req)

	if w.Result().StatusCode != http.StatusNotFound {
		t.Errorf("expected 404 for missing todo, got %d", w.Result().StatusCode)
	}
}

func TestCreateTodo_EmptyTitle(t *testing.T) {
	body := []byte(`{"title":"","done":false}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()

	TodoHandler(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 for empty title, got %d", w.Result().StatusCode)
	}
}
