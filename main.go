package main

import (
	"encoder/factory"
	"fmt"
)

func main() {
	var fact factory.Factory

	md5 := fact.CreateFac("MD5")
	if md5 != nil {
		fmt.Println("md5 encode =", md5.Encode("haha1235656sdsd"))
	}

	base64 := fact.CreateFac("BASE64")
	if base64 != nil {
		fmt.Println("base64 encode =", base64.Encode("haha1235656sdsd"))
	}

}
