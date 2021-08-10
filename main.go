package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

var cacheDir string

func main() {
	addr := "localhost:8080"
	flag.StringVar(&cacheDir, "cache", "", "cache dir")
	flag.Parse()

	if len(cacheDir) < 1 {
		log.Printf("use -cache=xxx set cache dir")
		return
	}

	log.Printf("proxy listen " + addr + ", cache at " + cacheDir)
	http.ListenAndServe(addr, &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct",
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org",
		},
		Cacher: goproxy.DirCacher(cacheDir),
	})
}
