/*
 *	Providing common function for file handler
 *  create_at: 16 Sep 2021 by Weifen
 */

package helpers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
)

//ReadJsonFile reads a json file by the absolute directory following fileDir
func ReadJsonFile(fileDir string) (map[string]string, error) {
	var result map[string]string
	f, err := os.Open(fileDir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byteValue, &result); err != nil {
		return nil, err
	}
	return result, nil
}

//ReadCsvFile reads a csv file by the absolute directory following fileDir
func ReadCsvFile(fileDir string) ([][]string, error) {
	f, err := os.Open(fileDir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	return csvReader, nil
}

//WriteJsonFile write a json file by the absolute directory as file_name.json
func WriteJsonFile(fileDir string, data interface{}) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(data); err != nil {
		return err
	}

	file, err := os.OpenFile(fileDir, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	if _, err = file.Write(buffer.Bytes()); err != nil {
		return err
	}
	return nil
}

//WriteCsvFile write a csv file by the absolute directory as file_name.csv
func WriteCsvFile(fileDir string, data [][]string) error {
	csvFile, err := os.Create(fileDir)
	if err != nil {
		return err
	}
	csvWriter := csv.NewWriter(csvFile)
	for _, row := range data {
		if err = csvWriter.Write(row); err != nil {
			return err
		}
	}
	csvWriter.Flush()
	csvFile.Close()
	return nil
}
