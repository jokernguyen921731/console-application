package service

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "/..")
)

func TestGetEmail(t *testing.T) {
	email, err := GetEmail(Root + "/example_build/email_template.json")
	assert.NoError(t, err)
	assert.NotEqual(t, email.From, nil)
	assert.Equal(t, email.To, "")
	assert.NotEqual(t, email.Subject, nil)
	assert.NotEqual(t, email.MimeType, nil)
	assert.NotEqual(t, email.Body, nil)
}

func TestGetListCustomer(t *testing.T) {
	customers, err := GetListCustomer(Root+"/example_build/customers.csv", Root+"/example_build/error.csv")
	assert.NoError(t, err)
	assert.NotEqual(t, len(customers), 0)
}

func TestMerge(t *testing.T) {
	email, err := GetEmail(Root + "/example_build/email_template.json")
	assert.NoError(t, err)
	customers, err := GetListCustomer(Root+"/example_build/customers.csv", Root+"/example_build/error.csv")
	assert.NoError(t, err)
	results, err := Merge(email, customers)
	assert.NoError(t, err)
	assert.NotEqual(t, len(results), 0)
}