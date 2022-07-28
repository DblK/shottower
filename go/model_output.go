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

	"github.com/creasty/defaults"
	"github.com/spf13/cast"
	"golang.org/x/exp/slices"
)

// Output - The output format, render range and type of media to generate.
type Output struct {

	// The output format and type of media file to generate. <ul>   <li>`mp4` - mp4 video file</li>   <li>`gif` - animated gif</li>   <li>`jpg` - jpg image file</li>   <li>`png` - png image file</li>   <li>`bmp` - bmp image file</li>   <li>`mp3` - mp3 audio file (audio only)</li> </ul>
	Format string `json:"format"`

	// The output resolution of the video or image. <ul>   <li>`preview` - 512px x 288px @ 15fps</li>   <li>`mobile` - 640px x 360px @ 25fps</li>   <li>`sd` - 1024px x 576px @ 25fps</li>   <li>`hd` - 1280px x 720px @ 25fps</li>   <li>`1080` - 1920px x 1080px @ 25fps</li>   <li>`360` - 640px x 360px @ 25fps</li>   <li>`480` - 848px x 480px @ 25fps</li>   <li>`540` - 960px x 540px @ 25fps</li>   <li>`720` - 1280px x 720px @ 25fps</li> </ul>
	Resolution string `json:"resolution,omitempty"`

	// The aspect ratio (shape) of the video or image. Useful for social media output formats. Options are: <ul>   <li>`16:9` - regular landscape/horizontal aspect ratio (default)</li>   <li>`9:16` - vertical/portrait aspect ratio</li>   <li>`1:1` - square aspect ratio</li>   <li>`4:5` - short vertical/portrait aspect ratio</li>   <li>`4:3` - legacy TV aspect ratio</li> </ul>
	AspectRatio string `json:"aspectRatio,omitempty"`

	Size *Size `json:"size,omitempty"`

	// Override the default frames per second. Useful for when the source footage is recorded at 30fps, i.e. on  mobile devices. Lower frame rates can be used to add cinematic quality (24fps) or to create smaller file size/faster render times or animated gifs (12 or 15fps). Default is 25fps. <ul>   <li>`12` - 12fps</li>   <li>`15` - 15fps</li>   <li>`24` - 24fps</li>   <li>`23.976` - 23.976fps</li>   <li>`25` - 25fps</li>   <li>`29.97` - 29.97fps</li>   <li>`30` - 30fps</li> </ul>
	Fps *float32 `json:"fps,omitempty" default:"25"`

	// Override the resolution and scale the video or image to render at a different size. When using scaleTo the asset should be edited at the resolution dimensions, i.e. use font sizes that look best at HD, then use scaleTo to output the file at SD and the text will be scaled to the correct size. This is useful if you want to create multiple asset sizes. <ul>   <li>`preview` - 512px x 288px @ 15fps</li>   <li>`mobile` - 640px x 360px @ 25fps</li>   <li>`sd` - 1024px x 576px @25fps</li>   <li>`hd` - 1280px x 720px @25fps</li>   <li>`1080` - 1920px x 1080px @25fps</li> </ul>
	ScaleTo string `json:"scaleTo,omitempty"`

	// Adjust the output quality of the video, image or audio. Adjusting quality affects  render speed, download speeds and storage requirements due to file size. The default `medium` provides the most optimized choice for all three  factors. <ul>   <li>`low` - slightly reduced quality, smaller file size</li>   <li>`medium` - optimized quality, render speeds and file size</li>   <li>`high` - slightly increased quality, larger file size</li> </ul>
	Quality string `json:"quality,omitempty" default:"medium"`

	// Loop settings for gif files. Set to `true` to loop, `false` to play only once.
	Repeat bool `json:"repeat,omitempty"`

	Range *Range `json:"range,omitempty"`

	Poster *Poster `json:"poster,omitempty"`

	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`

	Destinations []interface{} `json:"destinations,omitempty"`
}

func NewOutput(data map[string]interface{}, destinations []interface{}) *Output {
	output := &Output{
		Format: data["format"].(string),
	}

	if data["resolution"] != nil {
		output.Resolution = data["resolution"].(string)
	}
	if data["aspectRatio"] != nil {
		output.AspectRatio = data["aspectRatio"].(string)
	}
	if data["size"] != nil {
		*output.Size = *NewSize(data["size"].(map[string]interface{}))
	}
	if data["fps"] != nil {
		fps := cast.ToFloat32(data["fps"].(float64))
		output.Fps = &fps
	}
	if data["scaleTo"] != nil {
		output.ScaleTo = data["scaleTo"].(string)
	}
	if data["quality"] != nil {
		output.Quality = data["quality"].(string)
	}
	if data["repeat"] != nil {
		output.Repeat = data["repeat"].(bool)
	}
	if data["range"] != nil {
		*output.Range = *NewRange(data["range"].(map[string]interface{}))
	}
	if data["poster"] != nil {
		*output.Poster = *NewPoster(data["poster"].(map[string]interface{}))
	}
	if data["thumbnail"] != nil {
		*output.Thumbnail = *NewThumbnail(data["thumbnail"].(map[string]interface{}))
	}

	output.Destinations = destinations
	return output
}

func (s *Output) UnmarshalJSON(data []byte) error {
	var obj map[string]interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	var destinations []interface{}
	if obj["destinations"] != nil {
		for _, dest := range obj["destinations"].([]interface{}) {
			var provider = dest.(map[string]interface{})["provider"].(string)
			destination := NewDestination(provider, dest.(map[string]interface{}))
			destinations = append(destinations, destination)
		}
	}

	*s = *NewOutput(obj, destinations)

	return nil
}

func (s *Output) checkEnumValues() error {
	formatValues := []string{"mp4", "gif", "jpg", "png", "bmp", "mp3"}
	if !slices.Contains(formatValues, s.Format) {
		return &EnumError{Schema: "Output", Field: "Format", Value: s.Format}
	}

	resolutionValues := []string{"preview", "mobile", "sd", "hd", "1080", "360", "480", "540", "720"}
	if s.Resolution != "" && !slices.Contains(resolutionValues, s.Resolution) {
		return &EnumError{Schema: "Output", Field: "Resolution", Value: s.Resolution}
	}

	aspectRatioValues := []string{"16:9", "9:16", "1:1", "4:5", "4:3"}
	if s.AspectRatio != "" && !slices.Contains(aspectRatioValues, s.AspectRatio) {
		return &EnumError{Schema: "Output", Field: "AspectRatio", Value: s.AspectRatio}
	}

	fpsValues := []float32{12, 15, 24, 23.976, 25, 29.97, 30}
	if !slices.Contains(fpsValues, *s.Fps) {
		return &EnumError{Schema: "Output", Field: "Fps", Value: s.Fps}
	}

	scaleToValues := []string{"preview", "mobile", "sd", "hd", "1080"}
	if s.ScaleTo != "" && !slices.Contains(scaleToValues, s.ScaleTo) {
		return &EnumError{Schema: "Output", Field: "ScaleTo", Value: s.ScaleTo}
	}

	qualityValues := []string{"low", "medium", "high"}
	if s.Quality != "" && !slices.Contains(qualityValues, s.Quality) {
		return &EnumError{Schema: "Output", Field: "Quality", Value: s.Quality}
	}

	return nil
}

// AssertOutputRequired checks if the required fields are not zero-ed
func AssertOutputRequired(obj *Output) error {
	elements := map[string]interface{}{
		"format": obj.Format,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Schema: "Output", Field: name}
		}
	}

	if err := defaults.Set(obj); err != nil {
		panic(err)
	}

	if err := obj.checkEnumValues(); err != nil {
		return err
	}

	if err := AssertSizeRequired(obj.Size); err != nil {
		return err
	}
	if err := AssertRangeRequired(obj.Range); err != nil {
		return err
	}
	if err := AssertPosterRequired(obj.Poster); err != nil {
		return err
	}
	if err := AssertThumbnailRequired(obj.Thumbnail); err != nil {
		return err
	}
	if obj.Destinations != nil {
		for i := range obj.Destinations {
			if err := AssertDestinationsRequired(obj.Destinations[i]); err != nil {
				return err
			}
		}
	}
	return nil
}

// AssertRecurseOutputRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Output (e.g. [][]Output), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseOutputRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aOutput, ok := obj.(Output)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertOutputRequired(&aOutput)
	})
}
