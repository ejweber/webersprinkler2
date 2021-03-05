// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: proto/sprinkler_to_server.proto

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

type SprinklerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SprinklerStatus *SprinklerStatus `protobuf:"bytes,1,opt,name=sprinkler_status,json=sprinklerStatus,proto3" json:"sprinkler_status,omitempty"`
}

func (x *SprinklerStatusRequest) Reset() {
	*x = SprinklerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprinklerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprinklerStatusRequest) ProtoMessage() {}

func (x *SprinklerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprinklerStatusRequest.ProtoReflect.Descriptor instead.
func (*SprinklerStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{0}
}

func (x *SprinklerStatusRequest) GetSprinklerStatus() *SprinklerStatus {
	if x != nil {
		return x.SprinklerStatus
	}
	return nil
}

type SprinklerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SprinklerStatusResponse) Reset() {
	*x = SprinklerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprinklerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprinklerStatusResponse) ProtoMessage() {}

func (x *SprinklerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprinklerStatusResponse.ProtoReflect.Descriptor instead.
func (*SprinklerStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{1}
}

type GetProgramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Program *Program `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty"`
}

func (x *GetProgramRequest) Reset() {
	*x = GetProgramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProgramRequest) ProtoMessage() {}

func (x *GetProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProgramRequest.ProtoReflect.Descriptor instead.
func (*GetProgramRequest) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetProgramRequest) GetProgram() *Program {
	if x != nil {
		return x.Program
	}
	return nil
}

type GetProgramResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Program *Program `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty"`
}

func (x *GetProgramResponse) Reset() {
	*x = GetProgramResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProgramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProgramResponse) ProtoMessage() {}

func (x *GetProgramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProgramResponse.ProtoReflect.Descriptor instead.
func (*GetProgramResponse) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{3}
}

func (x *GetProgramResponse) GetProgram() *Program {
	if x != nil {
		return x.Program
	}
	return nil
}

type GetScheduledProgramsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetScheduledProgramsRequest) Reset() {
	*x = GetScheduledProgramsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScheduledProgramsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScheduledProgramsRequest) ProtoMessage() {}

func (x *GetScheduledProgramsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScheduledProgramsRequest.ProtoReflect.Descriptor instead.
func (*GetScheduledProgramsRequest) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{4}
}

type GetScheduledProgramsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduledPrograms []*ScheduledProgram `protobuf:"bytes,1,rep,name=scheduled_programs,json=scheduledPrograms,proto3" json:"scheduled_programs,omitempty"`
}

func (x *GetScheduledProgramsResponse) Reset() {
	*x = GetScheduledProgramsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScheduledProgramsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScheduledProgramsResponse) ProtoMessage() {}

func (x *GetScheduledProgramsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScheduledProgramsResponse.ProtoReflect.Descriptor instead.
func (*GetScheduledProgramsResponse) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetScheduledProgramsResponse) GetScheduledPrograms() []*ScheduledProgram {
	if x != nil {
		return x.ScheduledPrograms
	}
	return nil
}

type RecieveCommandsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecieveCommandsRequest) Reset() {
	*x = RecieveCommandsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecieveCommandsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecieveCommandsRequest) ProtoMessage() {}

func (x *RecieveCommandsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecieveCommandsRequest.ProtoReflect.Descriptor instead.
func (*RecieveCommandsRequest) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{6}
}

type SprinklerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SprinklerStatus) Reset() {
	*x = SprinklerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprinklerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprinklerStatus) ProtoMessage() {}

