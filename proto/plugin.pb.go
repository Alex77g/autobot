// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// TextFormat ...
type Message_TextFormat int32

const (
	Message_PLAIN_TEXT Message_TextFormat = 0
)

var Message_TextFormat_name = map[int32]string{
	0: "PLAIN_TEXT",
}

var Message_TextFormat_value = map[string]int32{
	"PLAIN_TEXT": 0,
}

func (x Message_TextFormat) String() string {
	return proto.EnumName(Message_TextFormat_name, int32(x))
}

func (Message_TextFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{11, 0}
}

// Errors to occur in the communication
type Error_Code int32

const (
	Error_UNKNOWN  Error_Code = 0
	Error_REGISTER Error_Code = 1
)

var Error_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "REGISTER",
}

var Error_Code_value = map[string]int32{
	"UNKNOWN":  0,
	"REGISTER": 1,
}

func (x Error_Code) String() string {
	return proto.EnumName(Error_Code_name, int32(x))
}

func (Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{12, 0}
}

// Event ...
//
// These are events that may occur.
type Event struct {
	// Plugin ...
	Plugin *Plugin `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	// UUID ...
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Event ...
	//
	// Types that are valid to be assigned to Event:
	//	*Event_Message
	//	*Event_Reply
	//	*Event_Private
	//	*Event_Error
	Event                isEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetPlugin() *Plugin {
	if m != nil {
		return m.Plugin
	}
	return nil
}

func (m *Event) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type isEvent_Event interface {
	isEvent_Event()
}

type Event_Message struct {
	Message *Message `protobuf:"bytes,10,opt,name=message,proto3,oneof"`
}

type Event_Reply struct {
	Reply *Message `protobuf:"bytes,11,opt,name=reply,proto3,oneof"`
}

type Event_Private struct {
	Private *Message `protobuf:"bytes,12,opt,name=private,proto3,oneof"`
}

type Event_Error struct {
	Error *Error `protobuf:"bytes,20,opt,name=error,proto3,oneof"`
}

func (*Event_Message) isEvent_Event() {}

func (*Event_Reply) isEvent_Event() {}

func (*Event_Private) isEvent_Event() {}

func (*Event_Error) isEvent_Event() {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetMessage() *Message {
	if x, ok := m.GetEvent().(*Event_Message); ok {
		return x.Message
	}
	return nil
}

func (m *Event) GetReply() *Message {
	if x, ok := m.GetEvent().(*Event_Reply); ok {
		return x.Reply
	}
	return nil
}

func (m *Event) GetPrivate() *Message {
	if x, ok := m.GetEvent().(*Event_Private); ok {
		return x.Private
	}
	return nil
}

func (m *Event) GetError() *Error {
	if x, ok := m.GetEvent().(*Event_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_Message)(nil),
		(*Event_Reply)(nil),
		(*Event_Private)(nil),
		(*Event_Error)(nil),
	}
}

// Action ...
//
// These are actions performed by the plugins.
type Action struct {
	// Action ...
	//
	// Types that are valid to be assigned to Action:
	//	*Action_Register
	//	*Action_Config
	//	*Action_Restart
	Action               isAction_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{1}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

type isAction_Action interface {
	isAction_Action()
}

type Action_Register struct {
	Register *Register `protobuf:"bytes,1,opt,name=register,proto3,oneof"`
}

type Action_Config struct {
	Config *Config `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type Action_Restart struct {
	Restart *Restart `protobuf:"bytes,3,opt,name=restart,proto3,oneof"`
}

func (*Action_Register) isAction_Action() {}

func (*Action_Config) isAction_Action() {}

func (*Action_Restart) isAction_Action() {}

func (m *Action) GetAction() isAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Action) GetRegister() *Register {
	if x, ok := m.GetAction().(*Action_Register); ok {
		return x.Register
	}
	return nil
}

func (m *Action) GetConfig() *Config {
	if x, ok := m.GetAction().(*Action_Config); ok {
		return x.Config
	}
	return nil
}

func (m *Action) GetRestart() *Restart {
	if x, ok := m.GetAction().(*Action_Restart); ok {
		return x.Restart
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Action) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Action_Register)(nil),
		(*Action_Config)(nil),
		(*Action_Restart)(nil),
	}
}

