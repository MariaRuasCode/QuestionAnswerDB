package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnswerInput struct {
	Content    string `json:"content"`
	Status     string `json:"status"`
	IsFollowup bool   `json:"is_followup"`
	Order      int    `json:"order"`
}

// Middleware-like input binding
func withAnswerInput(handler func(*gin.Context, AnswerInput)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input AnswerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		handler(c, input)
	}
}

func EndQuestionHandler(c *gin.Context, input AnswerInput) {
	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	

	c.JSON(http.StatusCreated, gin.H{"message": "Answer saved"})
}

func ConfigQuestionEndpoints(r *gin.Engine) {
	v1 := r.Group("/api/v1/questions")
	{
		v1.POST("/:id/end", withAnswerInput(EndQuestionHandler))
	}
}
