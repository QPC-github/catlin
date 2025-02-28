// Copyright © 2021 The Tekton Authors.
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
package validate

import (
	"testing"

	"github.com/tektoncd/catlin/pkg/app"
	"github.com/tektoncd/catlin/pkg/test"
	"gotest.tools/v3/assert"
)

func TestValidate(t *testing.T) {

	testParams := []struct {
		name      string
		args      []string
		wantError bool
		want      string
	}{
		{
			name:      "single filepath",
			args:      []string{"./testdata/task/maven/0.1"},
			wantError: false,
			want:      "",
		},
		{
			name:      "single filepath - with directory versioning flag",
			args:      []string{"./testdata/task/maven/0.1", "--versioning", "directory"},
			wantError: false,
			want:      "",
		},
		{
			name:      "single filepath git versioning",
			args:      []string{"./testdata/task/maven-git-versioning", "--versioning", "git"},
			wantError: false,
			want:      "",
		},
		{
			name:      "multiple filepath",
			args:      []string{"./testdata/task/maven/0.1", "./testdata/task/npm/0.1"},
			wantError: false,
			want:      "",
		},
		{
			name:      "multiple filepath git versioning",
			args:      []string{"./testdata/task/maven-git-versioning", "./testdata/task/npm-git-versioning", "--versioning", "git"},
			wantError: false,
			want:      "",
		},
	}

	for _, tp := range testParams {
		t.Run(tp.name, func(t *testing.T) {
			cli := app.New()
			validate := Command(cli)
			got, err := test.ExecuteCommand(validate, tp.args...)
			if !tp.wantError {
				if err != nil {
					t.Errorf("Unexpected error")
				}
				assert.Equal(t, tp.want, got)
			}
		})
	}
}

func TestValidateError(t *testing.T) {

	testParams := []struct {
		name      string
		args      []string
		wantError bool
		want      string
	}{
		{
			name:      "single filepath",
			args:      []string{"./testdata/task/black/0.1"},
			wantError: true,
			want:      "Error: ./testdata/task/black/0.1/black.yaml failed validation\n",
		},
		{
			name:      "versioning mismatch - git versioning",
			args:      []string{"./testdata/task/maven/0.1", "--versioning", "git"},
			wantError: true,
			want:      "Error: ./testdata/task/maven/0.1/maven.yaml failed validation\n",
		},
		{
			name:      "versioning mismatch - directory versioning",
			args:      []string{"./testdata/task/maven-git-versioning", "--versioning", "directory"},
			wantError: true,
			want:      "Error: ./testdata/task/maven-git-versioning/maven-git-versioning.yaml failed validation\n",
		},
	}

	for _, tp := range testParams {
		t.Run(tp.name, func(t *testing.T) {
			cli := app.New()
			validate := Command(cli)
			got, err := test.ExecuteCommand(validate, tp.args...)
			if tp.wantError {
				if err == nil {
					t.Errorf("Unexpected error")
				}
				assert.Equal(t, tp.want, got)
			}
		})
	}
}
