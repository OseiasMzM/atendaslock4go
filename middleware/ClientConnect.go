package middleware

import (
	"atendaslock4go/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateClientConnect struct {
	CreateClientConnectResponse
	Computer string `json:"computer"`
	User     string `json:"user"`
}
type CreateClientConnectResponse struct {
	ID        uint `json:"id"`
	CreatedAt int  `json:"createdAt"`
	UpdatedAt int  `json:"updatedAt,omitempty"`
	DeletedAt int  `json:"deletedAt,omitempty"`
}

// TODO -> Gabrel, please check the data we are passing in the request body is correct, if not, leave a comment here.

func ClientConnect(ctx *gin.Context) {

	request := CreateClientConnect{}

	ctx.BindJSON(&request)

	timestampCreatedAt := CreateClientConnectResponse{}
	timestampCreatedAt.toTimestamp()

	if err := request.Validate(); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	var bodyRequest = CreateClientConnect{
		Computer: request.Computer,
		User:     request.User,
		CreateClientConnectResponse: CreateClientConnectResponse{
			CreatedAt: timestampCreatedAt.CreatedAt,
		},
	}

	ctx.Set("bodyRequestKEY", bodyRequest)
	ctx.Next()
}

func (r *CreateClientConnect) Validate() error {
	if r.User == "" && r.Computer == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Computer == "" {
		return errParamIsRequired("computer", "string")
	}
	if r.User == "" {
		return errParamIsRequired("user", "string")
	}
	return nil
}

func (c *CreateClientConnectResponse) toTimestamp() int {
	now := time.Now().Unix()
	c.CreatedAt = int(now)
	return c.CreatedAt
}
