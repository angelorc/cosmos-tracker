package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/angelorc/cosmos-tracker/client"
	"github.com/angelorc/cosmos-tracker/config"
	"github.com/angelorc/cosmos-tracker/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server start",
		Short: "Start rest-server",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(args[1])
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			log.Printf("connecting to bitsong grpc server...\n")
			bitsongClient, err := client.NewClient(cfg.Bitsong.GRPC, cfg.Bitsong.Denom)
			if err != nil {
				fmt.Errorf("grpc conn error: %w", err)
			}
			log.Printf("bitsong grpc server connected...\n")
			defer bitsongClient.Close()

			log.Printf("connecting to osmosis grpc server...\n")
			osmosisClient, err := client.NewClient(cfg.Osmosis.GRPC, cfg.Osmosis.Denom)
			if err != nil {
				fmt.Errorf("grpc conn error: %w", err)
			}
			log.Printf("osmosis grpc server connected...\n")
			defer bitsongClient.Close()

			logger, _ := zap.NewProductionConfig().Build()
			defer logger.Sync()

			chains := &client.Chains{
				Bitsong: bitsongClient,
				Osmosis: osmosisClient,
			}

			s := server.NewServer(chains, logger)
			eg, _ := errgroup.WithContext(ctx)

			log.Printf("starting server %s...", cfg.Server.Address)
			eg.Go(func() error {
				if err := s.Start(cfg.Server.Address); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Printf("error starting server: %v", err)
					return err
				}
				return nil
			})
			log.Printf("server started %s...", cfg.Server.Address)

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit

			logger.Info("gracefully shutting down")
			if err := s.ShutdownWithTimeout(10 * time.Second); err != nil {
				return fmt.Errorf("shutdown server: %w", err)
			}

			cancel()
			if err := eg.Wait(); !errors.Is(err, context.Canceled) {
				return err
			}

			return nil

		},
	}

	return cmd
}
