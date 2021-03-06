// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

// Input object with additional configuration
type InputSettings struct {
	Url             string                       `json:"url,omitempty"`
	OverlaySettings InputSettingsOverlaySettings `json:"overlay_settings,omitempty"`
	Type            string                       `json:"type,omitempty"`
	TextType        string                       `json:"text_type,omitempty"`
	LanguageCode    string                       `json:"language_code,omitempty"`
	Name            string                       `json:"name,omitempty"`
	ClosedCaptions  bool                         `json:"closed_captions,omitempty"`
	Passthrough     string                       `json:"passthrough,omitempty"`
}
