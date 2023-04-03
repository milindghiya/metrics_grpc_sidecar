package main

import (
    "log"
    "net/http"
    mtc "github.com/milindghiya/metrics_grpc_sidecar/metrics_client"
    "github.com/gorilla/mux"
    rw "github.com/milindghiya/response_writer"
)

var (
    metricsClient = newMetricsClient() 
)

func newMetricsClient() mtc.MetricsClient {
    mc := mtc.MetricsClient{Enabled: false}
    mc.Connect("localhost:50051",3)
    mc.AddCounterMetric("num_requests",[]string{"container_name"}, "number of requests")
    mc.Wait()
    return mc
}

func commonMiddleware(next http.Handler) http.Handler {
    
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
        metricsClient.IncrementCounter("num_requests", map[string]string{"container_name":"default_http_server"})
        metricsClient.Wait()
        log.Println()
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
