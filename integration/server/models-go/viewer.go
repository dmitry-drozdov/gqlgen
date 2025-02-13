package models

import "github.com/dmitry-drozdov/gqlgen/integration/server/remote_api"

type Viewer struct {
	User *remote_api.User
}
