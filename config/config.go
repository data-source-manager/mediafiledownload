package config

type Config struct {
	Name          string      `json:"name,optional"`
	Mode          string      `json:"mode,optional"`
	Log           LogConfig   `json:"Log"`
	FileSavePath  string      `json:"file_save_path,optional"`
	FileSaveModel string      `json:"file_save_model,optional"`
	Redis         RedisConfig `json:"Redis"`
	Oss           OssConfig   `json:"Oss,optional"`
}

type LogConfig struct {
	Model      string `json:"model,omitempty" default:"dev"`
	Level      string ` json:"level" default:"warn" `
	FileName   string ` json:"file_name"`
	MaxSize    int    ` json:"max_size"`
	MaxAge     int    ` json:"max_age"`
	MaxBackups int    ` json:"max_backups"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Pwd      string `json:"pwd"`
	Db       int    `json:"db"`
	MediaKey string `json:"mediaKey"`
}

type OssConfig struct {
	AccessKeyId     string `json:"access_key_id,optional"`
	AccessKeySecret string `json:"access_key_secret,optional"`
	Endpoint        string `json:"endpoint,optional"`
	UniqueKey       string `json:"unique_key,optional"`
}
