/*
 *	Merging service combines customer information to email template
 *  and save result to file
 *  create_at: 16 Sep 2021 by Weifen
 */

package service

import (
	"console-application/src/helpers"
	"console-application/src/model"
	"fmt"
	"strings"
	"time"
)

// Replace change every labels of json_email_template into data using customer information
func Replace(body string, res model.Customer, cus model.Customer) string {
	strT := fmt.Sprintf("{{%s}}", res.Title)
	strF := fmt.Sprintf("{{%s}}", res.FirstName)
	strL := fmt.Sprintf("{{%s}}", res.LastName)
	strNow := fmt.Sprintf("{{%s}}", "TODAY")
	dt := time.Now()
	month := dt.Month().String()[:3]
	strToday := fmt.Sprintf("%d %s %d", dt.Day(), month, dt.Year())

	replacer := strings.NewReplacer(strT, cus.Title, strF, cus.FirstName, strL, cus.LastName, strNow, strToday)
	output := replacer.Replace(body)
	return output
}

//GetEmail get data from the json_email_template file into a list EmailTemplate
func GetEmail(fileDir string) (model.EmailTemplate, error) {
	var emailTemplate model.EmailTemplate
	emailData, err := helpers.ReadJsonFile(fileDir)
	if err != nil {
		return emailTemplate, err
	}
	emailTemplate = model.EmailTemplate{
		From:     emailData["from"],
		Subject:  emailData["subject"],
		MimeType: emailData["mimeType"],
		Body:     emailData["body"],
	}
	return emailTemplate, nil
}

//GetListCustomer get data from the csv_customer file into a list Customer
func GetListCustomer(fileDir string, csvErrorFile string) ([]model.Customer, error) {
	var listCustomer []model.Customer
	var listCustomerError [][]string
	customers, err := helpers.ReadCsvFile(fileDir)
	if err != nil {
		return nil, err
	}
	for i, cus := range customers {
		if i == 0 {
			listCustomerError = append(listCustomerError, cus)
		}

		if cus[3] == "" {
			listCustomerError = append(listCustomerError, cus)
		} else {
			customer := model.Customer{
				Title:     cus[0],
				FirstName: cus[1],
				LastName:  cus[2],
				EmailAddr: cus[3],
			}
			listCustomer = append(listCustomer, customer)
		}
	}

	if len(listCustomerError) > 1 {
		// Write error.csv here
		err = helpers.WriteCsvFile(csvErrorFile, listCustomerError)
		if err != nil {
			return nil, err
		}
	}
	return listCustomer, nil
}

//Merge merges data from the customers file to a list json_email_template
func Merge(email model.EmailTemplate, customers []model.Customer) ([]model.EmailTemplate, error) {
	var res []model.EmailTemplate
	for i := 1; i < len(customers); i++ {
		emailResult := model.EmailTemplate{
			From:     email.From,
			To:       customers[i].EmailAddr,
			Subject:  email.Subject,
			MimeType: email.MimeType,
			Body:     Replace(email.Body, customers[0], customers[i]),
		}
		res = append(res, emailResult)
	}
	return res, nil
}

//MergeEmailHandler gets input files include: json_email_template and csv_customer files,
//then, merges data from customers file to email template and writes result_data to output files
func MergeEmailHandler(jsonInput string, csvInput string, dirOutput string, csvError string) error {
	email, err := GetEmail(jsonInput)
	if err != nil {
		return err
	}

	customers, err := GetListCustomer(csvInput, csvError)
	if err != nil {
		return err
	}

	resultEmails, err := Merge(email, customers)
	if err != nil {
		return err
	}

	if err = helpers.WriteJsonFile(dirOutput+"emails.json", resultEmails); err != nil {
		return err
	}
	return nil
}
