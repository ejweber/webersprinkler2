// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: proto/common.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DayAndTime_Day int32

const (
	DayAndTime_SUNDAY    DayAndTime_Day = 0
	DayAndTime_MONDAY    DayAndTime_Day = 1
	DayAndTime_TUESDAY   DayAndTime_Day = 2
	DayAndTime_WEDNESDAY DayAndTime_Day = 3
	DayAndTime_THURSDAY  DayAndTime_Day = 4
	DayAndTime_FRIDAY    DayAndTime_Day = 5
	DayAndTime_SATURDAY  DayAndTime_Day = 6
)

// Enum value maps for DayAndTime_Day.
var (
	DayAndTime_Day_name = map[int32]string{
		0: "SUNDAY",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
	}
	DayAndTime_Day_value = map[string]int32{
		"SUNDAY":    0,
		"MONDAY":    1,
		"TUESDAY":   2,
		"WEDNESDAY": 3,
		"THURSDAY":  4,
		"FRIDAY":    5,
		"SATURDAY":  6,
	}
)

func (x DayAndTime_Day) Enum() *DayAndTime_Day {
	p := new(DayAndTime_Day)
	*p = x
	return p
}

func (x DayAndTime_Day) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DayAndTime_Day) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_common_proto_enumTypes[0].Descriptor()
}

func (DayAndTime_Day) Type() protoreflect.EnumType {
	return &file_proto_common_proto_enumTypes[0]
}

func (x DayAndTime_Day) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DayAndTime_Day.Descriptor instead.
func (DayAndTime_Day) EnumDescriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{2, 0}
}

type ZoneTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zone uint32 `protobuf:"varint,1,opt,name=zone,proto3" json:"zone,omitempty"`
	Secs uint32 `protobuf:"varint,2,opt,name=secs,proto3" json:"secs,omitempty"`
}

func (x *ZoneTime) Reset() {
	*x = ZoneTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneTime) ProtoMessage() {}

func (x *ZoneTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneTime.ProtoReflect.Descriptor instead.
func (*ZoneTime) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneTime) GetZone() uint32 {
	if x != nil {
		return x.Zone
	}
	return 0
}

func (x *ZoneTime) GetSecs() uint32 {
	if x != nil {
		return x.Secs
	}
	return 0
}

type Program struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ZoneTime []*ZoneTime `protobuf:"bytes,2,rep,name=zone_time,json=zoneTime,proto3" json:"zone_time,omitempty"`
}

func (x *Program) Reset() {
	*x = Program{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Program) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Program) ProtoMessage() {}

func (x *Program) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Program.ProtoReflect.Descriptor instead.
func (*Program) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{1}
}

func (x *Program) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Program) GetZoneTime() []*ZoneTime {
	if x != nil {
		return x.ZoneTime
	}
	return nil
}

type DayAndTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day              DayAndTime_Day `protobuf:"varint,1,opt,name=day,proto3,enum=DayAndTime_Day" json:"day,omitempty"`
	SecsFromMidnight uint32         `protobuf:"varint,2,opt,name=secs_from_midnight,json=secsFromMidnight,proto3" json:"secs_from_midnight,omitempty"`
	RemainingRepeats uint32         `protobuf:"varint,3,opt,name=remaining_repeats,json=remainingRepeats,proto3" json:"remaining_repeats,omitempty"`
}

func (x *DayAndTime) Reset() {
	*x = DayAndTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DayAndTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DayAndTime) ProtoMessage() {}

func (x *DayAndTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DayAndTime.ProtoReflect.Descriptor instead.
func (*DayAndTime) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{2}
}

func (x *DayAndTime) GetDay() DayAndTime_Day {
	if x != nil {
		return x.Day
	}
	return DayAndTime_SUNDAY
}

func (x *DayAndTime) GetSecsFromMidnight() uint32 {
	if x != nil {
		return x.SecsFromMidnight
	}
	return 0
}

func (x *DayAndTime) GetRemainingRepeats() uint32 {
	if x != nil {
		return x.RemainingRepeats
	}
	return 0
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DaysAndTimes []*DayAndTime `protobuf:"bytes,1,rep,name=days_and_times,json=daysAndTimes,proto3" json:"days_and_times,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{3}
}

func (x *Schedule) GetDaysAndTimes() []*DayAndTime {
	if x != nil {
		return x.DaysAndTimes
	}
	return nil
}

type ScheduledProgram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Program  *Program  `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty"`
	Schedule *Schedule `protobuf:"bytes,2,opt,name=schedule,proto3" json:"schedule,omitempty"`
}

