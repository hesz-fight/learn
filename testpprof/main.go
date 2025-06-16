package main

import (
	"expvar"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	// 注册GC统计变量
	expvar.Publish("gc_stats", expvar.Func(func() interface{} {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		return map[string]interface{}{
			"num_gc":      memStats.NumGC,
			"gc_pause_ns": memStats.PauseNs,
			"total_gc_ns": memStats.PauseTotalNs,
		}
	}))

	// 启动HTTP服务（默认 :6060）
	http.ListenAndServe(":6060", nil)
}
