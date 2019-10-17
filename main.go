package main

import (
	"strings"
	"log"
	"github.com/GitHamburg/pinyin-golang/pinyin"
)

func main()  {
	netName := "本地连接 2"
	netIfName := pinyin.NewDict().Convert(strings.Replace(netName, " ", "_", -1), "_").None2()
	iface := "iface=" + netIfName
	log.Println(netName," - ",netIfName)
	log.Println(iface)
}