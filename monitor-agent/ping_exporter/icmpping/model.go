package icmpping

import (
	"sync"
	"github.com/WeBankPartners/open-monitor/monitor-agent/ping_exporter/funcs"
)

var (
	TestModel bool
	localIp    []string   // 存储本机IP
	timeout int
	//lastResultMap = make(map[string]funcs.PingResultObj)  // 保存上次结果，和本次进行对比

	//resultMap = make(map[string]funcs.PingResultObj)   // 保存本次结果
	resultMapLock = new(sync.RWMutex)

	resultListNew []*funcs.PingResultObj

	//retryMap = make(map[string]int)   // 保存需要重试的IP
	//retryMapLock = new(sync.RWMutex)

	successIpList []string
	successListLock = new(sync.RWMutex)

	//retryIpList  []string
	//retryListLock = new(sync.RWMutex)
)

func addSuccessIp(ip string) {
	successListLock.Lock()
	defer successListLock.Unlock()
	successIpList = append(successIpList, ip)
}

func GetSuccessIp() []string {
	successListLock.RLock()
	defer successListLock.RUnlock()
	return successIpList
}

func ClearSuccessIp() {
	successListLock.Lock()
	successIpList = []string{}
	successListLock.Unlock()
}

//func addRetryIp(ip string) {
//	retryListLock.Lock()
//	defer retryListLock.Unlock()
//	retryIpList = append(retryIpList, ip)
//}
//
//func IsInRetryIp(ip string) bool {
//	exist := false
//	retryListLock.RLock()
//	for _,v := range retryIpList {
//		if v == ip {
//			exist = true
//			break
//		}
//	}
//	retryListLock.RUnlock()
//	return exist
//}
//
//func GetRetryIp() []string {
//	retryListLock.RLock()
//	defer retryListLock.RUnlock()
//	return retryIpList
//}
//
//func ClearRetryIp(){
//	retryListLock.Lock()
//	retryIpList = []string{}
//	retryListLock.Unlock()
//}

func writeResultMap(ip string, re int, useTime float64) {
	//resultMapLock.Lock()
	//defer resultMapLock.Unlock()
	//resultMap[ip] = funcs.PingResultObj{UpDown:re, UseTime:useTime}
	resultMapLock.Lock()
	resultListNew = append(resultListNew, &funcs.PingResultObj{Ip:ip, UpDown:re, UseTime:useTime})
	resultMapLock.Unlock()
}

func readResultMap() map[string]funcs.PingResultObj {
	//resultMapLock.RLock()
	//defer resultMapLock.RUnlock()
	//return resultMap
	tmpResultMap := make(map[string]funcs.PingResultObj)
	resultMapLock.RLock()
	for _,v := range resultListNew {
		tmpResultMap[v.Ip] = funcs.PingResultObj{UpDown:v.UpDown, UseTime:v.UseTime}
	}
	resultMapLock.RUnlock()
	return tmpResultMap
}

func clearResultMap()  {
	resultMapLock.Lock()
	//resultMap = make(map[string]funcs.PingResultObj)
	resultListNew = []*funcs.PingResultObj{}
	resultMapLock.Unlock()
}

//func GetRetryMap(ip string, n int) int {
//	retryMapLock.Lock()
//	defer retryMapLock.Unlock()
//	if v,ok := retryMap[ip]; ok{
//		retryMap[ip] = v+n
//		return v+n
//	}else{
//		retryMap[ip] = n
//		return n
//	}
//}
//
//func ClearRetryMap() {
//	retryMapLock.Lock()
//	defer retryMapLock.Unlock()
//	for k,_ := range retryMap{
//		retryMap[k] = 0
//	}
//}

func containString(ip string, ips []string) bool {
	for _,v := range ips {
		if ip == v {
			return true
		}
	}
	return false
}