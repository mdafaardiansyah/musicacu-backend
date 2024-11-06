package spotify

//this folder spotify can moved to external folder instead of internal
import (
	"github.com/mdafaardiansyah/musicacu-backend/internal/configs"
	"github.com/mdafaardiansyah/musicacu-backend/pkg/httpclient"
	"time"
)

type outbound struct {
	cfg         *configs.Config
	client      httpclient.HTTPClient
	AccessToken string
	TokenType   string
	ExpiredAt   time.Time
}

func NewSpotifyOutbound(cfg *configs.Config, client httpclient.HTTPClient) *outbound {
	return &outbound{
		cfg:    cfg,
		client: client,
	}
}
