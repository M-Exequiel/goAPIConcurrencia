package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/goConcurrencyAPI/src/api/services/user"
	"github.com/mercadolibre/goConcurrencyAPI/src/api/utils/apierrors"
	"net/http"
	"strconv"
)

const (
	paramUserId = "idUser"
)

func GetUser(context *gin.Context) {
	idUser := context.Param(paramUserId)
	i, err := strconv.ParseInt(idUser, 10, 64)
	if err != nil {
		apiErr := &apierrors.APIError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		context.JSON(apiErr.Status, apiErr)
		return
	}
	finalResult, apiErr:= user.GetUserFromAPI(i)
	if apiErr != nil {
		context.JSON(apiErr.Status, apiErr)
		return
	}
	fmt.Println("Resultado final: ", finalResult)
	context.JSON(200, finalResult)
}
