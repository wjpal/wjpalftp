package wjpalftp

import iftp "github.com/wjpal/wjpalftp/internal"

type WJPALIFtp interface {
	Open(url string) error
	Login(user, password string) error
	Mkdir(dirPath string) error
	FileSize(filename string) (int64, error)
	ReadFile(filename string) ([]byte, error)
	ReadFileFrom(filename string, offset uint64) ([]byte, error)
	UploadData(filename string, data []byte) error
	UploadLocalFile(filename, localFilePath string) error
	UploadDataFrom(filename string, data []byte, offset uint64) error
	AppendData(filename string, data []byte) error
	Rename(filenameold, filenamenew string) error
	RemoveDir(filename string) error
	RemoveDirRecur(filename string) error
	Delete(filename string) error
	Walk(dir string) (string, error) // list all file in dir. separator ;

}

func WJAPL_NewFtpClient() WJPALIFtp {
	return iftp.NewFtpClient()
}
