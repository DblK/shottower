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

import (
	"context"
	"errors"
	"net/http"

	"github.com/DblK/shottower/config"
)

// DownloadAPIService is a service that implements the logic for the DownloadAPIServicer
// This service should implement the business logic for every endpoint for the DownloadApi API.
// Include any external packages or services that will be required by this service.
type DownloadAPIService struct {
	config  config.ShottowerConfiguration
	editAPI EditAPIServicer
}

// NewDownloadAPIService creates a default api service
func NewDownloadAPIService(cfg config.ShottowerConfiguration, editAPI EditAPIServicer) DownloadAPIServicer {
	return &DownloadAPIService{
		config:  cfg,
		editAPI: editAPI,
	}
}

// DownloadByID - Download Asset by ID
func (s *DownloadAPIService) DownloadByID(ctx context.Context, id string) (ImplResponse, error) {
	var renderQueue *RenderQueue
	for _, n := range s.editAPI.GetQueue() {
		if n.ID == id {
			renderQueue = n
		}
	}

	if renderQueue == nil {
		return Response(http.StatusNotFound, nil), errors.New("Can not find id: '" + id + "'")
	}

	return Response(200, renderQueue.FileName), nil
}
