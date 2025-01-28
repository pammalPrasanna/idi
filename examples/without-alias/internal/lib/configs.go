package lib

import (
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	port           int
	env            string
	baseURL        string
	trustedOrigins []string
	contextTimeout time.Duration

	jwtSecret string

	tokenExpiration time.Duration
}

const (
	defaultPort           int           = 4200
	defaultENV            string        = "development"
	defaultBaseURL        string        = "http://localhost:4200"
	defaultContextTimeout time.Duration = time.Second * 3

	defaultJWTSecret string = "762898e6b788045334ab11d18e3a9c21e4d91d0399a767f636b78172a72324cc"

	defaultTokenExpiration time.Duration = time.Hour
)

var (
	port                           = os.Getenv("PORT")
	env                            = os.Getenv("ENVIRONMENT")
	baseURL                        = os.Getenv("BASE_URL")
	trustedOrigins                 = os.Getenv("TRUSTED_ORIGINS")
	defaulttrustedOrigins []string = []string{"*"}
	contextTimeout                 = os.Getenv("CONTEXT_TIMEOUT")

	jwtSecret = os.Getenv("JWT_SECRET")

	tokenExpiration = os.Getenv("TOKEN_EXPIRATION")
)

func configure() (cfg *config, err error) {
	cfg = &config{
		port:           defaultPort,
		env:            defaultENV,
		baseURL:        defaultBaseURL,
		contextTimeout: defaultContextTimeout,

		jwtSecret: defaultJWTSecret,

		tokenExpiration: defaultTokenExpiration,
	}
	if port != "" {
		if i, err := strconv.Atoi(port); err == nil {
			cfg.port = i
		} else {
			return nil, err
		}
	}

	if env != "" {
		cfg.env = env
	}

	if baseURL != "" {
		cfg.baseURL = baseURL
	}

	if contextTimeout != "" {
		if i, err := strconv.Atoi(contextTimeout); err == nil {
			cfg.contextTimeout = time.Duration(i) * time.Second
		} else {
			return nil, err
		}
	}

	if trustedOrigins != "" {
		origins := strings.Split(trustedOrigins, ",")
		for i, url := range origins {
			origins[i] = strings.TrimSpace(url)
		}
		cfg.trustedOrigins = origins
	}

	if jwtSecret != "" {
		cfg.jwtSecret = jwtSecret
	}

	if tokenExpiration != "" {
		if i, err := strconv.Atoi(tokenExpiration); err == nil {
			cfg.tokenExpiration = time.Duration(i) * time.Hour
		} else {
			return nil, err
		}
	}

	return cfg, nil
}
