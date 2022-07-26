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

// AssetResponseAttributes - The list of asset attributes and their values.
type AssetResponseAttributes struct {

	// The unique id of the hosted asset in UUID format.
	ID string `json:"id"`

	// The owner id of the render task.
	Owner string `json:"owner"`

	// The region the asset is hosted, currently only `au` (Australia).
	Region string `json:"region,omitempty"`

	// The original render id that created the asset in UUID format. Multiple assets can share the same render id.
	RenderID string `json:"renderId,omitempty"`

	// The third party id of an asset transferred to an external provider, i.e. Mux, YouTube or S3. If the provider is Shotstack, the providerID is the same as the asset id.
	ProviderID string `json:"providerId,omitempty"`

	// The asset file name.
	Filename string `json:"filename"`

	// The asset file name.
	URL string `json:"url,omitempty"`

	// The status of the asset. <ul>   <li>`importing` - the asset is being copied to the hosting service</li>   <li>`ready` - the asset is ready to be served to users</li>   <li>`failed` - the asset failed to copy or delete</li>   <li>`deleted` - the asset has been deleted</li> </ul>
	Status string `json:"status"`

	// The time the asset was created.
	Created string `json:"created"`

	// The time the asset status was last updated.
	Updated string `json:"updated"`
}

// AssertAssetResponseAttributesRequired checks if the required fields are not zero-ed
func AssertAssetResponseAttributesRequired(obj AssetResponseAttributes) error {
	elements := map[string]interface{}{
		"id":       obj.ID,
		"owner":    obj.Owner,
		"filename": obj.Filename,
		"status":   obj.Status,
		"created":  obj.Created,
		"updated":  obj.Updated,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Schema: "Asset Response Attributes", Field: name}
		}
	}

	return nil
}

// AssertRecurseAssetResponseAttributesRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AssetResponseAttributes (e.g. [][]AssetResponseAttributes), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAssetResponseAttributesRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAssetResponseAttributes, ok := obj.(AssetResponseAttributes)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAssetResponseAttributesRequired(aAssetResponseAttributes)
	})
}
