package wjpalftp

import (
	"fmt"
	"testing"
)

func TestWJAPL_NewFtpClient(t *testing.T) {
	inst := WJAPL_NewFtpClient()
	err := inst.Open("ip:port.....")
	if err != nil {
		t.Errorf("ftp client open error: " + err.Error())
		return
	} else {
		fmt.Println("connected")
	}
	err = inst.Login("user...", "****")
	if err != nil {
		t.Errorf("ftp login error: " + err.Error())
		return
	} else {
		fmt.Println("login successful")
	}

	data1, err := inst.Walk("/")
	if err != nil {
		t.Errorf("ftp walk error: " + err.Error())
		return
	} else {
		fmt.Println(data1)
	}

	err = inst.Mkdir("/selfdata/")
	if err != nil {
		fmt.Println("mkdir error: ", err)
		fmt.Println("continue.......")
	}

	data2, err := inst.ReadFile("/selfdata/test1.txt")
	if err != nil {
		t.Errorf("ftp read file error: " + err.Error())
		return
	} else {
		fmt.Println("load data: " + string(data2))
	}

	/*err = inst.UploadData("/selfdata/test1.txt", []byte("test data upload 1"))
	if err != nil {
		t.Errorf("ftp upload data error: " + err.Error())
		return
	} else {
		fmt.Println("upload successful")
	}*/

}
