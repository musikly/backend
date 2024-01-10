package soundcloud

import (
	"github.com/musikly/backend/config"
	"net/http"

	"github.com/musikly/backend/internal/logger"
)

// SoundCloud struct
type SoundCloud struct {
	Client *http.Client
	Logger logger.Service
	Config config.Config
}

// NewService function
func NewService(cfg config.Config, l logger.Service) *SoundCloud {
	return &SoundCloud{
		Client: &http.Client{},
		Logger: l,
		Config: cfg,
	}
}
