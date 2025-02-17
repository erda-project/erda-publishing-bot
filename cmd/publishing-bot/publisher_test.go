/*
Copyright 2018 The Kubernetes Authors.

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

package main

import (
	"reflect"
	"testing"
)

func Test_updateEnv(t *testing.T) {
	tests := []struct {
		name string
		env  []string
		want []string
	}{
		//{"no PATH", []string{"FOO=42"}, []string{"FOO=42", "PATH=/usr/local/go"}},
		{"no PATH", []string{"FOO=42", "PATH=/bin", "BAR=1"}, []string{"FOO=42", "PATH=/usr/local/go:/bin", "BAR=1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateEnv(tt.env, "PATH", prependPath("/usr/local/go"), "/usr/local/go"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_publish_run(t *testing.T) {
}