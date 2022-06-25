package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/xproto"
)

var (
	flagRequests    int
	flagGOMAXPROCS  int
	flagCpuProfName string
	flagMemProfName string
)

func init() {
	flag.IntVar(&flagRequests, "requests", 100000, "Number of atoms to intern.")
	flag.Parse()
}

func seqNames(n int) []string {
	names := make([]string, n)
	for i := range names {
		names[i] = fmt.Sprintf("NAME%d", i)
	}
	return names
}

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	names := seqNames(flagRequests)

	fcpu, err := os.Create(flagCpuProfName)
	if err != nil {
		log.Fatal(err)
	}
	defer fcpu.Close()
	pprof.StartCPUProfile(fcpu)
	defer pprof.StopCPUProfile()

	start := time.Now()
	cookies := make([]xproto.InternAtomCookie, flagRequests)
	for i := 0; i < flagRequests; i++ {
		cookies[i] = xproto.InternAtom(X,
			false, uint16(len(names[i])), names[i])
	}
	for _, cookie := range cookies {
		cookie.Reply()
	}
}
