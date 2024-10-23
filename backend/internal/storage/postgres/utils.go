package postgres

import (
	"github.com/google/uuid"
)

func uuidToString(uuid uuid.UUID) string {
	str := "uuid('" + uuid.String() + "')"
	return str
}

func uuidsToString(uuids []uuid.UUID) string {
	if uuids == nil || len(uuids) == 0 {
		return ""
	}

	str := ""
	str += uuidToString(uuids[0])
	for i := 1; i < len(uuids); i++ {
		str += ", " + uuidToString(uuids[i])
	}
	return str
}
