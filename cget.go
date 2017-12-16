/**
 * 2017.12.15
 * xiongjianbo
 */
package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"flag"
	"os"
	"os/exec"
)
/**
 * 判断文件或目录是否存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * 获取配置对象
 */
type Curlconf struct{
	Base struct {
		User string
		Password string
		Apiuri string
		Port string
	}
}
func getConf(fileName string) (Curlconf) {
	config := Curlconf{}
	err := gcfg.ReadFileInto(&config, fileName)
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	return config
}

/**
 * 入口
 */
func main(){
	//从命令行获取参数 -url和-json
	url := flag.String("url", "/", "url")
	json := flag.String("json", "", "json")
	flag.Parse()

	//从/etc/curltool.conf中获取配置
	config := getConf("/etc/curltool.conf")

	var cmdString string;
	cmdString = "curl -u "+config.Base.User+":"+config.Base.Password+" -H 'Content-Type: application/json' -XGET '"+config.Base.Apiuri+":"+config.Base.Port+"'"+*url;
	if *json != "" {
		cmdString += " -d '"+*json+"'"
	}

	//如果没有安装curl 先装上curl再说
	isext,_ := PathExists("/usr/bin/curl")
	if !isext{
		cmd := exec.Command("sh","-c","sudo apt install curl")
		out, _ := cmd.Output()
		fmt.Println(string(out))
	}

	//执行最终的curl命令
	cmd := exec.Command("sh","-c",cmdString)
	out, _ := cmd.Output()
	fmt.Println(string(out))
}