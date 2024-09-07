// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: ynisonstate/playable.proto

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

type Playable_PlayableType int32

const (
	Playable_UNSPECIFIED Playable_PlayableType = 0
	Playable_TRACK       Playable_PlayableType = 1
	// Локальный файл.
	Playable_LOCAL_TRACK Playable_PlayableType = 2
	// "Бесконечный" трек. Например, для генеративки или fm-радио.
	Playable_INFINITE Playable_PlayableType = 3
	// Видео клип
	Playable_VIDEO_CLIP Playable_PlayableType = 4
)

// Enum value maps for Playable_PlayableType.
var (
	Playable_PlayableType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "TRACK",
		2: "LOCAL_TRACK",
		3: "INFINITE",
		4: "VIDEO_CLIP",
	}
	Playable_PlayableType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"TRACK":       1,
		"LOCAL_TRACK": 2,
		"INFINITE":    3,
		"VIDEO_CLIP":  4,
	}
)

func (x Playable_PlayableType) Enum() *Playable_PlayableType {
	p := new(Playable_PlayableType)
	*p = x
	return p
}

func (x Playable_PlayableType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Playable_PlayableType) Descriptor() protoreflect.EnumDescriptor {
	return file_ynisonstate_playable_proto_enumTypes[0].Descriptor()
}

func (Playable_PlayableType) Type() protoreflect.EnumType {
	return &file_ynisonstate_playable_proto_enumTypes[0]
}

func (x Playable_PlayableType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Playable_PlayableType.Descriptor instead.
func (Playable_PlayableType) EnumDescriptor() ([]byte, []int) {
	return file_ynisonstate_playable_proto_rawDescGZIP(), []int{0, 0}
}

type VideoClipInfo_RecommendationType int32

const (
	VideoClipInfo_UNSPECIFIED VideoClipInfo_RecommendationType = 0
	// видео клип пришёл от рекоментадельного бэкенда
	VideoClipInfo_RECOMMENDED VideoClipInfo_RecommendationType = 1
	// видео клип явно выбран пользователем
	VideoClipInfo_ON_DEMAND VideoClipInfo_RecommendationType = 2
)

// Enum value maps for VideoClipInfo_RecommendationType.
var (
	VideoClipInfo_RecommendationType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "RECOMMENDED",
		2: "ON_DEMAND",
	}
	VideoClipInfo_RecommendationType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"RECOMMENDED": 1,
		"ON_DEMAND":   2,
	}
)

func (x VideoClipInfo_RecommendationType) Enum() *VideoClipInfo_RecommendationType {
	p := new(VideoClipInfo_RecommendationType)
	*p = x
	return p
}

func (x VideoClipInfo_RecommendationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VideoClipInfo_RecommendationType) Descriptor() protoreflect.EnumDescriptor {
	return file_ynisonstate_playable_proto_enumTypes[1].Descriptor()
}

func (VideoClipInfo_RecommendationType) Type() protoreflect.EnumType {
	return &file_ynisonstate_playable_proto_enumTypes[1]
}

func (x VideoClipInfo_RecommendationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VideoClipInfo_RecommendationType.Descriptor instead.
func (VideoClipInfo_RecommendationType) EnumDescriptor() ([]byte, []int) {
	return file_ynisonstate_playable_proto_rawDescGZIP(), []int{1, 0}
}

// Проигрываемая сущность.
// Может быть треком, видео и т.п.
type Playable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор сущности.
	PlayableId string `protobuf:"bytes,1,opt,name=playable_id,json=playableId,proto3" json:"playable_id,omitempty"`
	// Опциональный идентификатор альбома.
	// Используется для составного идентификатора playable при playable_type == .TRACK.
	//
	// TODO: Нужен рефакторинг с переходом на oneof для типов playable и идентификаторов,
	// относящихся к этим типам.
	AlbumIdOptional *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=album_id_optional,json=albumIdOptional,proto3" json:"album_id_optional,omitempty"`
	// Тип сущности.
	PlayableType Playable_PlayableType `protobuf:"varint,3,opt,name=playable_type,json=playableType,proto3,enum=ynison_state.Playable_PlayableType" json:"playable_type,omitempty"`
	// Фром для play-audio.
	From string `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	// Заголовок.
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	// Опциональная ссылка на обложку.
	// Может содержать плейсхолдер для размера в аватарнице.
	CoverUrlOptional *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=cover_url_optional,json=coverUrlOptional,proto3" json:"cover_url_optional,omitempty"`
	// Дополнительные параметры плейабла, зависит от типа плейабла
	//
	// Types that are assignable to AdditionalInfoOptional:
	//
	//	*Playable_VideoClipInfo
	//	*Playable_TrackInfo
	AdditionalInfoOptional isPlayable_AdditionalInfoOptional `protobuf_oneof:"additional_info_optional"`
}

