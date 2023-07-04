package main

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	cfg = pflag.StringP("config", "c", "./config.yaml",
		"api server config file path.")
)

func LoadConfig() {
	viper.SetConfigFile(*cfg)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed, err:%v", err)
	}
	viper.AutomaticEnv()
}

var count int64
var lock sync.RWMutex

func main() {
	flag.Parse()
	LoadConfig()

	target := viper.GetString("TARGET")
	port := viper.GetString("PORT")

	url, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			lock.RLock()
			lock.RUnlock()
			atomic.AddInt64(&count, 1)
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host
			req.Host = url.Host
		},
		ModifyResponse: func(r *http.Response) error {

			return nil
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		authFromHeader := r.Header.Get("x-api-key")
		if authFromHeader == "" {
			w.WriteHeader(401)
			return
		}
		proxy.ServeHTTP(w, r)
	})

	log.Println("Listen on port:" + port)
	log.Println("Running...")

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func removeBearer(str string) string {
	if strings.HasPrefix(str, "Bearer") {
		return str[7:]
	}
	return str
}
