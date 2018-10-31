package main

import (
    "multiHttp"    
)

func main(){
    res := multiHttp.Get("http://www.baidu.com")
    panic(res)
}