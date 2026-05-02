package models

import (
	"html/template"
)

type AdminLanguage struct {
	Locale      string
	Description string
}

type TestLanguage struct {
	Locale      string
	TwoLetterCode string
}

type VSPTBand struct {
	Locale string
	Level string
	Low int
	High int
}

type PreestWeight struct {
	Sa float64
	Vspt float64
	Coe float64
}

type PreestAssignment struct {
	Pe float64
	BookletId int
}


type SAWeight struct {
	Skill string
	WordId string
	Weight int
}

type Basket struct {
	Id int
	Type string
	ParentTestletId int
	ParentTestletPosition int
	Skill string
	Label string
	Prompt string
	GapText string
	Weight int
	MediaType string
	TextMedia string
	FileMedia string
}

type Item struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Skill string `json:"skill"`
	Position int `json:"position"`
	SubSkill string `json:"subskill"`
	Text template.HTML `json:"text"`
	Weight int `json:"weight"`
}

type Answer struct {
	Id int `json:"id"`
	ItemId int `json:"itemId"`
	Text string `json:"text"`
	Correct int `json:"correct"`
}

type ItemGrade struct {
	Tl string `json:"tl"`
	Skill string `json:"skill"`
	BookletId int `json:"bookletId"`
	RawScore int `json:"rawScore"`
	Ppe float64 `json:"ppe"`
	Se float64 `json:"se"`
	Grade int `json:"grade"`
}

type SAStatement struct {
	Locale string
	Skill string
	WordId string
	Statement string
}

type VSPTWord struct {
	WordId string
	Word string
	Valid int
	Weight int
}

type SAGrade struct {
	Skill string
	Rsc int
	Ppe float64
	Se float64
	Grade int
}
