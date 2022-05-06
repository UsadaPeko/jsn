package jsn_test

import (
	"github.com/UsadaPeko/jsn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	jsonObject := jsn.Init()

	jsonObject.Set("key", "val")

	val, ok := jsonObject.StringVal("key")
	assert.True(t, ok)
	assert.Equal(t, "val", val)
}

func TestNew(t *testing.T) {
	jsonObject, err := jsn.New(`{"key": "val"}`)
	assert.NoError(t, err)

	val, ok := jsonObject.StringVal("key")
	assert.True(t, ok)
	assert.Equal(t, "val", val)
}

func TestNewEmptyJSON(t *testing.T) {
	jsonObject, err := jsn.New(`{}`)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonObject)
}

func TestInvalidFormat(t *testing.T) {
	jsonObject, err := jsn.New(`{InvalidFormat`)
	assert.Error(t, err)
	assert.Empty(t, jsonObject)
}

func TestIntVal(t *testing.T) {
	jsonObject, err := jsn.New(`{"key1": 10}`)
	jsonObject.Set("key2", 20)
	assert.NoError(t, err)

	val, ok := jsonObject.IntVal("key1")
	assert.True(t, ok)
	assert.Equal(t, 10, val)

	val, ok = jsonObject.IntVal("key2")
	assert.True(t, ok)
	assert.Equal(t, 20, val)

	val, ok = jsonObject.IntVal("not-exist-key")
	assert.False(t, ok)
	assert.Empty(t, val)
}

func TestStringVal(t *testing.T) {
	jsonObject, err := jsn.New(`{"key": "val"}`)
	assert.NoError(t, err)

	val, ok := jsonObject.StringVal("key")
	assert.True(t, ok)
	assert.Equal(t, "val", val)

	val, ok = jsonObject.StringVal("not-exist-key")
	assert.False(t, ok)
	assert.Empty(t, val)
}

func TestGetVal_TypeMissMatch(t *testing.T) {
	jsonObject, err := jsn.New(`{"integer": 10, "string": "val"}`)
	assert.NoError(t, err)

	sv, ok := jsonObject.StringVal("integer")
	assert.False(t, ok)
	assert.Empty(t, sv)

	iv, ok := jsonObject.IntVal("string")
	assert.False(t, ok)
	assert.Empty(t, iv)
}
