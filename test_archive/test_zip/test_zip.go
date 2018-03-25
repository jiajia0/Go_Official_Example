package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

/**
 * @Author jiajia0
 * @Date 2018/3/25 15:57
 * @@Describe 对zip文件的操作 https://golang.org/pkg/archive/zip/
 **/

func ExampleWriter() {
	/*
		// 创建一个Buffer 缓冲区来写入文档
		buf := new(bytes.Buffer)

		// 创建一个将zip写入buf中的Writer返回，这里传入刚才创建的buf，所以会将内容写入buf中
		w := zip.NewWriter(buf)
	*/

	// 这里使用一个zip文件进行测试
	const zipName = "test_archive/test_zip/test_writer.zip"
	// 创建该文件
	fzip, cerr := os.Create(zipName)
	if cerr != nil {
		log.Fatal(cerr)
	}

	// 使用刚才创建的zip文件作为参数，会返回一个用来向zip文件中写入内容的Writer
	w := zip.NewWriter(fzip)
	// 创建一个文件结构体，并且初始化两个文件信息
	var files = []struct {
		Name, Body string
	}{
		{"hello.txt", "This is hello.txt"},
		{"new.txt", "This is new.txt"},
	}

	// 遍历刚才我们定义的两个文件，返回一个索引和内容
	for _, file := range files {
		// 创建对应的文件，返回一个io.Writer用来执行写操作和一个err
		f, err := w.Create(file.Name)
		// 如果此时发生了错误，err的值不为nil，则将其打印
		if err != nil {
			log.Fatal(err)
		}

		// 将Body写入对应的Name文件中,存入zip中
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 确保文件关闭
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleReader() {
	// 打开一个zip文档用来进行读取操作，返回一个ReadCloser，实现了Reader接口
	r, err := zip.OpenReader("test_archive/test_zip/test_writer.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// 遍历所有的文件
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		// 打开该文件，rc是一个ReadCloser
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		// 将rc中的内容拷贝到控制台中
		_, err = io.Copy(os.Stdout, rc)
		if err != nil {
			log.Fatal(err)
		}
		// 关闭该ReadCloser
		rc.Close()
		fmt.Println()
	}
}

// 不太明白要做什么...
func ExampleWriter_RegisterCompressor() {
	// Override the default Deflate compressor with a higher compression level.

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Register a custom Deflate compressor.
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// Proceed to add files to w.
}

func main() {
	// ExampleWriter()
	ExampleReader()
}
