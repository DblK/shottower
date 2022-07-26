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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// EditAPIController binds http requests to an api service and writes the service results to the http response
type EditAPIController struct {
	service      EditAPIServicer
	errorHandler ErrorHandler
	debug        bool
}

// EditAPIOption for how the controller is set up.
type EditAPIOption func(*EditAPIController)

// WithEditAPIErrorHandler inject ErrorHandler into controller
func WithEditAPIErrorHandler(h ErrorHandler) EditAPIOption {
	return func(c *EditAPIController) {
		c.errorHandler = h
	}
}

// NewEditAPIController creates a default api controller
func NewEditAPIController(s EditAPIServicer, opts ...EditAPIOption) Router {
	controller := &EditAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
		debug:        true, // Dev only
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EditAPIController
func (c *EditAPIController) Routes() Routes {
	return Routes{
		{
			"DeleteTemplate",
			strings.ToUpper("Delete"),
			"/stage/templates/{id}",
			c.DeleteTemplate,
		},
		{
			"GetRender",
			strings.ToUpper("Get"),
			"/stage/render/{id}",
			c.GetRender,
		},
		{
			"GetTemplate",
			strings.ToUpper("Get"),
			"/stage/templates/{id}",
			c.GetTemplate,
		},
		{
			"GetTemplates",
			strings.ToUpper("Get"),
			"/stage/templates",
			c.GetTemplates,
		},
		{
			"PostRender",
			strings.ToUpper("Post"),
			"/stage/render",
			c.PostRender,
		},
		{
			"PostTemplate",
			strings.ToUpper("Post"),
			"/stage/templates",
			c.PostTemplate,
		},
		{
			"PostTemplateRender",
			strings.ToUpper("Post"),
			"/stage/templates/render",
			c.PostTemplateRender,
		},
		{
			"Probe",
			strings.ToUpper("Get"),
			"/stage/probe/{url}",
			c.Probe,
		},
		{
			"PutTemplate",
			strings.ToUpper("Put"),
			"/stage/templates/{id}",
			c.PutTemplate,
		},
	}
}

// DeleteTemplate - Delete Template
func (c *EditAPIController) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]

	result, err := c.service.DeleteTemplate(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetRender - Get Render Status
func (c *EditAPIController) GetRender(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	idParam := params["id"]

	dataParam := true
	if _, found := mux.Vars(r)["data"]; found {
		var err error
		dataParam, err = parseBoolParameter(query.Get("data"))
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}
	fmt.Println(query.Get("merged"))
	mergedParam := false
	if _, found := mux.Vars(r)["merged"]; found {
		var err error
		mergedParam, err = parseBoolParameter(query.Get("merged"))
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}
	result, err := c.service.GetRender(r.Context(), idParam, dataParam, mergedParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTemplate - Retrieve Template
func (c *EditAPIController) GetTemplate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]

	result, err := c.service.GetTemplate(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTemplates - List Templates
func (c *EditAPIController) GetTemplates(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetTemplates(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// PostRender - Render Asset
func (c *EditAPIController) PostRender(w http.ResponseWriter, r *http.Request) {
	editParam := Edit{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&editParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	if c.debug {
		queryFile, err := json.MarshalIndent(editParam, "", "  ")
		if err == nil {
			_ = ioutil.WriteFile("query.json", queryFile, 0600)
		}
	}

	var min float32 = 99999999999
	var max float32
	for tIndex, track := range editParam.Timeline.Tracks {
		for cIndex, clip := range track.Clips {
			// Find min and max timing for the timeline
			if clip.Start < min {
				min = clip.Start
			}
			if clip.Start+clip.Length > max {
				max = clip.Start + clip.Length
			}

			// Convert clip Asset type
			switch clip.Asset.Type {
			case "video":
				var videoAsset = &VideoAsset{
					Type:     clip.Asset.Type,
					Src:      clip.Asset.Src,
					Trim:     clip.Asset.Trim,
					Volume:   clip.Asset.Volume,
					Crop:     clip.Asset.Crop,
					Subtitle: clip.Asset.Subtitle,
				}
				editParam.Timeline.Tracks[tIndex].Clips[cIndex].TypedAsset = *videoAsset
			case "image":
				var imageAsset = &ImageAsset{
					Type: clip.Asset.Type,
					Src:  clip.Asset.Src,
					Crop: clip.Asset.Crop,
				}
				editParam.Timeline.Tracks[tIndex].Clips[cIndex].TypedAsset = *imageAsset
			default:
				fmt.Println("Asset type not handled!!", clip.Asset.Type)
			}
		}
	}

	// TMP Dev display
	// for tIndex, track := range editParam.Timeline.Tracks {
	// 	fmt.Println("Track ClipSize", tIndex, len(track.Clips))
	// 	for cIndex, clip := range track.Clips {
	// 		fmt.Println(tIndex, cIndex, clip.Asset.Src, clip.Start, clip.Length)
	// 	}
	// }

	if err := AssertEditRequired(&editParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostRender(r.Context(), editParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// PostTemplate - Create Template
func (c *EditAPIController) PostTemplate(w http.ResponseWriter, r *http.Request) {
	templateParam := Template{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&templateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTemplateRequired(templateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostTemplate(r.Context(), templateParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// PostTemplateRender - Render Template
func (c *EditAPIController) PostTemplateRender(w http.ResponseWriter, r *http.Request) {
	templateRenderParam := TemplateRender{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&templateRenderParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTemplateRenderRequired(templateRenderParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostTemplateRender(r.Context(), templateRenderParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// Probe - Inspect Media
func (c *EditAPIController) Probe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	urlParam := params["url"]

	result, err := c.service.Probe(r.Context(), urlParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// PutTemplate - Update Template
func (c *EditAPIController) PutTemplate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]

	templateParam := Template{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&templateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTemplateRequired(templateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PutTemplate(r.Context(), idParam, templateParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
