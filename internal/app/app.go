package app

import (
	"context"
	"log/slog"
	"loghanteh-project/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config config.Config
	Logger *slog.Logger

	Context context.Context
	Cancel  context.CancelFunc

	Router *gin.Engine
	Server *http.Server
}
