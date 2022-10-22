package dto

type Record struct {
	SentenceId    string `json:"sentenceId" xorm:"sentenceId"`
	Text          string `json:"text" xorm:"text"`
	VieSentenceId string `json:"vieSentenceId" xorm:"vieSentenceId"`
	VieText       string `json:"vieText" xorm:"vieText"`
	AudioUrl      string `json:"audioUrl" xorm:"audioUrl"`
}

type Records []Record
