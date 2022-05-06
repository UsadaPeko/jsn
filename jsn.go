package jsn

import "encoding/json"

type JSON struct {
	m map[string]interface{}
}

func Init() *JSON {
	return &JSON{
		m: map[string]interface{}{},
	}
}

func New(s string) (*JSON, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return nil, err
	}
	return &JSON{m}, nil
}

func (j *JSON) StringVal(key string) (string, bool) {
	v, ok := j.m[key]
	if !ok {
		return "", false
	}

	stringVal, ok := v.(string)
	if !ok {
		return "", false
	}

	return stringVal, true
}

func (j *JSON) ListOfStringVal(key string) ([]string, bool) {
	v, ok := j.m[key]
	if !ok {
		return []string{}, false
	}

	stringVal, ok := v.([]string)
	if !ok {
		return []string{}, false
	}

	return stringVal, true
}

func (j *JSON) IntVal(key string) (int, bool) {
	v, ok := j.m[key]
	if !ok {
		return 0, false
	}

	intVal, ok := v.(int)
	if ok {
		return intVal, true
	}

	float64Val, ok := v.(float64)
	if ok {
		return int(float64Val), true
	}

	return 0, false
}

func (j *JSON) Set(key string, val interface{}) {
	j.m[key] = val
}

func (j *JSON) String() string {
	bytes, err := json.Marshal(j.m)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
