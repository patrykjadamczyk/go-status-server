package server

import (
	"github.com/patrykjadamczyk/go-status-server/common/memdb"
	"github.com/patrykjadamczyk/go-status-server/config"
)

type MainFunction func(appConfig config.Configuration, appDB memdb.DB)

type Server struct {
	AppConfiguration config.Configuration
	AppDatabase      memdb.DB
	ServerMainFunc   MainFunction
}

func (server *Server) Run() {
	server.ServerMainFunc(server.AppConfiguration, server.AppDatabase)
}

func MakeServer(appConfig config.Configuration, appDB memdb.DB, serverFunction MainFunction) Server {
	server := Server{
		AppConfiguration: appConfig,
		AppDatabase:      appDB,
		ServerMainFunc:   serverFunction,
	}
	return server
}
