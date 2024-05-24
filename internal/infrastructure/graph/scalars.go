package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"time"
)

func MarshalUUID(v uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(v.String())
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	id, err := uuid.Parse(v.(string))
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

func MarshalDateTime(v time.Time) graphql.Marshaler {
	return graphql.MarshalString(v.String())
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, v.(string))
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
