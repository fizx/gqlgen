package models

import "github.com/fizx/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
