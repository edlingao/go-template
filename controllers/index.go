package controllers

import (
	"exercise-app/components"
	"exercise-app/models"
	"exercise-app/utils"
	view_error "exercise-app/view/error"
	view_example "exercise-app/view/example"
	view_index "exercise-app/view/index"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexController struct{}

var indexModel models.Index

func (i IndexController) Show(c echo.Context) error {
	model, err := models.IndexModel.GetIndex("Name")
	indexModel = model

	if err != nil {
		log.Print(err)
		return utils.Render(c, view_error.Error(err), http.StatusBadGateway)
	}

	return utils.Render(
		c,
		view_index.Index(
			indexModel.Name(),
			indexModel.GetCount(),
		),
		http.StatusOK,
	)
}

func (i IndexController) ChangeName(c echo.Context) error {
	newName := c.Param("name")

	indexModel.SetName(newName)

	return utils.Render(
		c,
		view_index.Index(
			indexModel.Name(),
			indexModel.GetCount(),
		),
		http.StatusOK,
	)
}

func (i IndexController) AddCount(c echo.Context) error {
	count := 1

	indexModel.AddCount(count)

	return utils.Render(c,
		components.Count(indexModel.GetCount()),
		http.StatusOK,
	)
}

func (i IndexController) Example(c echo.Context) error {
  return utils.Render(c,
    view_example.Example(),
    http.StatusOK,
  )
}
