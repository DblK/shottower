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

// TemplateRender - Render a template by it's id and optional merge fields.
type TemplateRender struct {

	// The id of the template to render in UUID format.
	ID string `json:"id"`

	// An array of key/value pairs that provides an easy way to create templates with placeholders. The placeholders can be used to find and replace keys with values. For example you can search for the placeholder `{{NAME}}` and replace it with the value `Jane`.
	Merge []MergeField `json:"merge,omitempty"`
}

// AssertTemplateRenderRequired checks if the required fields are not zero-ed
func AssertTemplateRenderRequired(obj TemplateRender) error {
	elements := map[string]interface{}{
		"id": obj.ID,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Schema: "Template Render", Field: name}
		}
	}

	for i := range obj.Merge {
		if err := AssertMergeFieldRequired(&obj.Merge[i]); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseTemplateRenderRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TemplateRender (e.g. [][]TemplateRender), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTemplateRenderRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTemplateRender, ok := obj.(TemplateRender)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTemplateRenderRequired(aTemplateRender)
	})
}
