package catalog

import (
	"encoding/base64"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ListOptions defines configurable parameters for list/pagination behavior.
type ListOptions struct {
	Limit   uint64 `json:"limit"`   // Number of items per page
	Page    uint64 `json:"page"`    // Current page (starting from 0)
	Preload bool   `json:"preload"` // Whether to preload related data
}

// GetOffset calculates the offset based on the page and limit.
func (f ListOptions) GetOffset() uint64 {
	return f.Page * f.Limit
}

// ListOption defines a function that modifies ListOptions.
type ListOption func(*ListOptions)

// WithMaxResults sets the limit if the provided value is greater than zero.
func WithMaxResults(limit uint64) ListOption {
	return func(o *ListOptions) {
		if limit > 0 {
			o.Limit = limit
		}
	}
}

// WithPreload enables preloading of related data.
func WithPreload() ListOption {
	return func(o *ListOptions) {
		o.Preload = true
	}
}

// WithNextPage sets the page number based on a base64-encoded page token.
func WithNextPage(nextPage string) ListOption {
	return func(o *ListOptions) {
		if nextPage != "" {
			pt, _ := decode(nextPage) // Ignoring decode error
			o.Page = pt.page
		}
	}
}

// pageToken represents the page number for encoding/decoding.
type pageToken struct {
	page uint64
}

// encode converts the page number to a base64 string.
func (p pageToken) encode() string {
	token := fmt.Sprintf("%d", p.page)
	return base64.StdEncoding.EncodeToString([]byte(token))
}

// decode parses a base64-encoded string into a pageToken.
func decode(token string) (pageToken, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return pageToken{}, err
	}
	var page uint64
	if _, err := fmt.Sscanf(string(decoded), "%d", &page); err != nil {
		return pageToken{}, err
	}
	return pageToken{page}, nil
}

// FindOptions defines options for data retrieval operations.
type FindOptions struct {
	Tx *sqlx.Tx // Optional SQL transaction
}

// FindOption defines a function that modifies FindOptions.
type FindOption func(*FindOptions)

// WithFindTx sets the transaction for a find operation.
func WithFindTx(tx *sqlx.Tx) FindOption {
	return func(o *FindOptions) {
		o.Tx = tx
	}
}

// UpdateOptions defines options for update operations.
type UpdateOptions struct {
	Tx *sqlx.Tx // Optional SQL transaction
}

// UpdateOption defines a function that modifies UpdateOptions.
type UpdateOption func(*UpdateOptions)

// WithUpdateTx sets the transaction for an update operation.
func WithUpdateTx(tx *sqlx.Tx) UpdateOption {
	return func(o *UpdateOptions) {
		o.Tx = tx
	}
}
