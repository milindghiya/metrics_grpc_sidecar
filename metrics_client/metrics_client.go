package metrics_client

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/milindghiya/metrics_grpc_sidecar/metricspb/metricspb_v1"
	"google.golang.org/grpc"
)

type MetricsClient struct {
	Enabled     bool
	conn        *grpc.ClientConn
	client      pb.MetricsClient
	grpcTimeout time.Duration
	wg          sync.WaitGroup
}

func (mc *MetricsClient) ObserveHistogram(name string, labels map[string]string, value float64) {

	if mc.Enabled {
		ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
		defer cancel()
		updateParams := pb.UpdateParams{Name: name, LabelValues: labels, Value: value}
		_, err := mc.client.HistogramObserve(ctx, &updateParams)
		handlerError(err)
	}
}

func (mc *MetricsClient) SetGauge(name string, labels map[string]string, value float64) {
	if mc.Enabled {
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			updateParams := pb.UpdateParams{Name: name, LabelValues: labels, Value: value}
			_, err := mc.client.GaugeSet(ctx, &updateParams)
			handlerError(err)
	}
}

func (mc *MetricsClient) IncrementCounter(name string, labels map[string]string) {
	if mc.Enabled {
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			updateParams := pb.UpdateParams{Name: name, LabelValues: labels}
			_, err := mc.client.CounterInc(ctx, &updateParams)
			handlerError(err)
	}
}

func handlerError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func (mc *MetricsClient) CloseConnection() {
	if mc.Enabled {
		mc.conn.Close()
	}
}

func (mc *MetricsClient) Connect(address string, grpcTimeout float64) {
	if mc.Enabled {
		fmt.Println("Blocking till grpc connection gets established")
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Grpc connection gets established successfully!!")
		mc.client = pb.NewMetricsClient(conn)
		mc.conn = conn
		mc.grpcTimeout = time.Duration(grpcTimeout)
	}
}

func (mc *MetricsClient) AddHistogramMetric(name string, labels []string, help string, buckets []float64) {
	if mc.Enabled {
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			histogramParams := pb.CreateHistogramParams{Name: name, Labels: labels, Help: help, Buckets: buckets}
			_, err := mc.client.CreateHistogram(ctx, &histogramParams)
			handlerError(err)
	}
}

func (mc *MetricsClient) AddGaugeMetric(name string, labels []string, help string) {
	if mc.Enabled {
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			gaugeParams := pb.CreateGaugeParams{Name: name, Labels: labels, Help: help}
			_, err := mc.client.CreateGauge(ctx, &gaugeParams)
			handlerError(err)
	}
}

func (mc *MetricsClient) AddCounterMetric(name string, labels []string, help string) {
	if mc.Enabled {
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			counterParams := pb.CreateCounterParams{Name: name, Labels: labels, Help: help}
			_, err := mc.client.CreateCounter(ctx, &counterParams)
			handlerError(err)
	}
}