func (x *Playable) Reset() {
	*x = Playable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ynisonstate_playable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playable) ProtoMessage() {}

func (x *Playable) ProtoReflect() protoreflect.Message {
	mi := &file_ynisonstate_playable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playable.ProtoReflect.Descriptor instead.
func (*Playable) Descriptor() ([]byte, []int) {
	return file_ynisonstate_playable_proto_rawDescGZIP(), []int{0}
}

func (x *Playable) GetPlayableId() string {
	if x != nil {
		return x.PlayableId
	}
	return ""
}

func (x *Playable) GetAlbumIdOptional() *wrapperspb.StringValue {
	if x != nil {
		return x.AlbumIdOptional
	}
	return nil
}

func (x *Playable) GetPlayableType() Playable_PlayableType {
	if x != nil {
		return x.PlayableType
	}
	return Playable_UNSPECIFIED
}

func (x *Playable) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Playable) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Playable) GetCoverUrlOptional() *wrapperspb.StringValue {
	if x != nil {
		return x.CoverUrlOptional
	}
	return nil
}

func (m *Playable) GetAdditionalInfoOptional() isPlayable_AdditionalInfoOptional {
	if m != nil {
		return m.AdditionalInfoOptional
	}
	return nil
}

func (x *Playable) GetVideoClipInfo() *VideoClipInfo {
	if x, ok := x.GetAdditionalInfoOptional().(*Playable_VideoClipInfo); ok {
		return x.VideoClipInfo
	}
	return nil
}

func (x *Playable) GetTrackInfo() *TrackInfo {
	if x, ok := x.GetAdditionalInfoOptional().(*Playable_TrackInfo); ok {
		return x.TrackInfo
	}
	return nil
}

type isPlayable_AdditionalInfoOptional interface {
	isPlayable_AdditionalInfoOptional()
}

type Playable_VideoClipInfo struct {
	VideoClipInfo *VideoClipInfo `protobuf:"bytes,7,opt,name=video_clip_info,json=videoClipInfo,proto3,oneof"`
}

type Playable_TrackInfo struct {
	TrackInfo *TrackInfo `protobuf:"bytes,8,opt,name=track_info,json=trackInfo,proto3,oneof"`
}

func (*Playable_VideoClipInfo) isPlayable_AdditionalInfoOptional() {}

func (*Playable_TrackInfo) isPlayable_AdditionalInfoOptional() {}

type VideoClipInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Это поле содержит информацию для аналитики, обозначает способ попадания видео клипа в очередь
	RecommendationType VideoClipInfo_RecommendationType `protobuf:"varint,1,opt,name=recommendation_type,json=recommendationType,proto3,enum=ynison_state.VideoClipInfo_RecommendationType" json:"recommendation_type,omitempty"`
}

func (x *VideoClipInfo) Reset() {
	*x = VideoClipInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ynisonstate_playable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoClipInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoClipInfo) ProtoMessage() {}

func (x *VideoClipInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ynisonstate_playable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoClipInfo.ProtoReflect.Descriptor instead.
func (*VideoClipInfo) Descriptor() ([]byte, []int) {
	return file_ynisonstate_playable_proto_rawDescGZIP(), []int{1}
}

func (x *VideoClipInfo) GetRecommendationType() VideoClipInfo_RecommendationType {
	if x != nil {
		return x.RecommendationType
	}
	return VideoClipInfo_UNSPECIFIED
}

// Дополнительная информация о треке
type TrackInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ключ источника из [WaveQueue.EntityOptions.track_sources] (queue.proto)
	// См. `WaveQueue.EntityOptions.TrackSourceWithKey`
	TrackSourceKey uint32 `protobuf:"varint,1,opt,name=track_source_key,json=trackSourceKey,proto3" json:"track_source_key,omitempty"`
	// Относится поле к трекам, полученным из рекомендаций, т.е. трекам в волновой очереди.
	// Любой полученный из рекомендаций трек относится к определенному `batchId`.
	// Потеря batchId крайне нежелательна. Рекомендациям этот id очень важен.
	// В волновой очереди также можно ставить треки следующими через "Играть следующим".
	// Этим трекам необходим `batchId`, который генерируется на клиенте, и который
	// нужно переносить.
	//
	// `batchId` *прослушанных* в рамках rotor-сессии треков можно не переносить.
	BatchIdOptional *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=batch_id_optional,json=batchIdOptional,proto3" json:"batch_id_optional,omitempty"`
}

