// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BlockHeader block header
//
// swagger:model blockHeader
type BlockHeader struct {

	// height
	Height int64 `json:"height,omitempty"`

	// nonce
	Nonce string `json:"nonce,omitempty"`

	// prev block hash
	PrevBlockHash string `json:"prevBlockHash,omitempty"`

	// sheet hash
	SheetHash string `json:"sheet_hash,omitempty"`

	// target
	Target string `json:"target,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`

	// trie hash
	TrieHash string `json:"trie_hash,omitempty"`
}

// Validate validates this block header
func (m *BlockHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeader) UnmarshalBinary(b []byte) error {
	var res BlockHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
