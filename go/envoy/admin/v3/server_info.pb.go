// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3/server_info.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ServerInfo_State int32

const (
	// Server is live and serving traffic.
	ServerInfo_LIVE ServerInfo_State = 0
	// Server is draining listeners in response to external health checks failing.
	ServerInfo_DRAINING ServerInfo_State = 1
	// Server has not yet completed cluster manager initialization.
	ServerInfo_PRE_INITIALIZING ServerInfo_State = 2
	// Server is running the cluster manager initialization callbacks (e.g., RDS).
	ServerInfo_INITIALIZING ServerInfo_State = 3
)

var ServerInfo_State_name = map[int32]string{
	0: "LIVE",
	1: "DRAINING",
	2: "PRE_INITIALIZING",
	3: "INITIALIZING",
}

var ServerInfo_State_value = map[string]int32{
	"LIVE":             0,
	"DRAINING":         1,
	"PRE_INITIALIZING": 2,
	"INITIALIZING":     3,
}

func (x ServerInfo_State) String() string {
	return proto.EnumName(ServerInfo_State_name, int32(x))
}

func (ServerInfo_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{0, 0}
}

type CommandLineOptions_IpVersion int32

const (
	CommandLineOptions_v4 CommandLineOptions_IpVersion = 0
	CommandLineOptions_v6 CommandLineOptions_IpVersion = 1
)

var CommandLineOptions_IpVersion_name = map[int32]string{
	0: "v4",
	1: "v6",
}

var CommandLineOptions_IpVersion_value = map[string]int32{
	"v4": 0,
	"v6": 1,
}

func (x CommandLineOptions_IpVersion) String() string {
	return proto.EnumName(CommandLineOptions_IpVersion_name, int32(x))
}

func (CommandLineOptions_IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{1, 0}
}

type CommandLineOptions_Mode int32

const (
	// Validate configs and then serve traffic normally.
	CommandLineOptions_Serve CommandLineOptions_Mode = 0
	// Validate configs and exit.
	CommandLineOptions_Validate CommandLineOptions_Mode = 1
	// Completely load and initialize the config, and then exit without running the listener loop.
	CommandLineOptions_InitOnly CommandLineOptions_Mode = 2
)

var CommandLineOptions_Mode_name = map[int32]string{
	0: "Serve",
	1: "Validate",
	2: "InitOnly",
}

var CommandLineOptions_Mode_value = map[string]int32{
	"Serve":    0,
	"Validate": 1,
	"InitOnly": 2,
}

func (x CommandLineOptions_Mode) String() string {
	return proto.EnumName(CommandLineOptions_Mode_name, int32(x))
}

func (CommandLineOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{1, 1}
}