func (x *SprinklerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprinklerStatus.ProtoReflect.Descriptor instead.
func (*SprinklerStatus) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{7}
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//	*Command_Start
	//	*Command_Stop
	//	*Command_UpdateScheduledPrograms
	Command isCommand_Command `protobuf_oneof:"command"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{8}
}

func (m *Command) GetCommand() isCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *Command) GetStart() *StartCommand {
	if x, ok := x.GetCommand().(*Command_Start); ok {
		return x.Start
	}
	return nil
}

func (x *Command) GetStop() *StopCommand {
	if x, ok := x.GetCommand().(*Command_Stop); ok {
		return x.Stop
	}
	return nil
}

func (x *Command) GetUpdateScheduledPrograms() *UpdateScheduledProgramsCommand {
	if x, ok := x.GetCommand().(*Command_UpdateScheduledPrograms); ok {
		return x.UpdateScheduledPrograms
	}
	return nil
}

type isCommand_Command interface {
	isCommand_Command()
}

type Command_Start struct {
	Start *StartCommand `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type Command_Stop struct {
	Stop *StopCommand `protobuf:"bytes,2,opt,name=stop,proto3,oneof"`
}

type Command_UpdateScheduledPrograms struct {
	UpdateScheduledPrograms *UpdateScheduledProgramsCommand `protobuf:"bytes,3,opt,name=update_scheduled_programs,json=updateScheduledPrograms,proto3,oneof"`
}

func (*Command_Start) isCommand_Command() {}

func (*Command_Stop) isCommand_Command() {}

func (*Command_UpdateScheduledPrograms) isCommand_Command() {}

type StartCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*StartCommand_Program
	//	*StartCommand_ZoneTime
	Type isStartCommand_Type `protobuf_oneof:"type"`
}

func (x *StartCommand) Reset() {
	*x = StartCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCommand) ProtoMessage() {}

func (x *StartCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCommand.ProtoReflect.Descriptor instead.
func (*StartCommand) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{9}
}

func (m *StartCommand) GetType() isStartCommand_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *StartCommand) GetProgram() *Program {
	if x, ok := x.GetType().(*StartCommand_Program); ok {
		return x.Program
	}
	return nil
}

func (x *StartCommand) GetZoneTime() *ZoneTime {
	if x, ok := x.GetType().(*StartCommand_ZoneTime); ok {
		return x.ZoneTime
	}
	return nil
}

type isStartCommand_Type interface {
	isStartCommand_Type()
}

type StartCommand_Program struct {
	Program *Program `protobuf:"bytes,1,opt,name=program,proto3,oneof"`
}

type StartCommand_ZoneTime struct {
	ZoneTime *ZoneTime `protobuf:"bytes,2,opt,name=zone_time,json=zoneTime,proto3,oneof"`
}

func (*StartCommand_Program) isStartCommand_Type() {}

func (*StartCommand_ZoneTime) isStartCommand_Type() {}

type StopCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopCommand) Reset() {
	*x = StopCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopCommand) ProtoMessage() {}

func (x *StopCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopCommand.ProtoReflect.Descriptor instead.
func (*StopCommand) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{10}
}

type UpdateScheduledProgramsCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateScheduledProgramsCommand) Reset() {
	*x = UpdateScheduledProgramsCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sprinkler_to_server_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScheduledProgramsCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScheduledProgramsCommand) ProtoMessage() {}

func (x *UpdateScheduledProgramsCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sprinkler_to_server_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScheduledProgramsCommand.ProtoReflect.Descriptor instead.
func (*UpdateScheduledProgramsCommand) Descriptor() ([]byte, []int) {
	return file_proto_sprinkler_to_server_proto_rawDescGZIP(), []int{11}
}

var File_proto_sprinkler_to_server_proto protoreflect.FileDescriptor

var file_proto_sprinkler_to_server_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c, 0x65,
	0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x16, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3b, 0x0a, 0x10, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x70, 0x72, 0x69,
	0x6e, 0x6b, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x73, 0x70, 0x72,
	0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x19, 0x0a, 0x17,
	0x53, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x12, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x11, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x52,
	0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x73,
	0x74, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12,
	0x5d, 0x0a, 0x19, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x09,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x66, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12,
	0x28, 0x0a, 0x09, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x00, 0x52,
	0x08, 0x7a, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x20, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x32, 0x98, 0x02, 0x0a, 0x11, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72,
	0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x65,
	0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x30, 0x01, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6a, 0x77, 0x65,
	0x62, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x62, 0x65, 0x72, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c,
	0x65, 0x72, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_sprinkler_to_server_proto_rawDescOnce sync.Once
	file_proto_sprinkler_to_server_proto_rawDescData = file_proto_sprinkler_to_server_proto_rawDesc
)

func file_proto_sprinkler_to_server_proto_rawDescGZIP() []byte {
	file_proto_sprinkler_to_server_proto_rawDescOnce.Do(func() {
		file_proto_sprinkler_to_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sprinkler_to_server_proto_rawDescData)
	})
	return file_proto_sprinkler_to_server_proto_rawDescData
}

