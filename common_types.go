package gochimp3

import (
	"fmt"
	"strings"
)

// APIError is what the what the api returns on error
type APIError struct {
	Type            string `json:"type,omitempty"`
	Title           string `json:"title,omitempty"`
	Status          int    `json:"status,omitempty"`
	Detail          string `json:"detail,omitempty"`
	Instance        string `json:"instance,omitempty"`
	ReferenceNumber string `json:"ref_no,omitempty"`
	Errors          []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}

func (err *APIError) String() string {
	return fmt.Sprintf("%d : %s : %s : %s : %s", err.Status, err.Type, err.Title, err.Detail, err.Errors)
}

func (err *APIError) Error() string {
	return err.String()
}

// HasError checks if this call had an error
func (err *APIError) HasError() bool {
	return err.Type != ""
}

// QueryParams defines the different params
type QueryParams interface {
	Params() map[string]string
}

// ExtendedQueryParams includes a count and offset
type ExtendedQueryParams struct {
	BasicQueryParams

	Count  int
	Offset int
}

func (q *ExtendedQueryParams) Params() map[string]string {
	m := q.BasicQueryParams.Params()
	m["count"] = fmt.Sprintf("%d", q.Count)
	m["offset"] = fmt.Sprintf("%d", q.Offset)
	return m
}

// BasicQueryParams basic filter queries
type BasicQueryParams struct {
	Status              string
	SortField           string
	SortDirection       string
	Fields              []string
	ExcludeFields       []string
	SkipMergeValidation bool
}

func (q *BasicQueryParams) Params() map[string]string {
	return map[string]string{
		"status":                q.Status,
		"sort_field":            q.SortField,
		"sort_dir":              q.SortDirection,
		"fields":                strings.Join(q.Fields, ","),
		"exclude_fields":        strings.Join(q.ExcludeFields, ","),
		"skip_merge_validation": fmt.Sprintf("%t", q.SkipMergeValidation),
	}
}

type withLinks struct {
	Link []Link `json:"_link"`
}

type baseList struct {
	TotalItems int    `json:"total_items"`
	Links      []Link `json:"_links"`
}

// Link references another object
type Link struct {
	Rel          string `json:"re"`
	Href         string `json:"href"`
	Method       string `json:"method"`
	TargetSchema string `json:"targetSchema"`
	Schema       string `json:"schema"`
}

// Address represents what it says
type Address struct {
	Address1     string  `json:"address1"`
	Address2     string  `json:"address2"`
	City         string  `json:"city"`
	Province     string  `json:"province"`
	ProvinceCode string  `json:"province_code"`
	PostalCode   string  `json:"postal_code"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"country_code"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

// Customer defines a mailchimp customer
type Customer struct {
	// Required
	ID string `json:"id"`

	// Optional
	EmailAddress string   `json:"email_address,omitempty"`
	OptInStatus  bool     `json:"opt_in_status,omitempty"`
	Company      string   `json:"company,omitempty"`
	FirstName    string   `json:"first_name,omitempty"`
	LastName     string   `json:"last_name,omitempty"`
	OrdersCount  int      `json:"orders_count,omitempty"`
	TotalSpent   float64  `json:"total_spent,omitempty"`
	Address      *Address `json:"address,omitempty"`

	// Response
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Links     []Link `json:"_links,omitempty"`
}

// LineItem defines a mailchimp cart or order line item
type LineItem struct {
	// Required
	ID               string  `json:"id"`
	ProductID        string  `json:"product_id"`
	ProductVariantID string  `json:"product_variant_id"`
	Quantity         int     `json:"quantity"`
	Price            float64 `json:"price"`

	// Optional
	ProductTitle        string `json:"product_title,omitempty"`
	ProductVariantTitle string `json:"product_variant_title,omitempty"`
}

// Contact defines a single contact
type Contact struct {
	Company     string `json:"company"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phone"`
}
