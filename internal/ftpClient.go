package iftp

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

type CFtpClient struct {
	serverUrl string
	userID    string
	pasword   string
	ftpConn   *ftp.ServerConn
}

func NewFtpClient() *CFtpClient {
	return &CFtpClient{}
}

func (pInst *CFtpClient) Open(url string) error {
	pInst.serverUrl = url
	client, err := ftp.Dial(url, ftp.DialWithTimeout(15*time.Second))
	if err != nil {
		return err
	}
	pInst.ftpConn = client

	return nil
}
func (pInst *CFtpClient) Login(user, password string) error {
	pInst.userID = user
	pInst.pasword = password

	err := pInst.ftpConn.Login(user, password)

	return err
}
func (pInst *CFtpClient) ReadFile(filename string) ([]byte, error) {
	fileHandler, err := pInst.ftpConn.Retr(filename)
	if err != nil {
		return nil, err
	}
	defer fileHandler.Close()

	buf, err := io.ReadAll(fileHandler)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
func (pInst *CFtpClient) UploadData(filename string, data []byte) error {
	err := pInst.ftpConn.Stor(filename, bytes.NewBuffer(data))

	return err
}
func (pInst *CFtpClient) UploadLocalFile(filename, localFilePath string) error {
	fileHandler, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	err = pInst.ftpConn.Stor(filename, fileHandler)
	return err
}

func (pInst *CFtpClient) Walk(dir string) (string, error) {
	fileList, err := pInst.ftpConn.List(dir)
	if err != nil {
		return "", err
	}
	strRet := ""
	for _, fileInfo := range fileList {
		strRet = strRet + fileInfo.Name + ";"
	}

	return strRet, nil
}
