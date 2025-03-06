package handlers

import (
	"fam/common"
	"fam/db"
	"fam/service"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

type VideoQueryHandler struct {
	videoQueryService service.VideoQueryService
}

func NewVideoQueryHandler(videoQueryService service.VideoQueryService) *VideoQueryHandler {
	return &VideoQueryHandler{videoQueryService: videoQueryService}
}

func (handler *VideoQueryHandler) GetVideos(ctx *gin.Context) {
	context := ctx.Request.Context()
	lastFetchedTime := ctx.Query("lastFetchedTime")
	size := ctx.Query("size")

	limit, err := strconv.Atoi(size)
	if err != nil {
		limit = math.MaxInt32 // default
	}
	page := &db.Page{
		LastFetchedTime: lastFetchedTime,
		Size:            limit,
	}
	result, err := handler.videoQueryService.GetVideos(&context, page)
	if err != nil {
		common.HandleError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, result)

}

func (handler *VideoQueryHandler) SearchVideos(ctx *gin.Context) {
	context := ctx.Request.Context()
	searchQuery := ctx.Query("query")
	lastFetchedTime := ctx.Query("lastFetchedTime")
	size := ctx.Query("size")

	limit, err := strconv.Atoi(size)
	if err != nil {
		limit = math.MaxInt32 // default
	}
	page := &db.Page{
		LastFetchedTime: lastFetchedTime,
		Size:            limit,
	}
	result, err := handler.videoQueryService.SearchVideos(&context, searchQuery, page)
	if err != nil {
		common.HandleError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
