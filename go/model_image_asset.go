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

import "github.com/creasty/defaults"

// ImageAsset - The ImageAsset is used to create video from images to compose an image. The src must be a publicly accessible URL to an image resource such as a jpg or png file.
type ImageAsset struct {

	// The type of asset - set to `image` for images.
	Type string `json:"type"`

	// The image source URL. The URL must be publicly accessible or include credentials.
	Src string `json:"src"`

	Crop *Crop `json:"crop,omitempty"`
}

func NewImageAsset(m map[string]interface{}) *ImageAsset {
	imageAsset := &ImageAsset{
		Type: m["type"].(string),
	}

	if m["src"] != nil {
		imageAsset.Src = m["src"].(string)
	}
	if m["crop"] != nil {
		imageAsset.Crop = NewCrop(m["crop"].(map[string]interface{}))
	}
	return imageAsset
}

// AssertImageAssetRequired checks if the required fields are not zero-ed
func AssertImageAssetRequired(obj ImageAsset) error {
	elements := map[string]interface{}{
		"type": obj.Type,
		"src":  obj.Src,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Schema: "Image Asset", Field: name}
		}
	}

	if err := defaults.Set(obj); err != nil {
		panic(err)
	}

	if obj.Src == "" {
		return &RequiredError{Schema: "Image Asset", Field: "src"}
	}

	if err := AssertCropRequired(obj.Crop); err != nil {
		return err
	}
	return nil
}

// AssertRecurseImageAssetRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ImageAsset (e.g. [][]ImageAsset), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseImageAssetRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aImageAsset, ok := obj.(ImageAsset)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertImageAssetRequired(aImageAsset)
	})
}
