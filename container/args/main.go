// Author: yangzq80@gmail.com
// Date: 2022/10/13
package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	name := flag.String("name", "default-name-1", "--name=name1")
	name2 := flag.String("name2", "default-name-2", "--name2=name2")

	//kind : write or read
	kind := flag.String("kind", "write", "--kind=write")

	flag.Parse()

	log.Println("cmd args:", *name, *name2, *kind)

	fn := "/tmpdir/tmp-dat1"
	if *kind == "write" {
		d1 := []byte(*kind + "---" + *name)
		err := os.WriteFile(fn, d1, 0644)
		check(err)
		log.Println("Write success", fn)
	} else {
		dat, err := os.ReadFile(fn)
		check(err)
		log.Println("Read file:", string(dat))
	}

	for i := 0; i < 5; i++ {
		log.Println("times:", i)
		time.Sleep(time.Second * 1)
	}
}

func check(e error) {
	if e != nil {
		log.Println(e.Error())
	}
}