func (x *TrackInfo) Reset() {
	*x = TrackInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ynisonstate_playable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackInfo) ProtoMessage() {}

func (x *TrackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ynisonstate_playable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackInfo.ProtoReflect.Descriptor instead.
func (*TrackInfo) Descriptor() ([]byte, []int) {
	return file_ynisonstate_playable_proto_rawDescGZIP(), []int{2}
}

func (x *TrackInfo) GetTrackSourceKey() uint32 {
	if x != nil {
		return x.TrackSourceKey
	}
	return 0
}

func (x *TrackInfo) GetBatchIdOptional() *wrapperspb.StringValue {
	if x != nil {
		return x.BatchIdOptional
	}
	return nil
}

var File_ynisonstate_playable_proto protoreflect.FileDescriptor

var file_ynisonstate_playable_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c,
	0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x79, 0x6e,
	0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x04, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x11, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x5f, 0x69, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x12, 0x48, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x79, 0x6e, 0x69, 0x73,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x12, 0x45, 0x0a, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6c, 0x69, 0x70,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x79, 0x6e,
	0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x43, 0x6c, 0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0d, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x43, 0x6c, 0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x59, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0e,
	0x0a, 0x0a, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x43, 0x4c, 0x49, 0x50, 0x10, 0x04, 0x42, 0x1a,
	0x0a, 0x18, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5f, 0x0a, 0x13,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x79, 0x6e, 0x69, 0x73,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c,
	0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a,
	0x12, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e,
	0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4d, 0x41,
	0x4e, 0x44, 0x10, 0x02, 0x22, 0x7f, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x11, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x66, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x41, 0x61, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x2d, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x6e, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ynisonstate_playable_proto_rawDescOnce sync.Once
	file_ynisonstate_playable_proto_rawDescData = file_ynisonstate_playable_proto_rawDesc
)

func file_ynisonstate_playable_proto_rawDescGZIP() []byte {
	file_ynisonstate_playable_proto_rawDescOnce.Do(func() {
		file_ynisonstate_playable_proto_rawDescData = protoimpl.X.CompressGZIP(file_ynisonstate_playable_proto_rawDescData)
	})
	return file_ynisonstate_playable_proto_rawDescData
}

var file_ynisonstate_playable_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ynisonstate_playable_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ynisonstate_playable_proto_goTypes = []interface{}{
	(Playable_PlayableType)(0),            // 0: ynison_state.Playable.PlayableType
	(VideoClipInfo_RecommendationType)(0), // 1: ynison_state.VideoClipInfo.RecommendationType
	(*Playable)(nil),                      // 2: ynison_state.Playable
	(*VideoClipInfo)(nil),                 // 3: ynison_state.VideoClipInfo
	(*TrackInfo)(nil),                     // 4: ynison_state.TrackInfo
	(*wrapperspb.StringValue)(nil),        // 5: google.protobuf.StringValue
}
var file_ynisonstate_playable_proto_depIdxs = []int32{
	5, // 0: ynison_state.Playable.album_id_optional:type_name -> google.protobuf.StringValue
	0, // 1: ynison_state.Playable.playable_type:type_name -> ynison_state.Playable.PlayableType
	5, // 2: ynison_state.Playable.cover_url_optional:type_name -> google.protobuf.StringValue
	3, // 3: ynison_state.Playable.video_clip_info:type_name -> ynison_state.VideoClipInfo
	4, // 4: ynison_state.Playable.track_info:type_name -> ynison_state.TrackInfo
	1, // 5: ynison_state.VideoClipInfo.recommendation_type:type_name -> ynison_state.VideoClipInfo.RecommendationType
	5, // 6: ynison_state.TrackInfo.batch_id_optional:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_ynisonstate_playable_proto_init() }
func file_ynisonstate_playable_proto_init() {
	if File_ynisonstate_playable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ynisonstate_playable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Playable); i {
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
		file_ynisonstate_playable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoClipInfo); i {
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
		file_ynisonstate_playable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackInfo); i {
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
	file_ynisonstate_playable_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Playable_VideoClipInfo)(nil),
		(*Playable_TrackInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ynisonstate_playable_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ynisonstate_playable_proto_goTypes,
		DependencyIndexes: file_ynisonstate_playable_proto_depIdxs,
		EnumInfos:         file_ynisonstate_playable_proto_enumTypes,
		MessageInfos:      file_ynisonstate_playable_proto_msgTypes,
	}.Build()
	File_ynisonstate_playable_proto = out.File
	file_ynisonstate_playable_proto_rawDesc = nil
	file_ynisonstate_playable_proto_goTypes = nil
	file_ynisonstate_playable_proto_depIdxs = nil
}
