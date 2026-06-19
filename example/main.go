package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	gologging "github.com/pucora/velonetics-gologging/v2"
	koanf "github.com/pucora/velonetics-koanf"
	martian "github.com/pucora/velonetics-martian/v2"
	"github.com/pucora/lura/v2/proxy"
	veloneticsgin "github.com/pucora/lura/v2/router/gin"
	"github.com/pucora/lura/v2/transport/http/client"
	"github.com/pucora/lura/v2/transport/http/server"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "/etc/pucora/configuration.json", "Path to the configuration filename")
	flag.Parse()

	parser := koanf.New()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, err := gologging.NewLogger(serviceConfig.ExtraConfig)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	backendFactory := martian.NewBackendFactory(logger, client.DefaultHTTPRequestExecutor(client.NewHTTPClient))

	routerFactory := veloneticsgin.NewFactory(veloneticsgin.Config{
		Engine:         gin.Default(),
		Logger:         logger,
		HandlerFactory: veloneticsgin.EndpointHandler,
		ProxyFactory:   proxy.NewDefaultFactory(backendFactory, logger),
		RunServer:      server.RunServer,
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case sig := <-sigs:
			logger.Info("Signal intercepted:", sig)
			cancel()
		case <-ctx.Done():
		}
	}()

	routerFactory.NewWithContext(ctx).Run(serviceConfig)

	// cancel()
}
