package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func delErr[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

func find[T comparable](arr []T, target T) bool {
	for _, t := range arr {
		if t == target {
			return true
		}
	}
	return false
}

var wihteList = []string{"go.mod", "main.go", "main.exe"}

func main() {
	var format string
	fmt.Print("将本文件放在所需要修改的文件夹中，如果不在请按 crtl + c 退出。\n请输入当前文件统一前缀，按回车结束。例如：{你的输入}-顺序 \n")
	fmt.Scanf("%s", &format)
	if format == "" {
		fmt.Println("请按照要求输入！")
		return
	}
	pathName := delErr(os.Getwd())
	idx := 1
	fileInfos := delErr(ioutil.ReadDir(pathName))
	for _, file := range fileInfos {
		fileName := file.Name()
		if find(wihteList, fileName) {
			continue
		}
		names := strings.Split(fileName, ".")
		suffix := names[len(names)-1]
		newName := format + "-" + strconv.Itoa(idx) + "." + suffix
		idx++
		if err := os.Rename(pathName+"\\"+fileName, newName); err != nil {
			fmt.Println("err : ", err)
		}
	}
}
