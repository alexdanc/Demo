package handlers

import (
	"OneProject/internal/TaskService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RequestBodyHandlers struct {
	service TaskService.RequestBodyService
}

func NewRequestBodyHandlers(s TaskService.RequestBodyService) *RequestBodyHandlers {
	return &RequestBodyHandlers{service: s}
}

func (h *RequestBodyHandlers) GetHandler(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *RequestBodyHandlers) PostHandler(c echo.Context) error {
	var req TaskService.RequestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Ошибка записи JSON",
		})
	}
	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Задача не может быть пустой",
		})
	}
	body, err := h.service.CreatesTask(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Произошла ошибка при создании задачи",
		})
	}

	return c.JSON(http.StatusCreated, body)
}

func (h *RequestBodyHandlers) PatchHandler(c echo.Context) error {
	id := c.Param("id")
	var body TaskService.RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "Ошибка парсинга JSON")
	}
	updatedTask, err := h.service.UpdateTask(id, body.Task)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedTask)
}

func (h *RequestBodyHandlers) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTaskByID(id); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
