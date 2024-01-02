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

var conf *ShottowerConfig

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

func Init(baseURL string, endpointType EndpointType) error {

	// viper.SetConfigName("config")

	// https://github.com/TakasBU/TakasBU/blob/main/initializers/loadEnv.go
	// https://github.com/laixhe/goimg/blob/c27cdb95bdf9b09f36589c2d465c7a6ebdd44b92/config/config.go
	// https://github.com/Bikram-Gyawali/LearningGoLang/blob/17f5c3a1048f4857e31652d3a139253be68ec9a9/day24_viper/utils/config.go
	// https://articles.wesionary.team/environment-variable-configuration-in-your-golang-project-using-viper-4e8289ef664d

	conf = &ShottowerConfig{
		rootURL:      baseURL,
		endpointType: endpointType,
	}

	return nil
}

func Get() *ShottowerConfig {
	return conf
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
