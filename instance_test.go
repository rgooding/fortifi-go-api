package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstance(t *testing.T) {
	instance, err := NewInstance("ORG:FORT:2047:anmbz", "TDd0RldhVTY0-GT-WGpnMmZiWWZy", "cGRBdzBYcUU5WFZmTDFTQk52Z3BlNUZKUnQzaktG")
	assert.NoError(t, err)

	instance.SetAPIHost("http://api.fortifi.me:9090")

	err = instance.InitAPI()
	assert.NoError(t, err)

	assert.NotEmpty(t, instance.authToken)
}
