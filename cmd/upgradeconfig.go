// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/palantir/godel-conjure-plugin/v4/conjureplugin/config"
	"github.com/palantir/godel/v2/framework/pluginapi"
)

var upgradeConfigCmd = pluginapi.CobraUpgradeConfigCmd(func(cfgBytes []byte) ([]byte, error) {
	return config.UpgradeConfig(cfgBytes)
})

func init() {
	rootCmd.AddCommand(upgradeConfigCmd)
}
