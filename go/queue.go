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
package openapi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"path/filepath"
	"time"
)

type ProcessingQueue struct {
	currentQueue *RenderQueue

	LocalResources map[string]*LocalResource
}

func NewProcessingQueuer() ProcessingQueuer {
	processingQueue := &ProcessingQueue{}
	processingQueue.LocalResources = make(map[string]*LocalResource)

	return processingQueue
}

func (s *ProcessingQueue) StartProcessQueue(editAPI EditAPIServicer) {
	fmt.Println("StartProcessQueue ", len(editAPI.GetQueue()))

	go time.AfterFunc(1*time.Second, func() {
		s.ProcessQueue(editAPI)
	})

	go time.AfterFunc(1*time.Hour, func() {
		s.CleanCache()
	})
}

func (s *ProcessingQueue) CleanCache() {
	for _, resource := range s.LocalResources {
		if !resource.KeepCache {
			// TODO: Delete local resource
			// resource.LocalURL

			s.LocalResources[resource.OriginalURL] = nil
		}
	}

	go time.AfterFunc(1*time.Hour, func() {
		s.CleanCache()
	})
}

func (s *ProcessingQueue) FindSourceClip(trackNumber int, clipNumber int) string {
	for _, resource := range s.LocalResources {
		for _, used := range resource.Used {
			if used.Track == trackNumber && used.Clip == clipNumber {
				return resource.LocalURL
			}
		}
	}
	return ""
}

func (s *ProcessingQueue) ProcessQueue(editAPI EditAPIServicer) {
	var queue = editAPI.GetQueue()
	fmt.Println("ProcessQueue (pending:", len(editAPI.GetQueuePending()), ") ", len(queue))

	if len(queue) != 0 && s.currentQueue == nil {
		if len(editAPI.GetQueuePending()) != 0 {
			s.currentQueue = editAPI.GetQueuePending()[0]
			s.currentQueue.FFMPEGCommand = NewFFMPEGCommand()

			// Set Resolution (Mandatory for Filler)
			if s.currentQueue.Data.Output.Resolution != "" {
				_ = s.currentQueue.FFMPEGCommand.SetOutputResolution(s.currentQueue.Data.Output.Resolution)
			} else if s.currentQueue.Data.Output.Size != nil {
				_ = s.currentQueue.FFMPEGCommand.SetOutputSize(s.currentQueue.Data.Output.Size.Width, s.currentQueue.Data.Output.Size.Height)
			}

			fmt.Println(s.currentQueue.ID)

			go s.FetchAssets(s.currentQueue)
		}
	}

	if s.currentQueue != nil {
		fmt.Println(s.currentQueue.Status.String())

		if s.currentQueue.InternalStatus == Fetched {
			params := s.GenerateParameters(s.currentQueue)
			s.currentQueue.Updated = time.Now()

			// Launch Rendering
			// fmt.Println(params)
			go s.ExecuteFFMpeg(params)
		}

		if s.currentQueue.InternalStatus == Rendered {
			s.currentQueue.Status = Saving
			s.currentQueue.InternalStatus = Saving
			s.currentQueue.FileName = s.currentQueue.FFMPEGCommand.GetOutputName()
		}

		if s.currentQueue.InternalStatus == Saving {
			// FIXME: Pass through Saving status at some point
			s.currentQueue.Status = Done
			s.currentQueue.InternalStatus = Done
		}

		if s.currentQueue.InternalStatus == Failed || s.currentQueue.InternalStatus == Done {
			s.currentQueue = nil
		}
	}

	go time.AfterFunc(1*time.Second, func() {
		s.ProcessQueue(editAPI)
	})
}

