// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: ynisonstate/player_queue_inject.proto

package ynisonstate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlayerQueueInject_Playable_PlayableType int32

const (
	PlayerQueueInject_Playable_UNSPECIFIED PlayerQueueInject_Playable_PlayableType = 0
	PlayerQueueInject_Playable_ALICE_SHOT  PlayerQueueInject_Playable_PlayableType = 1
	PlayerQueueInject_Playable_AD          PlayerQueueInject_Playable_PlayableType = 2
	PlayerQueueInject_Playable_PREROLL     PlayerQueueInject_Playable_PlayableType = 3
)

// Enum value maps for PlayerQueueInject_Playable_PlayableType.
var (
	PlayerQueueInject_Playable_PlayableType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ALICE_SHOT",
		2: "AD",
		3: "PREROLL",
	}
	PlayerQueueInject_Playable_PlayableType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"ALICE_SHOT":  1,
		"AD":          2,
		"PREROLL":     3,
	}
)

func (x PlayerQueueInject_Playable_PlayableType) Enum() *PlayerQueueInject_Playable_PlayableType {
	p := new(PlayerQueueInject_Playable_PlayableType)
	*p = x
	return p
}

func (x PlayerQueueInject_Playable_PlayableType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerQueueInject_Playable_PlayableType) Descriptor() protoreflect.EnumDescriptor {
	return file_ynisonstate_player_queue_inject_proto_enumTypes[0].Descriptor()
}

func (PlayerQueueInject_Playable_PlayableType) Type() protoreflect.EnumType {
	return &file_ynisonstate_player_queue_inject_proto_enumTypes[0]
}

