package seely

import (
	"context"
	"reflect"
	"testing"

	"github.com/openlyinc/pointy"
	"github.com/syfun/go-graphql"
)

func TestMemberRecordService_Search(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	want := graphql.JSON{
		"count": 1,
		"data": []*MemberRecord{
			{
				ID:         "1",
				Start:      1578722460000,
				Device:     &Device{ID: "2", UUID: "09154c62-dc4d-4a9f-ac48-18eee4ef70da", Name: "android"},
				AccessType: "NORMAL",
				Member:     &Member{ID: "20", Name: "Katy Perry", SerialNumber: pointy.String("test")},
			},
		},
	}
	mux.want("memberRecords", want)
	selection := `
id
start
device {
  id
  uuid
  name
}
accessType
member {
  id
  name
  serialNumber
}
`
	count, records, err := client.MemberRecord.Search(context.Background(), selection, &Page{NoPage: pointy.Bool(true)}, nil)
	if err != nil {
		t.Errorf("MemberRecord.Search returned error : %v", err)
	}
	got := graphql.JSON{"count": count, "data": records}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MemberRecord.Search returned: %+v, want %+v", got, want)
	}
}
