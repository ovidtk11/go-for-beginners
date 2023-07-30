package model

import (
	"encoding/json"
	"log"
)

type ultraman struct {
	Name          string `json:"name"`
	Tool          string `json:"tool"`
	SpecialAttack string `json:"specialAttack"`
	TransformKit  string `json:"transformKit"`
}

func (u *ultraman) ovi() string {
	//TODO implement me
	panic("implement me")
}

func (u *ultraman) Json() string {
	bytes, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		log.Println("Error: ", err)
	}
	return string(bytes)
}

func (u *ultraman) GetName() string {
	return u.Name
}

func (u *ultraman) GetTool() string {
	return u.Tool
}

func (u *ultraman) GetSpecialAttack() string {
	return u.SpecialAttack
}

func (u *ultraman) GetTransformKit() string {
	return u.TransformKit
}

func (u *ultraman) SetName(s string) {
	u.Name = s
}

func (u *ultraman) SetTool(s string) {
	u.Tool = s
}

func (u *ultraman) SetSpecialAttack(s string) {
	u.SpecialAttack = s
}

func (u *ultraman) SetTransformKit(s string) {
	u.TransformKit = s
}

type IUltraman interface {
	Json() string
	IGetter
	ISetter
	ovi() string
}

type IGetter interface {
	GetName() string
	GetTool() string
	GetSpecialAttack() string
	GetTransformKit() string
}

type ISetter interface {
	SetName(s string)
	SetTool(s string)
	SetSpecialAttack(s string)
	SetTransformKit(s string)
}

func NewUltraman() IUltraman {
	return &ultraman{}
}