func (s *ProcessingQueue) ExecuteFFMpeg(params []string) {
	s.currentQueue.InternalStatus = Rendering
	cmd := exec.Command("ffmpeg", params...)

	// FIXME: Dev debug
	// fmt.Println(strings.Join(cmd.Args, "||"))

	var out bytes.Buffer
	cmd.Stdout = &out
	var outErr bytes.Buffer
	cmd.Stderr = &outErr

	// Dev debug
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	err := cmd.Run()

	s.currentQueue.Updated = time.Now()
	if err != nil {
		fmt.Println(outErr.String())
		s.currentQueue.Status = Failed
		s.currentQueue.InternalStatus = Failed
		log.Fatal(err)
		// fmt.Println(err)
	} else {
		fmt.Println("FFMPEG ended!")
		fmt.Println(out.String())
		s.currentQueue.InternalStatus = Rendered
	}
}

func (s *ProcessingQueue) GenerateParameters(queue *RenderQueue) []string {
	queue.Status = Rendering
	queue.InternalStatus = Generating

	_ = queue.FFMPEGCommand.ToFFMPEG(queue, s)

	queue.InternalStatus = Generated

	return queue.FFMPEGCommand.ToString()
}

func (s *ProcessingQueue) FetchVideoAssets(trackNumber int, clipNumber int, clip Clip, useCache bool) bool {
	var hasError bool

	var asset = clip.Asset.(*VideoAsset)

	if s.LocalResources[asset.Src] == nil {
		var fileName string
		url, _ := url.Parse(asset.Src)

		if url.Scheme == "file" {
			fileName = asset.Src[7:]
		} else {
			var err error
			fileName, err = s.DownloadFile(asset.Src)
			if err != nil {
				fmt.Println("Error while downloading asset", err)
				hasError = true
			}
		}

		if !hasError {
			fmt.Println("Asset downloaded: "+asset.Src, fileName)
			localResource := &LocalResource{
				Downloaded:  time.Now(),
				OriginalURL: asset.Src,
				LocalURL:    fileName,
				KeepCache:   useCache,
			}
			s.LocalResources[asset.Src] = localResource
		}
	}

	if !hasError {
		s.LocalResources[asset.Src].Used = append(
			s.LocalResources[asset.Src].Used,
			&LocalResourceTrackInfo{
				Track: trackNumber,
				Clip:  clipNumber,
			})
	}

	return hasError
}

func (s *ProcessingQueue) FetchAssets(queue *RenderQueue) {
	queue.Status = Fetching
	queue.InternalStatus = Fetching

	var hasError bool

	useCache := queue.Data.Timeline.Cache

	for tIndex, track := range queue.Data.Timeline.Tracks {
		for cIndex, clip := range track.Clips {
			// fmt.Println(tIndex, cIndex, clip.Asset.Type)

			var typeAsset = GetAssetType(clip.Asset)
			switch typeAsset { // nolint:exhaustive
			case VideoAssetType:
				hasError = s.FetchVideoAssets(tIndex, cIndex, clip, useCache)

			// case "image":
			// 	fmt.Println("Image")
			// 	if clip.IsFiller {
			// 		// queue.LocalResources = append(queue.LocalResources, queue.FFMPEGCommand.GenerateFiller(""))
			// 	} else {
			// 		fmt.Println("TODO: Download asset")
			// 	}
			default:
				fmt.Println("Unhandled asset type", typeAsset.String())
			}
		}
	}

	// Dev Only
	// fmt.Println(queue.LocalResources)

	if !hasError {
		queue.InternalStatus = Fetched
		queue.Status = Fetching
	} else {
		queue.InternalStatus = Failed
		queue.Status = Failed
	}
}

func (s *ProcessingQueue) DownloadFile(url string) (string, error) {
	file, err := ioutil.TempFile("", "asset*"+filepath.Ext(url))
	if err != nil {
		fmt.Println("Error temp file", err)
		return "", err
	}
	defer file.Close()
	fmt.Println("----" + file.Name())

	// Get the data
	resp, err := http.Get(url) // nolint:gosec,G107
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

	// Create the file
	// out, err := os.Create(file.Name())
	// if err != nil {
	// 	return err
	// }
	// defer out.Close()

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	return file.Name(), err
}
