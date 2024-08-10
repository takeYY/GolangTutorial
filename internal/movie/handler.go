package movie

import (
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
		commandService *QueryService
	}
	Handler struct {
		ReaderHandler ReaderHandler
		WriterHandler WriterHandler
	}
)

// 新しいHandlerを作成
func NewHandler(rdb *config.Database, wdb *config.Database) *Handler {
	queryService := NewQueryService(rdb)
	commandService := NewQueryService(wdb)

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
	// Create, Update, Delete
	// e.POST("/movies", h.WriterHandler.CreateMovie)
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

	movie, err := rh.queryService.GetMovieDetails(&id)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.ErrorResponse{
			Message: "movie is not found",
		})
	}

	return c.JSON(http.StatusOK, movie)
}

// 映画を作成
/*
func (wh *WriterHandler) CreateMovie(c echo.Context) error {
	return c.JSON(http.StatusOK, movie)
}
*/