func (x PlayerQueueInject_Playable_PlayableType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerQueueInject_Playable_PlayableType.Descriptor instead.
func (PlayerQueueInject_Playable_PlayableType) EnumDescriptor() ([]byte, []int) {
	return file_ynisonstate_player_queue_inject_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Состояние проигрывания ижектируемой в очередь сущности.
// Инжектироваться может шот, преролл и проч.
type PlayerQueueInject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Статус воспроизведения.
	PlayingStatus *PlayingStatus `protobuf:"bytes,1,opt,name=playing_status,json=playingStatus,proto3" json:"playing_status,omitempty"`
	// Доп. проигрываемая сущность.
	Playable *PlayerQueueInject_Playable `protobuf:"bytes,2,opt,name=playable,proto3" json:"playable,omitempty"`
	// Версия последнего изменения состояния.
	Version *UpdateVersion `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *PlayerQueueInject) Reset() {
	*x = PlayerQueueInject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ynisonstate_player_queue_inject_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerQueueInject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerQueueInject) ProtoMessage() {}

func (x *PlayerQueueInject) ProtoReflect() protoreflect.Message {
	mi := &file_ynisonstate_player_queue_inject_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerQueueInject.ProtoReflect.Descriptor instead.
func (*PlayerQueueInject) Descriptor() ([]byte, []int) {
	return file_ynisonstate_player_queue_inject_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerQueueInject) GetPlayingStatus() *PlayingStatus {
	if x != nil {
		return x.PlayingStatus
	}
	return nil
}

func (x *PlayerQueueInject) GetPlayable() *PlayerQueueInject_Playable {
	if x != nil {
		return x.Playable
	}
	return nil
}

func (x *PlayerQueueInject) GetVersion() *UpdateVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

type PlayerQueueInject_Playable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор сущности.
	PlayableId string `protobuf:"bytes,1,opt,name=playable_id,json=playableId,proto3" json:"playable_id,omitempty"`
	// Тип сущности.
	PlayableType PlayerQueueInject_Playable_PlayableType `protobuf:"varint,2,opt,name=playable_type,json=playableType,proto3,enum=ynison_state.PlayerQueueInject_Playable_PlayableType" json:"playable_type,omitempty"`
	// Заголовок.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// Опциональная ссылка на обложку.
	// Может содержать плейсхолдер для размера в аватарнице.
	CoverUrl *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
}

func (x *PlayerQueueInject_Playable) Reset() {
	*x = PlayerQueueInject_Playable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ynisonstate_player_queue_inject_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerQueueInject_Playable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerQueueInject_Playable) ProtoMessage() {}

func (x *PlayerQueueInject_Playable) ProtoReflect() protoreflect.Message {
	mi := &file_ynisonstate_player_queue_inject_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerQueueInject_Playable.ProtoReflect.Descriptor instead.
func (*PlayerQueueInject_Playable) Descriptor() ([]byte, []int) {
	return file_ynisonstate_player_queue_inject_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PlayerQueueInject_Playable) GetPlayableId() string {
	if x != nil {
		return x.PlayableId
	}
	return ""
}

func (x *PlayerQueueInject_Playable) GetPlayableType() PlayerQueueInject_Playable_PlayableType {
	if x != nil {
		return x.PlayableType
	}
	return PlayerQueueInject_Playable_UNSPECIFIED
}

func (x *PlayerQueueInject_Playable) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PlayerQueueInject_Playable) GetCoverUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.CoverUrl
	}
	return nil
}

var File_ynisonstate_player_queue_inject_proto protoreflect.FileDescriptor

var file_ynisonstate_player_queue_inject_proto_rawDesc = []byte{
	0x0a, 0x25, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x20, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x03, 0x0a, 0x11, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x42, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x44, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x79, 0x6e, 0x69,
	0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x1a, 0x9e, 0x02, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x5a,
	0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6c,
	0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x39, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x44, 0x0a, 0x0c, 0x50,
	0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x41, 0x4c, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x48, 0x4f, 0x54, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02,
	0x41, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x52, 0x4f, 0x4c, 0x4c, 0x10,
	0x03, 0x42, 0x67, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x42, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2d, 0x79, 0x6e, 0x69,
	0x73, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79,
	0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ynisonstate_player_queue_inject_proto_rawDescOnce sync.Once
	file_ynisonstate_player_queue_inject_proto_rawDescData = file_ynisonstate_player_queue_inject_proto_rawDesc
)

func file_ynisonstate_player_queue_inject_proto_rawDescGZIP() []byte {
	file_ynisonstate_player_queue_inject_proto_rawDescOnce.Do(func() {
		file_ynisonstate_player_queue_inject_proto_rawDescData = protoimpl.X.CompressGZIP(file_ynisonstate_player_queue_inject_proto_rawDescData)
	})
	return file_ynisonstate_player_queue_inject_proto_rawDescData
}

var file_ynisonstate_player_queue_inject_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ynisonstate_player_queue_inject_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ynisonstate_player_queue_inject_proto_goTypes = []interface{}{
	(PlayerQueueInject_Playable_PlayableType)(0), // 0: ynison_state.PlayerQueueInject.Playable.PlayableType
	(*PlayerQueueInject)(nil),                    // 1: ynison_state.PlayerQueueInject
	(*PlayerQueueInject_Playable)(nil),           // 2: ynison_state.PlayerQueueInject.Playable
	(*PlayingStatus)(nil),                        // 3: ynison_state.PlayingStatus
	(*UpdateVersion)(nil),                        // 4: ynison_state.UpdateVersion
	(*wrapperspb.StringValue)(nil),               // 5: google.protobuf.StringValue
}
var file_ynisonstate_player_queue_inject_proto_depIdxs = []int32{
	3, // 0: ynison_state.PlayerQueueInject.playing_status:type_name -> ynison_state.PlayingStatus
	2, // 1: ynison_state.PlayerQueueInject.playable:type_name -> ynison_state.PlayerQueueInject.Playable
	4, // 2: ynison_state.PlayerQueueInject.version:type_name -> ynison_state.UpdateVersion
	0, // 3: ynison_state.PlayerQueueInject.Playable.playable_type:type_name -> ynison_state.PlayerQueueInject.Playable.PlayableType
	5, // 4: ynison_state.PlayerQueueInject.Playable.cover_url:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ynisonstate_player_queue_inject_proto_init() }
func file_ynisonstate_player_queue_inject_proto_init() {
	if File_ynisonstate_player_queue_inject_proto != nil {
		return
	}
	file_ynisonstate_update_version_proto_init()
	file_ynisonstate_playing_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ynisonstate_player_queue_inject_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerQueueInject); i {
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
		file_ynisonstate_player_queue_inject_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerQueueInject_Playable); i {
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
			RawDescriptor: file_ynisonstate_player_queue_inject_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ynisonstate_player_queue_inject_proto_goTypes,
		DependencyIndexes: file_ynisonstate_player_queue_inject_proto_depIdxs,
		EnumInfos:         file_ynisonstate_player_queue_inject_proto_enumTypes,
		MessageInfos:      file_ynisonstate_player_queue_inject_proto_msgTypes,
	}.Build()
	File_ynisonstate_player_queue_inject_proto = out.File
	file_ynisonstate_player_queue_inject_proto_rawDesc = nil
	file_ynisonstate_player_queue_inject_proto_goTypes = nil
	file_ynisonstate_player_queue_inject_proto_depIdxs = nil
}
