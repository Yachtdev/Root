package command

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"server/internal/config"
	"server/internal/server"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           `yachtdev-map-server`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			viper.SetEnvPrefix(`yachtdev-map-server`)
			viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
			viper.AutomaticEnv()
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return fmt.Errorf(`[yachtdev-map-server] failed to bind command line arguments: %+v`, err)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log, _ := zap.NewProduction()
			slog := log.Sugar()

			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			viper.AddConfigPath("./config/server")

			err := viper.ReadInConfig()
			if err != nil {
				panic(fmt.Errorf("[yachtdev-map-server] fatal error config file: %w", err))
			}

			cfg := config.NewConfig(viper.GetViper())

			slog.Infof(`[yachtdev-map-server] Starting API-server %s:%s`, cfg.Http.Host, cfg.Http.Port)
			srv := server.NewServer(cfg)

			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				s := <-sigCh
				slog.Infof("[yachtdev-map-server] got signal %v, attempting graceful shutdown", s)
				_ = srv.Shutdown(cmd.Context())
				wg.Done()
			}()

			go func() {
				slog.Infof(`[yachtdev-map-server] Serve http-server %s:%s`, cfg.HttpHost(), cfg.HttpPort())
				err := srv.ListenAndServe()
				if err != nil {
					slog.Errorf(`[yachtdev-map-server] error while serving: %+v`, err)
				}
			}()

			wg.Wait()
			slog.Infof("[yachtdev-map-server] clean shutdown")

			return nil
		},
	}

	return cmd
}
