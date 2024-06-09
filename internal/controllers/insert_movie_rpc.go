package controllers

import (
	"github.com/sev-2/raiden"
)

type InsertMovieRPC struct { // Add your request data
}

type InsertMovieRPCResponse struct {
	Message string `json:"message"`
}

type InsertMovieRPCController struct {
	raiden.ControllerBase
	Http    string `path:"/movie" type:"rpc"`
	Payload *InsertMovieRPC
	Result  InsertMovieRPCResponse
}

func (c *InsertMovieRPCController) Post(ctx raiden.Context) error { // Add your code here
	c.Result.Message = "Welcome to raiden"
	return ctx.SendJson(c.Result)
}
