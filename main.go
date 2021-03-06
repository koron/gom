package main

import (
	"flag"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		fmt.Println(" Tasks:")
		fmt.Println("   gom build")
		fmt.Println("   gom install")
		fmt.Println("   gom test")
		fmt.Println("   gom gen-travis-yml")
		os.Exit(1)
	}
	goms, err := parseGomfile("Gomfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, "gom: ", err)
		os.Exit(1)
	}

	sc := make(chan os.Signal, 10)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	go func() {
		<-sc
		ct.ResetColor()
		os.Exit(0)
	}()
	switch flag.Arg(0) {
	case "install":
		err = install(goms)
	case "build":
		err = build()
	case "test":
		err = test()
	case "gen-travis-yml":
		err = gen_travis_yml()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "gom: ", err)
		os.Exit(1)
	}
}
