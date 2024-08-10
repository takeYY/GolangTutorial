package movie

import (
	"context"
	"golang_tutorial/config"
	"golang_tutorial/internal/common"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// 映画エンドポイントのハンドラを持つ構造体
type (
	ReaderHandler struct {
		queryService *QueryService
	}
	WriterHandler struct {
		commandService *CommandService
	}
	Handler struct {
		ReaderHandler ReaderHandler
		WriterHandler WriterHandler
	}
)

// 新しいHandlerを作成
func NewHandler(cfg *config.Config) *Handler {
	queryService := NewQueryService(cfg)
	commandService := NewCommandService(cfg)

	return &Handler{
		ReaderHandler: ReaderHandler{
			queryService: queryService,
		},
		WriterHandler: WriterHandler{
			commandService: commandService,
		},
	}
}

// 映画エンドポイントのルーティングを設定
func (h *Handler) RegisterRoutes(e *echo.Echo) {
	// Read
	e.GET("/movies/:id", h.ReaderHandler.GetMovie)
	e.GET("/movies", h.ReaderHandler.GetMovies)
	// Create, Update, Delete
	e.POST("/movie", h.WriterHandler.CreateMovie)
}

// 映画IDで映画情報を取得
func (rh *ReaderHandler) GetMovie(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{
			Message: "invalid id",
		})
	}
	id := int32(id64)

	movie, err := rh.queryService.GetMovieDetails(context.Background(), &id)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.ErrorResponse{
			Message: "movie is not found",
		})
	}

	return c.JSON(http.StatusOK, movie)
}

func (rh *ReaderHandler) GetMovies(c echo.Context) error {
	movies, err := rh.queryService.GetMovies(context.Background())
	if err != nil {
		return c.JSON(http.StatusNotFound, common.ErrorResponse{
			Message: "movies are not found",
		})
	}

	return c.JSON(http.StatusOK, movies)
}

// 映画を作成
func (wh *WriterHandler) CreateMovie(c echo.Context) error {
	var m NewMovie
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	movie, err := wh.commandService.CreateMovie(context.Background(), &m)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.ErrorResponse{
			Message: "movie can not created",
		})
	}

	return c.JSON(http.StatusOK, movie)
}
