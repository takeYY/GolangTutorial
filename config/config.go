package config

import "os"

// アプリケーションの設定を格納する構造体
type (
	server struct {
		Port string
	}
	Database struct {
		Addr     string
		User     string
		Password string
		Name     string
	}
	Config struct {
		Server         server
		ReaderDatabase Database
		WriterDatabase Database
	}
)

// 設定ファイルを読み込み、Config構造体を返す
func LoadConfig() (*Config, error) {
	config := &Config{
		Server: server{
			Port: os.Getenv("SERVER_PORT"),
		},
		ReaderDatabase: Database{
			Addr:     os.Getenv("READER_DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PWD"),
			Name:     os.Getenv("DB_NAME"),
		},
		WriterDatabase: Database{
			Addr:     os.Getenv("WRITER_DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PWD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}

	return config, nil
}
