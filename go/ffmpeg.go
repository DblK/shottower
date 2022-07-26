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
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strings"

	"github.com/spf13/cast"
)

type FFMPEGSource struct {
	path string
}

type FFMPEGTrack struct {
	video []string
	audio []string
}

type FFMPEG struct {
	src                  []FFMPEGSource
	defaultParams        bool
	tracks               []FFMPEGTrack
	size                 Size
	format               string
	fps                  float32
	hasOverlay           bool
	backgroundColor      string
	overlayFillerCounter int
	fillerCounter        int
	outputName           string
	duration             float32
}

func NewFFMPEGCommand() FFMPEGCommand {
	return &FFMPEG{}
}

func (s *FFMPEG) AddDefaultParams() error {
	s.defaultParams = true
	return nil
}

func (s *FFMPEG) AddSource(fileName string) error {
	var newSource = &FFMPEGSource{
		path: fileName,
	}
	s.src = append(s.src, *newSource)
	return nil
}

func (s *FFMPEG) GetOutputName() string {
	return s.outputName
}

func (s *FFMPEG) GetDuration() float32 {
	return s.duration
}

func (s *FFMPEG) AddClip(trackNumber int, params string) error {
	s.tracks[trackNumber].video = append(s.tracks[trackNumber].video, params)
	return nil
}

func (s *FFMPEG) AddAudioClip(trackNumber int, params string) error {
	s.tracks[trackNumber].audio = append(s.tracks[trackNumber].audio, params)
	return nil
}

