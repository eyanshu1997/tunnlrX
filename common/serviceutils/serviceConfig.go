package serviceutils

type LogConfig struct {
	Level        string `json:"level"`
	FilePath     string `json:"file_path"`
	IncludeStdio bool   `json:"include_stdio"`
}

type ServiceConfig struct {
	Log LogConfig `json:"log"`
}
