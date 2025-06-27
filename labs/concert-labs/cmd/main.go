package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"example.com/internal/config"
	"example.com/internal/postgres"
	"example.com/internal/resource"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type cliOpts struct {
	configPath   string
	migrationDir string
}

type serverOpts struct {
	envPrefix string
}

func newServerStart(opts *cliOpts, serverOpts *serverOpts) *cobra.Command {
	command := &cobra.Command{
		Use: "start",
		RunE: func(c *cobra.Command, args []string) error {
			conf, err := config.New(opts.configPath, serverOpts.envPrefix)
			if err != nil {
				log.Fatal().Err(err).Msg("Unable to load config file")
			}

			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			signal.Notify(ch, syscall.SIGTERM)
			go func() {
				oscall := <-ch
				log.Warn().Msgf("System call:%+v", oscall)
				cancel()
			}()

			log.Debug().Msgf("Running migration on %s", opts.migrationDir)
			if err := postgres.Migrate(opts.migrationDir, conf.DB.GetUrl(), true); err != nil {
				log.Fatal().Err(err).Msg("unable to run migration")
			}
		},
	}
	return command
}

func newServerSeed(opts *cliOpts, serverOpts *serverOpts) *cobra.Command {
	command := &cobra.Command{
		Use:   "seed",
		Short: "seed db",
		RunE: func(c *cobra.Command, args []string) error {
			conf, err := config.New(opts.configPath, serverOpts.envPrefix)
			if err != nil {
				log.Fatal().Err(err).Msg("unable to load config file")
			}

			ctx := log.With().Logger().WithContext(context.Background())
			ctx, cancel := context.WithCancel(ctx)
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			signal.Notify(ch, syscall.SIGTERM)
			go func() {
				oscall := <-ch
				log.Warn().Msgf("system call:%+v", oscall)
				cancel()
			}()

			clients := &resource.Client{
				DB: postgres.NewSQLx(conf.DB),
			}
			return catalogSvc.Seed(ctx)
		},
	}
	return command
}

func newServer(opts *cliOpts) *cobra.Command {
	serverOpts := &serverOpts{}
	command := &cobra.Command{
		Use:   "server",
		Short: "server subcommands",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
		},
	}
	command.AddCommand(
		newServerStart(opts, serverOpts),
		newServerSeed(opts, serverOpts),
	)

	command.PersistentFlags().StringVar(&serverOpts.envPrefix, "env-prefix", "CONCERT_SERVER", "config prefix")
	return command
}

func main() {
	opts := &cliOpts{}
	cli := &cobra.Command{
		Use:   "concert",
		Short: "all cli commands for concert service",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
		},
	}
	cli.AddCommand(
		newServer(opts),
	)
	cli.PersistentFlags().StringVar(&opts.configPath, "config", "/etc/concert/configs/config.yaml", "path to config file")
	cli.PersistentFlags().StringVar(&opts.migrationDir, "migration", "/etc/concert/migrations", "migration directory")

	err := cli.Execute()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to start server")
	}
}
