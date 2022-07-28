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

// YoutubeDestination - Send rendered videos to [Youtube](https://www.youtube.com/) video hosting and streaming service. Add the `youtube` destination provider to send the output video to Youtube. Youtube credentials are required and added inside Options for now in the request.
type YoutubeDestination struct {

	// The destination to send rendered assets to - set to `youtube` for Youtube.
	Provider string `json:"provider"`

	Options YoutubeDestinationOptions `json:"options,omitempty"`
}

func NewYoutubeDestination(obj map[string]interface{}) *YoutubeDestination {
	destination := &YoutubeDestination{
		Provider: obj["provider"].(string),
	}

	if obj["options"] != nil {
		destination.Options = *NewYoutubeDestinationOptions(obj["options"].(map[string]interface{}))
	}

	return destination
}

// AssertYoutubeDestinationRequired checks if the required fields are not zero-ed
func AssertYoutubeDestinationRequired(obj *YoutubeDestination) error {
	elements := map[string]interface{}{
		"provider": obj.Provider,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertYoutubeDestinationOptionsRequired(obj.Options); err != nil {
		return err
	}
	return nil
}

// AssertRecurseYoutubeDestinationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of YoutubeDestination (e.g. [][]YoutubeDestination), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseYoutubeDestinationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aYoutubeDestination, ok := obj.(YoutubeDestination)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertYoutubeDestinationRequired(&aYoutubeDestination)
	})
}
