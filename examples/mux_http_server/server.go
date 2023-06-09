package main

import (
    "log"
    "net/http"
    mtc "github.com/milindghiya/metrics_grpc_sidecar/metrics_client"
    "github.com/gorilla/mux"
    rw "github.com/milindghiya/metrics_grpc_sidecar/response_writer"
    "time"
)

var (
    metricsClient = newMetricsClient() 
)

func newMetricsClient() mtc.MetricsClient {
    mc := mtc.MetricsClient{Enabled: true}
    mc.Connect("localhost:50051",3)
    mc.AddCounterMetric("http_requests_total",[]string{"host_name","request_path","http_method","status_code"}, "This metric tracks the total number of requests made to the server over a period of time. It can help us identify trends and spikes in traffic.")
    mc.AddHistogramMetric("http_request_duration_seconds",[]string{"host_name","request_path","http_method","status_code"}, "Time taken to serve HTTP requests.",  []float64{0.01, 0.1, 0.5, 1, 2.5, 5, 10, 50})
    mc.AddHistogramMetric("http_request_size_bytes",[]string{"host_name","request_path","http_method","status_code"}, "Size of HTTP requests.",  []float64{1, 100, 500, 1000, 5000, 10000},)
    mc.AddHistogramMetric("http_response_size_bytes",[]string{"host_name","request_path","http_method","status_code"}, "Size of HTTP responses.",  []float64{1, 100, 500, 1000, 5000, 10000},)
    return mc
}



func commonMiddleware(next http.Handler) http.Handler {
    
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        respWriter := rw.NewResponseWriter(w)
        next.ServeHTTP(respWriter, r)
        route := mux.CurrentRoute(r)
        path, _  := route.GetPathTemplate()
        duration := time.Since(respWriter.StartTime())
        metricsClient.ObserveHistogram("http_request_size_bytes", map[string]string{"host_name":r.Host,"request_path": path,"http_method": r.Method, "status_code": respWriter.StatusCodeStr() }, float64(r.ContentLength))
        metricsClient.ObserveHistogram("http_response_size_bytes", map[string]string{"host_name":r.Host,"request_path": path,"http_method": r.Method, "status_code": respWriter.StatusCodeStr() }, float64(respWriter.Count()))
        metricsClient.IncrementCounter("http_requests_total", map[string]string{"host_name":r.Host,"request_path": path,"http_method": r.Method, "status_code": respWriter.StatusCodeStr() })
        metricsClient.ObserveHistogram("http_request_duration_seconds", map[string]string{"host_name":r.Host,"request_path": path,"http_method": r.Method, "status_code": respWriter.StatusCodeStr() }, float64(duration.Seconds()))        
	})
}


func main() {
    r := mux.NewRouter()    
    r.Use(commonMiddleware)
    r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })
    r.HandleFunc("/hello/{id}", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world with id!"))
    })
    log.Fatal(http.ListenAndServe(":3333", r))
}
