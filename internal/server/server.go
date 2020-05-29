package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	schema "github.com/TylerGrey/tenants/internal/graphql"
	"github.com/TylerGrey/tenants/internal/mysql"
	"github.com/TylerGrey/tenants/internal/mysql/repo"
	"github.com/TylerGrey/tenants/internal/server/handler"
	"github.com/TylerGrey/tenants/internal/server/resolver"
	"github.com/graph-gophers/graphql-go"
	"github.com/rs/cors"
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

	// Handler 설정
	h := &handler.GraphQL{
		Schema: graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{
			ReviewRepo: reviewRepo,
			BldgRepo:   bldgRepo,
		}),
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(os.Getenv("PROJECT_ROOT_PATH")+"/web")))
	mux.Handle("/graphiql", &handler.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h)

	// CORS 설정
	op := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Access-Control-Allow-Headers", "DeviceInfo", "Authorization", "X-Requested-With"},
		AllowCredentials: false,
	}
	handler := cors.New(op).Handler(mux)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errc <- fmt.Errorf("%s", <-c)
	}()

	log.Printf("Listening for requests on %s", *s.Addr)

	go func() {
		errc <- http.ListenAndServe(*s.Addr, handler)
	}()

	return <-errc
}