func (x *ScheduledProgram) Reset() {
	*x = ScheduledProgram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduledProgram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduledProgram) ProtoMessage() {}

func (x *ScheduledProgram) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduledProgram.ProtoReflect.Descriptor instead.
func (*ScheduledProgram) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{4}
}

func (x *ScheduledProgram) GetProgram() *Program {
	if x != nil {
		return x.Program
	}
	return nil
}

func (x *ScheduledProgram) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

type ScheduledPrograms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduledPrograms []*ScheduledProgram `protobuf:"bytes,1,rep,name=scheduled_programs,json=scheduledPrograms,proto3" json:"scheduled_programs,omitempty"`
}

func (x *ScheduledPrograms) Reset() {
	*x = ScheduledPrograms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduledPrograms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduledPrograms) ProtoMessage() {}

func (x *ScheduledPrograms) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduledPrograms.ProtoReflect.Descriptor instead.
func (*ScheduledPrograms) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{5}
}

func (x *ScheduledPrograms) GetScheduledPrograms() []*ScheduledProgram {
	if x != nil {
		return x.ScheduledPrograms
	}
	return nil
}

var File_proto_common_proto protoreflect.FileDescriptor

var file_proto_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x73, 0x65, 0x63, 0x73, 0x22, 0x45, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x7a, 0x6f, 0x6e, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x5a, 0x6f, 0x6e,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x7a, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xed, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x79, 0x41, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x44, 0x61,
	0x79, 0x41, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x2e, 0x44, 0x61, 0x79, 0x52, 0x03, 0x64, 0x61,
	0x79, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x63, 0x73, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6d,
	0x69, 0x64, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x73,
	0x65, 0x63, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x64, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x2b, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x73, 0x22, 0x61, 0x0a, 0x03,
	0x44, 0x61, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54,
	0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x45, 0x44, 0x4e,
	0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x48, 0x55, 0x52, 0x53,
	0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x49, 0x44, 0x41, 0x59, 0x10,
	0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59, 0x10, 0x06, 0x22,
	0x3d, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x0e, 0x64,
	0x61, 0x79, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x44, 0x61, 0x79, 0x41, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x0c, 0x64, 0x61, 0x79, 0x73, 0x41, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x5d,
	0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x25, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x55, 0x0a,
	0x11, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x40, 0x0a, 0x12, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x52, 0x11, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6a, 0x77, 0x65, 0x62, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x62, 0x65, 0x72,
	0x73, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_proto_rawDescOnce sync.Once
	file_proto_common_proto_rawDescData = file_proto_common_proto_rawDesc
)

func file_proto_common_proto_rawDescGZIP() []byte {
	file_proto_common_proto_rawDescOnce.Do(func() {
		file_proto_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_proto_rawDescData)
	})
	return file_proto_common_proto_rawDescData
}

var file_proto_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_common_proto_goTypes = []interface{}{
	(DayAndTime_Day)(0),       // 0: DayAndTime.Day
	(*ZoneTime)(nil),          // 1: ZoneTime
	(*Program)(nil),           // 2: Program
	(*DayAndTime)(nil),        // 3: DayAndTime
	(*Schedule)(nil),          // 4: Schedule
	(*ScheduledProgram)(nil),  // 5: ScheduledProgram
	(*ScheduledPrograms)(nil), // 6: ScheduledPrograms
}
var file_proto_common_proto_depIdxs = []int32{
	1, // 0: Program.zone_time:type_name -> ZoneTime
	0, // 1: DayAndTime.day:type_name -> DayAndTime.Day
	3, // 2: Schedule.days_and_times:type_name -> DayAndTime
	2, // 3: ScheduledProgram.program:type_name -> Program
	4, // 4: ScheduledProgram.schedule:type_name -> Schedule
	5, // 5: ScheduledPrograms.scheduled_programs:type_name -> ScheduledProgram
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_common_proto_init() }
func file_proto_common_proto_init() {
	if File_proto_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneTime); i {
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
		file_proto_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Program); i {
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
		file_proto_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DayAndTime); i {
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
		file_proto_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_proto_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduledProgram); i {
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
		file_proto_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduledPrograms); i {
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
			RawDescriptor: file_proto_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_proto_goTypes,
		DependencyIndexes: file_proto_common_proto_depIdxs,
		EnumInfos:         file_proto_common_proto_enumTypes,
		MessageInfos:      file_proto_common_proto_msgTypes,
	}.Build()
	File_proto_common_proto = out.File
	file_proto_common_proto_rawDesc = nil
	file_proto_common_proto_goTypes = nil
	file_proto_common_proto_depIdxs = nil
}
