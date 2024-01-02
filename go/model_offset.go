/*
shottower
Copyright (C) 2022-2023 Rémy Boulanouar

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * Shottower
 *
 * Shottower is the open source version of Shotstack which is a video, image and audio editing service that allows for the automated generation of videos, images and audio using JSON and a RESTful API.  You arrange and configure an edit and POST it to the API which will render your media and provide a file  location when complete.  For more details visit [shottower](https://github.com/DblK/shottower) or checkout our [getting started](https://shotstack.io/docs/guide/) documentation.  There are two main API's, one for editing and generating assets (Edit API) and one for managing hosted assets (Serve API).  The Edit API base URL is: <b>http://0.0.0.0:4000/{version}</b>  The Serve API base URL is: <b>http://0.0.0.0:4000/serve/{version}</b>
 *
 * API version: stage
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import "github.com/spf13/cast"

// Offset - Offsets the position of an asset horizontally or vertically by a relative distance.
type Offset struct {

	// Offset an asset on the horizontal axis (left or right), range varies from -1 to 1. Positive numbers move the asset right, negative left. For all assets except titles the distance moved is relative to the width  of the viewport - i.e. an X offset of 0.5 will move the asset half the  screen width to the right.
	X float32 `json:"x,omitempty"`

	// Offset an asset on the vertical axis (up or down), range varies from -1 to 1. Positive numbers move the asset up, negative down. For all assets except titles the distance moved is relative to the height  of the viewport - i.e. an Y offset of 0.5 will move the asset up half the  screen height.
	Y float32 `json:"y,omitempty"`
}

func NewOffset(m map[string]interface{}) *Offset {
	offset := &Offset{}

	if m["x"] != nil {
		offset.X = cast.ToFloat32(m["x"].(float64))
	}
	if m["y"] != nil {
		offset.Y = cast.ToFloat32(m["y"].(float64))
	}
	return offset
}

func (s *Offset) checkEnumValues() error {
	if s.X < -1 || s.X > 1 {
		return &EnumError{Schema: "Offset", Field: "X", Value: s.X}
	}
	if s.Y < -1 || s.Y > 1 {
		return &EnumError{Schema: "Offset", Field: "Y", Value: s.Y}
	}

	return nil
}

// AssertOffsetRequired checks if the required fields are not zero-ed
func AssertOffsetRequired(obj *Offset) error {
	if obj == nil {
		return nil
	}

	if err := obj.checkEnumValues(); err != nil {
		return err
	}

	return nil
}

// AssertRecurseOffsetRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Offset (e.g. [][]Offset), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseOffsetRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aOffset, ok := obj.(Offset)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertOffsetRequired(&aOffset)
	})
}
