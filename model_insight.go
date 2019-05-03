// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type Insight struct {
	TotalWatchTime      int64   `json:"total_watch_time,omitempty"`
	TotalViews          int64   `json:"total_views,omitempty"`
	NegativeImpactScore float32 `json:"negative_impact_score,omitempty"`
	Metric              float64 `json:"metric,omitempty"`
	FilterValue         string  `json:"filter_value,omitempty"`
	FilterColumn        string  `json:"filter_column,omitempty"`
}
