package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/agent"
	"github.com/deepakdinesh1123/mezcala/pkgs/backend/config"
	"github.com/deepakdinesh1123/mezcala/pkgs/backend/mq"
	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
	"github.com/deepakdinesh1123/mezcala/pkgs/backend/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
)

type Server struct {
	js        jetstream.JetStream
	agent     *agent.Agent
	envConfig *config.EnvConfig
	logger    *zerolog.Logger
	db        store.Store
}

func NewServer(ctx context.Context, envConfig *config.EnvConfig, logger *zerolog.Logger) (*http.Server, error) {

	logger.Debug().Msg("Starting NATS server")

	var nc *nats.Conn
	var err error
	if envConfig.EMBEDDED_NATS {
		nc, err = mq.StartEmbeddedNatsServer()
		if err != nil {
			return nil, err
		}
	} else {
		nc, err = nats.Connect(fmt.Sprintf("nats://%s:%s", envConfig.NATS_HOST, envConfig.NATS_PORT))
		if err != nil {
			return nil, err
		}
	}

	logger.Debug().Msg("NATS server started")

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}

	logger.Debug().Msg("NATS Jetstream created")

	err = mq.SetupJetStream(ctx, js)
	if err != nil {
		return nil, err
	}

	logger.Debug().Msg("NATS Jetstream setup complete")

	ag, err := agent.NewAgent(ctx, js, logger)

	if err != nil {
		return nil, err
	}

	logger.Debug().Msg("Agent created")

	errCh := ag.StartAsync(ctx)

	go func() {
		if err := <-errCh; err != nil {
			logger.Fatal().Err(err).Msg("Agent failed to start")
		}
	}()

	logger.Info().Msg("Agent started")

	db, err := store.NewDBStore(ctx, envConfig)

	srv := &Server{
		js:        js,
		envConfig: envConfig,
		logger:    logger,
		agent:     ag,
		db:        db,
	}

	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	handler := spec.NewStrictHandler(srv, []spec.StrictMiddlewareFunc{})
	spec.HandlerFromMux(handler, r)
	server := &http.Server{
		Addr:    envConfig.HOST + ":" + fmt.Sprint(envConfig.PORT),
		Handler: handlers.CORS(corsOptions, corsMethods, corsHeaders)(r),
	}

	logger.Info().Msgf("Starting server %s:%d", envConfig.HOST, envConfig.PORT)
	return server, nil
}

var _ spec.StrictServerInterface = (*Server)(nil)
