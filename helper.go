package kuaidailigo

import (
	"encoding/json"
	"errors"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	return tt.MarshalJSON()
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// TODO(https://go.dev/issue/47353): Properly unescape a JSON string.
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Time.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	var err error
	// 2021-10-23 00:55:09
	tt, err := time.Parse(time.DateTime, string(data))
	if err != nil {
		return err
	}
	*t = Time(tt)
	return nil
}

func toParams(p any) map[string]string {
	params := make(map[string]string)
	d, _ := json.Marshal(p)
	json.Unmarshal(d, &params)
	return params
}
