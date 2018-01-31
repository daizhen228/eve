// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devconfig.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MapServer struct {
	NameOrIp   string `protobuf:"bytes,1,opt,name=NameOrIp" json:"NameOrIp,omitempty"`
	Credential string `protobuf:"bytes,2,opt,name=Credential" json:"Credential,omitempty"`
}

func (m *MapServer) Reset()                    { *m = MapServer{} }
func (m *MapServer) String() string            { return proto.CompactTextString(m) }
func (*MapServer) ProtoMessage()               {}
func (*MapServer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *MapServer) GetNameOrIp() string {
	if m != nil {
		return m.NameOrIp
	}
	return ""
}

func (m *MapServer) GetCredential() string {
	if m != nil {
		return m.Credential
	}
	return ""
}

type ZedServer struct {
	HostName string   `protobuf:"bytes,1,opt,name=HostName" json:"HostName,omitempty"`
	EID      []string `protobuf:"bytes,2,rep,name=EID" json:"EID,omitempty"`
}

func (m *ZedServer) Reset()                    { *m = ZedServer{} }
func (m *ZedServer) String() string            { return proto.CompactTextString(m) }
func (*ZedServer) ProtoMessage()               {}
func (*ZedServer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ZedServer) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *ZedServer) GetEID() []string {
	if m != nil {
		return m.EID
	}
	return nil
}

type DeviceLispDetails struct {
	LispMapServers         []*MapServer `protobuf:"bytes,1,rep,name=LispMapServers" json:"LispMapServers,omitempty"`
	LispInstance           uint32       `protobuf:"varint,2,opt,name=LispInstance" json:"LispInstance,omitempty"`
	EID                    string       `protobuf:"bytes,4,opt,name=EID" json:"EID,omitempty"`
	EIDHashLen             uint32       `protobuf:"varint,5,opt,name=EIDHashLen" json:"EIDHashLen,omitempty"`
	ZedServers             []*ZedServer `protobuf:"bytes,6,rep,name=ZedServers" json:"ZedServers,omitempty"`
	EidAllocationPrefix    []byte       `protobuf:"bytes,8,opt,name=EidAllocationPrefix,proto3" json:"EidAllocationPrefix,omitempty"`
	EidAllocationPrefixLen uint32       `protobuf:"varint,9,opt,name=EidAllocationPrefixLen" json:"EidAllocationPrefixLen,omitempty"`
	ClientAddr             string       `protobuf:"bytes,10,opt,name=ClientAddr" json:"ClientAddr,omitempty"`
}

func (m *DeviceLispDetails) Reset()                    { *m = DeviceLispDetails{} }
func (m *DeviceLispDetails) String() string            { return proto.CompactTextString(m) }
func (*DeviceLispDetails) ProtoMessage()               {}
func (*DeviceLispDetails) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *DeviceLispDetails) GetLispMapServers() []*MapServer {
	if m != nil {
		return m.LispMapServers
	}
	return nil
}

func (m *DeviceLispDetails) GetLispInstance() uint32 {
	if m != nil {
		return m.LispInstance
	}
	return 0
}

func (m *DeviceLispDetails) GetEID() string {
	if m != nil {
		return m.EID
	}
	return ""
}

func (m *DeviceLispDetails) GetEIDHashLen() uint32 {
	if m != nil {
		return m.EIDHashLen
	}
	return 0
}

func (m *DeviceLispDetails) GetZedServers() []*ZedServer {
	if m != nil {
		return m.ZedServers
	}
	return nil
}

func (m *DeviceLispDetails) GetEidAllocationPrefix() []byte {
	if m != nil {
		return m.EidAllocationPrefix
	}
	return nil
}

func (m *DeviceLispDetails) GetEidAllocationPrefixLen() uint32 {
	if m != nil {
		return m.EidAllocationPrefixLen
	}
	return 0
}

func (m *DeviceLispDetails) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

