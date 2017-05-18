package proxy

import (
	"net/http"
	"net/http/httputil"
	"strings"
	"net/url"
	"log"
	"io/ioutil"
)

type HandleProxy struct {
}

//实现Handler的接口
func (h *HandleProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	host, ip, prefixPath := handleCfgAndServer(r.Host, r.URL.Path)

	//配置文件中不存在该host，转到欢迎页
	if host == "" {
		w.Write([]byte("<h1>Welcome to go2rap!</h1>"))
		log.Printf("host:%v does not exist in the configuration file", r.Host)
		return
	}

	remote, err := url.Parse("http://" + host)

	if err != nil {
		log.Fatalln("Parse url", err)
	}

	log.Println("request:" + host + prefixPath + singleJoiningSlash(remote.Path, r.URL.Path))

	if r.Method == "GET" {
		log.Printf("method: %v param:%v", r.Method, r.URL.Query().Encode())
		log.Println()
	}

	if r.Method == "POST" {
		x, _ := ioutil.ReadAll(r.Body)

		log.Printf("method: %v param:%v", r.Method, []byte(x))

	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			//设置主机
			req.Host = host
			req.URL.Host = ip
			req.URL.Scheme = remote.Scheme
			//设置路径
			req.URL.Path = prefixPath + singleJoiningSlash(remote.Path, r.URL.Path);
			//设置参数
			req.PostForm = r.PostForm
			req.URL.RawQuery = r.URL.RawQuery
			req.Form = r.Form

		},
	}

	proxy.ServeHTTP(w, r)
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

//处理请求的host、路径和配置，返回实际请求host和ip
func handleCfgAndServer(host string, path string) (string, string, string) {

	serverMap, serverBMap := getParsedCfg()
	ip, okHost := serverMap[host]
	serverB, okCondition := serverBMap[host]

	if okHost {
		if okCondition {
			for _, x := range serverB.paths {
				//如果条件符合
				if strings.Contains(path, x) {
					//打印命中条件的转发
					log.Printf("hit conditions：to host:%v", serverB.host)
					return serverB.host, serverB.ip, serverB.prefixPath
				}
			}
		}
		return host, ip, ""
	}
	return "", "", ""
}
