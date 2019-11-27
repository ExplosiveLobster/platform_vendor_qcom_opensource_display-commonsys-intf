// Copyright (C) 2019 Pavel Dubrova <pashadubrova@gmail.com>
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

package gralloc_selector

import (
	"android/soong/android"
	"android/soong/cc"
)

func init() {
	android.RegisterModuleType("gralloc_selector", gralloc_DefaultsFactory)
}

func gralloc_DefaultsFactory() android.Module {
	module := cc.DefaultsFactory()
	android.AddLoadHook(module, gralloc_Defaults)

	return module
}

func gralloc_Defaults(ctx android.LoadHookContext) {
	type props struct {
		Export_include_dirs []string
	}

	p := &props{}
	p.Export_include_dirs = exportDir(ctx)

	ctx.AppendProperties(p)
}

func exportDir(ctx android.BaseContext) []string {
	var export_include_dirs []string
	gralloc_ver := ctx.Config().VendorConfig("gralloc").Bool("use_v1")

	if gralloc_ver {
		export_include_dirs = append(export_include_dirs, "gralloc_headers_legacy")
	} else {
		export_include_dirs = append(export_include_dirs, "gralloc_headers")
	}

	return export_include_dirs
}
