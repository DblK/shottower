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

type FFMPEGCommand interface {
	AddSource(string) error
	AddDefaultParams() error
	ToString() []string
	AddTrack(int) error
	CloseTrack(int) error
	AddClip(int, string) error
	ClipTrim(int, int, int, float32, float32) string
	ClipImage(int, int, int, float32, float32) string
	ClipMerge(int, int, int, []string) string
	ClipRaw(int, int, int) string
	ClipResize(int, int, int, float32) string
	ClipFillerOverlay(int, int, int, string) string
	ClipSubtitleBurn(int, int, int, int) string

	AddAudioClip(int, string) error
	ClipAudioVolume(int, int, int, float32) string
	ClipAudioDelay(int, int, int, float32) string
	ClipAudioTrim(int, int, int, float32, float32) string
	ClipAudioMerge(int, int, int, []string) string

	SetOutputResolution(string) error
	SetOutputSize(*int32, *int32) error
	SetOutputFormat(string) error
	SetDefaultBackground(string) error
	GetResolution() string
	GenerateFiller(string) string
	GenerateBackground() string
	ToFFMPEG(*RenderQueue) error
	GetOutputName() string
	GetDuration() float32
}
