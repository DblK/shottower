/*
shottower
Copyright (C) 2022 Rémy Boulanouar

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

// LumaAsset - The LumaAsset is used to create luma matte masks, transitions and effects between other assets. A luma matte is a grey scale image or animated video where the black areas are transparent and the white areas solid. The luma matte animation should be provided as an mp4 video file. The src must be a publicly accessible URL to the file.
type LumaAsset struct {

	// The type of asset - set to `luma` for luma mattes.
	Type string `json:"type"`

	// The luma matte source URL. The URL must be publicly accessible or include credentials.
	Src string `json:"src"`

	// The start trim point of the luma matte clip, in seconds (defaults to 0). Videos will start from the in trim point. A luma matte video will play until the file ends or the Clip length is reached.
	Trim float32 `json:"trim,omitempty"`
}

func NewLumaAsset(m map[string]interface{}) *LumaAsset {
	lumaAsset := &LumaAsset{
		Type: m["type"].(string),
	}

	if m["src"] != nil {
		lumaAsset.Src = m["src"].(string)
	}
	if m["trim"] != nil {
		lumaAsset.Trim = cast.ToFloat32(m["trim"].(float64))
	}
	return lumaAsset
}

// AssertLumaAssetRequired checks if the required fields are not zero-ed
func AssertLumaAssetRequired(obj LumaAsset) error {
	elements := map[string]interface{}{
		"type": obj.Type,
		"src":  obj.Src,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Schema: "Luma", Field: name}
		}
	}

	return nil
}

// AssertRecurseLumaAssetRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LumaAsset (e.g. [][]LumaAsset), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLumaAssetRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLumaAsset, ok := obj.(LumaAsset)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLumaAssetRequired(aLumaAsset)
	})
}
