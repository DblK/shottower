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

import "time"

type RenderQueue struct {
	// The id of the render task in UUID format.
	ID string `json:"id"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`

	Data     Edit `json:"data"`
	Status   RenderResponseStatus
	FileName string

	InternalStatus RenderResponseStatus
	LocalResources []string
	FFMPEGCommand  FFMPEGCommand
}
