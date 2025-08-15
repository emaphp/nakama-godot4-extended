package storage

import (
	"context"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	STORAGE_COLLECTION_GLOBAL string = "global_data"
	STORAGE_KEY_NAMES         string = "names"
)

type GlobalDataNames struct {
	Names []string `json:"names"`
}

func ReadNames(ctx context.Context, nk runtime.NakamaModule) ([]string, error) {
	read := &runtime.StorageRead{
		Collection: STORAGE_COLLECTION_GLOBAL,
		Key:        STORAGE_KEY_NAMES,
	}

	records, err := nk.StorageRead(ctx, []*runtime.StorageRead{read})
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	data := []byte(records[0].Value)
	var names GlobalDataNames
	if err := json.Unmarshal(data, &names); err != nil {
		return nil, runtime.NewError("error unmarshaling data", 13)
	}

	return names.Names, nil
}

func WriteNames(ctx context.Context, nk runtime.NakamaModule, names []string) error {
	data := &GlobalDataNames{
		Names: names,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return runtime.NewError("error marshaling data", 13)
	}

	write := &runtime.StorageWrite{
		Collection:      STORAGE_COLLECTION_GLOBAL,
		Key:             STORAGE_KEY_NAMES,
		Value:           string(bytes),
		PermissionRead:  2,
		PermissionWrite: 0,
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{write})
	if err != nil {
		return runtime.NewError("error saving data", 13)
	}

	return nil
}
