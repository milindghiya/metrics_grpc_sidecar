PHONY: all
all: metrics_collector metrics_producer mux_http_server


.PHONY: metrics_collector
metrics_collector:
	echo ">> Building metrics_collector"
	go build -o metrics_collector ./cmd/main.go

metrics_producer:
	echo ">> Building metrics_producer"
	go build -o metrics_producer ./examples/metrics_producer.go

mux_http_server:
	echo ">> Building mux_http_server"
	go build -o mux_http_server ./examples/mux_http_server/server.go

clean:
	rm metrics_collector metrics_producer mux_http_server