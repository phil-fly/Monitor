package Monitor

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"shadow/define"
)

//面包屑监控

func Monitor_server(value interface{}) error {
	switch value.(type) {
	case define.ServiceNode :break
	default :
		return errors.New("传参错误")
	}
	honeyconfig := value.(define.ServiceNode)
	var webAddr string
	if honeyconfig.Port != "" {
		webAddr = "0.0.0.0:"+honeyconfig.Port
	}

	http.HandleFunc("/upload", upload)
	http.ListenAndServe(webAddr, nil)
	return nil
}

func upload(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	if r.Method == "POST"{
		//把上传的文件存储在内存和临时文件中
		r.ParseMultipartForm(32 << 20)
		//获取文件句柄，然后对文件进行存储等处理
		file, handler, err := r.FormFile("uploadfile")
		if err != nil{
			fmt.Println("form file err: ", err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		//创建上传的目的文件
		f, err := os.OpenFile("./Monitor/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil{
			fmt.Println("open file err: ", err)
			return
		}
		defer f.Close()
		//拷贝文件
		io.Copy(f, file)
	}
}