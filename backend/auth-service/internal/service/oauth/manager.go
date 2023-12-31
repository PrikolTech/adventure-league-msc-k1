package oauth

import (
	"time"

	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
)

type ManagerOptions struct {
	ClientID     string
	ClientDomain string
	Key          []byte
}

func NewManager(opts ManagerOptions) (*manage.Manager, error) {
	generator, err := NewJWTAccessGenerate(opts.Key)
	if err != nil {
		return nil, err
	}

	clientStore := store.NewClientStore()
	clientStore.Set(opts.ClientID, &models.Client{
		ID:     opts.ClientID,
		Domain: opts.ClientDomain,
		Public: true,
	})

	tokenStore, err := store.NewMemoryTokenStore()
	if err != nil {
		return nil, err
	}

	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)
	manager.MapAccessGenerate(generator)

	passwordTokenConfig := &manage.Config{AccessTokenExp: time.Minute * 15, RefreshTokenExp: time.Hour * 24 * 7, IsGenerateRefresh: true}

	refreshTokenConfig := &manage.RefreshingConfig{AccessTokenExp: time.Minute * 15, RefreshTokenExp: time.Hour * 24 * 7, IsGenerateRefresh: true, IsRemoveAccess: true, IsRemoveRefreshing: true}

	manager.SetPasswordTokenCfg(passwordTokenConfig)
	manager.SetRefreshTokenCfg(refreshTokenConfig)

	return manager, nil
}
