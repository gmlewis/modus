package modinfo

import (
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-version"
)

func TestCollectModuleInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		sourceDir string
		want      *ModuleInfo
	}{
		{
			name:      "simple",
			sourceDir: "testdata/simple",
			want: &ModuleInfo{
				ModulePath:      "gmlewis/modus/examples/simple-example",
				ModusSDKVersion: version.Must(version.NewVersion("40.11.0")),
			},
		},
		{
			name:      "simple with path",
			sourceDir: "testdata/examples/simple",
			want: &ModuleInfo{
				ModulePath:      "gmlewis/modus/examples/simple-example",
				ModusSDKVersion: version.Must(version.NewVersion("50.1.0")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &config.Config{SourceDir: tt.sourceDir}
			got, err := CollectModuleInfo(config)
			if err != nil {
				t.Fatalf("CollectModuleInfo() error = %v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CollectModuleInfo() mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
