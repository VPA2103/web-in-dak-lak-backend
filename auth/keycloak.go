package auth

import (
	"context"

	"github.com/coreos/go-oidc"
)

var Verifier *oidc.IDTokenVerifier

func InitKeycloak() error {
	provider, err := oidc.NewProvider(context.Background(), "http://localhost:8180/realms/golang-backend")
	if err != nil {
		return err
	}

	Verifier = provider.Verifier(&oidc.Config{ClientID: "golang-backend-auth"})
	return nil
}
