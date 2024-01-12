package usrv

import "encoding/json"

func AdaptJSON[I any](raw json.RawMessage) (*I, error) {
	var in I
	if err := json.Unmarshal(raw, &in); err != nil {
		return nil, err
	}
	return &in, nil
}
