package dependencies

import (
	"database/sql"
	"fmt"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/cmd/handlers"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/internal/auth"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/central_server/internal/directory"
	"os"
)

type Config struct {
	Host string
	Port string
}

func InitialConfig() (Config, error) {
	return Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}, nil
}
func SetDependencies(db *sql.DB) *handlers.ApiRest {
	var repoAuth auth.RepositoryAuth
	var repoDir directory.DirRepository
	switch {
	case db != nil:
		fmt.Println("using db connexion for repositories...")
	default:
		fmt.Println("default dependencies...creating peersMap & indexTable as map structure")
		peersMap := make(map[string]auth.Peer)
		indexTable := make(map[string]string)
		repoAuth = auth.NewDefaultRepo(peersMap)
		repoDir = directory.NewDefaultRepo(&indexTable)

	}
	authService := auth.NewServiceClient(repoAuth)
	directoryService := directory.NewServiceClient(repoDir)
	handler := handlers.NewApiRest(authService, directoryService)
	fmt.Println("settings done...")
	return handler

}
