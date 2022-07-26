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

type LocalResource struct {
	Downloaded       time.Time `json:"downloaded"`
	OriginalURL      string
	LocalURL         string
	KeepCache        bool
	IsRemoteResource bool
	Used             []*LocalResourceTrackInfo
}

type LocalResourceTrackInfo struct {
	Track   int
	Clip    int
	Handled bool // Temp Property until all resource are handled (this allow fetching without handling)
}