func (s *FFMPEG) ClipAudioVolume(sourceClip int, trackNumber int, clipNumber int, volume float32) string {
	// fmt.Println("ClipAudioVolume", trackNumber, clipNumber)
	return "[" + cast.ToString(sourceClip) + ":a]" +
		"volume=" + cast.ToString(volume) +
		s.trackName("audio", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipAudioDelay(sourceClip int, trackNumber int, clipNumber int, delay float32) string {
	// fmt.Println("ClipAudioDelay", trackNumber, clipNumber)
	return "[" + cast.ToString(sourceClip) + ":a]" +
		"adelay=" + cast.ToString(delay) +
		"|" + cast.ToString(delay) +
		s.trackName("audio", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipAudioTrim(sourceClip int, trackNumber int, clipNumber int, start float32, length float32) string {
	return "[" +
		cast.ToString(sourceClip) +
		":a] atrim=start=" +
		cast.ToString(start) +
		":end=" +
		cast.ToString(length) +
		", asetpts=PTS-STARTPTS " + s.trackName("audio", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipAudioMerge(sourceClip int, trackNumber int, clipNumber int, effects []string) string {
	var trackName = s.trackName("audio", trackNumber, clipNumber, -1)
	var numEffects = len(effects) - 1
	// fmt.Println("ClipAudioMerge: len #", len(effects))

	// fmt.Println("---")
	// fmt.Println("ClipAudioMerge (Before)", effects)
	// fmt.Println("---")

	var clipEffect = ""
	var previousTrackName = ""
	// Replace sources by previous result
	for i, effect := range effects {
		var currentTrackName = s.trackName("audio", trackNumber, clipNumber, i)
		replacedEffect := effect
		// fmt.Println("--", i, trackName, currentTrackName, clipEffect)
		// fmt.Println("--", i, trackName, previousTrackName, effect)
		if i != 0 {
			// Replace source from previous
			replacedEffect = strings.ReplaceAll(replacedEffect, "["+cast.ToString(sourceClip)+":a]", previousTrackName)
			replacedEffect = strings.ReplaceAll(replacedEffect, "["+cast.ToString(sourceClip)+"]", previousTrackName)
		}

		if i != numEffects {
			// Replace destination
			replacedEffect = strings.ReplaceAll(replacedEffect, trackName, currentTrackName)
			if previousTrackName != "" {
				// fmt.Println(i, trackName, previousTrackName, currentTrackName)
				replacedEffect = strings.ReplaceAll(replacedEffect, "["+cast.ToString(sourceClip)+":a]", previousTrackName)
				replacedEffect = strings.ReplaceAll(replacedEffect, "["+cast.ToString(sourceClip)+"]", previousTrackName)
			}
		}

		clipEffect = clipEffect + replacedEffect
		previousTrackName = currentTrackName
	}
	if clipEffect == "" && len(effects) != 0 {
		clipEffect = effects[0]
	}
	// fmt.Println("---")
	// fmt.Println("ClipAudioMerge", clipEffect)
	// fmt.Println("---")
	return clipEffect
}

func (s *FFMPEG) SetOutputFormat(format string) error {
	s.format = format
	return nil
}

func (s *FFMPEG) SetOutputFps(fps float32) error {
	s.fps = fps
	return nil
}

func (s *FFMPEG) SetDefaultBackground(color string) error {
	s.backgroundColor = color
	return nil
}

func (s *FFMPEG) SetOutputSize(width *int32, height *int32) error {
	s.size.Height = height
	s.size.Width = width

	return nil
}

func (s *FFMPEG) SetOutputResolution(format string) error {
	var height, width int32
	var fps float32
	switch format {
	case "preview":
		width = 512
		height = 288
		fps = 15
	case "mobile":
		width = 640
		height = 360
		fps = 25
	case "sd":
		width = 1024
		height = 576
		fps = 25
	case "hd":
		width = 1280
		height = 720
		fps = 25
	case "1080":
		width = 1920
		height = 1080
		fps = 25
	case "360":
		width = 640
		height = 360
		fps = 25
	case "480":
		width = 848
		height = 480
		fps = 25
	case "540":
		width = 960
		height = 540
		fps = 25
	case "720":
		width = 1280
		height = 720
		fps = 25
	}

	s.size.Height = &height
	s.size.Width = &width
	s.fps = fps

	return nil
}

func (s *FFMPEG) GetResolution() string {
	return cast.ToString(s.size.Width) + "x" + cast.ToString(s.size.Height)
}

func (*FFMPEG) trackName(trackType string, trackNumber int, clipNumber int, tempIndex int) string {
	var prefix = "v"
	if trackType == "audio" {
		prefix = "a"
	} else if trackType == "subtitle" {
		prefix = "s"
	}

	if tempIndex == -1 {
		if clipNumber == -1 {
			return "[" + prefix + "track" + cast.ToString(trackNumber) + "]"
		}
		return "[" + prefix + "track" +
			cast.ToString(trackNumber) +
			"c" +
			cast.ToString(clipNumber) +
			"]"
	}

	if clipNumber == -1 {
		return "[" + prefix + "track" + cast.ToString(trackNumber) + "]"
	}
	return "[" + prefix + "track" +
		cast.ToString(trackNumber) +
		"c" +
		cast.ToString(clipNumber) +
		"p" +
		cast.ToString(tempIndex) +
		"]"
}

func (s *FFMPEG) ClipFiller(trackNumber int, clipNumber int, start float32, length float32) string {
	s.fillerCounter = s.fillerCounter + 1

	return "[filler" +
		cast.ToString(s.fillerCounter) +
		// ":v] trim=start=" +
		"] trim=start=" +
		cast.ToString(start) +
		":end=" +
		cast.ToString(length) +
		", setpts=PTS-STARTPTS ,format=yuva420p " + s.trackName("video", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipImage(sourceClip int, trackNumber int, clipNumber int, start float32, length float32) string {
	return "[" +
		cast.ToString(sourceClip) +
		":v] trim=start=" +
		cast.ToString(start) +
		":end=" +
		cast.ToString(length) +
		", setpts=PTS-STARTPTS ,format=yuva420p " + s.trackName("video", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) computeOverlayPosition(position string) string {
	if position == "" {
		return "x=0:y:0"
	}
	var x = "(main_w-overlay_w)/2"
	var y = "(main_h-overlay_h)/2"

	if strings.Contains(strings.ToLower(position), "left") {
		x = "0"
	}
	if strings.Contains(strings.ToLower(position), "right") {
		x = "main_w-overlay_w"
	}
	if strings.Contains(strings.ToLower(position), "top") {
		y = "0"
	}
	if strings.Contains(strings.ToLower(position), "bottom") {
		y = "main_h-overlay_h"
	}

	return "x=" + x + ":y=" + y
}

func (s *FFMPEG) ClipSubtitleBurn(sourceClip int, trackNumber int, clipNumber int, index int) string {
	return "subtitles='" + s.src[sourceClip].path + "':" +
		"stream_index=" + cast.ToString(index) + s.trackName("subtitle", trackNumber, clipNumber, -1) + "; " +
		"[" + cast.ToString(sourceClip) + ":v] " + s.trackName("subtitle", trackNumber, clipNumber, -1) +
		" overlay " + s.trackName("video", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipFillerOverlay(sourceClip int, trackNumber int, clipNumber int, position string) string {
	s.hasOverlay = true
	s.overlayFillerCounter = s.overlayFillerCounter + 1
	return "[overfiller" + cast.ToString(s.overlayFillerCounter) + "] [" +
		cast.ToString(sourceClip) +
		"] overlay=shortest=1:" + s.computeOverlayPosition(position) + " " + s.trackName("video", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipTrim(sourceClip int, trackNumber int, clipNumber int, start float32, length float32) string {
	return "[" +
		cast.ToString(sourceClip) +
		":v] trim=start=" +
		cast.ToString(start) +
		":end=" +
		cast.ToString(length) +
		", setpts=PTS-STARTPTS " + s.trackName("video", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipResize(sourceClip int, trackNumber int, clipNumber int, scaleRatio float32) string {
	return "[" +
		cast.ToString(sourceClip) +
		"] scale=w=" + cast.ToString(cast.ToFloat32(*s.size.Width)*scaleRatio) + ":" +
		"h=" + cast.ToString(cast.ToFloat32(*s.size.Height)*scaleRatio) + " " +
		s.trackName("video", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipRaw(sourceClip int, trackNumber int, clipNumber int) string {
	// TODO: Handle audio part
	return "[" + cast.ToString(sourceClip) + ":v] concat=n=1:v=1 " +
		// "[" + cast.ToString(clipNumber) + ":a] " +
		s.trackName("video", trackNumber, clipNumber, -1) + ";"
}

func (s *FFMPEG) ClipMerge(sourceClip int, trackNumber int, clipNumber int, effects []string) string {
	var trackName = s.trackName("video", trackNumber, clipNumber, -1)
	var numEffects = len(effects) - 1
	// fmt.Println("ClipMerge: len #", len(effects))

	// fmt.Println("---")
	// fmt.Println("ClipMerge (Before)", effects)
	// fmt.Println("---")

	var clipEffect = ""
	var previousTrackName = ""
	// Replace sources by previous result
	for i, effect := range effects {
		var currentTrackName = s.trackName("video", trackNumber, clipNumber, i)
		replacedEffect := effect
		// fmt.Println("--", i, trackName, currentTrackName, clipEffect)
		// fmt.Println("--", i, trackName, previousTrackName, effect)
		if i != 0 {
			// Replace source from previous
			replacedEffect = strings.ReplaceAll(replacedEffect, "["+cast.ToString(sourceClip)+"]", previousTrackName)
			replacedEffect = strings.ReplaceAll(replacedEffect, "["+cast.ToString(sourceClip)+":v]", previousTrackName[0:len(previousTrackName)-1]+"]")
		}

		if i != numEffects {
			// Replace destination
			replacedEffect = strings.ReplaceAll(replacedEffect, trackName, currentTrackName)
			if previousTrackName != "" {
				// fmt.Println(i, trackName, previousTrackName, currentTrackName)
				replacedEffect = strings.ReplaceAll(replacedEffect, "["+cast.ToString(sourceClip)+"]", previousTrackName)
				replacedEffect = strings.ReplaceAll(replacedEffect, "["+cast.ToString(sourceClip)+":v]", previousTrackName[0:len(previousTrackName)-1]+"]")
			}
		}

		clipEffect = clipEffect + replacedEffect
		previousTrackName = currentTrackName
	}
	if clipEffect == "" && len(effects) != 0 {
		clipEffect = effects[0]
	}
	// fmt.Println("---")
	// fmt.Println("ClipMerge", clipEffect)
	// fmt.Println("---")
	return clipEffect
}

func (s *FFMPEG) GenerateBackground() string {
	return "color=c=" + s.backgroundColor + ":s=" + s.GetResolution() + ":d=999999"
}

func (s *FFMPEG) GenerateFiller(color string) string {
	if color == "" {
		color = "yellow@.0"
	}
	return "color=c=" + color + ":s=" + s.GetResolution() + ":d=999999,format=yuva420p"
}

func (s *FFMPEG) AddTrack(trackNumber int) error {
	newTrack := &FFMPEGTrack{
		video: make([]string, 0),
		audio: make([]string, 0),
	}
	s.tracks = append(s.tracks, *newTrack)
	return nil
}
func (s *FFMPEG) CloseTrack(trackNumber int) error {
	var allVideoClips string
	var allAudioClips string

	// Video
	for _, clip := range s.tracks[trackNumber].video {
		// fmt.Println("CloseTrack", trackNumber, len(s.tracks[trackNumber].video), i, clip == "")
		if clip != "" {
			clipName := clip[strings.LastIndex(clip, "["):]
			if clipName[len(clipName)-1:] == ";" {
				clipName = clipName[0 : len(clipName)-1]
			}
			allVideoClips = allVideoClips + clipName + " "
		}
	}
	// Concat all clips within the track
	if allVideoClips != "" {
		// Add start/end for the clip
		s.tracks[trackNumber].video = append(s.tracks[trackNumber].video, allVideoClips+
			" concat=n="+cast.ToString(len(strings.Split(allVideoClips, " "))-1)+
			":v=1 [vtrack"+cast.ToString(trackNumber)+
			"];")
	}

	// Audio
	for _, clip := range s.tracks[trackNumber].audio {
		// fmt.Println("closeTrack [Audio]", trackNumber, clip)
		// fmt.Println("CloseTrack [Audio]", trackNumber, len(s.tracks[trackNumber].audio), i, clip == "")
		if clip != "" {
			clipName := clip[strings.LastIndex(clip, "["):]
			if clipName[len(clipName)-1:] == ";" {
				clipName = clipName[0 : len(clipName)-1]
			}
			allAudioClips = allAudioClips + clipName + " "
		}
	}
	// Concat all clips within the track
	if allAudioClips != "" {
		// Add start/end for the clip
		s.tracks[trackNumber].audio = append(s.tracks[trackNumber].audio, allAudioClips+
			" amix=inputs="+cast.ToString(len(strings.Split(allAudioClips, " "))-1)+
			" [atrack"+cast.ToString(trackNumber)+
			"];")
	}
	return nil
}

func (s *FFMPEG) generateOutputName() string {
	file, err := ioutil.TempFile("", "*."+s.format)
	if err != nil {
		fmt.Println("Error temp file", err)
		return ""
	}
	defer file.Close()
	s.outputName = file.Name()
	return file.Name()
}

func (s *FFMPEG) ToFFMPEG(queue *RenderQueue) error {
	_ = s.AddDefaultParams()
	_ = s.SetOutputFormat(queue.Data.Output.Format)
	if queue.Data.Output.Fps != nil {
		_ = s.SetOutputFps(*queue.Data.Output.Fps)
	}
	_ = s.SetDefaultBackground(queue.Data.Timeline.Background)

	// Handle Sources
	var sourceClip = 0
	s.fillerCounter = 0

	for trackNumber, track := range queue.Data.Timeline.Tracks {
		var lastStart float32
		var clipNumber = 0

		_ = s.AddTrack(trackNumber)
		for _, clip := range track.Clips {
			// for cIndex, clip := range track.Clips {
			// fmt.Println(cIndex)

			if clip.Start > lastStart {
				// fmt.Println("Add filler clip")
				// Add filler Clip
				_ = s.AddClip(
					trackNumber,
					s.ClipFiller(trackNumber, clipNumber, 0, clip.Start-lastStart),
				)
				clipNumber = clipNumber + 1
			}
			// fmt.Println(sourceClip, s.fillerCounter)
			_ = s.AddSource(queue.LocalResources[sourceClip])

			_ = clip.ToFFMPEG(s, sourceClip, trackNumber, clipNumber)

			sourceClip = sourceClip + 1
			clipNumber = clipNumber + 1
			lastStart = clip.Start + clip.Length
		}
		if lastStart > s.duration {
			s.duration = lastStart
		}
		// TODO: Ensure, we do not need a padding after last clip to match whole timeline
		_ = s.CloseTrack(trackNumber)
	}

	return nil
}

func (s *FFMPEG) ToString() []string {
	var parameters = make([]string, 0)
	if s.defaultParams {
		parameters = append(parameters, "-hide_banner")
		parameters = append(parameters, "-loglevel")
		// parameters = append(parameters, "panic")
		parameters = append(parameters, "debug")
		parameters = append(parameters, "-y")
	}

	// Handle source
	var maxSource = len(s.src) - 1
	for _, source := range s.src {
		parameters = append(parameters, "-i")
		parameters = append(parameters, source.path)
	}
	// Add filler for in between clip
	if s.fillerCounter > 0 {
		parameters = append(parameters, "-f")
		parameters = append(parameters, "lavfi")
		parameters = append(parameters, "-i")
		parameters = append(parameters, s.GenerateFiller("yellow@.0"))
	}

	// Add filler special source for overlay
	if s.overlayFillerCounter > 0 {
		parameters = append(parameters, "-f")
		parameters = append(parameters, "lavfi")
		parameters = append(parameters, "-i")
		parameters = append(parameters, s.GenerateFiller("pink@.0"))
	}

	// Add background source
	parameters = append(parameters, "-f")
	parameters = append(parameters, "lavfi")
	parameters = append(parameters, "-i")
	parameters = append(parameters, s.GenerateBackground())

	// Add FPS output
	parameters = append(parameters, "-r")
	parameters = append(parameters, cast.ToString(s.fps))

	// Handle filter complex
	parameters = append(parameters, "-filter_complex")

	var filterComplex string
	var addedSources int

	// Add filler in between stream
	if s.fillerCounter > 0 {
		addedSources = 1
		for i := 1; i <= s.fillerCounter; i++ {
			filterComplex = filterComplex +
				"[" + cast.ToString(maxSource+addedSources) + "] concat=n=1:v=1,setpts=PTS-STARTPTS,format=yuva420p [filler" + cast.ToString(i) + "];"
		}
	}

	// Add filler overlay stream
	if s.overlayFillerCounter > 0 {
		addedSources = addedSources + 1

		for i := 1; i <= s.overlayFillerCounter; i++ {
			filterComplex = filterComplex +
				"[" + cast.ToString(maxSource+addedSources) + "] concat=n=1:v=1,setpts=PTS-STARTPTS,format=yuva420p [overfiller" + cast.ToString(i) + "];"
		}
	}
	// Add background stream
	addedSources = addedSources + 1
	filterComplex = filterComplex + "[" + cast.ToString(maxSource+addedSources) + "] concat=n=1:v=1,setpts=PTS-STARTPTS,format=yuv420p [bg];"

	// Add all tracks infos
	for _, track := range s.tracks {
		filterComplex = filterComplex + strings.Join(track.video, " ") + strings.Join(track.audio, " ")
	}

	// Overlay of all tracks
	var vTracks = " [bg]"
	var curOverlayName string
	var previousOverlayName string
	for i := len(s.tracks) - 1; i >= 0; i-- {
		curOverlayName = "[overlay" + cast.ToString(math.Abs(cast.ToFloat64(i-(len(s.tracks)-1)))) + "]"
		if previousOverlayName != "" {
			vTracks = vTracks + previousOverlayName
		}
		vTracks = vTracks + "[vtrack" + cast.ToString(i) + "] overlay=shortest=1:x=0:y=0 "
		if i == 0 {
			vTracks = vTracks + "[vtracks];"
		} else {
			vTracks = vTracks + curOverlayName + ";"
		}
		previousOverlayName = curOverlayName
	}
	filterComplex = filterComplex + vTracks

	// Handle audio tracks
	var aTracks string
	for i, track := range s.tracks {
		// fmt.Println(track.audio, len(track.audio))
		if len(track.audio) != 0 && track.audio[0] != "" {
			aTracks = aTracks + "[atrack" + cast.ToString(i) + "] "
		}
	}
	filterComplex = filterComplex + aTracks +
		" amix=inputs=" + cast.ToString(len(strings.Split(aTracks, " "))-1) + " [atracks];"

	// Remove pending ";" if exists
	if filterComplex[len(filterComplex)-1:] == ";" {
		filterComplex = filterComplex[0 : len(filterComplex)-1]
	}
	parameters = append(parameters, filterComplex)

	// Map result
	parameters = append(parameters, "-map")
	parameters = append(parameters, "[vtracks]")
	parameters = append(parameters, "-map")
	parameters = append(parameters, "[atracks]")

	// Handle output
	parameters = append(parameters, "-s")
	parameters = append(parameters, s.GetResolution())

	parameters, err := s.getOutputFormat(parameters)
	if err != nil {
		return make([]string, 0)
	}

	// FIXME: Deprecated field (fps_mode??)
	parameters = append(parameters, "-vsync") // https://stackoverflow.com/questions/18064604/frame-rate-very-high-for-a-muxer-not-efficiently-supporting-it
	parameters = append(parameters, "2")

	var outputName = s.generateOutputName()

	parameters = append(parameters, outputName)

	// FIXME: Dev Only debug
	fmt.Println(strings.Join(parameters, "|"))
	// fmt.Println(outputName)

	return parameters
}

func (s *FFMPEG) getOutputFormat(parameters []string) ([]string, error) {
	switch s.format {
	case "mp4":
		parameters = append(parameters, "-codec:v")
		parameters = append(parameters, "libx264")
		parameters = append(parameters, "-preset")
		parameters = append(parameters, "ultrafast") // slow
		// parameters = append(parameters, "-codec:v")
		// parameters = append(parameters, "vp9")
	default:
		return make([]string, 0), errors.New("format not handled")
	}

	return parameters, nil
}
