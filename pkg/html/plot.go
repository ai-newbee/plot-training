package html

import (
	"bufio"
	"dl-base/pkg/config"
	"dl-base/pkg/sample"
	"html/template"
	"log"
	"os"
)

func Render(samples sample.XY) {
	pwd, err := os.Getwd()
	log.Printf("pwd:%v", pwd)

	tpl := template.Must(template.ParseGlob(config.PojectRoot + "/pkg/html/tpl/*.gohtml"))
	filePath := config.PojectRoot + "/" + config.StaticFolderName + "/plot.html"
	f, err := os.Create(filePath)

	defer f.Close()
	if err != nil {
		log.Fatalln("f, err := os.Create(filePath)", err)
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	tpl.ExecuteTemplate(write, "plot.gohtml", samples)
	write.Flush()
	log.Printf("File Created Successfully %s \n", filePath)
}
