package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os/exec"
	"runtime"
)

//How to pretty print a Golang structure? [duplicate]
// https://stackoverflow.com/questions/56242013/how-to-pretty-print-a-golang-structure

// [Golang] Pretty Print Variable (struct, map, array, slice)
// https://siongui.github.io/2016/01/30/go-pretty-print-variable/

func Dev(c *cli.Context) error {
	// github.com/mattn/go-sqlite3

	//fmt.Println(rows)
	//applog.Default().Info("ssssss")
	//applog.Default().Info("sfsfssfs")
	//
	//v := redis_conn.GetRdsDB().Get(redis_conn.BGCtx, "aa").Val()
	//print(v)

	//var a complex64 = complex(1, 2)
	//s := fmt.Sprint( a)
	//fmt.Println(s)

	//arrayString := os.Environ() //获取系统变量
	//fmt.Println(arrayString)
	//apputil.PrettyPrint(arrayString)
	//fmt.Println(os.Getenv("ZSH"))
	return nil
}

func OpenBrowser(c *cli.Context) error {

	url := "http://www.baidu.com"

	openbrowser(url)

	return nil
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