// Proto representation of the value returned by /server_info, containing
// server version/server status information.
// [#next-free-field: 7]
type ServerInfo struct {
	// Server version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// State of the server.
	State ServerInfo_State `protobuf:"varint,2,opt,name=state,proto3,enum=envoy.admin.v3.ServerInfo_State" json:"state,omitempty"`
	// Uptime since current epoch was started.
	UptimeCurrentEpoch *duration.Duration `protobuf:"bytes,3,opt,name=uptime_current_epoch,json=uptimeCurrentEpoch,proto3" json:"uptime_current_epoch,omitempty"`
	// Uptime since the start of the first epoch.
	UptimeAllEpochs *duration.Duration `protobuf:"bytes,4,opt,name=uptime_all_epochs,json=uptimeAllEpochs,proto3" json:"uptime_all_epochs,omitempty"`
	// Hot restart version.
	HotRestartVersion string `protobuf:"bytes,5,opt,name=hot_restart_version,json=hotRestartVersion,proto3" json:"hot_restart_version,omitempty"`
	// Command line options the server is currently running with.
	CommandLineOptions   *CommandLineOptions `protobuf:"bytes,6,opt,name=command_line_options,json=commandLineOptions,proto3" json:"command_line_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetState() ServerInfo_State {
	if m != nil {
		return m.State
	}
	return ServerInfo_LIVE
}

func (m *ServerInfo) GetUptimeCurrentEpoch() *duration.Duration {
	if m != nil {
		return m.UptimeCurrentEpoch
	}
	return nil
}

func (m *ServerInfo) GetUptimeAllEpochs() *duration.Duration {
	if m != nil {
		return m.UptimeAllEpochs
	}
	return nil
}

func (m *ServerInfo) GetHotRestartVersion() string {
	if m != nil {
		return m.HotRestartVersion
	}
	return ""
}

func (m *ServerInfo) GetCommandLineOptions() *CommandLineOptions {
	if m != nil {
		return m.CommandLineOptions
	}
	return nil
}

// [#next-free-field: 29]
type CommandLineOptions struct {
	// See :option:`--base-id` for details.
	BaseId uint64 `protobuf:"varint,1,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	// See :option:`--concurrency` for details.
	Concurrency uint32 `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// See :option:`--config-path` for details.
	ConfigPath string `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	// See :option:`--config-yaml` for details.
	ConfigYaml string `protobuf:"bytes,4,opt,name=config_yaml,json=configYaml,proto3" json:"config_yaml,omitempty"`
	// See :option:`--allow-unknown-static-fields` for details.
	AllowUnknownStaticFields bool `protobuf:"varint,5,opt,name=allow_unknown_static_fields,json=allowUnknownStaticFields,proto3" json:"allow_unknown_static_fields,omitempty"`
	// See :option:`--reject-unknown-dynamic-fields` for details.
	RejectUnknownDynamicFields bool `protobuf:"varint,26,opt,name=reject_unknown_dynamic_fields,json=rejectUnknownDynamicFields,proto3" json:"reject_unknown_dynamic_fields,omitempty"`
	// See :option:`--admin-address-path` for details.
	AdminAddressPath string `protobuf:"bytes,6,opt,name=admin_address_path,json=adminAddressPath,proto3" json:"admin_address_path,omitempty"`
	// See :option:`--local-address-ip-version` for details.
	LocalAddressIpVersion CommandLineOptions_IpVersion `protobuf:"varint,7,opt,name=local_address_ip_version,json=localAddressIpVersion,proto3,enum=envoy.admin.v3.CommandLineOptions_IpVersion" json:"local_address_ip_version,omitempty"`
	// See :option:`--log-level` for details.
	LogLevel string `protobuf:"bytes,8,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// See :option:`--component-log-level` for details.
	ComponentLogLevel string `protobuf:"bytes,9,opt,name=component_log_level,json=componentLogLevel,proto3" json:"component_log_level,omitempty"`
	// See :option:`--log-format` for details.
	LogFormat string `protobuf:"bytes,10,opt,name=log_format,json=logFormat,proto3" json:"log_format,omitempty"`
	// See :option:`--log-format-escaped` for details.
	LogFormatEscaped bool `protobuf:"varint,27,opt,name=log_format_escaped,json=logFormatEscaped,proto3" json:"log_format_escaped,omitempty"`
	// See :option:`--log-path` for details.
	LogPath string `protobuf:"bytes,11,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	// See :option:`--service-cluster` for details.
	ServiceCluster string `protobuf:"bytes,13,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	// See :option:`--service-node` for details.
	ServiceNode string `protobuf:"bytes,14,opt,name=service_node,json=serviceNode,proto3" json:"service_node,omitempty"`
	// See :option:`--service-zone` for details.
	ServiceZone string `protobuf:"bytes,15,opt,name=service_zone,json=serviceZone,proto3" json:"service_zone,omitempty"`
	// See :option:`--file-flush-interval-msec` for details.
	FileFlushInterval *duration.Duration `protobuf:"bytes,16,opt,name=file_flush_interval,json=fileFlushInterval,proto3" json:"file_flush_interval,omitempty"`
	// See :option:`--drain-time-s` for details.
	DrainTime *duration.Duration `protobuf:"bytes,17,opt,name=drain_time,json=drainTime,proto3" json:"drain_time,omitempty"`
	// See :option:`--parent-shutdown-time-s` for details.
	ParentShutdownTime *duration.Duration `protobuf:"bytes,18,opt,name=parent_shutdown_time,json=parentShutdownTime,proto3" json:"parent_shutdown_time,omitempty"`
	// See :option:`--mode` for details.
	Mode CommandLineOptions_Mode `protobuf:"varint,19,opt,name=mode,proto3,enum=envoy.admin.v3.CommandLineOptions_Mode" json:"mode,omitempty"`
	// See :option:`--disable-hot-restart` for details.
	DisableHotRestart bool `protobuf:"varint,22,opt,name=disable_hot_restart,json=disableHotRestart,proto3" json:"disable_hot_restart,omitempty"`
	// See :option:`--enable-mutex-tracing` for details.
	EnableMutexTracing bool `protobuf:"varint,23,opt,name=enable_mutex_tracing,json=enableMutexTracing,proto3" json:"enable_mutex_tracing,omitempty"`
	// See :option:`--restart-epoch` for details.
	RestartEpoch uint32 `protobuf:"varint,24,opt,name=restart_epoch,json=restartEpoch,proto3" json:"restart_epoch,omitempty"`
	// See :option:`--cpuset-threads` for details.
	CpusetThreads bool `protobuf:"varint,25,opt,name=cpuset_threads,json=cpusetThreads,proto3" json:"cpuset_threads,omitempty"`
	// See :option:`--disable-extensions` for details.
	DisabledExtensions   []string `protobuf:"bytes,28,rep,name=disabled_extensions,json=disabledExtensions,proto3" json:"disabled_extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandLineOptions) Reset()         { *m = CommandLineOptions{} }
func (m *CommandLineOptions) String() string { return proto.CompactTextString(m) }
func (*CommandLineOptions) ProtoMessage()    {}
func (*CommandLineOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{1}
}

func (m *CommandLineOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLineOptions.Unmarshal(m, b)
}
func (m *CommandLineOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLineOptions.Marshal(b, m, deterministic)
}
func (m *CommandLineOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLineOptions.Merge(m, src)
}
func (m *CommandLineOptions) XXX_Size() int {
	return xxx_messageInfo_CommandLineOptions.Size(m)
}
func (m *CommandLineOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLineOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLineOptions proto.InternalMessageInfo

func (m *CommandLineOptions) GetBaseId() uint64 {
	if m != nil {
		return m.BaseId
	}
	return 0
}

func (m *CommandLineOptions) GetConcurrency() uint32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *CommandLineOptions) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *CommandLineOptions) GetConfigYaml() string {
	if m != nil {
		return m.ConfigYaml
	}
	return ""
}

