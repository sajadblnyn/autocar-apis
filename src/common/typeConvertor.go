package common

import "encoding/json"

func ConvertType[T any](tr any) (*T, error) {
	b, er := json.Marshal(tr)
	if er != nil {
		return nil, er
	}

	var res T
	er = json.Unmarshal(b, &res)
	if er != nil {
		return nil, er
	}

	return &res, nil

}
