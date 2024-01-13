package handler

import (
	"strconv"
	"todoist/model"
	"todoist/view/todo"

	"github.com/anthdm/slick"
)

type TodoHandler struct {
	Todos []model.Todo
}

func (h *TodoHandler) Index(c *slick.Context) error {
	return c.Render(todo.Index(h.Todos))
}

func (h *TodoHandler) Toggle(c *slick.Context) error {
	index, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.Text(404, "Not found")
	}

	todo := h.Todos[index]
	todo.Completed = !todo.Completed
	h.Todos[index] = todo

	return c.Redirect("/", 301)
}

func (h *TodoHandler) Store(c *slick.Context) error {
	t := model.Todo{Title: c.FormValue("title"), Completed: false}
	h.Todos = append(h.Todos, t)

	return c.Redirect("/", 301)
}

func (h *TodoHandler) Destroy(c *slick.Context) error {
	index, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.Text(404, "Not found")
	}

	h.Todos = append(h.Todos[:index], h.Todos[index+1:]...)

	return c.Redirect("/", 301)
}
