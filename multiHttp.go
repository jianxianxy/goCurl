//批量发送http请求
package main

import (
	"strings"
	"io"
	"bufio"
	"os"
    "fmt"
    "net/http"
    "io/ioutil"
	"net/url"
)

func main(){
    para := paramLine("id.txt")
    for _,val := range para{
        fmt.Println(val)
    }
    /*
    param := make(map[string]string)
    param["id"] = "123"
    param["name"] = "leiluo"

    body := post("http://www.v.com/test.php",param)
    fmt.Println(body)
    */
}
/* GET 请求
 * uri http地址
 */ 
func get(uri string) string{
    rep,err := http.Get(uri)
    if err != nil{
        return "404"
    }
    defer rep.Body.Close()
    body,err := ioutil.ReadAll(rep.Body)
    if err != nil{
        return "400"
    }
    return string(body)
}
/* POST 请求
 * uri http地址
 * param POST参数
 */ 
func post(uri string,param map[string]string) string{
    query := url.Values{}
    for key,val := range param{
        query.Set(key,val)    
    }
    rep, err := http.PostForm(uri,query)
    if err != nil {
        return "404"
    }
    defer rep.Body.Close()
    body, err := ioutil.ReadAll(rep.Body)
    if err != nil {
        return "400"
    }
    return string(body)
}

/* 解析文件成map切片用于post
 * fpath 文件路径
 */ 
func paramFile(fpath string) []map[string]string{
    file,err := os.Open(fpath)
    if err != nil{
        panic("文件不存在")
    }
    defer file.Close()
    rd := bufio.NewReader(file)
    paret := make([]map[string]string,0)
    for{
        //line,err := rd.ReadString('\n') //读取行,如果文件末尾没有空行，最后一行不返回
        line,_,err := rd.ReadLine()
        if err != nil || err == io.EOF{
            break
        }
        lnstr := string(line)
        lnarr := strings.Split(lnstr," ")
        quarr := make([]string,0)
        for _,val := range lnarr{
            v := strings.TrimSpace(val)
            if len(v) > 0{
                quarr = append(quarr,v)    
            }
        }
        if len(quarr)%2 != 0{
            panic("参数数量异常内容:"+lnstr)
        }
        qumap := make(map[string]string)
        for i := 0;i < len(quarr)/2;i++{
            qumap[quarr[i*2]] = quarr[i*2+1]    
        }
        paret = append(paret,qumap)
    }
    return paret
}

/* 解析文件成string切片用于get
 * fpath 文件路径
 */ 
func paramLine(fpath string) []string{
    file,err := os.Open(fpath)
    if err != nil{
        panic("文件不存在")
    }
    defer file.Close()
    rd := bufio.NewReader(file)
    paret := make([]string,0)
    for{
        //line,err := rd.ReadString('\n') //读取行,如果文件末尾没有空行，最后一行不返回
        line,_,err := rd.ReadLine()
        if err != nil || err == io.EOF{
            break
        }
        qstr := strings.TrimSpace(string(line))
        paret = append(paret,qstr)
    }
    return paret
}