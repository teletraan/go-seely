package seely

import (
	"context"
	"reflect"
	"testing"

	"github.com/syfun/go-graphql"
)

func TestMemberService_Search(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	want := graphql.JSON{
		"count": 1,
		"data": []*Member{
			{
				ID:     "1",
				Name:   "Jack",
				Status: "NORMAL",
			},
		},
	}

	mux.want("members", want)

	selection := `
id
name
status
`
	count, members, err := client.Member.Search(context.Background(), selection, page, nil)
	if err != nil {
		t.Errorf("Member.Search returned error : %v", err)
	}
	got := graphql.JSON{"count": count, "data": members}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Member.Search returned records: %+v, want %+v", got, want)
	}
}
