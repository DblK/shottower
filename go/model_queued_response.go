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

// QueuedResponse - The response received after a [render request](#render-asset) or [template render](#render-template) is submitted. The render task is queued for rendering and a unique render id is returned.
type QueuedResponse struct {

	// `true` if successfully queued, else `false`.
	Success bool `json:"success"`

	// `Created`, `Bad Request` or an error message.
	Message string `json:"message"`

	Response QueuedResponseData `json:"response"`
}

// AssertQueuedResponseRequired checks if the required fields are not zero-ed
func AssertQueuedResponseRequired(obj QueuedResponse) error {
	elements := map[string]interface{}{
		"success":  obj.Success,
		"message":  obj.Message,
		"response": obj.Response,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Schema: "Queued Response", Field: name}
		}
	}

	if err := AssertQueuedResponseDataRequired(obj.Response); err != nil {
		return err
	}
	return nil
}

// AssertRecurseQueuedResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of QueuedResponse (e.g. [][]QueuedResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQueuedResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQueuedResponse, ok := obj.(QueuedResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQueuedResponseRequired(aQueuedResponse)
	})
}
