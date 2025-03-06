package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleError(err error, ctx *gin.Context) {
	log.Println(err.Error())
	var badStateError *BadStateError
	okBadStateError := errors.As(err, &badStateError)
	if okBadStateError {
		StateError(ctx, err.Error())
		return
	}
	var alreadyExist *AlreadyExist
	okAlreadyExistError := errors.As(err, &alreadyExist)
	if okAlreadyExistError {
		ConflictError(ctx, err.Error())
		return
	}
	var internalError *InternalError
	okInternalError := errors.As(err, &internalError)
	if okInternalError {
		ServerInternalError(ctx, err.Error())
		return
	}
	var unauthorizedError *UnauthorizedError
	okUnauthorizedError := errors.As(err, &unauthorizedError)
	if okUnauthorizedError {
		Unauthorized(ctx, "unauthorized", err.Error())
		return
	}
	var notFound *NotFound
	notFoundError := errors.As(err, &notFound)
	if notFoundError {
		NotFoundError(ctx, err.Error())
		return
	}
	ServerInternalError(ctx, err.Error())
}
