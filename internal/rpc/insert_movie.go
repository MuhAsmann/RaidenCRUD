package rpc

import (
	"github.com/sev-2/raiden"
)

type InsertMovieParams struct {
	Title string `json:"title" column:"name:title;type:text"`
	Description string `json:"description" column:"name:description;type:text"`
}
type InsertMovieResult interface{}

type InsertMovie struct {
	raiden.RpcBase
	Params   *InsertMovieParams `json:"-"`
	Return   InsertMovieResult `json:"-"`
}

func (r *InsertMovie) GetName() string {
	return "insert_movie"
}

func  (r *InsertMovie) UseParamPrefix() bool {
	return false
}

func (r *InsertMovie) GetReturnType() raiden.RpcReturnDataType {
	return raiden.RpcReturnDataTypeVoid
}

func (r *InsertMovie) GetRawDefinition() string {
	return `BEGIN INSERT INTO movies (:title, :description) VALUES (:title, :description); END;`
}