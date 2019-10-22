package main

import (
	"context"
	"testing"

	"github.com/lucsky/cuid"
	"github.com/stretchr/testify/assert"
	{{name}}pb "github.com/upnid/protobuf/gen/go/upnid/{{name}}/v1"
)

var s Server = Server{
	clients: Clients{
		Db: initDatabase(context.Background(), "mongodb://localhost:27017").Database("upnid"),
	},
}

func createBasicEntry(old_value string, new_value string, user_id string, entity_id string) (*{{name}}pb.CreateEntryResponse, error) {
	req := &{{name}}pb.CreateEntryRequest{
		EntityId: entity_id,
		OldValue: old_value,
		NewValue: new_value,
		UserId:   user_id,
	}

	res, err := s.CreateEntry(context.Background(), req)
	return res, err
}

func TestCreateEntryChanges(t *testing.T) {
	table := []struct {
		old     string
		new     string
		changes []*{{name}}pb.Change
	}{
		{`{"foo": "bar"}`, `{"foo": "zero"}`, []*{{name}}pb.Change{
			&{{name}}pb.Change{
				Path: []string{"foo"},
				Type: "update",
				From: "bar",
				To:   "zero",
			},
		}},
		{`{"foo": "bar"}`, `{}`, []*{{name}}pb.Change{
			&{{name}}pb.Change{
				Path: []string{"foo"},
				Type: "delete",
				From: "bar",
				To:   "<nil>",
			},
		}},
		{`{"foo": "bar"}`, `{"foo": "bar", "baz": "bat"}`, []*{{name}}pb.Change{
			&{{name}}pb.Change{
				Path: []string{"baz"},
				Type: "create",
				From: "<nil>",
				To:   "bat",
			},
		}},
	}

	for _, test_case := range table {
		res, err := createBasicEntry(test_case.old, test_case.new, "us_cs2tqoe1c00b4dr035x72o51h", "ac_ck1tboo0g0000dr095x7ao5nh")

		assert.Nil(t, err)
		assert.Equal(t, test_case.changes, res.Entry.Changes)
	}
}

func TestGetEntity(t *testing.T) {
	user := "us_" + cuid.New()
	entity := "ac_" + cuid.New()

	entries := []struct {
		old    string
		new    string
		user   string
		entity string
	}{
		{`{"foo": "bar"}`, `{"foo": "zero"}`, user, entity},
		{`{"foo": "bar"}`, `{"foo": "bat"}`, user, entity},
		{`{"foo": "bar"}`, `{"foo": "1"}`, user, entity},
		{`{"foo": "bar"}`, `{}`, user, entity},
		{`{"foo": "bar"}`, `{"foo": "123"}`, user, entity},
		{`{"foo": "bar"}`, `{"foo": "123"}`, user, "ac_" + cuid.New()},
		{`{"foo": "bar"}`, `{"foo": "123"}`, user, "ac_" + cuid.New()},
		{`{"foo": "bar"}`, `{"foo": "123"}`, user, "ac_" + cuid.New()},
	}

	for _, entry := range entries {
		_, err := createBasicEntry(entry.old, entry.new, entry.user, entry.entity)

		assert.Nil(t, err)
	}

	res, err := s.GetEntity(context.Background(), &{{name}}pb.GetEntityRequest{
		Entity: entity,
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Entries), 5)
}
