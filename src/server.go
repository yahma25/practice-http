package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	var httpServer http.Server
	// '/' 경로로 브라우저 or curl 커맨드 등 클라이언트가 접속했을 때 호출
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888") // 80번은 http 기본 포트고, 대체 포트로 8000, 8080, 8888 등이 있지만, 18888 사용
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