// Register ...
//
// Register action for the server.
type Register struct {
	Plugin               *Plugin  `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{2}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetPlugin() *Plugin {
	if m != nil {
		return m.Plugin
	}
	return nil
}

// Restart ...
//
// Restart action from the server.
type Restart struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Restart) Reset()         { *m = Restart{} }
func (m *Restart) String() string { return proto.CompactTextString(m) }
func (*Restart) ProtoMessage()    {}
func (*Restart) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{3}
}

func (m *Restart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restart.Unmarshal(m, b)
}
func (m *Restart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restart.Marshal(b, m, deterministic)
}
func (m *Restart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restart.Merge(m, src)
}
func (m *Restart) XXX_Size() int {
	return xxx_messageInfo_Restart.Size(m)
}
func (m *Restart) XXX_DiscardUnknown() {
	xxx_messageInfo_Restart.DiscardUnknown(m)
}

var xxx_messageInfo_Restart proto.InternalMessageInfo

// Registered ...
//
// Registered reply for the server
type Registered struct {
	// Config ...
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Registered) Reset()         { *m = Registered{} }
func (m *Registered) String() string { return proto.CompactTextString(m) }
func (*Registered) ProtoMessage()    {}
func (*Registered) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{4}
}

func (m *Registered) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registered.Unmarshal(m, b)
}
func (m *Registered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registered.Marshal(b, m, deterministic)
}
func (m *Registered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registered.Merge(m, src)
}
func (m *Registered) XXX_Size() int {
	return xxx_messageInfo_Registered.Size(m)
}
func (m *Registered) XXX_DiscardUnknown() {
	xxx_messageInfo_Registered.DiscardUnknown(m)
}

var xxx_messageInfo_Registered proto.InternalMessageInfo

func (m *Registered) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

// Config ...
//
// Config is a plugin config from the server
type Config struct {
	// cluster id ...
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// cluster url ...
	ClusterUrl string `protobuf:"bytes,2,opt,name=cluster_url,json=clusterUrl,proto3" json:"cluster_url,omitempty"`
	// inbox ...
	Inbox string `protobuf:"bytes,3,opt,name=inbox,proto3" json:"inbox,omitempty"`
	// outbox ...
	Outbox string `protobuf:"bytes,4,opt,name=outbox,proto3" json:"outbox,omitempty"`
	// debug ...
	Debug bool `protobuf:"varint,10,opt,name=debug,proto3" json:"debug,omitempty"`
	// verbose ...
	Verbose              bool     `protobuf:"varint,11,opt,name=verbose,proto3" json:"verbose,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{5}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Config) GetClusterUrl() string {
	if m != nil {
		return m.ClusterUrl
	}
	return ""
}

func (m *Config) GetInbox() string {
	if m != nil {
		return m.Inbox
	}
	return ""
}

func (m *Config) GetOutbox() string {
	if m != nil {
		return m.Outbox
	}
	return ""
}

func (m *Config) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *Config) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

