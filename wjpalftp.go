package wjpalftp

import iftp "github.com/wjpal/wjpalftp/internal"

type WJPALIFtp interface {
	Open(url string) error
	Login(user, password string) error
	ReadFile(filename string) ([]byte, error)
	UploadData(filename string, data []byte) error
	UploadLocalFile(filename, localFilePath string) error
	Walk(dir string) (string, error) // list all file in dir. separator ;
}

func WJAPL_NewFtpClient() WJPALIFtp {
	return iftp.NewFtpClient()
}
