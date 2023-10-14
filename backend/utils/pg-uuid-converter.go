package utils

import (
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertToPGUUID(googleUUIDStr string) (pgtype.UUID, error) {
	googleUUID, err := uuid.Parse(googleUUIDStr)
	if err != nil {
		return pgtype.UUID{}, err
	}

	var pgUUID pgtype.UUID
	pgUUID.Bytes, err = parseUUID(googleUUID.String())
	if err != nil {
		return pgtype.UUID{}, err
	}
	pgUUID.Valid = true

	return pgUUID, nil
}

func parseUUID(src string) ([16]byte, error) {
	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		return [16]byte{}, fmt.Errorf("cannot parse UUID %v", src)
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return [16]byte{}, err
	}

	var dst [16]byte
	copy(dst[:], buf)
	return dst, err
}
