package seely

import (
	"context"
	"fmt"

	"github.com/syfun/go-graphql"
)

// MemberRecordService represents service about member.
type MemberRecordService service

func (mrs *MemberRecordService) Search(ctx context.Context, selection string, page *Page, filter *MemberRecordFilter) (int, []*MemberRecord, error) {
	resp, err := mrs.client.graphqlClient.Do(
		ctx,
		fmt.Sprintf(searchMemberRecordQuery, selection),
		"memberRecords",
		graphql.JSON{
			"page":   page,
			"filter": filter,
		},
	)
	if err != nil {
		return 0, nil, fmt.Errorf("search member record error: %w", err)
	}

	var data struct {
		Count int             `json:"count"`
		Data  []*MemberRecord `json:"data"`
	}
	if err := resp.Guess("memberRecords", &data); err != nil {
		return 0, nil, fmt.Errorf("search member record error: %w", err)
	}
	return data.Count, data.Data, nil
}
