// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type ListIncidentsResponse struct {
	Data          []Incident `json:"data,omitempty"`
	TotalRowCount int64      `json:"total_row_count,omitempty"`
	Timeframe     []int64    `json:"timeframe,omitempty"`
}
