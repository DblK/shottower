/*
shottower
Copyright (C) 2022-2023 Rémy Boulanouar

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
package config

type EndpointType int64

const (
	V1 EndpointType = iota
	Stage
)

func (s EndpointType) String() string {
	switch s {
	case V1:
		return "v1"
	case Stage:
		return "stage"
	default:
		return "unknown"
	}
}

type ShottowerConfiguration interface {
	GetRootURL() string
	GetRenderBaseURL() string
	GetServeBaseURL() string
	GetDownloadBaseURL() string
	GetEndpointType() EndpointType
}

type ShottowerConfig struct {
	rootURL      string
	endpointType EndpointType
}

func NewShottowerConfig(baseURL string, endpointType EndpointType) ShottowerConfiguration {
	return &ShottowerConfig{
		rootURL:      baseURL,
		endpointType: endpointType,
	}
}

func (s *ShottowerConfig) GetRootURL() string {
	return s.rootURL
}

func (s *ShottowerConfig) GetRenderBaseURL() string {
	return s.rootURL + "/" + s.endpointType.String()
}

func (s *ShottowerConfig) GetServeBaseURL() string {
	return s.rootURL + "/serve/" + s.endpointType.String()
}

func (s *ShottowerConfig) GetDownloadBaseURL() string {
	return s.rootURL + "/dl/" + s.endpointType.String()
}

func (s *ShottowerConfig) GetEndpointType() EndpointType {
	return s.endpointType
}
