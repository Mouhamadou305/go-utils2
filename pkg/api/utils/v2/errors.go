package v2

import (
	"fmt"

	"github.com/Mouhamadou305/go-utils2/pkg/api/models"
)

// ErrWithStatusCode message
const ErrWithStatusCode = "error with status code %d"

func handleErrStatusCode(statusCode int, body []byte) *models.Error {
	respErr := &models.Error{}
	if err := respErr.FromJSON(body); err == nil && respErr != nil {
		return respErr
	}

	return buildErrorResponse(fmt.Sprintf(ErrWithStatusCode, statusCode))
}
