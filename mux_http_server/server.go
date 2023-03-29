package main

import (
    "log"
    "net/http"
    mtc "github.com/milindghiya/metrics_grpc_sidecar/metrics_client"
    "github.com/gorilla/mux"
)


func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("before request")
        
		next.ServeHTTP(w, r)
        
        log.Println("after request")
	})
}


func main() {
	
    r := mux.NewRouter()    
    
    mc := mtc.MetricsClient{enabled: false}
    mc.Connect("localhost:50051",3,)
    mc.AddCounterMetric("num_requests",[]string{"container_name"}, "number of requests")
    r.Use(commonMiddleware)
    r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })
    
    
    log.Fatal(http.ListenAndServe(":3333", r))
}
