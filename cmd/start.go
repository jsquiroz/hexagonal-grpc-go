package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"

	"github.com/jsquiroz/hexagonal-grpc-go/pkg/application/role"
	"github.com/jsquiroz/hexagonal-grpc-go/pkg/config"
	grpc_role "github.com/jsquiroz/hexagonal-grpc-go/pkg/infrastructure/delivery/grpc/proto/role"
	"github.com/jsquiroz/hexagonal-grpc-go/pkg/infrastructure/delivery/handler"
	"github.com/jsquiroz/hexagonal-grpc-go/pkg/infrastructure/storage/psql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var cfgFile string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start permission microservice",
	Run: func(cmd *cobra.Command, args []string) {
		runserver()
	},
}

func init() {
	startCmd.Flags().StringVar(&cfgFile, "config", "config.yml", "Configuration file path")
	startCmd.MarkPersistentFlagRequired("config")
	viper.BindPFlag("config", startCmd.Flags().Lookup("config"))

	rootCmd.AddCommand(startCmd)
}

func runserver() {
	cnf := config.LoadVariables()

	creds, err := credentials.NewServerTLSFromFile(cnf.CertPem, cnf.CertKey)
	if err != nil {
		log.Fatalf("could not read certificates: %v", err)
	}

	srv := grpc.NewServer(grpc.Creds(creds))

	ss := new(psql.Storage)

	addserv := role.NewService(ss)
	handler.NewRoleServerGrpc(srv, addserv)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cnf.GRPCPort))
	if err != nil {
		log.Fatalf("could not listen to %d: %v", cnf.GRPCPort, err)
	}

	log.WithFields(log.Fields{
		"gRPC Port": cnf.GRPCPort,
		"HTTP Port": cnf.HTTPPort,
	}).Info("Server is running")

	go createRestServer(cnf.GRPCPort, cnf.HTTPPort, cnf.CertPem)

	log.Fatal(srv.Serve(l))

}

func createRestServer(gRPCPort, httpPort uint, cakeyFile string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cacert, _ := ioutil.ReadFile(cakeyFile)
	certPool := x509.NewCertPool()

	ok := certPool.AppendCertsFromPEM([]byte(cacert))
	if !ok {
		log.Fatal("Bad certificates")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName: fmt.Sprintf("localhost:%d", gRPCPort),
		RootCAs:    certPool,
	})

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	err := grpc_role.RegisterRoleServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", gRPCPort), opts)
	if err != nil {
		log.Fatalf("Canno start server %v", err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}
