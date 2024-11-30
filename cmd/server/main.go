package main

import (
	"fmt"
	"hypersonic/internal/infrastructure/api/http"
	"hypersonic/internal/infrastructure/datasource/filesystem"
	"hypersonic/internal/interface-adapter/handler/graphql/graph"
	"hypersonic/internal/usecase/search"
	"log/slog"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	err := run()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	option := getOption()
	deps := loadDependencies(option.DataDirPath)

	s := http.NewServer(
		fmt.Sprintf("%s:%s", option.Host, option.Port),
		deps.Usecase.GraphQL,
	)
	return s.Serve()
}

type option struct {
	Host        string
	Port        string
	DataDirPath string
}

func getOption() option {
	host := pflag.String("host", "127.0.0.1", "")
	port := pflag.StringP("port", "p", "8080", "")
	dataDirPath := pflag.StringP("data", "d", "/var/lib/hypersonic", "")
	pflag.Parse()

	return option{
		*host,
		*port,
		*dataDirPath,
	}
}

type dependencies struct {
	Usecase depsUsecase
}
type depsUsecase struct {
	GraphQL graph.Dependencies
}

func loadDependencies(dataDirPath string) dependencies {
	repository := filesystem.NewRepository(dataDirPath)
	return dependencies{
		Usecase: depsUsecase{
			graph.Dependencies{
				Search: search.New(repository),
			},
		},
	}
}
