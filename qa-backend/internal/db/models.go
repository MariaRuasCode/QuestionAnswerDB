package db

type AnswerStatus string

const (
	StatusNormal       AnswerStatus = "normal"
	StatusSpoiler      AnswerStatus = "spoiler"
	StatusNoComment    AnswerStatus = "no_comment"
	StatusHoldOntoThat AnswerStatus = "hold_onto_that"
)

type Answer struct {
	ID         int
	QuestionID int
	Content    string
	Status     AnswerStatus
	Timestamp  string
	IsFollowup bool
	Order      int
}
