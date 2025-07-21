package guid

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

type GUID struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

func MarshalGUID(guid GUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(fmt.Sprintf("%s:%d", guid.Type, guid.ID)))
	})
}

func UnmarshalGUID(v any) (GUID, error) {
	switch v := v.(type) {
	case string:
		s := strings.Split(v, ":")
		if len(s) != 2 {
			return GUID{}, fmt.Errorf("invalid GUID format: %s", v)
		}
		id, err := strconv.Atoi(s[1])
		if err != nil {
			return GUID{}, fmt.Errorf("invalid GUID format: %s", v)
		}
		return GUID{Type: s[0], ID: id}, nil
	default:
		return GUID{}, fmt.Errorf("%T is not a GUID", v)
	}
}

func ConvertStringToGUID(s string) (*GUID, error) {
	guid, err := UnmarshalGUID(s)
	if err != nil {
		return nil, err
	}
	return &guid, nil
}

func ConvertGUIDToInt(guid *GUID) (int, error) {
	return guid.ID, nil
}

func ConvertGUIDToIntForTodo(guid *GUID) (int, error) {
	return guid.ID, nil
}

func ConvertStringToGUIDForNeq(s string) (*GUID, error) {
	guid, err := UnmarshalGUID(s)
	if err != nil {
		return nil, err
	}
	return &guid, nil
}
