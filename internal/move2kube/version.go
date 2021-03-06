/*
Copyright IBM Corporation 2020

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package move2kube

import (
	"gopkg.in/yaml.v3"

	"github.com/konveyor/move2kube/types/info"
)

// GetVersion returns the version
func GetVersion(long bool) string {
	if !long {
		return info.GetVersion()
	}
	v := info.GetVersionInfo()
	ver, _ := yaml.Marshal(v)
	return string(ver)
}
