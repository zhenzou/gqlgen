package model

import "github.com/99designs/gqlgen/_examples/guid/guid"

type Todo struct {
	ID   *guid.GUID `json:"id"`
	Text string     `json:"text"`
	Done bool       `json:"done"`
}

type TodosFilterInput struct {
	IDEq  int                 `json:"IDEq,omitempty"`
	IDNeq int                 `json:"IDNeq,omitempty"`
	IDIn  []*guid.GUID        `json:"IDIn,omitempty"`
	IDNin []*guid.GUID        `json:"IDNin,omitempty"`
	IDLt  *guid.GUID          `json:"IDLt,omitempty"`
	IDLte *guid.GUID          `json:"IDLte,omitempty"`
	IDGt  *guid.GUID          `json:"IDGt,omitempty"`
	IDGte *guid.GUID          `json:"IDGte,omitempty"`
	And   []*TodosFilterInput `json:"AND,omitempty"`
	Or    []*TodosFilterInput `json:"OR,omitempty"`
	Not   []*TodosFilterInput `json:"NOT,omitempty"`
}
