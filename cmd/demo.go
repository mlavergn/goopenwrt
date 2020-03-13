package main

import (
	"log"

	openwrt "github.com/mlavergn/goopenwrt"
)

func main() {
	ref := openwrt.NewOpenWRT()
	log.Println(ref.GetWANIPV4())
}
