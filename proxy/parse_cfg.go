package proxy

import (
	"go2rap/config"
)

//根据配置信息拿到易于程序处理的数据类型
func getParsedCfg() (map[string]string, map[string]ServerB) {
	cfgJson, error := config.ReadCfg("go2rap.json")

	if error != nil {

	}

	//服务器Map，key=host，value=proxy
	serverMap := make(map[string]string)

	//服务器Map，key=name，value=host
	serverNameMap := make(map[string]string)

	for _, x := range cfgJson.Servers {
		serverMap[x.Host] = x.Proxy
		serverNameMap[x.Name] = x.Host
	}

	//serverB的map，key=host，vale=ServerB{host,ip,paths}
	serverBMap := make(map[string]ServerB)
	for _, x := range cfgJson.Conditions {
		hostA, okA := serverNameMap[x.ServerA]
		hostB, okB := serverNameMap[x.ServerB]

		//条件中存在该server
		if okA && okB {
			ipB := serverMap[hostB]
			serverBMap[hostA] = ServerB{host: hostB, ip: ipB, paths: x.Path}
		}
	}

	return serverMap, serverBMap
}

//配置文件结构
type ServerB struct {
	host  string
	ip    string
	paths []string
}