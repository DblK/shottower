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

// Transformation - Apply one or more transformations to a clip. Transformations alter the visual properties of a clip and can be combined to create new shapes and effects.
type Transformation struct {
	Rotate RotateTransformation `json:"rotate,omitempty"`

	Skew SkewTransformation `json:"skew,omitempty"`

	Flip FlipTransformation `json:"flip,omitempty"`
}

func NewTransformation(m map[string]interface{}) *Transformation {
	transform := &Transformation{}

	if m["rotate"] != nil {
		transform.Rotate = *NewRotateTransformation(m["rotate"].(map[string]interface{}))
	}
	if m["skew"] != nil {
		transform.Skew = *NewSkewTransformation(m["skew"].(map[string]interface{}))
	}
	if m["flip"] != nil {
		transform.Flip = *NewFlipTransformation(m["flip"].(map[string]interface{}))
	}
	return transform
}

// AssertTransformationRequired checks if the required fields are not zero-ed
func AssertTransformationRequired(obj *Transformation) error {
	if obj == nil {
		return nil
	}

	if err := AssertRotateTransformationRequired(obj.Rotate); err != nil {
		return err
	}
	if err := AssertSkewTransformationRequired(obj.Skew); err != nil {
		return err
	}
	if err := AssertFlipTransformationRequired(obj.Flip); err != nil {
		return err
	}
	return nil
}

// AssertRecurseTransformationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Transformation (e.g. [][]Transformation), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTransformationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTransformation, ok := obj.(Transformation)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTransformationRequired(&aTransformation)
	})
}
