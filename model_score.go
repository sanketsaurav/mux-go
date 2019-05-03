// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type Score struct {
	WatchTime int64    `json:"watch_time,omitempty"`
	ViewCount int64    `json:"view_count,omitempty"`
	Name      string   `json:"name,omitempty"`
	Value     float64  `json:"value,omitempty"`
	Metric    string   `json:"metric,omitempty"`
	Items     []Metric `json:"items,omitempty"`
}
