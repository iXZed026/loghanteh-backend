package response

import (
	"fmt"
	"net/http"

	appI18n "loghanteh-project/internal/i18n"

	"github.com/gin-gonic/gin"
)

func Success(
	c *gin.Context,
	messageKey string,
	data any,
) {
	lang := c.GetHeader("Accept-Language")

	message := appI18n.TranslateSuccess(
		messageKey,
		lang,
	)

	fmt.Println("Accept-Language:", lang)
	fmt.Println("messageKey:", messageKey)
	fmt.Println("message:", message)

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": message,
			"data":    data,
		},
	)
}
