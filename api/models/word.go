package models

type Word struct {
	Word  string `json:"word"`
	Point int32  `json:"point"`
}

type GetWordResponse struct {
	Words []*Word `json:"words"`
	Count int32   `json:"count"`
}
