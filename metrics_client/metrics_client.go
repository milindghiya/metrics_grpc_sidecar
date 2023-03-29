package metrics_client

import (
	"google.golang.org/grpc"
	pb "github.com/milindghiya/metrics_grpc_sidecar/metricspb/metricspb_v1"
	"context"
	"time"
)

type MetricsClient struct {
	enabled bool
	conn *grpc.ClientConn
	client pb.MetricsClient
	grpcTimeoutInSeconds time.Duration
}

func (mc *MetricsClient) ObserveHistogram(name string, labels map[string]string, value float64) {
	if mc.enabled {
		ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeoutInSeconds * time.Second)
		defer cancel()
		updateParams := pb.UpdateParams{Name: name, LabelValues: labels, Value: value}
		mc.client.HistogramObserve(ctx, &updateParams)
	}
}


func (mc *MetricsClient) SetGauge(name string, labels map[string]string, value float64) {
	if mc.enabled {
		ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeoutInSeconds * time.Second)
		defer cancel()
		updateParams := pb.UpdateParams{Name: name, LabelValues: labels, Value: value}
		mc.client.GaugeSet(ctx, &updateParams)
	}
}

func (mc *MetricsClient) IncrementCounter(name string, labels map[string]string) {
	if mc.enabled {
		ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeoutInSeconds * time.Second)
		defer cancel()
		updateParams := pb.UpdateParams{Name: name, LabelValues: labels}
		mc.client.CounterInc(ctx, &updateParams)
	}
}

func (mc *MetricsClient) CloseConnection() {
	if mc.enabled {
		mc.conn.Close()
	}
}

func (mc *MetricsClient) Connect(address string, grpcTimeoutInSeconds float64, isEnabled bool) error {
	if isEnabled {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			//print error
			return err
		}
		mc.client = pb.NewMetricsClient(conn)
		mc.conn = conn
		mc.enabled = true
		mc.grpcTimeoutInSeconds = time.Duration(grpcTimeoutInSeconds)
	}
	return nil
}

func (mc *MetricsClient) AddHistogramMetric(name string, labels []string, help string, buckets []float64) {
	if mc.enabled {
		ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeoutInSeconds * time.Second)
		defer cancel()
		histogramParams := pb.CreateHistogramParams{Name:name, Labels:labels, Help:help, Buckets:buckets}
		mc.client.CreateHistogram(ctx, &histogramParams)
	}
}

func (mc *MetricsClient) AddGaugeMetric(name string, labels []string, help string) {
	if mc.enabled {
		ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeoutInSeconds * time.Second)
		defer cancel()
		gaugeParams := pb.CreateGaugeParams{Name: name, Labels: labels, Help: help}
		mc.client.CreateGauge(ctx, &gaugeParams)
	}
}


func (mc *MetricsClient) AddCounterMetric(name string, labels []string, help string) {
	if mc.enabled {
		ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeoutInSeconds * time.Second)
		defer cancel()
		counterParams := pb.CreateCounterParams{Name: name, Labels: labels, Help: help}
		mc.client.CreateCounter(ctx, &counterParams)
	}
}
