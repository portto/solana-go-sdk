package client

import (
	"context"
	"errors"
)

// GetFirstAvailableBlock returns the slot of the lowest confirmed block that has not been purged from the ledger
func (s *Client) GetFirstAvailableBlock(ctx context.Context) (uint64, error) {
	res := struct {
		GeneralResponse
		Result uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getFirstAvailableBlock", []interface{}{}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
