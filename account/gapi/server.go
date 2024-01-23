package gapi

import (
	db "pharmago/account/db/sqlc"
	"pharmago/utils"
)

type ServerGRPC struct {
	config utils.Config
	store  *db.StoreAccout
}

func NewServer(config utils.Config, store *db.StoreAccout) (*ServerGRPC, error) {
	server := &ServerGRPC{
		config: config,
		store:  store,
	}

	return server, nil
}
