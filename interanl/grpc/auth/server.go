package auth

import ssov1 "github.com/DragonSlayer1995/protos/gen/go/sso"

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}
