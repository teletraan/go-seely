package seely

import (
	"context"
	"errors"
	"fmt"

	"github.com/syfun/go-graphql"
)

// MemberService represents service about member.
type MemberService service

// ErrMemberNotFound represent wanted member not found.
var ErrMemberNotFound = errors.New("member not found")

// Search members.
func (ms *MemberService) Search(ctx context.Context, selection string, page *Page, filter *MemberFilter) (int, []*Member, error) {
	resp, err := ms.client.graphqlClient.Do(
		ctx,
		fmt.Sprintf(searchMemberQuery, selection),
		"members",
		graphql.JSON{
			"page":   page,
			"filter": filter,
		},
	)
	if err != nil {
		return 0, nil, fmt.Errorf("search member error: %w", err)
	}

	var data struct {
		Count int       `json:"count"`
		Data  []*Member `json:"data"`
	}
	if err := resp.Guess("members", &data); err != nil {
		return 0, nil, fmt.Errorf("search member error: %w", err)
	}
	return data.Count, data.Data, nil
}

// Create seely member with GSMARegistrant.
func (ms *MemberService) Create(ctx context.Context, input *CreateMemberInput) error {
	if _, err := ms.client.graphqlClient.Do(
		ctx,
		createMemberQuery,
		"createMember",
		graphql.JSON{"member": input},
	); err != nil {
		return fmt.Errorf("create member(email: %v) error: %w", input.Email, err)
	}
	return nil
}

// Update seely member with GSMARegistrant.
func (ms *MemberService) Update(ctx context.Context, IDs []string, input *UpdateMemberInput) error {
	if _, err := ms.client.graphqlClient.Do(
		ctx,
		updateMemberQuery,
		"updateMember",
		graphql.JSON{"ids": IDs, "member": input},
	); err != nil {
		return fmt.Errorf("update member(ids: %v) error: %w", IDs, err)
	}
	return nil
}

// Delete member.
func (ms *MemberService) Delete(ctx context.Context, IDs []string) error {
	if _, err := ms.client.graphqlClient.Do(
		ctx,
		deleteMemberQuery,
		"deleteMember",
		graphql.JSON{"ids": IDs},
	); err != nil {
		return fmt.Errorf("delete member(ids: %v) error: %w", IDs, err)
	}
	return nil
}
