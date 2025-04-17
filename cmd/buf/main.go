// Copyright 2020-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (

	// _ "unsafe"

	"github.com/bufbuild/buf/private/buf/cmd/buf"
)

// //go:linkname checkd github.com/bufbuild/buf/private/usage.check
// func checkd() error {
// 	return nil
// }

// //go:linkname myReplacement github.com/bufbuild/buf/private/usage.check
// func myReplacement() error {
// 	return nil
// }

// //go:linkname walltime runtime.walltime
// func walltime() (sec int64, nsec int32)

func main() {
	// os.Args = append(os.Args, "dep", "update", "../..")
	// os.Args = append(os.Args, "generate", "--template=../../buf.gen.yaml")
	buf.Main("buf")
}
