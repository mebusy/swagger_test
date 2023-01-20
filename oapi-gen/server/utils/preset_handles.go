package utils

import (
	"fmt"
	"net/http"
	"syscall"
)

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

var gitcommit string

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// time.Sleep( 10 * time.Second ) // test gracefully shutdown
	fmt.Fprintf(w, "git commit: %s\n", gitcommit)
	fmt.Fprintf(w, "net host: %s\n", GetIP())

	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Fprintf(w, "rlimt cur:%d , max: %d \n", rLimit.Cur, rLimit.Max)
	fmt.Fprintf(w, "x-for:%s, x-real-ip:%s, r.Host:%s", r.Header.Get("X-Forwarded-For"), r.Header.Get("X-Real-Ip"), r.Host)
}
