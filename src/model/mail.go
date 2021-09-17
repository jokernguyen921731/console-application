/*
 *	Modeling email and customer for saving to the database if needed
 *  create_at: 16 Sep 2021 by Weifen
 */

package model

type EmailTemplate struct {
	From     string
	To       string
	Subject  string
	MimeType string
	Body     string
}

type Customer struct {
	Title     string
	FirstName string
	LastName  string
	EmailAddr string
}