var file_proto_sprinkler_to_server_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_sprinkler_to_server_proto_goTypes = []interface{}{
	(*SprinklerStatusRequest)(nil),         // 0: SprinklerStatusRequest
	(*SprinklerStatusResponse)(nil),        // 1: SprinklerStatusResponse
	(*GetProgramRequest)(nil),              // 2: GetProgramRequest
	(*GetProgramResponse)(nil),             // 3: GetProgramResponse
	(*GetScheduledProgramsRequest)(nil),    // 4: GetScheduledProgramsRequest
	(*GetScheduledProgramsResponse)(nil),   // 5: GetScheduledProgramsResponse
	(*RecieveCommandsRequest)(nil),         // 6: RecieveCommandsRequest
	(*SprinklerStatus)(nil),                // 7: SprinklerStatus
	(*Command)(nil),                        // 8: Command
	(*StartCommand)(nil),                   // 9: StartCommand
	(*StopCommand)(nil),                    // 10: StopCommand
	(*UpdateScheduledProgramsCommand)(nil), // 11: UpdateScheduledProgramsCommand
	(*Program)(nil),                        // 12: Program
	(*ScheduledProgram)(nil),               // 13: ScheduledProgram
	(*ZoneTime)(nil),                       // 14: ZoneTime
}
var file_proto_sprinkler_to_server_proto_depIdxs = []int32{
	7,  // 0: SprinklerStatusRequest.sprinkler_status:type_name -> SprinklerStatus
	12, // 1: GetProgramRequest.program:type_name -> Program
	12, // 2: GetProgramResponse.program:type_name -> Program
	13, // 3: GetScheduledProgramsResponse.scheduled_programs:type_name -> ScheduledProgram
	9,  // 4: Command.start:type_name -> StartCommand
	10, // 5: Command.stop:type_name -> StopCommand
	11, // 6: Command.update_scheduled_programs:type_name -> UpdateScheduledProgramsCommand
	12, // 7: StartCommand.program:type_name -> Program
	14, // 8: StartCommand.zone_time:type_name -> ZoneTime
	0,  // 9: SprinklerToServer.SendStatus:input_type -> SprinklerStatusRequest
	2,  // 10: SprinklerToServer.GetProgram:input_type -> GetProgramRequest
	4,  // 11: SprinklerToServer.GetScheduledPrograms:input_type -> GetScheduledProgramsRequest
	6,  // 12: SprinklerToServer.RecieveCommands:input_type -> RecieveCommandsRequest
	1,  // 13: SprinklerToServer.SendStatus:output_type -> SprinklerStatusResponse
	3,  // 14: SprinklerToServer.GetProgram:output_type -> GetProgramResponse
	5,  // 15: SprinklerToServer.GetScheduledPrograms:output_type -> GetScheduledProgramsResponse
	8,  // 16: SprinklerToServer.RecieveCommands:output_type -> Command
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_sprinkler_to_server_proto_init() }
func file_proto_sprinkler_to_server_proto_init() {
	if File_proto_sprinkler_to_server_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_sprinkler_to_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprinklerStatusRequest); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprinklerStatusResponse); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProgramRequest); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProgramResponse); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScheduledProgramsRequest); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScheduledProgramsResponse); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecieveCommandsRequest); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprinklerStatus); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartCommand); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopCommand); i {
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
		file_proto_sprinkler_to_server_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScheduledProgramsCommand); i {
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
	file_proto_sprinkler_to_server_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Command_Start)(nil),
		(*Command_Stop)(nil),
		(*Command_UpdateScheduledPrograms)(nil),
	}
	file_proto_sprinkler_to_server_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*StartCommand_Program)(nil),
		(*StartCommand_ZoneTime)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_sprinkler_to_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sprinkler_to_server_proto_goTypes,
		DependencyIndexes: file_proto_sprinkler_to_server_proto_depIdxs,
		MessageInfos:      file_proto_sprinkler_to_server_proto_msgTypes,
	}.Build()
	File_proto_sprinkler_to_server_proto = out.File
	file_proto_sprinkler_to_server_proto_rawDesc = nil
	file_proto_sprinkler_to_server_proto_goTypes = nil
	file_proto_sprinkler_to_server_proto_depIdxs = nil
}