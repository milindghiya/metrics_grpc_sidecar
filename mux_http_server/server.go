package main

import (
    "log"
    "net/http"
    mtc "github.com/milindghiya/metrics_grpc_sidecar/metrics_client"
    "github.com/gorilla/mux"
)

var (
    metricsClient = newMetricsClient() 
)

func newMetricsClient() mtc.MetricsClient {
    mc := mtc.MetricsClient{Enabled: true}
    err := mc.Connect("localhost:50051",3)
    if err != nil {
        log.Panic(err)
    }
    mc.AddCounterMetric("num_requests",[]string{"container_name"}, "number of requests")
    return mc
}

func commonMiddleware(next http.Handler) http.Handler {
    
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("before request")
		next.ServeHTTP(w, r)
        metricsClient.IncrementCounter("num_requests", map[string]string{"container_name":"default_http_server"})
        log.Println("after request")
	})
}


func main() {
	
    r := mux.NewRouter()    
    
    r.Use(commonMiddleware)
    r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })
    
    log.Fatal(http.ListenAndServe(":3333", r))
}
