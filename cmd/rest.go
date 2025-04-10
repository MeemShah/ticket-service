package cmd

import (
	"ticket-service/cache"
	"ticket-service/config"
	web "ticket-service/controller"
	"ticket-service/controller/handlers"
	"ticket-service/controller/middlewares"
	"ticket-service/controller/utils"
	"ticket-service/logger"
	"ticket-service/ticket"

	"github.com/spf13/cobra"
)

var serveRestCmd = &cobra.Command{
	Use:   "serve-rest",
	Short: "start a rest server",
	RunE:  serveRest,
}

func serveRest(cmd *cobra.Command, args []string) error {
	cnf := config.GetConfig()

	utils.InitValidator()
	logger.SetupLogger(cnf.ServiceName)
	readRedisClinet, err := cache.NewRedisClient(cnf.ReadRedisUrl, cnf.EnableSSLMode)
	if err != nil {
		return err
	}

	writeRedisClinet, err := cache.NewRedisClient(cnf.WriteRedisUrl, cnf.EnableSSLMode)
	if err != nil {
		return err
	}

	cache := cache.NewCache(readRedisClinet, writeRedisClinet)

	ticket_svc := ticket.NewService(cnf, cache)

	handlers := handlers.NewHandlers(cnf, ticket_svc)

	middleware := middlewares.NewMiddleware(cnf, cache)
	server := web.NewServer(cnf, handlers, middleware)
	server.Run()
	server.Wg.Wait()

	return nil
}
