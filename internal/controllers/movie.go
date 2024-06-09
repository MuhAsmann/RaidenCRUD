package controllers

import (
	"raidencrud/internal/models"

	"github.com/sev-2/raiden"
)

type MovieRequest struct {
	Id int64 `path:"id"`
}

type MovieController struct {
	raiden.ControllerBase
	Http  string `path:"/movies" type:"rest"`
	Model *models.Movie
}
