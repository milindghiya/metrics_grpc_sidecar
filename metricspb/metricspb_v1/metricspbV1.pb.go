// Metrics Sidecar Protobuf
//
// Provide GRPC interface for various containers to access the metric publish service.
// The client containers can invoke this interface to create instances of metrics
// with unique type, name and labels associated. This is specially useful
// for multi container pod to expose a single /metrics end point with each container
// making its own metrics available on /metrics

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: metricspbV1.proto

package metricspb_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCounterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Help   string   `protobuf:"bytes,3,opt,name=help,proto3" json:"help,omitempty"`
}

func (x *CreateCounterParams) Reset() {
	*x = CreateCounterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricspbV1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCounterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCounterParams) ProtoMessage() {}

func (x *CreateCounterParams) ProtoReflect() protoreflect.Message {
	mi := &file_metricspbV1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCounterParams.ProtoReflect.Descriptor instead.
func (*CreateCounterParams) Descriptor() ([]byte, []int) {
	return file_metricspbV1_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCounterParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCounterParams) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateCounterParams) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

type CreateGaugeParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Help   string   `protobuf:"bytes,3,opt,name=help,proto3" json:"help,omitempty"`
}

func (x *CreateGaugeParams) Reset() {
	*x = CreateGaugeParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricspbV1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGaugeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGaugeParams) ProtoMessage() {}

func (x *CreateGaugeParams) ProtoReflect() protoreflect.Message {
	mi := &file_metricspbV1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGaugeParams.ProtoReflect.Descriptor instead.
func (*CreateGaugeParams) Descriptor() ([]byte, []int) {
	return file_metricspbV1_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGaugeParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGaugeParams) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateGaugeParams) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

type CreateHistogramParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels  []string  `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Help    string    `protobuf:"bytes,3,opt,name=help,proto3" json:"help,omitempty"`
	Buckets []float64 `protobuf:"fixed64,4,rep,packed,name=buckets,proto3" json:"buckets,omitempty"`
}

func (x *CreateHistogramParams) Reset() {
	*x = CreateHistogramParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricspbV1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHistogramParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistogramParams) ProtoMessage() {}

func (x *CreateHistogramParams) ProtoReflect() protoreflect.Message {
	mi := &file_metricspbV1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistogramParams.ProtoReflect.Descriptor instead.
func (*CreateHistogramParams) Descriptor() ([]byte, []int) {
	return file_metricspbV1_proto_rawDescGZIP(), []int{2}
}

func (x *CreateHistogramParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateHistogramParams) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateHistogramParams) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *CreateHistogramParams) GetBuckets() []float64 {
	if x != nil {
		return x.Buckets
	}
	return nil
}

type UpdateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LabelValues map[string]string `protobuf:"bytes,2,rep,name=label_values,json=labelValues,proto3" json:"label_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value       float64           `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateParams) Reset() {
	*x = UpdateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricspbV1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateParams) ProtoMessage() {}

func (x *UpdateParams) ProtoReflect() protoreflect.Message {
	mi := &file_metricspbV1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateParams.ProtoReflect.Descriptor instead.
func (*UpdateParams) Descriptor() ([]byte, []int) {
	return file_metricspbV1_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateParams) GetLabelValues() map[string]string {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

func (x *UpdateParams) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricspbV1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_metricspbV1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_metricspbV1_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_metricspbV1_proto protoreflect.FileDescriptor

var file_metricspbV1_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x70, 0x62, 0x56, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x22, 0x53, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x75, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x65, 0x6c, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x22,
	0x71, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x3e, 0x0a, 0x10, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x45, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9b, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x30, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x75,
	0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x12, 0x0d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x08, 0x47, 0x61, 0x75, 0x67, 0x65, 0x53, 0x65, 0x74, 0x12, 0x0d, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x10, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x12, 0x0d, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x3b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x70, 0x62, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metricspbV1_proto_rawDescOnce sync.Once
	file_metricspbV1_proto_rawDescData = file_metricspbV1_proto_rawDesc
)

func file_metricspbV1_proto_rawDescGZIP() []byte {
	file_metricspbV1_proto_rawDescOnce.Do(func() {
		file_metricspbV1_proto_rawDescData = protoimpl.X.CompressGZIP(file_metricspbV1_proto_rawDescData)
	})
	return file_metricspbV1_proto_rawDescData
}

var file_metricspbV1_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metricspbV1_proto_goTypes = []interface{}{
	(*CreateCounterParams)(nil),   // 0: CreateCounterParams
	(*CreateGaugeParams)(nil),     // 1: CreateGaugeParams
	(*CreateHistogramParams)(nil), // 2: CreateHistogramParams
	(*UpdateParams)(nil),          // 3: UpdateParams
	(*Response)(nil),              // 4: Response
	nil,                           // 5: UpdateParams.LabelValuesEntry
}
var file_metricspbV1_proto_depIdxs = []int32{
	5, // 0: UpdateParams.label_values:type_name -> UpdateParams.LabelValuesEntry
	0, // 1: Metrics.CreateCounter:input_type -> CreateCounterParams
	1, // 2: Metrics.CreateGauge:input_type -> CreateGaugeParams
	2, // 3: Metrics.CreateHistogram:input_type -> CreateHistogramParams
	3, // 4: Metrics.CounterInc:input_type -> UpdateParams
	3, // 5: Metrics.GaugeSet:input_type -> UpdateParams
	3, // 6: Metrics.HistogramObserve:input_type -> UpdateParams
	4, // 7: Metrics.CreateCounter:output_type -> Response
	4, // 8: Metrics.CreateGauge:output_type -> Response
	4, // 9: Metrics.CreateHistogram:output_type -> Response
	4, // 10: Metrics.CounterInc:output_type -> Response
	4, // 11: Metrics.GaugeSet:output_type -> Response
	4, // 12: Metrics.HistogramObserve:output_type -> Response
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metricspbV1_proto_init() }
func file_metricspbV1_proto_init() {
	if File_metricspbV1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metricspbV1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCounterParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metricspbV1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGaugeParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metricspbV1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHistogramParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metricspbV1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metricspbV1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metricspbV1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metricspbV1_proto_goTypes,
		DependencyIndexes: file_metricspbV1_proto_depIdxs,
		MessageInfos:      file_metricspbV1_proto_msgTypes,
	}.Build()
	File_metricspbV1_proto = out.File
	file_metricspbV1_proto_rawDesc = nil
	file_metricspbV1_proto_goTypes = nil
	file_metricspbV1_proto_depIdxs = nil
}
