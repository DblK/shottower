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
package openapi_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	openapi "github.com/DblK/shottower/go"
)

var _ = Describe("Ffmpeg", func() {
	Describe("ClipMerge", func() {
		var ff openapi.FFMPEGCommand
		BeforeEach(func() {
			ff = openapi.NewFFMPEGCommand()
		})

		It("Test merge with only one effect", func() {
			var sourceClip = 0
			var trackNumber = 0
			var clipNumber = 0
			var effects []string
			effects = append(effects, "[0:v] trim=start=0:end=4.96, setpts=PTS-STARTPTS [vtrack0c0];")

			res := ff.ClipMerge(sourceClip, trackNumber, clipNumber, effects)
			Expect(res).To(Equal("[0:v] trim=start=0:end=4.96, setpts=PTS-STARTPTS [vtrack0c0];"))
		})

		It("Test merge with only two effects (video only)", func() {
			var sourceClip = 0
			var trackNumber = 0
			var clipNumber = 0
			var effects []string
			effects = append(effects, "[0:v] trim=start=0:end=4.96, setpts=PTS-STARTPTS [vtrack0c0];")
			effects = append(effects, "[0:v] scale=w=1024:h=576 [vtrack0c0];")

			res := ff.ClipMerge(sourceClip, trackNumber, clipNumber, effects)
			Expect(res).To(Equal("[0:v] trim=start=0:end=4.96, setpts=PTS-STARTPTS [vtrack0c0p0];[vtrack0c0p0] scale=w=1024:h=576 [vtrack0c0];"))
		})

		It("Test merge with only three effects (video only)", func() {
			var sourceClip = 1
			var trackNumber = 0
			var clipNumber = 1
			var effects []string

			effects = append(effects, "[1:v] trim=start=5:end=9.96, setpts=PTS-STARTPTS [vtrack0c1];")
			effects = append(effects, "[1] scale=w=448:h=252 [vtrack0c1];")
			effects = append(effects, "[filler] [1:v] overlay=shortest [vtrack0c1];")

			res := ff.ClipMerge(sourceClip, trackNumber, clipNumber, effects)
			Expect(res).To(Equal("[1:v] trim=start=5:end=9.96, setpts=PTS-STARTPTS [vtrack0c1p0];[vtrack0c1p0] scale=w=448:h=252 [vtrack0c1p1];[filler] [vtrack0c1p1] overlay=shortest [vtrack0c1];"))
		})

		It("Test merge with only two effects (full stream)", func() {
			var sourceClip = 0
			var trackNumber = 0
			var clipNumber = 0
			var effects []string
			effects = append(effects, "[0] trim=start=0:end=4.96, setpts=PTS-STARTPTS [vtrack0c0];")
			effects = append(effects, "[0] scale=w=1024:h=576 [vtrack0c0];")

			res := ff.ClipMerge(sourceClip, trackNumber, clipNumber, effects)
			Expect(res).To(Equal("[0] trim=start=0:end=4.96, setpts=PTS-STARTPTS [vtrack0c0p0];[vtrack0c0p0] scale=w=1024:h=576 [vtrack0c0];"))
		})
	})
	Describe("SetOutputResolution", func() {
		var ff openapi.FFMPEGCommand
		BeforeEach(func() {
			ff = openapi.NewFFMPEGCommand()
		})

		It("Set Output resolution to 'preview'", func() {
			_ = ff.SetOutputResolution("preview")
			Expect(ff.GetResolution()).To(Equal("512x288"))
		})
		It("Set Output resolution to 'mobile'", func() {
			_ = ff.SetOutputResolution("mobile")
			Expect(ff.GetResolution()).To(Equal("640x360"))
		})
		It("Set Output resolution to 'sd'", func() {
			_ = ff.SetOutputResolution("sd")
			Expect(ff.GetResolution()).To(Equal("1024x576"))
		})
		It("Set Output resolution to 'hd'", func() {
			_ = ff.SetOutputResolution("hd")
			Expect(ff.GetResolution()).To(Equal("1280x720"))
		})
		It("Set Output resolution to '1080'", func() {
			_ = ff.SetOutputResolution("1080")
			Expect(ff.GetResolution()).To(Equal("1920x1080"))
		})
		It("Set Output resolution to '360'", func() {
			_ = ff.SetOutputResolution("360")
			Expect(ff.GetResolution()).To(Equal("640x360"))
		})
		It("Set Output resolution to '480'", func() {
			_ = ff.SetOutputResolution("480")
			Expect(ff.GetResolution()).To(Equal("848x480"))
		})
		It("Set Output resolution to '540'", func() {
			_ = ff.SetOutputResolution("540")
			Expect(ff.GetResolution()).To(Equal("960x540"))
		})
		It("Set Output resolution to '720'", func() {
			_ = ff.SetOutputResolution("720")
			Expect(ff.GetResolution()).To(Equal("1280x720"))
		})
	})
	Describe("SetOutputSize", func() {
		var ff openapi.FFMPEGCommand
		BeforeEach(func() {
			ff = openapi.NewFFMPEGCommand()
		})

		It("Set Output size", func() {
			width := int32(30)
			height := int32(20)
			_ = ff.SetOutputSize(&width, &height)
			Expect(ff.GetResolution()).To(Equal("30x20"))
		})
	})
	FDescribe("OverlayAllTracks", func() {
		var ff openapi.FFMPEGCommand
		var missingVideoTracks []string
		BeforeEach(func() {
			ff = openapi.NewFFMPEGCommand()
			missingVideoTracks = []string{}
		})

		It("Overlay one handled tracks out of one", func() {

			_ = ff.AddTrack(0)
			res := ff.OverlayAllTracks(missingVideoTracks)
			Expect(res).To(Equal(" [bg][vtrack0] overlay=shortest=1:x=0:y=0 [vtracks];"))
		})
		It("Overlay one handled tracks out of two", func() {
			missingVideoTracks = append(missingVideoTracks, "[vtrack0]")
			_ = ff.AddTrack(0)
			_ = ff.AddTrack(1)
			res := ff.OverlayAllTracks(missingVideoTracks)
			Expect(res).To(Equal(" [bg][vtrack1] overlay=shortest=1:x=0:y=0 [vtracks];"))
		})
		It("Overlay two handled tracks out of two", func() {
			_ = ff.AddTrack(0)
			_ = ff.AddTrack(1)
			res := ff.OverlayAllTracks(missingVideoTracks)
			Expect(res).To(Equal(" [bg][vtrack1] overlay=shortest=1:x=0:y=0 [overlay0];[overlay0][vtrack0] overlay=shortest=1:x=0:y=0 [vtracks];"))
		})
		It("Overlay two handled tracks out of three", func() {
			missingVideoTracks = append(missingVideoTracks, "[vtrack1]")
			_ = ff.AddTrack(0)
			_ = ff.AddTrack(1)
			_ = ff.AddTrack(2)
			res := ff.OverlayAllTracks(missingVideoTracks)
			Expect(res).To(Equal(" [bg][vtrack2] overlay=shortest=1:x=0:y=0 [overlay0];[overlay0][vtrack0] overlay=shortest=1:x=0:y=0 [vtracks];"))
		})
	})
	Describe("SetDefaultBackground/GenerateBackground", func() {
		var ff openapi.FFMPEGCommand
		BeforeEach(func() {
			ff = openapi.NewFFMPEGCommand()
			_ = ff.SetOutputResolution("480")
		})

		It("Generate a background", func() {
			_ = ff.SetDefaultBackground("#FFFFFF")

			res := ff.GenerateBackground()
			Expect(res).To(Equal("color=c=#FFFFFF:s=848x480:d=999999"))
		})
	})
	Describe("GenerateFiller", func() {
		var ff openapi.FFMPEGCommand
		BeforeEach(func() {
			ff = openapi.NewFFMPEGCommand()
			_ = ff.SetOutputResolution("480")
		})

		It("Generate a background with a color", func() {
			res := ff.GenerateFiller("#FFFFFF")
			Expect(res).To(Equal("color=c=#FFFFFF:s=848x480:d=999999,format=yuva420p"))
		})
		It("Generate a background without color", func() {
			res := ff.GenerateFiller("")
			Expect(res).To(Equal("color=c=yellow@.0:s=848x480:d=999999,format=yuva420p"))
		})
	})
	Describe("ClipAudioMerge", func() {
		var ff openapi.FFMPEGCommand
		BeforeEach(func() {
			ff = openapi.NewFFMPEGCommand()
		})

		It("Test merge with only one effect", func() {
			var sourceClip = 0
			var trackNumber = 0
			var clipNumber = 0
			var effects []string
			effects = append(effects, "[0:a] atrim=start=0:end=4.96, asetpts=PTS-STARTPTS [atrack0c0];")

			res := ff.ClipAudioMerge(sourceClip, trackNumber, clipNumber, effects)
			Expect(res).To(Equal("[0:a] atrim=start=0:end=4.96, asetpts=PTS-STARTPTS [atrack0c0];"))
		})

		It("Test merge with only two effects (audio only)", func() {
			var sourceClip = 0
			var trackNumber = 0
			var clipNumber = 0
			var effects []string
			effects = append(effects, "[0:a] atrim=start=0:end=4.96, asetpts=PTS-STARTPTS [atrack0c0];")
			effects = append(effects, "[0:a] volume=0.5 [atrack0c0];")

			res := ff.ClipAudioMerge(sourceClip, trackNumber, clipNumber, effects)
			Expect(res).To(Equal("[0:a] atrim=start=0:end=4.96, asetpts=PTS-STARTPTS [atrack0c0p0];[atrack0c0p0] volume=0.5 [atrack0c0];"))
		})

		It("Test merge with only three effects (audio only)", func() {
			var sourceClip = 1
			var trackNumber = 0
			var clipNumber = 1
			var effects []string

			effects = append(effects, "[1:a] atrim=start=5:end=9.96, asetpts=PTS-STARTPTS [atrack0c1];")
			effects = append(effects, "[1] volume=0.5 [atrack0c1];")
			effects = append(effects, "[filler] [1:a] amix=inputs=2 [atrack0c1];")

			res := ff.ClipAudioMerge(sourceClip, trackNumber, clipNumber, effects)
			Expect(res).To(Equal("[1:a] atrim=start=5:end=9.96, asetpts=PTS-STARTPTS [atrack0c1p0];[atrack0c1p0] volume=0.5 [atrack0c1p1];[filler] [atrack0c1p1] amix=inputs=2 [atrack0c1];"))
		})
	})
})
