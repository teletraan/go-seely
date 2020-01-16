package seely

import (
	"github.com/syfun/go-graphql"
	"github.com/teletraan/httpx/auth"
)

type service struct {
	client *Client
}

// Client represents seely client.
type Client struct {
	graphqlClient *graphql.Client

	common       service
	Member       *MemberService
	File         *FileService
	MemberRecord *MemberRecordService
}

// New create a seely client with url, account and password.
func New(url, account, password string) *Client {
	tokenSource := &JWTTokenSource{account: account, password: password}
	transport := auth.TokenAuthTransport{
		Source: tokenSource,
	}
	graphqlClient := graphql.New(url, transport.Client())
	tokenSource.graphqlClient = graphqlClient

	client := &Client{
		graphqlClient: graphqlClient,
	}
	client.common.client = client
	client.Member = (*MemberService)(&client.common)
	client.File = (*FileService)(&client.common)
	client.MemberRecord = (*MemberRecordService)(&client.common)

	return client
}
