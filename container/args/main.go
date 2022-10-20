// Author: yangzq80@gmail.com
// Date: 2022/10/13
package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	content := flag.String("content", "default-content-1", "--content=default-content-1")
	fileName := flag.String("fileName", "config.json", "--fileName=config.json")

	//kind : write or read
	kind := flag.String("kind", "write", "--kind=write")

	flag.Parse()

	log.Println("cmd args:", *content, *fileName, *kind)

	fn := "/tmpdir/" + *fileName
	if *kind == "write" {
		d1 := []byte(*content)
		err := os.WriteFile(fn, d1, 0644)
		check(err)
		log.Println("Write success", fn)
	} else {
		dat, err := os.ReadFile(fn)
		check(err)
		log.Println("Read file:", string(dat))
	}
	/*for i := 0; i < 5; i++ {
		log.Println("times:", i)
		time.Sleep(time.Second * 1)
	}*/
}

func check(e error) {
	if e != nil {
		log.Println(e.Error())
	}
}
