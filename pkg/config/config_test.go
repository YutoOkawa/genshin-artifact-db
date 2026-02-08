package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.Port != DefaultPort {
		t.Errorf("expected port %s, got %s", DefaultPort, cfg.Port)
	}

	if cfg.DataFilePath != DefaultDataFilePath {
		t.Errorf("expected data file path %s, got %s", DefaultDataFilePath, cfg.DataFilePath)
	}
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name           string
		configContent  string
		expectedConfig *Config
		expectError    bool
	}{
		{
			name: "ShouldLoadConfigSuccessfully",
			configContent: `port: ":9090"
data_file_path: "/custom/path/data.json"
`,
			expectedConfig: &Config{
				Port:         ":9090",
				DataFilePath: "/custom/path/data.json",
			},
			expectError: false,
		},
		{
			name: "ShouldUseDefaultValuesForEmptyFields",
			configContent: `port: ""
data_file_path: ""
`,
			expectedConfig: &Config{
				Port:         DefaultPort,
				DataFilePath: DefaultDataFilePath,
			},
			expectError: false,
		},
		{
			name:          "ShouldUseDefaultConfigForEmptyFile",
			configContent: "",
			expectedConfig: &Config{
				Port:         DefaultPort,
				DataFilePath: DefaultDataFilePath,
			},
			expectError: false,
		},
		{
			name: "ShouldUseDefaultPortWhenOnlyDataFilePathIsSet",
			configContent: `data_file_path: "/custom/data.json"
`,
			expectedConfig: &Config{
				Port:         DefaultPort,
				DataFilePath: "/custom/data.json",
			},
			expectError: false,
		},
		{
			name:           "ShouldReturnErrorForInvalidYAML",
			configContent:  "invalid: yaml: content:",
			expectedConfig: nil,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 一時ディレクトリに設定ファイルを作成
			tmpDir := t.TempDir()
			configPath := filepath.Join(tmpDir, "config.yaml")

			if err := os.WriteFile(configPath, []byte(tt.configContent), 0644); err != nil {
				t.Fatalf("failed to write config file: %v", err)
			}

			cfg, err := LoadConfig(configPath)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if diff := cmp.Diff(tt.expectedConfig, cfg); diff != "" {
				t.Errorf("config mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLoadConfigFileNotFound(t *testing.T) {
	cfg, err := LoadConfig("/non/existent/path/config.yaml")

	if err != nil {
		t.Errorf("expected no error for non-existent file, got: %v", err)
	}

	if cfg.Port != DefaultPort {
		t.Errorf("expected default port %s, got %s", DefaultPort, cfg.Port)
	}

	if cfg.DataFilePath != DefaultDataFilePath {
		t.Errorf("expected default data file path %s, got %s", DefaultDataFilePath, cfg.DataFilePath)
	}
}
