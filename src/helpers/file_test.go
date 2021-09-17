package helpers

import (
	"console-application/src/model"
	"encoding/csv"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "/..")
)

func TestReadJsonFile(t *testing.T) {
	f, err := os.Open(Root + "/example_build/email_template.json")
	assert.NoError(t, err)
	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	var result map[string]string
	err = json.Unmarshal(byteValue, &result)
	assert.NoError(t, err)
	assert.NotEqual(t, len(result), 0)

}

func TestReadCsvFile(t *testing.T) {
	f, err := os.Open(Root + "/example_build/customers.csv")
	assert.NoError(t, err)
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	assert.NoError(t, err)
	assert.NotEqual(t, len(records), 0)
}

func TestWriteJsonFile(t *testing.T) {
	dataTest := model.Customer{}
	err := WriteJsonFile(Root+"/example_build/output_emails/emails.json", dataTest)
	assert.NoError(t, err)
}

func TestWriteCsvFile(t *testing.T) {
	var listCustomerError [][]string
	err := WriteCsvFile(Root+"/example_build/errors.csv", listCustomerError)
	assert.NoError(t, err)
}
