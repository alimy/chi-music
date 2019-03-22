package cmd

import (
	"github.com/alimy/chi-music/cmd"
	"github.com/alimy/chi-music/models"
	"github.com/spf13/cobra"
	"github.com/unisx/logus"
	"net/http"
	"time"
)

const (
	listenAddrDefault   = "127.0.0.1:8013" // default listen address
	certFilePathDefault = "cert.pem"       // certificate file default path
	keyFilePathDefault  = "key.pem"        // key file used in https server default path
)

var (
	address     string
	certFile    string
	keyFile     string
	enableHttps bool
	inDebug     bool
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start to mirMusic service",
		Long:  "this cmd will start a https server to provide ginMusic service",
		Run:   serveRun,
	}

	// Parse flags for serveCmd
	serveCmd.Flags().StringVarP(&address, "addr", "a", listenAddrDefault, "service listen address")
	serveCmd.Flags().StringVarP(&certFile, "cert", "c", certFilePathDefault, "certificate path used in https connect")
	serveCmd.Flags().StringVarP(&keyFile, "key", "k", keyFilePathDefault, "key path used in https connect")
	serveCmd.Flags().BoolVarP(&enableHttps, "https", "s", false, "whether use https serve connect")
	serveCmd.Flags().BoolVarP(&inDebug, "debug", "d", false, "whether in debug mode")

	// Register serveCmd as sub-command
	cmd.Register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	setup()

	// Setup http.Server
	server := &http.Server{
		Handler: newChi(),
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start http.Server
	if enableHttps {
		logus.Info("start listen and serve", logus.String("address", address))
		server.ListenAndServeTLS(certFile, keyFile)
	} else {
		logus.Info("listen and serve",
			logus.String("address", address),
			logus.Bool("enableHttps", enableHttps))
		server.ListenAndServe()
	}
}

func setup() {
	if !inDebug {
		logus.InProduction()
	}

	// initial models with MemoryProfile
	if err := models.Register(models.MemoryProfile); err != nil {
		panic(err)
	}
}
