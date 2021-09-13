package html

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"plot-training/pkg/config"
	"plot-training/pkg/sample"
	"plot-training/pkg/vanilla"
)

type VO struct {
	Samples sample.XY
	Records []vanllia.LostAndW
}

func Render3dClass(fileName string) {
	pwd, err := os.Getwd()
	log.Printf("pwd:%v", pwd)

	tpl := template.Must(template.ParseGlob(config.PojectRoot + "/pkg/html/tpl/*.gohtml"))
	filePath := config.PojectRoot + "/" + config.StaticFolderName + "/" + fileName
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

	tpl.ExecuteTemplate(write, "3d-class.gohtml", nil)
	write.Flush()
	log.Printf("File Created Successfully %s \n", filePath)
}
func Render(samples sample.XY, records []vanllia.LostAndW, fileName string) {
	pwd, err := os.Getwd()
	log.Printf("pwd:%v", pwd)

	tpl := template.Must(template.ParseGlob(config.PojectRoot + "/pkg/html/tpl/*.gohtml"))
	filePath := config.PojectRoot + "/" + config.StaticFolderName + "/" + fileName
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

	tpl.ExecuteTemplate(write, "plot.gohtml", VO{samples, records})
	tpl.ExecuteTemplate(write, "3d-class.gohtml", nil)
	write.Flush()
	log.Printf("File Created Successfully %s \n", filePath)
}
