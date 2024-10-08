// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: ynisonstate/playing_status.proto

package ynisonstate

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

// Статус воспроизведения.
type PlayingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Прогресс проигрываемой сущности. Значение в интервале [0; длина] в миллисекундах.
	// Для infinite-очередей равен 0.
	ProgressMs int64 `protobuf:"varint,1,opt,name=progress_ms,json=progressMs,proto3" json:"progress_ms,omitempty"`
	// Длительность проигрываемой сущности в миллисекундах.
	// Для infinite-очередей равна 0.
	DurationMs int64 `protobuf:"varint,2,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	// Был ли трек поставлен на паузу.
	Paused bool `protobuf:"varint,3,opt,name=paused,proto3" json:"paused,omitempty"`
	// Скорость воспроизведения.
	PlaybackSpeed float64 `protobuf:"fixed64,4,opt,name=playback_speed,json=playbackSpeed,proto3" json:"playback_speed,omitempty"`
	// Версия последнего изменения статуса воспроизведения.
	Version *UpdateVersion `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *PlayingStatus) Reset() {
	*x = PlayingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ynisonstate_playing_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayingStatus) ProtoMessage() {}

func (x *PlayingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_ynisonstate_playing_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayingStatus.ProtoReflect.Descriptor instead.
func (*PlayingStatus) Descriptor() ([]byte, []int) {
	return file_ynisonstate_playing_status_proto_rawDescGZIP(), []int{0}
}

func (x *PlayingStatus) GetProgressMs() int64 {
	if x != nil {
		return x.ProgressMs
	}
	return 0
}

func (x *PlayingStatus) GetDurationMs() int64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *PlayingStatus) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

func (x *PlayingStatus) GetPlaybackSpeed() float64 {
	if x != nil {
		return x.PlaybackSpeed
	}
	return 0
}

func (x *PlayingStatus) GetVersion() *UpdateVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

var File_ynisonstate_playing_status_proto protoreflect.FileDescriptor

var file_ynisonstate_playing_status_proto_rawDesc = []byte{
	0x0a, 0x20, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c,
	0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x1a, 0x20, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x4d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x67, 0x0a, 0x1f,
	0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x2e, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x01, 0x5a, 0x42, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x72, 0x75, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2d, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x2f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ynisonstate_playing_status_proto_rawDescOnce sync.Once
	file_ynisonstate_playing_status_proto_rawDescData = file_ynisonstate_playing_status_proto_rawDesc
)

func file_ynisonstate_playing_status_proto_rawDescGZIP() []byte {
	file_ynisonstate_playing_status_proto_rawDescOnce.Do(func() {
		file_ynisonstate_playing_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_ynisonstate_playing_status_proto_rawDescData)
	})
	return file_ynisonstate_playing_status_proto_rawDescData
}

var file_ynisonstate_playing_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ynisonstate_playing_status_proto_goTypes = []interface{}{
	(*PlayingStatus)(nil), // 0: ynison_state.PlayingStatus
	(*UpdateVersion)(nil), // 1: ynison_state.UpdateVersion
}
var file_ynisonstate_playing_status_proto_depIdxs = []int32{
	1, // 0: ynison_state.PlayingStatus.version:type_name -> ynison_state.UpdateVersion
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ynisonstate_playing_status_proto_init() }
func file_ynisonstate_playing_status_proto_init() {
	if File_ynisonstate_playing_status_proto != nil {
		return
	}
	file_ynisonstate_update_version_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ynisonstate_playing_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayingStatus); i {
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
			RawDescriptor: file_ynisonstate_playing_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ynisonstate_playing_status_proto_goTypes,
		DependencyIndexes: file_ynisonstate_playing_status_proto_depIdxs,
		MessageInfos:      file_ynisonstate_playing_status_proto_msgTypes,
	}.Build()
	File_ynisonstate_playing_status_proto = out.File
	file_ynisonstate_playing_status_proto_rawDesc = nil
	file_ynisonstate_playing_status_proto_goTypes = nil
	file_ynisonstate_playing_status_proto_depIdxs = nil
}
