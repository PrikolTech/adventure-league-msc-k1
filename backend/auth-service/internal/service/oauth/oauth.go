package oauth

import (
	"auth-service/internal/entity"
	"auth-service/internal/net"
	"context"
	"strings"

	"github.com/go-oauth2/oauth2/v4"
)

type oauth struct {
	manager oauth2.Manager
	user    net.User
}

func NewOAuth(user net.User, opts ManagerOptions) (*oauth, error) {
	manager, err := NewManager(opts)
	if err != nil {
		return nil, err
	}
	return &oauth{manager, user}, nil
}

func (o *oauth) GenerateToken(clientID string, email string, password string) (*entity.Token, error) {
	id, roles, err := o.user.Authenticate(email, password)
	if err != nil {
		return nil, err
	}

	var builder strings.Builder
	for _, role := range roles {
		builder.WriteString(role.Title + ",")
	}

	scope := builder.String()
	if scope != "" {
		scope = scope[:len(scope)-1]
	}

	gt := oauth2.PasswordCredentials
	tgr := &oauth2.TokenGenerateRequest{
		ClientID: clientID,
		UserID:   id,
		Scope:    scope,
	}
	ti, err := o.manager.GenerateAccessToken(context.Background(), gt, tgr)
	if err != nil {
		return nil, err
	}

	token := &entity.Token{ti.GetAccess(), ti.GetAccessExpiresIn(), ti.GetRefresh(), ti.GetRefreshExpiresIn(), ti.GetUserID()}
	return token, nil
}

func (o *oauth) RefreshToken(clientID string, refreshToken string) (*entity.Token, error) {
	tgr := &oauth2.TokenGenerateRequest{
		ClientID: clientID,
		Refresh:  refreshToken,
	}

	ti, err := o.manager.RefreshAccessToken(context.Background(), tgr)
	if err != nil {
		return nil, err
	}

	token := &entity.Token{ti.GetAccess(), ti.GetAccessExpiresIn(), ti.GetRefresh(), ti.GetRefreshExpiresIn(), ti.GetUserID()}
	return token, nil
}

func (o *oauth) RemoveToken(refreshToken string) error {
	return o.manager.RemoveRefreshToken(context.Background(), refreshToken)
}