// Plugin ...
//
// These are all info for a plugin.
type Plugin struct {
	// name ...
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// identifier ...
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// version ...
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// path ...
	Path                 string   `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Plugin) Reset()         { *m = Plugin{} }
func (m *Plugin) String() string { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()    {}
func (*Plugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{6}
}

func (m *Plugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plugin.Unmarshal(m, b)
}
func (m *Plugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plugin.Marshal(b, m, deterministic)
}
func (m *Plugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plugin.Merge(m, src)
}
func (m *Plugin) XXX_Size() int {
	return xxx_messageInfo_Plugin.Size(m)
}
func (m *Plugin) XXX_DiscardUnknown() {
	xxx_messageInfo_Plugin.DiscardUnknown(m)
}

var xxx_messageInfo_Plugin proto.InternalMessageInfo

func (m *Plugin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plugin) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Plugin) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Plugin) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// User ...
//
// This is a user interaction ...
type User struct {
	// ID ...
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name ...
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Team ...
	Team                 *Team    `protobuf:"bytes,3,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{7}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

// Team ...
//
// This is a team interaction ...
type Team struct {
	// ID ...
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name ...
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{8}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Channel ...
//
// This is the channel of the interaction ...
type Channel struct {
	// ID ...
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name ...
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{9}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Channel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Recipient ...
//
// This is the recipient of an interaction ...
type Recipient struct {
	// ID ...
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name ...
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Team ...
	Team                 *Team    `protobuf:"bytes,3,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recipient) Reset()         { *m = Recipient{} }
func (m *Recipient) String() string { return proto.CompactTextString(m) }
func (*Recipient) ProtoMessage()    {}
func (*Recipient) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{10}
}

func (m *Recipient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipient.Unmarshal(m, b)
}
func (m *Recipient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipient.Marshal(b, m, deterministic)
}
func (m *Recipient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipient.Merge(m, src)
}
func (m *Recipient) XXX_Size() int {
	return xxx_messageInfo_Recipient.Size(m)
}
func (m *Recipient) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipient.DiscardUnknown(m)
}

var xxx_messageInfo_Recipient proto.InternalMessageInfo

func (m *Recipient) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Recipient) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Recipient) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

// Message
//
// Information from a chat.
type Message struct {
	// ID ...
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type ...
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Channel ...
	Channel *Channel `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	// From ...
	From *User `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	// Recipient ...
	Recipient *Recipient `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// isBot ...
	IsBot bool `protobuf:"varint,6,opt,name=is_bot,json=isBot,proto3" json:"is_bot,omitempty"`
	// isDirectMessage ...
	IsDirectMessage bool `protobuf:"varint,7,opt,name=is_direct_message,json=isDirectMessage,proto3" json:"is_direct_message,omitempty"`
	// Timestamp ...
	Timestamp *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// TextFormat ...
	TextFormat Message_TextFormat `protobuf:"varint,20,opt,name=text_format,json=textFormat,proto3,enum=proto.Message_TextFormat" json:"text_format,omitempty"`
	// Text ...
	Text                 string   `protobuf:"bytes,21,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{11}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Message) GetChannel() *Channel {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *Message) GetFrom() *User {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Message) GetRecipient() *Recipient {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Message) GetIsBot() bool {
	if m != nil {
		return m.IsBot
	}
	return false
}

func (m *Message) GetIsDirectMessage() bool {
	if m != nil {
		return m.IsDirectMessage
	}
	return false
}

func (m *Message) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Message) GetTextFormat() Message_TextFormat {
	if m != nil {
		return m.TextFormat
	}
	return Message_PLAIN_TEXT
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// Error
//
// These are standard error codes.
type Error struct {
	Code                 Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=proto.Error_Code" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{12}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() Error_Code {
	if m != nil {
		return m.Code
	}
	return Error_UNKNOWN
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Empty ...
//
// Empty reply.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{13}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("proto.Message_TextFormat", Message_TextFormat_name, Message_TextFormat_value)
	proto.RegisterEnum("proto.Error_Code", Error_Code_name, Error_Code_value)
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterType((*Action)(nil), "proto.Action")
	proto.RegisterType((*Register)(nil), "proto.Register")
	proto.RegisterType((*Restart)(nil), "proto.Restart")
	proto.RegisterType((*Registered)(nil), "proto.Registered")
	proto.RegisterType((*Config)(nil), "proto.Config")
	proto.RegisterType((*Plugin)(nil), "proto.Plugin")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Team)(nil), "proto.Team")
	proto.RegisterType((*Channel)(nil), "proto.Channel")
	proto.RegisterType((*Recipient)(nil), "proto.Recipient")
	proto.RegisterType((*Message)(nil), "proto.Message")
	proto.RegisterType((*Error)(nil), "proto.Error")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor_22a625af4bc1cc87) }

var fileDescriptor_22a625af4bc1cc87 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0x8d, 0x43, 0x6c, 0xc7, 0x93, 0xdc, 0x10, 0x56, 0x70, 0xe5, 0x8b, 0xee, 0xbd, 0x50, 0xab,
	0xb4, 0x08, 0x09, 0xa3, 0xc2, 0x4b, 0xd5, 0x37, 0xa0, 0x69, 0x41, 0xb4, 0x69, 0xb4, 0x0d, 0x6a,
	0xdf, 0x22, 0x27, 0xde, 0x84, 0x6d, 0x63, 0xaf, 0xb5, 0x5e, 0xa3, 0xf0, 0x1d, 0x95, 0xfa, 0x09,
	0xfd, 0xbf, 0xfe, 0x41, 0xb5, 0xeb, 0x5d, 0x07, 0x2a, 0x90, 0x78, 0xe8, 0x93, 0x77, 0xce, 0x9c,
	0x99, 0x3d, 0x33, 0x3b, 0x1e, 0x68, 0x67, 0xf3, 0x62, 0x46, 0xd3, 0x30, 0xe3, 0x4c, 0x30, 0x64,
	0xab, 0xcf, 0xe6, 0xd6, 0x8c, 0xb1, 0xd9, 0x9c, 0x1c, 0x28, 0x6b, 0x5c, 0x4c, 0x0f, 0x04, 0x4d,
	0x48, 0x2e, 0xa2, 0x24, 0x2b, 0x79, 0xc1, 0x4f, 0x0b, 0xec, 0xde, 0x35, 0x49, 0x05, 0xda, 0x01,
	0xa7, 0xcc, 0xe0, 0x5b, 0xdb, 0xd6, 0x6e, 0xeb, 0xf0, 0xaf, 0x92, 0x11, 0x0e, 0x14, 0x88, 0xb5,
	0x13, 0x21, 0x68, 0x14, 0x05, 0x8d, 0xfd, 0xfa, 0xb6, 0xb5, 0xeb, 0x61, 0x75, 0x46, 0x7b, 0xe0,
	0x26, 0x24, 0xcf, 0xa3, 0x19, 0xf1, 0x41, 0xc5, 0x76, 0x74, 0xec, 0xfb, 0x12, 0x3d, 0xab, 0x61,
	0x43, 0x40, 0xcf, 0xc0, 0xe6, 0x24, 0x9b, 0xdf, 0xf8, 0xad, 0x07, 0x98, 0xa5, 0x5b, 0xe6, 0xcc,
	0x38, 0xbd, 0x8e, 0x04, 0xf1, 0xdb, 0x0f, 0xe5, 0xd4, 0x04, 0xf4, 0x14, 0x6c, 0xc2, 0x39, 0xe3,
	0xfe, 0xba, 0x62, 0xb6, 0x35, 0xb3, 0x27, 0x31, 0x99, 0x51, 0x39, 0x4f, 0x5c, 0xb0, 0x89, 0xac,
	0x34, 0xf8, 0x6e, 0x81, 0x73, 0x3c, 0x11, 0x94, 0xa5, 0x68, 0x1f, 0x9a, 0x9c, 0xcc, 0x68, 0x2e,
	0x08, 0xd7, 0x65, 0xaf, 0xea, 0x60, 0xac, 0xe1, 0xb3, 0x1a, 0xae, 0x28, 0xe8, 0x39, 0x38, 0x13,
	0x96, 0x4e, 0xe9, 0x4c, 0x95, 0xbf, 0xec, 0xd1, 0xa9, 0x02, 0xcf, 0x6a, 0x58, 0xbb, 0xa5, 0x7a,
	0x2e, 0xfb, 0xcc, 0x85, 0xbf, 0x72, 0x47, 0x3d, 0x2e, 0x51, 0xa9, 0x5e, 0x13, 0x4e, 0x9a, 0xe0,
	0x44, 0x4a, 0x4d, 0xf0, 0x02, 0x9a, 0xe6, 0xda, 0x47, 0x3e, 0x47, 0xe0, 0x81, 0xab, 0x53, 0x06,
	0x47, 0x00, 0x26, 0x9a, 0xc4, 0x32, 0x5e, 0x4b, 0xb5, 0xee, 0x91, 0x6a, 0x84, 0x06, 0x3f, 0x2c,
	0x70, 0x4a, 0x08, 0xfd, 0x07, 0x30, 0x99, 0x17, 0x32, 0x7c, 0x44, 0x63, 0x15, 0xe5, 0x61, 0x4f,
	0x23, 0xe7, 0x31, 0xda, 0x82, 0x96, 0x71, 0x17, 0x7c, 0xae, 0xdf, 0xdf, 0x44, 0x5c, 0xf2, 0x39,
	0x5a, 0x07, 0x9b, 0xa6, 0x63, 0xb6, 0x50, 0x15, 0x7b, 0xb8, 0x34, 0xd0, 0xdf, 0xe0, 0xb0, 0x42,
	0x48, 0xb8, 0xa1, 0x60, 0x6d, 0x49, 0x76, 0x4c, 0xc6, 0xc5, 0x4c, 0x4d, 0x4c, 0x13, 0x97, 0x06,
	0xf2, 0xc1, 0xbd, 0x26, 0x7c, 0xcc, 0x72, 0xa2, 0xe6, 0xa3, 0x89, 0x8d, 0x19, 0x7c, 0x01, 0x67,
	0x50, 0x4d, 0x60, 0x1a, 0x25, 0x44, 0x2b, 0x54, 0x67, 0xf4, 0x3f, 0x00, 0x8d, 0x49, 0x2a, 0xe8,
	0x94, 0x12, 0x6e, 0xb4, 0x2d, 0x11, 0x9d, 0x37, 0xa7, 0x2c, 0xd5, 0xea, 0x8c, 0x29, 0xb3, 0x65,
	0x91, 0xb8, 0xd2, 0xea, 0xd4, 0x39, 0xb8, 0x80, 0xc6, 0x65, 0x4e, 0x38, 0xea, 0x40, 0xbd, 0xea,
	0x44, 0x9d, 0xc6, 0xd5, 0xcd, 0xf5, 0x5b, 0x37, 0x6f, 0x41, 0x43, 0x90, 0x28, 0xd1, 0xcf, 0xdc,
	0xd2, 0x5d, 0x1e, 0x92, 0x28, 0xc1, 0xca, 0x11, 0xec, 0x41, 0x43, 0x5a, 0x8f, 0x49, 0x16, 0xec,
	0x83, 0x7b, 0x7a, 0x15, 0xa5, 0x29, 0x99, 0x3f, 0x8a, 0x3e, 0x00, 0x0f, 0x93, 0x09, 0xcd, 0xa8,
	0xfc, 0x7f, 0xff, 0x88, 0xd8, 0x6f, 0x2b, 0xe0, 0xea, 0x1f, 0xec, 0xbe, 0x84, 0xe2, 0x26, 0xab,
	0x12, 0xca, 0x33, 0xda, 0x05, 0x77, 0x52, 0x0a, 0xfe, 0x6d, 0xce, 0x75, 0x19, 0xd8, 0xb8, 0xe5,
	0xd5, 0x53, 0xce, 0x12, 0xd5, 0xe7, 0xe5, 0xd5, 0xb2, 0xcd, 0x58, 0x39, 0x50, 0x08, 0x1e, 0x37,
	0xc5, 0xf8, 0xb6, 0x62, 0x75, 0xab, 0x9f, 0x46, 0xe3, 0x78, 0x49, 0x41, 0x1b, 0xe0, 0xd0, 0x7c,
	0x34, 0x66, 0xc2, 0x77, 0xca, 0x09, 0xa2, 0xf9, 0x09, 0x13, 0x68, 0x0f, 0xd6, 0x68, 0x3e, 0x8a,
	0x29, 0x27, 0x13, 0x31, 0x32, 0x5b, 0xc9, 0x55, 0x8c, 0x55, 0x9a, 0xbf, 0x56, 0xb8, 0xa9, 0xf0,
	0x25, 0x78, 0xd5, 0x3e, 0xd4, 0x9b, 0x6b, 0x33, 0x2c, 0x37, 0x66, 0x68, 0x36, 0x66, 0x38, 0x34,
	0x0c, 0xbc, 0x24, 0xa3, 0x57, 0xd0, 0x12, 0x64, 0x21, 0x46, 0x53, 0xc6, 0x93, 0x48, 0xa8, 0xbd,
	0xd3, 0x39, 0xfc, 0xe7, 0xee, 0x86, 0x0a, 0x87, 0x64, 0x21, 0xde, 0x28, 0x02, 0x06, 0x51, 0x9d,
	0x55, 0x1f, 0xc9, 0x42, 0xf8, 0x1b, 0xba, 0x8f, 0x64, 0x21, 0x82, 0x7f, 0x01, 0x96, 0x6c, 0xd4,
	0x01, 0x18, 0xbc, 0x3b, 0x3e, 0xef, 0x8f, 0x86, 0xbd, 0xcf, 0xc3, 0x6e, 0x2d, 0xf8, 0x0a, 0xb6,
	0xda, 0x65, 0x68, 0x07, 0x1a, 0x13, 0x16, 0x97, 0xa3, 0xdf, 0x39, 0x5c, 0xbb, 0xbd, 0xe7, 0xc2,
	0x53, 0x16, 0x13, 0xac, 0xdc, 0x72, 0xda, 0x4d, 0xe5, 0xe5, 0x63, 0x19, 0x33, 0x78, 0x02, 0x0d,
	0xc9, 0x43, 0x2d, 0x70, 0x2f, 0xfb, 0x17, 0xfd, 0x0f, 0x9f, 0xfa, 0xdd, 0x1a, 0x6a, 0x43, 0x13,
	0xf7, 0xde, 0x9e, 0x7f, 0x1c, 0xf6, 0x70, 0xd7, 0x0a, 0x5c, 0xb0, 0x7b, 0x49, 0x26, 0x6e, 0xc6,
	0x8e, 0xca, 0x7e, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x36, 0x46, 0xaf, 0x5b, 0x59, 0x06, 0x00,
	0x00,
}
