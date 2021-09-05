package utils

import (
	"errors"
	"log"
	"os"
)

const DirSplitChar = "/"
const EOL = '\n'
const EolStr = string(EOL)

//CreateDirNX 创建一个目录如果它不存在
func CreateDirNX(path string) error {
	log.Println(path)
	s, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, 0755)
			if err != nil {
				return err
			}
		}
	}
	if s == nil {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	}
	if !s.IsDir() {
		return errors.New(path + "is not dir")
	}
	return nil
}

//CreateFileNX 打开一个文件(如果他不存在会创建,仅可写/追加写)
func CreateFileNX(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

//WriteFileNX 写数据到一个文件(如果他不存在会创建,仅追加写)
func WriteFileNX(path string, data string) {
	file := CreateFileNX(path)
	_, err := file.WriteString(data)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
}

//WriteByteNX 写byte数据到一个文件(如果他不存在会创建,仅追加写)
func WriteByteNX(path string, data []byte) {
	file := CreateFileNX(path)
	_, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
}

//WriteFileNXln 换行写数据到一个文件(如果他不存在会创建,仅追加写)
func WriteFileNXln(path string, data string) {
	file := CreateFileNX(path)
	_, err := file.WriteString(data + EolStr)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
}

//WriteByteNXln 换行写byte数据到一个文件(如果他不存在会创建,仅追加写)
func WriteByteNXln(path string, data []byte) {
	file := CreateFileNX(path)
	data = append(data, EOL)
	_, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
}
