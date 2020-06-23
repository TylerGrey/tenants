package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TylerGrey/tenants/internal/graphql/generated"
	"github.com/TylerGrey/tenants/internal/graphql/resolver"
	"github.com/TylerGrey/tenants/internal/mysql"
	"github.com/TylerGrey/tenants/internal/mysql/repo"
)

// Server API Server
type Server struct {
	Addr *string
}

// Start run server
func (s Server) Start() error {
	// mysql 설정
	mysqlMaster, mysqlReplica, err := mysql.IntializeDatabase(os.Getenv("MYSQL_DB_NAME"))
	if err != nil {
		log.Println("db initialize error", err.Error())
		panic(err)
	}
	reviewRepo := repo.NewReviewRepository(mysqlMaster, mysqlReplica)
	bldgRepo := repo.NewBldgRepository(mysqlMaster, mysqlReplica)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		ReviewRepo: reviewRepo,
		BldgRepo:   bldgRepo,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errc <- fmt.Errorf("signal: %s", <-c)
	}()

	go func() {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", *s.Addr)
		errc <- http.ListenAndServe(*s.Addr, nil)
	}()

	return <-errc
}
