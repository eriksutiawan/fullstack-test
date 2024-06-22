package delivery

import (
	"backend/internal/app/submission/usecase"
	"backend/internal/pkg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmissionRegisterHandler(c *gin.Context, submission usecase.ISubmissionService) {
	var dto usecase.RegisterDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := submission.Register(c, dto)
	fmt.Println("masuk : ", err)

	if err != nil {
		switch err.Error() {
		case pkg.BadRequestError.Error():
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission created successfully"})
}
