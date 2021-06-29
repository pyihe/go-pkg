package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io/ioutil"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"strconv"
	"syscall"

	"github.com/pyihe/go-pkg/rands"
)

var (
	mailChecker  = regexp.MustCompile(`^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$`)
	phoneChecker = regexp.MustCompile(`^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$`)
)

//校验邮箱格式
func CheckMailFormat(mail string) bool {
	return mailChecker.MatchString(mail)
}

//校验电话号码格式
func CheckPhoneFormat(phone string) bool {
	return phoneChecker.MatchString(phone)
}

//生成一个1-100的随机数, 用于简单的判断概率
func LessThanIn100(per int) bool {
	if per < 1 || per > 100 {
		panic("input must between 1 and 100")
	}
	return per >= rands.Int(1, 100)
}

//如果监听到系统中断信号，则执行onNotify()
func Notify(onNotify func()) {
	//SIGHUP		终端控制进程结束(终端连接断开)
	//SIGQUIT		用户发送QUIT字符(Ctrl+/)触发
	//SIGTERM		结束程序(可以被捕获、阻塞或忽略)
	//SIGINT		用户发送INTR字符(Ctrl+C)触发
	//SIGSTOP		停止进程(不能被捕获、阻塞或忽略)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGHUP:
			if onNotify == nil {
				return
			}
			onNotify()
		default:
			return
		}
	}
}

//判断src中是否有元素ele
func Contain(src interface{}, ele interface{}) bool {
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(src)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(ele, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

//将嵌套的map[string]interface全部转换成一层
func Interface2Map(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data.(map[string]interface{}) {
		switch v := v.(type) {
		case map[string]interface{}:
			for i, u := range v {
				result[i] = u
			}
		default:
			result[k] = v
		}
	}
	return result
}

//gzip解压
func UnGZIP(content []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, content)
	if err != nil {
		return nil, err
	}
	zipReader, err := gzip.NewReader(buffer)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()
	result, err := ioutil.ReadAll(zipReader)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ConvertToBinary 十进制转换为二进制字符串
func ConvertToBinary(data int) string {
	result := ""
	for ; data > 0; data = data / 2 {
		n := data % 2
		result = strconv.Itoa(n) + result
	}
	return result
}
