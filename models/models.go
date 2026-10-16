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
	Locale string `json:"locale"`
	Level string `json:"level"`
	Low int `json:"low"`
	High int `json:"high"`
}

type PreestWeight struct {
	Sa float64 `json:"sa"`
	Vspt float64 `json:"vspt"`
	Coe float64 `json:"coe"`
}

type PreestAssignment struct {
	Pe float64 `json:"pe"`
	BookletId int `json:"bookletId"`
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
	Locale string `json:"locale"`
	Skill string `json:"skill"`
	WordId string
	Statement string
}

type VSPTWord struct {
	WordId string `json:"word_id"`
	Word string `json:"word"`
	Valid int `json:"valid"`
	Weight int `json:"weight"`
}

type SAGrade struct {
	Skill string `json:"skill"`
	Rsc int `json:"rsc"`
	Ppe float64 `json:"ppe"`
	Se float64 `json:"se"`
	Grade int `json:"grade"`
}
