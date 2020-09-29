package main

import (
	"fmt"
	"os"
	pdf "github.com/unidoc/unidoc/pdf/model"
	"flag"
)

var (
	inputPath  string
	outputPath string
	from       int
	to         int
)

func init() {
	flag.StringVar(&inputPath, "i", "", "被切割的文件路径")
	flag.IntVar(&from, "f", 0, "起始页")
	flag.IntVar(&to, "t", 0, "结束页")
	flag.StringVar(&outputPath, "o", "", "输出文件路径")

}

/*
https://unidoc.io/examples/pages/pdf-split

切割文件

go run pdf-split.go  -i "K:\ProgramData\百度云同步盘\BK\《SRE》\SRE.pdf" -f=53 -t=64 -o "K:\ProgramData\百度云同步盘\BK\《SRE》\SRE-02 Google 生产环境：SRE 视角.pdf"
pdf-split  -i "K:\ProgramData\百度云同步盘\BK\《SRE》\SRE.pdf" -f=53 -t=64 -o "K:\ProgramData\百度云同步盘\BK\《SRE》\SRE - 02 Google 生产环境：SRE 视角.pdf"
 */
func main() {
	flag.Parse()

	if inputPath == "" || from <= 0 || to <= 0 || outputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if from > to {
		fmt.Printf("起始页(-f) 不能大于 结束页(-t)")
		os.Exit(1)
	}

	err := splitPdf(inputPath, outputPath, from, to)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputPath)
}

/*
inputPath  文件地址
outputPath 切割后后文件地址
pageFrom   起始页
pageTo     结束页
 */
func splitPdf(inputPath string, outputPath string, pageFrom int, pageTo int) error {
	pdfWriter := pdf.NewPdfWriter()

	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	if numPages < pageTo {
		return err
	}

	for i := pageFrom; i <= pageTo; i++ {
		pageNum := i

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		err = pdfWriter.AddPage(page)
		if err != nil {
			return err
		}
	}

	fWrite, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}