func (m *CommandLineOptions) GetAllowUnknownStaticFields() bool {
	if m != nil {
		return m.AllowUnknownStaticFields
	}
	return false
}

func (m *CommandLineOptions) GetRejectUnknownDynamicFields() bool {
	if m != nil {
		return m.RejectUnknownDynamicFields
	}
	return false
}

func (m *CommandLineOptions) GetAdminAddressPath() string {
	if m != nil {
		return m.AdminAddressPath
	}
	return ""
}

func (m *CommandLineOptions) GetLocalAddressIpVersion() CommandLineOptions_IpVersion {
	if m != nil {
		return m.LocalAddressIpVersion
	}
	return CommandLineOptions_v4
}

func (m *CommandLineOptions) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetComponentLogLevel() string {
	if m != nil {
		return m.ComponentLogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormat() string {
	if m != nil {
		return m.LogFormat
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormatEscaped() bool {
	if m != nil {
		return m.LogFormatEscaped
	}
	return false
}

func (m *CommandLineOptions) GetLogPath() string {
	if m != nil {
		return m.LogPath
	}
	return ""
}

func (m *CommandLineOptions) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *CommandLineOptions) GetServiceNode() string {
	if m != nil {
		return m.ServiceNode
	}
	return ""
}

func (m *CommandLineOptions) GetServiceZone() string {
	if m != nil {
		return m.ServiceZone
	}
	return ""
}

func (m *CommandLineOptions) GetFileFlushInterval() *duration.Duration {
	if m != nil {
		return m.FileFlushInterval
	}
	return nil
}

func (m *CommandLineOptions) GetDrainTime() *duration.Duration {
	if m != nil {
		return m.DrainTime
	}
	return nil
}

func (m *CommandLineOptions) GetParentShutdownTime() *duration.Duration {
	if m != nil {
		return m.ParentShutdownTime
	}
	return nil
}

func (m *CommandLineOptions) GetMode() CommandLineOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return CommandLineOptions_Serve
}

func (m *CommandLineOptions) GetDisableHotRestart() bool {
	if m != nil {
		return m.DisableHotRestart
	}
	return false
}

func (m *CommandLineOptions) GetEnableMutexTracing() bool {
	if m != nil {
		return m.EnableMutexTracing
	}
	return false
}

func (m *CommandLineOptions) GetRestartEpoch() uint32 {
	if m != nil {
		return m.RestartEpoch
	}
	return 0
}

func (m *CommandLineOptions) GetCpusetThreads() bool {
	if m != nil {
		return m.CpusetThreads
	}
	return false
}

func (m *CommandLineOptions) GetDisabledExtensions() []string {
	if m != nil {
		return m.DisabledExtensions
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.admin.v3.ServerInfo_State", ServerInfo_State_name, ServerInfo_State_value)
	proto.RegisterEnum("envoy.admin.v3.CommandLineOptions_IpVersion", CommandLineOptions_IpVersion_name, CommandLineOptions_IpVersion_value)
	proto.RegisterEnum("envoy.admin.v3.CommandLineOptions_Mode", CommandLineOptions_Mode_name, CommandLineOptions_Mode_value)
	proto.RegisterType((*ServerInfo)(nil), "envoy.admin.v3.ServerInfo")
	proto.RegisterType((*CommandLineOptions)(nil), "envoy.admin.v3.CommandLineOptions")
}

func init() { proto.RegisterFile("envoy/admin/v3/server_info.proto", fileDescriptor_8736ae14312a45ee) }

var fileDescriptor_8736ae14312a45ee = []byte{
	// 1037 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x52, 0xdb, 0x46,
	0x14, 0x8e, 0xc1, 0x10, 0xfb, 0xf0, 0x27, 0x36, 0x90, 0x08, 0x48, 0xa8, 0xe3, 0x4c, 0x1a, 0x2e,
	0x12, 0xbb, 0x85, 0x4e, 0xa6, 0x93, 0x4e, 0x2f, 0x08, 0x98, 0x54, 0x94, 0x10, 0x46, 0x50, 0x66,
	0x9a, 0x9b, 0x9d, 0x45, 0x5a, 0xdb, 0x9b, 0xae, 0x76, 0x35, 0xd2, 0xca, 0xc1, 0x7d, 0x82, 0x3e,
	0x43, 0xdf, 0xa7, 0x6f, 0xd2, 0x8b, 0x3e, 0x46, 0x67, 0xcf, 0xca, 0x3f, 0x84, 0xce, 0xd0, 0x2b,
	0x5b, 0xdf, 0xf9, 0xce, 0xa7, 0xf3, 0xa7, 0x0f, 0x1a, 0x5c, 0x0d, 0xf4, 0xb0, 0xcd, 0xe2, 0x44,
	0xa8, 0xf6, 0x60, 0xaf, 0x9d, 0xf3, 0x6c, 0xc0, 0x33, 0x2a, 0x54, 0x57, 0xb7, 0xd2, 0x4c, 0x1b,
	0x4d, 0x96, 0x91, 0xd1, 0x42, 0x46, 0x6b, 0xb0, 0xb7, 0xb9, 0xdd, 0xd3, 0xba, 0x27, 0x79, 0x1b,
	0xa3, 0x57, 0x45, 0xb7, 0x1d, 0x17, 0x19, 0x33, 0x42, 0x2b, 0xc7, 0xdf, 0x7c, 0x5a, 0xc4, 0x29,
	0x6b, 0x33, 0xa5, 0xb4, 0x41, 0x38, 0x6f, 0x0f, 0x78, 0x96, 0x0b, 0xad, 0x84, 0xea, 0x95, 0x94,
	0x67, 0xe5, 0x4b, 0xa7, 0x38, 0x31, 0x4f, 0x33, 0x1e, 0x4d, 0xe9, 0x34, 0xff, 0x99, 0x05, 0x38,
	0xc7, 0x6a, 0x02, 0xd5, 0xd5, 0xc4, 0x87, 0xfb, 0xa5, 0x8e, 0x5f, 0x69, 0x54, 0x76, 0xea, 0xe1,
	0xe8, 0x91, 0xbc, 0x86, 0xb9, 0xdc, 0x30, 0xc3, 0xfd, 0x99, 0x46, 0x65, 0x67, 0x79, 0xb7, 0xd1,
	0xba, 0x59, 0x70, 0x6b, 0x22, 0xd2, 0x3a, 0xb7, 0xbc, 0xd0, 0xd1, 0xc9, 0xcf, 0xb0, 0x56, 0xa4,
	0x46, 0x24, 0x9c, 0x46, 0x45, 0x96, 0x71, 0x65, 0x28, 0x4f, 0x75, 0xd4, 0xf7, 0x67, 0x1b, 0x95,
	0x9d, 0x85, 0xdd, 0x8d, 0x96, 0xeb, 0xb3, 0x35, 0xea, 0xb3, 0x75, 0x58, 0xf6, 0x19, 0x12, 0x97,
	0x76, 0xe0, 0xb2, 0x3a, 0x36, 0x89, 0x74, 0x60, 0xb5, 0x14, 0x63, 0x52, 0x3a, 0xa1, 0xdc, 0xaf,
	0xde, 0xa5, 0xb4, 0xe2, 0x72, 0xf6, 0xa5, 0x44, 0x95, 0x9c, 0xb4, 0xe0, 0x41, 0x5f, 0x1b, 0x9a,
	0xf1, 0xdc, 0xb0, 0xcc, 0xd0, 0x51, 0xc7, 0x73, 0xd8, 0xf1, 0x6a, 0x5f, 0x9b, 0xd0, 0x45, 0x2e,
	0xcb, 0xde, 0x2f, 0x60, 0x2d, 0xd2, 0x49, 0xc2, 0x54, 0x4c, 0xa5, 0x50, 0x9c, 0xea, 0x14, 0xc7,
	0xe9, 0xcf, 0xe3, 0x9b, 0x9b, 0x5f, 0x8e, 0xe2, 0xc0, 0x71, 0x4f, 0x84, 0xe2, 0x1f, 0x1c, 0x33,
	0x24, 0xd1, 0x2d, 0xac, 0xf9, 0x0e, 0xe6, 0x70, 0x52, 0xa4, 0x06, 0xd5, 0x93, 0xe0, 0xb2, 0xe3,
	0xdd, 0x23, 0x8b, 0x50, 0x3b, 0x0c, 0xf7, 0x83, 0xd3, 0xe0, 0xf4, 0x9d, 0x57, 0x21, 0x6b, 0xe0,
	0x9d, 0x85, 0x1d, 0x1a, 0x9c, 0x06, 0x17, 0xc1, 0xfe, 0x49, 0xf0, 0xd1, 0xa2, 0x33, 0xc4, 0x83,
	0xc5, 0x1b, 0xc8, 0xec, 0x9b, 0xe7, 0x7f, 0xfe, 0xf5, 0xc7, 0x76, 0x03, 0xb6, 0x6f, 0x94, 0xb1,
	0xcb, 0x64, 0xda, 0x67, 0x53, 0x6b, 0x69, 0xfe, 0x0d, 0x40, 0x6e, 0x97, 0x46, 0x1e, 0xc1, 0xfd,
	0x2b, 0x96, 0x73, 0x2a, 0x62, 0x5c, 0x79, 0x35, 0x9c, 0xb7, 0x8f, 0x41, 0x4c, 0x1a, 0xb0, 0x10,
	0x69, 0xe5, 0xb6, 0x16, 0x0d, 0x71, 0xef, 0x4b, 0xe1, 0x34, 0x44, 0xbe, 0x42, 0x46, 0x57, 0xf4,
	0x68, 0xca, 0x8c, 0x5b, 0x69, 0x3d, 0x04, 0x07, 0x9d, 0x31, 0xd3, 0x9f, 0x22, 0x0c, 0x59, 0x22,
	0x71, 0x53, 0x63, 0xc2, 0xaf, 0x2c, 0x91, 0xe4, 0x47, 0xd8, 0x62, 0x52, 0xea, 0xcf, 0xb4, 0x50,
	0xbf, 0x29, 0xfd, 0x59, 0x51, 0x7b, 0x34, 0x22, 0xa2, 0x5d, 0xc1, 0x65, 0x9c, 0xe3, 0x46, 0x6a,
	0xa1, 0x8f, 0x94, 0x5f, 0x1c, 0xe3, 0x1c, 0x09, 0x47, 0x18, 0x27, 0xfb, 0xf0, 0x24, 0xe3, 0x9f,
	0x78, 0x64, 0xc6, 0xf9, 0xf1, 0x50, 0xb1, 0x64, 0x22, 0xb0, 0x89, 0x02, 0x9b, 0x8e, 0x54, 0x2a,
	0x1c, 0x3a, 0x4a, 0x29, 0xf1, 0x12, 0x08, 0x4e, 0x8c, 0xb2, 0x38, 0xce, 0x78, 0x9e, 0xbb, 0x56,
	0xe6, 0xb1, 0x52, 0x0f, 0x23, 0xfb, 0x2e, 0x80, 0x0d, 0x71, 0xf0, 0xa5, 0x8e, 0x98, 0x1c, 0xb3,
	0x45, 0x3a, 0x3e, 0x9f, 0xfb, 0xf8, 0x61, 0xbc, 0xbc, 0xfb, 0x1a, 0x5a, 0x41, 0x5a, 0x5e, 0x56,
	0xb8, 0x8e, 0x6a, 0xe5, 0x1b, 0xc6, 0x30, 0xd9, 0x82, 0xba, 0xd4, 0x3d, 0x2a, 0xf9, 0x80, 0x4b,
	0xbf, 0x86, 0xb5, 0xd4, 0xa4, 0xee, 0x9d, 0xd8, 0x67, 0x7b, 0xbd, 0x91, 0x4e, 0x52, 0xad, 0xec,
	0xc7, 0x34, 0xa1, 0xd5, 0xdd, 0xf5, 0x8e, 0x43, 0x27, 0x23, 0xfe, 0x13, 0x00, 0xcb, 0xea, 0xea,
	0x2c, 0x61, 0xc6, 0x07, 0xa4, 0x59, 0xf9, 0x23, 0x04, 0xec, 0x00, 0x26, 0x61, 0xca, 0xf3, 0x88,
	0xa5, 0x3c, 0xf6, 0xb7, 0x70, 0x70, 0xde, 0x98, 0xd6, 0x71, 0x38, 0xd9, 0x00, 0x5b, 0x88, 0x1b,
	0xd2, 0x82, 0x73, 0x08, 0xa9, 0xdd, 0xb2, 0x5f, 0xc0, 0x8a, 0xf5, 0x35, 0x11, 0x71, 0x1a, 0xc9,
	0x22, 0x37, 0x3c, 0xf3, 0x97, 0x90, 0xb1, 0x5c, 0xc2, 0x07, 0x0e, 0x25, 0x4f, 0x61, 0x71, 0x44,
	0x54, 0x3a, 0xe6, 0xfe, 0x32, 0xb2, 0x16, 0x4a, 0xec, 0x54, 0xc7, 0x7c, 0x9a, 0xf2, 0xbb, 0x56,
	0xdc, 0x5f, 0xb9, 0x41, 0xf9, 0xa8, 0x15, 0x27, 0x01, 0x3c, 0xe8, 0x0a, 0xc9, 0x69, 0x57, 0x16,
	0x79, 0x9f, 0x0a, 0x65, 0x78, 0x36, 0x60, 0xd2, 0xf7, 0xee, 0x72, 0x83, 0x55, 0x9b, 0x75, 0x64,
	0x93, 0x82, 0x32, 0x87, 0x7c, 0x0f, 0x10, 0x67, 0x4c, 0x28, 0x6a, 0x6d, 0xc2, 0x5f, 0xbd, 0x4b,
	0xa1, 0x8e, 0xe4, 0x0b, 0x91, 0xa0, 0xbb, 0xa5, 0x0c, 0x5d, 0x2d, 0xef, 0x17, 0x26, 0xb6, 0x17,
	0x88, 0x1a, 0xe4, 0x4e, 0x77, 0x73, 0x69, 0xe7, 0x65, 0x16, 0x8a, 0xfd, 0x00, 0xd5, 0xc4, 0xce,
	0xe3, 0x01, 0x1e, 0xd2, 0x8b, 0xff, 0x71, 0x48, 0xef, 0x75, 0xcc, 0x43, 0x4c, 0xb2, 0x57, 0x11,
	0x8b, 0x9c, 0x5d, 0x49, 0x4e, 0xa7, 0xbc, 0xcd, 0x7f, 0x88, 0x7b, 0x5c, 0x2d, 0x43, 0x3f, 0x8d,
	0xad, 0x8d, 0x7c, 0x03, 0x6b, 0x5c, 0x21, 0x3d, 0x29, 0x0c, 0xbf, 0xa6, 0x26, 0x63, 0x91, 0x50,
	0x3d, 0xff, 0x11, 0x26, 0x10, 0x17, 0x7b, 0x6f, 0x43, 0x17, 0x2e, 0x42, 0x9e, 0xc1, 0xd2, 0xc8,
	0x31, 0x9d, 0x85, 0xfb, 0xe8, 0x08, 0x8b, 0x25, 0xe8, 0x1c, 0xfa, 0x39, 0x2c, 0x47, 0x69, 0x91,
	0x73, 0x43, 0x4d, 0x3f, 0xe3, 0x2c, 0xce, 0xfd, 0x0d, 0x14, 0x5c, 0x72, 0xe8, 0x85, 0x03, 0x49,
	0x7b, 0x5c, 0x6d, 0x4c, 0xf9, 0xb5, 0xe1, 0x2a, 0x47, 0x43, 0x7d, 0xdc, 0x98, 0xdd, 0xa9, 0x87,
	0x64, 0x14, 0xea, 0x8c, 0x23, 0xcd, 0x2d, 0xa8, 0x4f, 0x3e, 0x8f, 0x79, 0x98, 0x19, 0x7c, 0xe7,
	0xdd, 0xc3, 0xdf, 0xd7, 0x5e, 0xa5, 0xf9, 0x0a, 0xaa, 0x76, 0x12, 0xa4, 0x0e, 0x73, 0xe8, 0x77,
	0xce, 0x49, 0x2f, 0x99, 0x14, 0x31, 0x33, 0xdc, 0xab, 0xd8, 0xa7, 0x40, 0x09, 0xf3, 0x41, 0xc9,
	0xa1, 0x37, 0xf3, 0xe6, 0x95, 0xf5, 0xcb, 0x1d, 0xf8, 0xfa, 0xbf, 0xfc, 0xf2, 0xf6, 0x90, 0x8f,
	0xab, 0xb5, 0x45, 0x6f, 0xe9, 0xb8, 0x5a, 0x5b, 0xf3, 0xd6, 0x8f, 0xab, 0xb5, 0x75, 0xef, 0x61,
	0x58, 0x4f, 0xd8, 0x35, 0x3a, 0x55, 0x1e, 0x7a, 0xf6, 0xaf, 0xbe, 0xfa, 0x44, 0x15, 0x4b, 0x38,
	0x95, 0x5c, 0xbd, 0xfd, 0x16, 0x1e, 0x0b, 0xed, 0x76, 0x97, 0x66, 0xfa, 0x7a, 0xf8, 0xc5, 0x1a,
	0xdf, 0xae, 0x4c, 0x2c, 0xf9, 0xcc, 0x9e, 0xc5, 0x59, 0xe5, 0x6a, 0x1e, 0xef, 0x63, 0xef, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x88, 0x31, 0x75, 0x27, 0x08, 0x00, 0x00,
}