// Device Operational Commands Semantic
// For rebooting device,  command=Reset, counter = counter+delta, desiredState = on
// For poweroff device,  command=Reset, counter = counter+delta, desiredState = off
// For backup at midnight, command=Backup, counter = counter+delta, desiredState=n/a, opsTime = mm/dd/yy:hh:ss
// Current implementation does support only single command outstanding for each type
// In future can be extended to have more scheduled events
//
type DeviceOpsCmd struct {
	Counter      uint32 `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
	DesiredState bool   `protobuf:"varint,3,opt,name=desiredState" json:"desiredState,omitempty"`
	// FIXME: change to timestamp, once we move to gogo proto
	OpsTime string `protobuf:"bytes,4,opt,name=opsTime" json:"opsTime,omitempty"`
}

func (m *DeviceOpsCmd) Reset()                    { *m = DeviceOpsCmd{} }
func (m *DeviceOpsCmd) String() string            { return proto.CompactTextString(m) }
func (*DeviceOpsCmd) ProtoMessage()               {}
func (*DeviceOpsCmd) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *DeviceOpsCmd) GetCounter() uint32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *DeviceOpsCmd) GetDesiredState() bool {
	if m != nil {
		return m.DesiredState
	}
	return false
}

func (m *DeviceOpsCmd) GetOpsTime() string {
	if m != nil {
		return m.OpsTime
	}
	return ""
}

type EdgeDevConfig struct {
	Id                 *UUIDandVersion      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DevConfigSha256    []byte               `protobuf:"bytes,2,opt,name=devConfigSha256,proto3" json:"devConfigSha256,omitempty"`
	DevConfigSignature []byte               `protobuf:"bytes,3,opt,name=devConfigSignature,proto3" json:"devConfigSignature,omitempty"`
	Apps               []*AppInstanceConfig `protobuf:"bytes,4,rep,name=apps" json:"apps,omitempty"`
	Networks           []*NetworkConfig     `protobuf:"bytes,5,rep,name=networks" json:"networks,omitempty"`
	Datastores         []*DatastoreConfig   `protobuf:"bytes,6,rep,name=datastores" json:"datastores,omitempty"`
	LispInfo           *DeviceLispDetails   `protobuf:"bytes,7,opt,name=lispInfo" json:"lispInfo,omitempty"`
	Base               []*BaseOSConfig      `protobuf:"bytes,8,rep,name=base" json:"base,omitempty"`
	Reboot             *DeviceOpsCmd        `protobuf:"bytes,9,opt,name=reboot" json:"reboot,omitempty"`
	Backup             *DeviceOpsCmd        `protobuf:"bytes,10,opt,name=backup" json:"backup,omitempty"`
}

func (m *EdgeDevConfig) Reset()                    { *m = EdgeDevConfig{} }
func (m *EdgeDevConfig) String() string            { return proto.CompactTextString(m) }
func (*EdgeDevConfig) ProtoMessage()               {}
func (*EdgeDevConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *EdgeDevConfig) GetId() *UUIDandVersion {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EdgeDevConfig) GetDevConfigSha256() []byte {
	if m != nil {
		return m.DevConfigSha256
	}
	return nil
}

func (m *EdgeDevConfig) GetDevConfigSignature() []byte {
	if m != nil {
		return m.DevConfigSignature
	}
	return nil
}

func (m *EdgeDevConfig) GetApps() []*AppInstanceConfig {
	if m != nil {
		return m.Apps
	}
	return nil
}

func (m *EdgeDevConfig) GetNetworks() []*NetworkConfig {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *EdgeDevConfig) GetDatastores() []*DatastoreConfig {
	if m != nil {
		return m.Datastores
	}
	return nil
}

func (m *EdgeDevConfig) GetLispInfo() *DeviceLispDetails {
	if m != nil {
		return m.LispInfo
	}
	return nil
}

func (m *EdgeDevConfig) GetBase() []*BaseOSConfig {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *EdgeDevConfig) GetReboot() *DeviceOpsCmd {
	if m != nil {
		return m.Reboot
	}
	return nil
}

func (m *EdgeDevConfig) GetBackup() *DeviceOpsCmd {
	if m != nil {
		return m.Backup
	}
	return nil
}

type ConfigRequest struct {
	ConfigHash string `protobuf:"bytes,1,opt,name=configHash" json:"configHash,omitempty"`
}

func (m *ConfigRequest) Reset()                    { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()               {}
func (*ConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *ConfigRequest) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

type ConfigResponse struct {
	Config     *EdgeDevConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	ConfigHash string         `protobuf:"bytes,2,opt,name=configHash" json:"configHash,omitempty"`
}

func (m *ConfigResponse) Reset()                    { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()               {}
func (*ConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *ConfigResponse) GetConfig() *EdgeDevConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ConfigResponse) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

func init() {
	proto.RegisterType((*MapServer)(nil), "MapServer")
	proto.RegisterType((*ZedServer)(nil), "ZedServer")
	proto.RegisterType((*DeviceLispDetails)(nil), "DeviceLispDetails")
	proto.RegisterType((*DeviceOpsCmd)(nil), "DeviceOpsCmd")
	proto.RegisterType((*EdgeDevConfig)(nil), "EdgeDevConfig")
	proto.RegisterType((*ConfigRequest)(nil), "ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "ConfigResponse")
}

func init() { proto.RegisterFile("devconfig.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x4f, 0x1b, 0x39,
	0x14, 0x55, 0x3e, 0x08, 0xc9, 0x25, 0x09, 0xac, 0x57, 0x5a, 0x8d, 0x90, 0x76, 0xc9, 0x8e, 0xd4,
	0x2a, 0xe2, 0xc1, 0x41, 0xa9, 0x8a, 0xd4, 0xb7, 0x02, 0x89, 0x4a, 0x24, 0x0a, 0xd5, 0xa4, 0x54,
	0x15, 0x6f, 0xce, 0xf8, 0x26, 0x58, 0x24, 0xf6, 0xd4, 0x76, 0xd2, 0x8a, 0x7f, 0xd5, 0xdf, 0xd2,
	0x3f, 0x54, 0xd9, 0xf3, 0x41, 0x12, 0xe8, 0xdb, 0xdc, 0x73, 0x8e, 0xef, 0x9c, 0xb9, 0xf7, 0x78,
	0x60, 0x9f, 0xe3, 0x2a, 0x56, 0x72, 0x2a, 0x66, 0x34, 0xd1, 0xca, 0xaa, 0xc3, 0x14, 0x58, 0x2c,
	0x94, 0xcc, 0x01, 0x96, 0x24, 0x1b, 0x0a, 0x32, 0x61, 0x06, 0x95, 0xd9, 0x3c, 0x25, 0xd1, 0x6e,
	0x00, 0x2d, 0x63, 0x95, 0x66, 0x33, 0x4c, 0xcb, 0xf0, 0x03, 0x34, 0x3e, 0xb2, 0x64, 0x8c, 0x7a,
	0x85, 0x9a, 0x1c, 0x42, 0xfd, 0x9a, 0x2d, 0xf0, 0x46, 0x8f, 0x92, 0xa0, 0xd4, 0x29, 0x75, 0x1b,
	0x51, 0x51, 0x93, 0xff, 0x00, 0x2e, 0x34, 0x72, 0x94, 0x56, 0xb0, 0x79, 0x50, 0xf6, 0xec, 0x1a,
	0x12, 0xbe, 0x83, 0xc6, 0x1d, 0xf2, 0xa7, 0x46, 0x97, 0xca, 0x58, 0x77, 0x38, 0x6f, 0x94, 0xd7,
	0xe4, 0x00, 0x2a, 0xc3, 0xd1, 0x20, 0x28, 0x77, 0x2a, 0xdd, 0x46, 0xe4, 0x1e, 0xc3, 0x5f, 0x65,
	0xf8, 0x6b, 0x80, 0x2b, 0x11, 0xe3, 0x95, 0x30, 0xc9, 0x00, 0x2d, 0x13, 0x73, 0x43, 0xfa, 0xd0,
	0x76, 0x65, 0xe1, 0xce, 0x04, 0xa5, 0x4e, 0xa5, 0xbb, 0xd7, 0x07, 0x5a, 0x40, 0xd1, 0x96, 0x82,
	0x84, 0xd0, 0x74, 0xc8, 0x48, 0x1a, 0xcb, 0x64, 0x8c, 0xde, 0x66, 0x2b, 0xda, 0xc0, 0xf2, 0xf7,
	0x57, 0xbd, 0x2d, 0xf7, 0xe8, 0x3e, 0x6d, 0x38, 0x1a, 0x5c, 0x32, 0x73, 0x7f, 0x85, 0x32, 0xd8,
	0xf1, 0x67, 0xd6, 0x10, 0x72, 0x0c, 0x50, 0x7c, 0x9a, 0x09, 0x6a, 0x99, 0x8b, 0x02, 0x8a, 0xd6,
	0x58, 0x72, 0x02, 0x7f, 0x0f, 0x05, 0x3f, 0x9b, 0xcf, 0x55, 0xcc, 0xac, 0x50, 0xf2, 0x93, 0xc6,
	0xa9, 0xf8, 0x11, 0xd4, 0x3b, 0xa5, 0x6e, 0x33, 0x7a, 0x89, 0x22, 0xa7, 0xf0, 0xcf, 0x0b, 0xb0,
	0x73, 0xd2, 0xf0, 0x4e, 0xfe, 0xc0, 0xfa, 0x85, 0xcc, 0x05, 0x4a, 0x7b, 0xc6, 0xb9, 0x0e, 0x20,
	0x5b, 0x48, 0x81, 0x84, 0x53, 0x68, 0xa6, 0x43, 0xbd, 0x49, 0xcc, 0xc5, 0x82, 0x93, 0x00, 0x76,
	0x63, 0xb5, 0x94, 0x16, 0x75, 0x36, 0x96, 0xbc, 0x74, 0x53, 0xe3, 0x68, 0x84, 0x46, 0x3e, 0xb6,
	0xcc, 0x62, 0x50, 0xe9, 0x94, 0xba, 0xf5, 0x68, 0x03, 0x73, 0xa7, 0x55, 0x62, 0x3e, 0x8b, 0x05,
	0x66, 0x93, 0xcb, 0xcb, 0xf0, 0x67, 0x05, 0x5a, 0x43, 0x3e, 0xc3, 0x01, 0xae, 0x2e, 0x7c, 0xd0,
	0xc8, 0x11, 0x94, 0x05, 0xf7, 0x7b, 0xdf, 0xeb, 0xef, 0xd3, 0xdb, 0xdb, 0xd1, 0x80, 0x49, 0xfe,
	0x05, 0xb5, 0x11, 0x4a, 0x46, 0x65, 0xc1, 0x49, 0xd7, 0xa7, 0x3b, 0x55, 0x8f, 0xef, 0x59, 0xff,
	0xed, 0xa9, 0xb7, 0xd4, 0x8c, 0xb6, 0x61, 0x42, 0x81, 0x3c, 0x41, 0x62, 0x26, 0x99, 0x5d, 0xea,
	0xd4, 0x60, 0x33, 0x7a, 0x81, 0x21, 0xaf, 0xa1, 0xca, 0x92, 0xc4, 0x04, 0x55, 0xbf, 0x24, 0x42,
	0xcf, 0x92, 0x62, 0xf1, 0xa9, 0x34, 0xf2, 0x3c, 0x39, 0x86, 0xba, 0x44, 0xfb, 0x5d, 0xe9, 0x07,
	0x13, 0xec, 0x78, 0x6d, 0x9b, 0x5e, 0xa7, 0x40, 0xa6, 0x2b, 0x78, 0x72, 0x02, 0xc0, 0x99, 0x65,
	0xee, 0xde, 0x60, 0xbe, 0xfe, 0x03, 0x3a, 0xc8, 0xa1, 0x4c, 0xbf, 0xa6, 0x21, 0x14, 0xea, 0x73,
	0x1f, 0xb9, 0xa9, 0x0a, 0x76, 0xfd, 0x18, 0x08, 0x7d, 0x16, 0xf0, 0xa8, 0xd0, 0x90, 0xff, 0xa1,
	0xea, 0xae, 0x6e, 0x50, 0xf7, 0xbd, 0x5b, 0xf4, 0x9c, 0x19, 0xbc, 0x19, 0xe7, 0x86, 0x1d, 0x45,
	0x5e, 0x41, 0x4d, 0xe3, 0x44, 0x29, 0xeb, 0x53, 0xe1, 0x44, 0xeb, 0xcb, 0x8d, 0x32, 0xd2, 0xc9,
	0x26, 0x2c, 0x7e, 0x58, 0x26, 0x3e, 0x10, 0xcf, 0x65, 0x29, 0x19, 0xf6, 0xa0, 0x95, 0x75, 0xc7,
	0x6f, 0x4b, 0x34, 0xd6, 0x85, 0x29, 0xfd, 0x4b, 0xb8, 0xcc, 0x67, 0x57, 0x76, 0x0d, 0x09, 0xbf,
	0x42, 0x3b, 0x3f, 0x60, 0x12, 0x25, 0x8d, 0x9b, 0x74, 0x2d, 0xe5, 0xb3, 0x45, 0xb7, 0xe9, 0x46,
	0x08, 0xa2, 0x8c, 0xdd, 0xea, 0x5c, 0xde, 0xee, 0x7c, 0xfe, 0x1e, 0x8e, 0x62, 0xb5, 0xa0, 0x8f,
	0xc8, 0x91, 0x33, 0x1a, 0xcf, 0xd5, 0x92, 0xd3, 0xa5, 0x41, 0xed, 0x5c, 0xa7, 0xff, 0xa8, 0xbb,
	0x7f, 0x67, 0xc2, 0xde, 0x2f, 0x27, 0x34, 0x56, 0x8b, 0x5e, 0xaa, 0xeb, 0xb1, 0x44, 0xf4, 0x1e,
	0xd3, 0x2e, 0x93, 0x9a, 0x57, 0xbd, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xea, 0x23, 0x81,
	0x32, 0x05, 0x00, 0x00,
}
