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
		mc.wg.Add(1)
		go func() {
			defer mc.wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			updateParams := pb.UpdateParams{Name: name, LabelValues: labels, Value: value}
			mc.client.HistogramObserve(ctx, &updateParams)
		}()
	}
}

func (mc *MetricsClient) SetGauge(name string, labels map[string]string, value float64) {
	if mc.Enabled {
		mc.wg.Add(1)
		go func() {
			defer mc.wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			updateParams := pb.UpdateParams{Name: name, LabelValues: labels, Value: value}
			mc.client.GaugeSet(ctx, &updateParams)
		}()
	}
}

func (mc *MetricsClient) IncrementCounter(name string, labels map[string]string) {
	if mc.Enabled {
		mc.wg.Add(1)
		go func() {
			defer mc.wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			updateParams := pb.UpdateParams{Name: name, LabelValues: labels}
			mc.client.CounterInc(ctx, &updateParams)
		}()
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
		mc.client = pb.NewMetricsClient(conn)
		mc.conn = conn
		mc.grpcTimeout = time.Duration(grpcTimeout)
	}
}

func (mc *MetricsClient) AddHistogramMetric(name string, labels []string, help string, buckets []float64) {
	if mc.Enabled {
		mc.wg.Add(1)
		go func() {
			defer mc.wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			histogramParams := pb.CreateHistogramParams{Name: name, Labels: labels, Help: help, Buckets: buckets}
			mc.client.CreateHistogram(ctx, &histogramParams)
		}()
	}
}

func (mc *MetricsClient) AddGaugeMetric(name string, labels []string, help string) {
	if mc.Enabled {
		mc.wg.Add(1)
		go func() {
			defer mc.wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			gaugeParams := pb.CreateGaugeParams{Name: name, Labels: labels, Help: help}
			mc.client.CreateGauge(ctx, &gaugeParams)
		}()
	}
}

func (mc *MetricsClient) AddCounterMetric(name string, labels []string, help string) {
	if mc.Enabled {
		mc.wg.Add(1)
		go func() {
			defer mc.wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), mc.grpcTimeout*time.Second)
			defer cancel()
			counterParams := pb.CreateCounterParams{Name: name, Labels: labels, Help: help}
			mc.client.CreateCounter(ctx, &counterParams)
		}()
	}
}

func (mc *MetricsClient) Wait() {
	if mc.Enabled {
		mc.wg.Wait()
	}
}