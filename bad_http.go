package main

import (
	"github.com/go-cmd/cmd"
	"io/ioutil"
	"net/http"
	"strings"
)

func sendDingMsg(msg string) {

	webHook := `https://oapi.dingtalk.com/robot/send?access_token=xxxxx`  //钉钉机器人api token
	content := `{"msgtype": "text",
		"text": {"content": "`+ msg + `"}
	}`

	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
	}
}

func runCommand(ccmd string) string {
	ccmd1 := strings.Replace(ccmd,"\n","", -1) // 注意，从web读取到的字符串结尾有个换行符，需要去掉。
	list := strings.Split(ccmd1, " ")
	envCmd := cmd.NewCmd(list[0], list[1:]...)
	status := envCmd.Start()
	finalStatus := <-status
	a := strings.Join(finalStatus.Stdout,"\n")
	//fmt.Println(a)

	return a
}

func getCommand()string{
	url := "http://xxxx/xxx.txt"  //命令文件，内容：ls -a -l 参数间以1个空隔为间隔
	res, err :=http.Get(url)
	if err != nil {
		return ""
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return ""
	}
	return string(robots)
}

func main() {
	ccmd := getCommand()
	ccmdrss := "Command Output:\n" + runCommand(ccmd)  // Command是设置的钉钉接口关键字，必须携带才能请求成功。
	sendDingMsg(ccmdrss)
	sayOk := "Command \n" + ccmd + " \nSuccess!"
	sendDingMsg(sayOk)


}