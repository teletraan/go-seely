package seely

import (
	"context"
	"fmt"
	"sync"

	"github.com/syfun/go-graphql"
	"github.com/teletraan/httpx/auth"
)

// JWTTokenSource reuse jwt token before token expiry.
type JWTTokenSource struct {
	account       string
	password      string
	graphqlClient *graphql.Client
	token         *auth.JWTToken
	mu            sync.Mutex
}

// Token implement TokenSource.
func (ts *JWTTokenSource) Token() (auth.Token, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	if ts.token.Valid() {
		return ts.token, nil
	}

	c := ts.graphqlClient.Copy(nil)

	resp, err := c.Do(
		context.Background(),
		fmt.Sprintf(loginQuery, "token"),
		"login",
		graphql.JSON{"account": ts.account, "password": ts.password},
	)
	if err != nil {
		return nil, fmt.Errorf("get token error: %w", err)
	}

	var authInfo AuthInfo
	if err := resp.Guess("login", &authInfo); err != nil {
		return nil, fmt.Errorf("get token error: %w", err)
	}

	t, err := auth.NewJWTToken(*authInfo.Token)
	if err != nil {
		return nil, fmt.Errorf("get token error: %w", err)
	}
	ts.token = t
	return t, nil
}
