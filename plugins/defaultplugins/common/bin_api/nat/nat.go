// Package nat represents the VPP binary API of the 'nat' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api/nat.api.json'
package nat

import "git.fd.io/govpp.git/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xdde46fca

// Nat44LbAddrPort represents the VPP binary API data type 'nat44_lb_addr_port'.
// Generated from '/usr/share/vpp/api/nat.api.json', line 1546:
//
//                "vl_api_nat44_lb_addr_port_t",
//                "locals",
//                0,
//
type Nat44LbAddrPort struct {
	Addr        []byte `struc:"[4]byte"`
	Port        uint16
	Probability uint8
}

func (*Nat44LbAddrPort) GetTypeName() string {
	return "nat44_lb_addr_port"
}
func (*Nat44LbAddrPort) GetCrcString() string {
	return "513cf9d0"
}

// NatControlPing represents the VPP binary API message 'nat_control_ping'.
//
type NatControlPing struct {
}

func (*NatControlPing) GetMessageName() string {
	return "nat_control_ping"
}
func (*NatControlPing) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatControlPing) GetCrcString() string {
	return "51077d14"
}
func NewNatControlPing() api.Message {
	return &NatControlPing{}
}

// NatControlPingReply represents the VPP binary API message 'nat_control_ping_reply'.
//
type NatControlPingReply struct {
	Retval      int32
	ClientIndex uint32
	VpePid      uint32
}

func (*NatControlPingReply) GetMessageName() string {
	return "nat_control_ping_reply"
}
func (*NatControlPingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatControlPingReply) GetCrcString() string {
	return "f6b0b8ca"
}
func NewNatControlPingReply() api.Message {
	return &NatControlPingReply{}
}

// NatShowConfig represents the VPP binary API message 'nat_show_config'.
//
type NatShowConfig struct {
}

func (*NatShowConfig) GetMessageName() string {
	return "nat_show_config"
}
func (*NatShowConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatShowConfig) GetCrcString() string {
	return "51077d14"
}
func NewNatShowConfig() api.Message {
	return &NatShowConfig{}
}

// NatShowConfigReply represents the VPP binary API message 'nat_show_config_reply'.
//
type NatShowConfigReply struct {
	Retval                          int32
	StaticMappingOnly               uint8
	StaticMappingConnectionTracking uint8
	Deterministic                   uint8
	TranslationBuckets              uint32
	TranslationMemorySize           uint32
	UserBuckets                     uint32
	UserMemorySize                  uint32
	MaxTranslationsPerUser          uint32
	OutsideVrfID                    uint32
	InsideVrfID                     uint32
}

func (*NatShowConfigReply) GetMessageName() string {
	return "nat_show_config_reply"
}
func (*NatShowConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatShowConfigReply) GetCrcString() string {
	return "7685fc1c"
}
func NewNatShowConfigReply() api.Message {
	return &NatShowConfigReply{}
}

// NatSetWorkers represents the VPP binary API message 'nat_set_workers'.
//
type NatSetWorkers struct {
	WorkerMask uint64
}

func (*NatSetWorkers) GetMessageName() string {
	return "nat_set_workers"
}
func (*NatSetWorkers) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatSetWorkers) GetCrcString() string {
	return "da926638"
}
func NewNatSetWorkers() api.Message {
	return &NatSetWorkers{}
}

// NatSetWorkersReply represents the VPP binary API message 'nat_set_workers_reply'.
//
type NatSetWorkersReply struct {
	Retval int32
}

func (*NatSetWorkersReply) GetMessageName() string {
	return "nat_set_workers_reply"
}
func (*NatSetWorkersReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatSetWorkersReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNatSetWorkersReply() api.Message {
	return &NatSetWorkersReply{}
}

// NatWorkerDump represents the VPP binary API message 'nat_worker_dump'.
//
type NatWorkerDump struct {
}

func (*NatWorkerDump) GetMessageName() string {
	return "nat_worker_dump"
}
func (*NatWorkerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatWorkerDump) GetCrcString() string {
	return "51077d14"
}
func NewNatWorkerDump() api.Message {
	return &NatWorkerDump{}
}

// NatWorkerDetails represents the VPP binary API message 'nat_worker_details'.
//
type NatWorkerDetails struct {
	WorkerIndex uint32
	LcoreID     uint32
	Name        []byte `struc:"[64]byte"`
}

func (*NatWorkerDetails) GetMessageName() string {
	return "nat_worker_details"
}
func (*NatWorkerDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatWorkerDetails) GetCrcString() string {
	return "2e3f9d4b"
}
func NewNatWorkerDetails() api.Message {
	return &NatWorkerDetails{}
}

// NatIpfixEnableDisable represents the VPP binary API message 'nat_ipfix_enable_disable'.
//
type NatIpfixEnableDisable struct {
	DomainID uint32
	SrcPort  uint16
	Enable   uint8
}

func (*NatIpfixEnableDisable) GetMessageName() string {
	return "nat_ipfix_enable_disable"
}
func (*NatIpfixEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatIpfixEnableDisable) GetCrcString() string {
	return "745dd24b"
}
func NewNatIpfixEnableDisable() api.Message {
	return &NatIpfixEnableDisable{}
}

// NatIpfixEnableDisableReply represents the VPP binary API message 'nat_ipfix_enable_disable_reply'.
//
type NatIpfixEnableDisableReply struct {
	Retval int32
}

func (*NatIpfixEnableDisableReply) GetMessageName() string {
	return "nat_ipfix_enable_disable_reply"
}
func (*NatIpfixEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatIpfixEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNatIpfixEnableDisableReply() api.Message {
	return &NatIpfixEnableDisableReply{}
}

// NatSetReass represents the VPP binary API message 'nat_set_reass'.
//
type NatSetReass struct {
	Timeout  uint32
	MaxReass uint16
	MaxFrag  uint8
	DropFrag uint8
	IsIP6    uint8
}

func (*NatSetReass) GetMessageName() string {
	return "nat_set_reass"
}
func (*NatSetReass) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatSetReass) GetCrcString() string {
	return "cb126174"
}
func NewNatSetReass() api.Message {
	return &NatSetReass{}
}

// NatSetReassReply represents the VPP binary API message 'nat_set_reass_reply'.
//
type NatSetReassReply struct {
	Retval int32
}

func (*NatSetReassReply) GetMessageName() string {
	return "nat_set_reass_reply"
}
func (*NatSetReassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatSetReassReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNatSetReassReply() api.Message {
	return &NatSetReassReply{}
}

// NatGetReass represents the VPP binary API message 'nat_get_reass'.
//
type NatGetReass struct {
}

func (*NatGetReass) GetMessageName() string {
	return "nat_get_reass"
}
func (*NatGetReass) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatGetReass) GetCrcString() string {
	return "51077d14"
}
func NewNatGetReass() api.Message {
	return &NatGetReass{}
}

// NatGetReassReply represents the VPP binary API message 'nat_get_reass_reply'.
//
type NatGetReassReply struct {
	Retval      int32
	IP4Timeout  uint32
	IP4MaxReass uint16
	IP4MaxFrag  uint8
	IP4DropFrag uint8
	IP6Timeout  uint32
	IP6MaxReass uint16
	IP6MaxFrag  uint8
	IP6DropFrag uint8
}

func (*NatGetReassReply) GetMessageName() string {
	return "nat_get_reass_reply"
}
func (*NatGetReassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatGetReassReply) GetCrcString() string {
	return "8102a0fb"
}
func NewNatGetReassReply() api.Message {
	return &NatGetReassReply{}
}

// NatReassDump represents the VPP binary API message 'nat_reass_dump'.
//
type NatReassDump struct {
}

func (*NatReassDump) GetMessageName() string {
	return "nat_reass_dump"
}
func (*NatReassDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatReassDump) GetCrcString() string {
	return "51077d14"
}
func NewNatReassDump() api.Message {
	return &NatReassDump{}
}

// NatReassDetails represents the VPP binary API message 'nat_reass_details'.
//
type NatReassDetails struct {
	IsIP4   uint8
	SrcAddr []byte `struc:"[16]byte"`
	DstAddr []byte `struc:"[16]byte"`
	FragID  uint32
	Proto   uint8
	FragN   uint8
}

func (*NatReassDetails) GetMessageName() string {
	return "nat_reass_details"
}
func (*NatReassDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatReassDetails) GetCrcString() string {
	return "ee46e2d4"
}
func NewNatReassDetails() api.Message {
	return &NatReassDetails{}
}

// Nat44AddDelAddressRange represents the VPP binary API message 'nat44_add_del_address_range'.
//
type Nat44AddDelAddressRange struct {
	FirstIPAddress []byte `struc:"[4]byte"`
	LastIPAddress  []byte `struc:"[4]byte"`
	VrfID          uint32
	TwiceNat       uint8
	IsAdd          uint8
}

func (*Nat44AddDelAddressRange) GetMessageName() string {
	return "nat44_add_del_address_range"
}
func (*Nat44AddDelAddressRange) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44AddDelAddressRange) GetCrcString() string {
	return "4a7d5c11"
}
func NewNat44AddDelAddressRange() api.Message {
	return &Nat44AddDelAddressRange{}
}

// Nat44AddDelAddressRangeReply represents the VPP binary API message 'nat44_add_del_address_range_reply'.
//
type Nat44AddDelAddressRangeReply struct {
	Retval int32
}

func (*Nat44AddDelAddressRangeReply) GetMessageName() string {
	return "nat44_add_del_address_range_reply"
}
func (*Nat44AddDelAddressRangeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44AddDelAddressRangeReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44AddDelAddressRangeReply() api.Message {
	return &Nat44AddDelAddressRangeReply{}
}

// Nat44AddressDump represents the VPP binary API message 'nat44_address_dump'.
//
type Nat44AddressDump struct {
}

func (*Nat44AddressDump) GetMessageName() string {
	return "nat44_address_dump"
}
func (*Nat44AddressDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44AddressDump) GetCrcString() string {
	return "51077d14"
}
func NewNat44AddressDump() api.Message {
	return &Nat44AddressDump{}
}

// Nat44AddressDetails represents the VPP binary API message 'nat44_address_details'.
//
type Nat44AddressDetails struct {
	IPAddress []byte `struc:"[4]byte"`
	TwiceNat  uint8
	VrfID     uint32
}

func (*Nat44AddressDetails) GetMessageName() string {
	return "nat44_address_details"
}
func (*Nat44AddressDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44AddressDetails) GetCrcString() string {
	return "512feae8"
}
func NewNat44AddressDetails() api.Message {
	return &Nat44AddressDetails{}
}

// Nat44InterfaceAddDelFeature represents the VPP binary API message 'nat44_interface_add_del_feature'.
//
type Nat44InterfaceAddDelFeature struct {
	IsAdd     uint8
	IsInside  uint8
	SwIfIndex uint32
}

func (*Nat44InterfaceAddDelFeature) GetMessageName() string {
	return "nat44_interface_add_del_feature"
}
func (*Nat44InterfaceAddDelFeature) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44InterfaceAddDelFeature) GetCrcString() string {
	return "9b1ac600"
}
func NewNat44InterfaceAddDelFeature() api.Message {
	return &Nat44InterfaceAddDelFeature{}
}

// Nat44InterfaceAddDelFeatureReply represents the VPP binary API message 'nat44_interface_add_del_feature_reply'.
//
type Nat44InterfaceAddDelFeatureReply struct {
	Retval int32
}

func (*Nat44InterfaceAddDelFeatureReply) GetMessageName() string {
	return "nat44_interface_add_del_feature_reply"
}
func (*Nat44InterfaceAddDelFeatureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44InterfaceAddDelFeatureReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44InterfaceAddDelFeatureReply() api.Message {
	return &Nat44InterfaceAddDelFeatureReply{}
}

// Nat44InterfaceDump represents the VPP binary API message 'nat44_interface_dump'.
//
type Nat44InterfaceDump struct {
}

func (*Nat44InterfaceDump) GetMessageName() string {
	return "nat44_interface_dump"
}
func (*Nat44InterfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44InterfaceDump) GetCrcString() string {
	return "51077d14"
}
func NewNat44InterfaceDump() api.Message {
	return &Nat44InterfaceDump{}
}

// Nat44InterfaceDetails represents the VPP binary API message 'nat44_interface_details'.
//
type Nat44InterfaceDetails struct {
	IsInside  uint8
	SwIfIndex uint32
}

func (*Nat44InterfaceDetails) GetMessageName() string {
	return "nat44_interface_details"
}
func (*Nat44InterfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44InterfaceDetails) GetCrcString() string {
	return "2b15e8e4"
}
func NewNat44InterfaceDetails() api.Message {
	return &Nat44InterfaceDetails{}
}

// Nat44InterfaceAddDelOutputFeature represents the VPP binary API message 'nat44_interface_add_del_output_feature'.
//
type Nat44InterfaceAddDelOutputFeature struct {
	IsAdd     uint8
	IsInside  uint8
	SwIfIndex uint32
}

func (*Nat44InterfaceAddDelOutputFeature) GetMessageName() string {
	return "nat44_interface_add_del_output_feature"
}
func (*Nat44InterfaceAddDelOutputFeature) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44InterfaceAddDelOutputFeature) GetCrcString() string {
	return "9b1ac600"
}
func NewNat44InterfaceAddDelOutputFeature() api.Message {
	return &Nat44InterfaceAddDelOutputFeature{}
}

// Nat44InterfaceAddDelOutputFeatureReply represents the VPP binary API message 'nat44_interface_add_del_output_feature_reply'.
//
type Nat44InterfaceAddDelOutputFeatureReply struct {
	Retval int32
}

func (*Nat44InterfaceAddDelOutputFeatureReply) GetMessageName() string {
	return "nat44_interface_add_del_output_feature_reply"
}
func (*Nat44InterfaceAddDelOutputFeatureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44InterfaceAddDelOutputFeatureReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44InterfaceAddDelOutputFeatureReply() api.Message {
	return &Nat44InterfaceAddDelOutputFeatureReply{}
}

// Nat44InterfaceOutputFeatureDump represents the VPP binary API message 'nat44_interface_output_feature_dump'.
//
type Nat44InterfaceOutputFeatureDump struct {
}

func (*Nat44InterfaceOutputFeatureDump) GetMessageName() string {
	return "nat44_interface_output_feature_dump"
}
func (*Nat44InterfaceOutputFeatureDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44InterfaceOutputFeatureDump) GetCrcString() string {
	return "51077d14"
}
func NewNat44InterfaceOutputFeatureDump() api.Message {
	return &Nat44InterfaceOutputFeatureDump{}
}

// Nat44InterfaceOutputFeatureDetails represents the VPP binary API message 'nat44_interface_output_feature_details'.
//
type Nat44InterfaceOutputFeatureDetails struct {
	IsInside  uint8
	SwIfIndex uint32
}

func (*Nat44InterfaceOutputFeatureDetails) GetMessageName() string {
	return "nat44_interface_output_feature_details"
}
func (*Nat44InterfaceOutputFeatureDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44InterfaceOutputFeatureDetails) GetCrcString() string {
	return "2b15e8e4"
}
func NewNat44InterfaceOutputFeatureDetails() api.Message {
	return &Nat44InterfaceOutputFeatureDetails{}
}

// Nat44AddDelStaticMapping represents the VPP binary API message 'nat44_add_del_static_mapping'.
//
type Nat44AddDelStaticMapping struct {
	IsAdd             uint8
	AddrOnly          uint8
	LocalIPAddress    []byte `struc:"[4]byte"`
	ExternalIPAddress []byte `struc:"[4]byte"`
	Protocol          uint8
	LocalPort         uint16
	ExternalPort      uint16
	ExternalSwIfIndex uint32
	VrfID             uint32
	TwiceNat          uint8
	Out2inOnly        uint8
}

func (*Nat44AddDelStaticMapping) GetMessageName() string {
	return "nat44_add_del_static_mapping"
}
func (*Nat44AddDelStaticMapping) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44AddDelStaticMapping) GetCrcString() string {
	return "eba57973"
}
func NewNat44AddDelStaticMapping() api.Message {
	return &Nat44AddDelStaticMapping{}
}

// Nat44AddDelStaticMappingReply represents the VPP binary API message 'nat44_add_del_static_mapping_reply'.
//
type Nat44AddDelStaticMappingReply struct {
	Retval int32
}

func (*Nat44AddDelStaticMappingReply) GetMessageName() string {
	return "nat44_add_del_static_mapping_reply"
}
func (*Nat44AddDelStaticMappingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44AddDelStaticMappingReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44AddDelStaticMappingReply() api.Message {
	return &Nat44AddDelStaticMappingReply{}
}

// Nat44StaticMappingDump represents the VPP binary API message 'nat44_static_mapping_dump'.
//
type Nat44StaticMappingDump struct {
}

func (*Nat44StaticMappingDump) GetMessageName() string {
	return "nat44_static_mapping_dump"
}
func (*Nat44StaticMappingDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44StaticMappingDump) GetCrcString() string {
	return "51077d14"
}
func NewNat44StaticMappingDump() api.Message {
	return &Nat44StaticMappingDump{}
}

// Nat44StaticMappingDetails represents the VPP binary API message 'nat44_static_mapping_details'.
//
type Nat44StaticMappingDetails struct {
	AddrOnly          uint8
	LocalIPAddress    []byte `struc:"[4]byte"`
	ExternalIPAddress []byte `struc:"[4]byte"`
	Protocol          uint8
	LocalPort         uint16
	ExternalPort      uint16
	ExternalSwIfIndex uint32
	VrfID             uint32
	TwiceNat          uint8
	Out2inOnly        uint8
}

func (*Nat44StaticMappingDetails) GetMessageName() string {
	return "nat44_static_mapping_details"
}
func (*Nat44StaticMappingDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44StaticMappingDetails) GetCrcString() string {
	return "d68ab782"
}
func NewNat44StaticMappingDetails() api.Message {
	return &Nat44StaticMappingDetails{}
}

// Nat44AddDelIdentityMapping represents the VPP binary API message 'nat44_add_del_identity_mapping'.
//
type Nat44AddDelIdentityMapping struct {
	IsAdd     uint8
	AddrOnly  uint8
	IPAddress []byte `struc:"[4]byte"`
	Protocol  uint8
	Port      uint16
	SwIfIndex uint32
	VrfID     uint32
}

func (*Nat44AddDelIdentityMapping) GetMessageName() string {
	return "nat44_add_del_identity_mapping"
}
func (*Nat44AddDelIdentityMapping) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44AddDelIdentityMapping) GetCrcString() string {
	return "a9af2636"
}
func NewNat44AddDelIdentityMapping() api.Message {
	return &Nat44AddDelIdentityMapping{}
}

// Nat44AddDelIdentityMappingReply represents the VPP binary API message 'nat44_add_del_identity_mapping_reply'.
//
type Nat44AddDelIdentityMappingReply struct {
	Retval int32
}

func (*Nat44AddDelIdentityMappingReply) GetMessageName() string {
	return "nat44_add_del_identity_mapping_reply"
}
func (*Nat44AddDelIdentityMappingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44AddDelIdentityMappingReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44AddDelIdentityMappingReply() api.Message {
	return &Nat44AddDelIdentityMappingReply{}
}

// Nat44IdentityMappingDump represents the VPP binary API message 'nat44_identity_mapping_dump'.
//
type Nat44IdentityMappingDump struct {
}

func (*Nat44IdentityMappingDump) GetMessageName() string {
	return "nat44_identity_mapping_dump"
}
func (*Nat44IdentityMappingDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44IdentityMappingDump) GetCrcString() string {
	return "51077d14"
}
func NewNat44IdentityMappingDump() api.Message {
	return &Nat44IdentityMappingDump{}
}

// Nat44IdentityMappingDetails represents the VPP binary API message 'nat44_identity_mapping_details'.
//
type Nat44IdentityMappingDetails struct {
	AddrOnly  uint8
	IPAddress []byte `struc:"[4]byte"`
	Protocol  uint8
	Port      uint16
	SwIfIndex uint32
	VrfID     uint32
}

func (*Nat44IdentityMappingDetails) GetMessageName() string {
	return "nat44_identity_mapping_details"
}
func (*Nat44IdentityMappingDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44IdentityMappingDetails) GetCrcString() string {
	return "1bbc8eef"
}
func NewNat44IdentityMappingDetails() api.Message {
	return &Nat44IdentityMappingDetails{}
}

// Nat44AddDelInterfaceAddr represents the VPP binary API message 'nat44_add_del_interface_addr'.
//
type Nat44AddDelInterfaceAddr struct {
	IsAdd     uint8
	TwiceNat  uint8
	SwIfIndex uint32
}

func (*Nat44AddDelInterfaceAddr) GetMessageName() string {
	return "nat44_add_del_interface_addr"
}
func (*Nat44AddDelInterfaceAddr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44AddDelInterfaceAddr) GetCrcString() string {
	return "61105dfa"
}
func NewNat44AddDelInterfaceAddr() api.Message {
	return &Nat44AddDelInterfaceAddr{}
}

// Nat44AddDelInterfaceAddrReply represents the VPP binary API message 'nat44_add_del_interface_addr_reply'.
//
type Nat44AddDelInterfaceAddrReply struct {
	Retval int32
}

func (*Nat44AddDelInterfaceAddrReply) GetMessageName() string {
	return "nat44_add_del_interface_addr_reply"
}
func (*Nat44AddDelInterfaceAddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44AddDelInterfaceAddrReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44AddDelInterfaceAddrReply() api.Message {
	return &Nat44AddDelInterfaceAddrReply{}
}

// Nat44InterfaceAddrDump represents the VPP binary API message 'nat44_interface_addr_dump'.
//
type Nat44InterfaceAddrDump struct {
}

func (*Nat44InterfaceAddrDump) GetMessageName() string {
	return "nat44_interface_addr_dump"
}
func (*Nat44InterfaceAddrDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44InterfaceAddrDump) GetCrcString() string {
	return "51077d14"
}
func NewNat44InterfaceAddrDump() api.Message {
	return &Nat44InterfaceAddrDump{}
}

// Nat44InterfaceAddrDetails represents the VPP binary API message 'nat44_interface_addr_details'.
//
type Nat44InterfaceAddrDetails struct {
	SwIfIndex uint32
	TwiceNat  uint8
}

func (*Nat44InterfaceAddrDetails) GetMessageName() string {
	return "nat44_interface_addr_details"
}
func (*Nat44InterfaceAddrDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44InterfaceAddrDetails) GetCrcString() string {
	return "4cdc575d"
}
func NewNat44InterfaceAddrDetails() api.Message {
	return &Nat44InterfaceAddrDetails{}
}

// Nat44UserDump represents the VPP binary API message 'nat44_user_dump'.
//
type Nat44UserDump struct {
}

func (*Nat44UserDump) GetMessageName() string {
	return "nat44_user_dump"
}
func (*Nat44UserDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44UserDump) GetCrcString() string {
	return "51077d14"
}
func NewNat44UserDump() api.Message {
	return &Nat44UserDump{}
}

// Nat44UserDetails represents the VPP binary API message 'nat44_user_details'.
//
type Nat44UserDetails struct {
	VrfID           uint32
	IPAddress       []byte `struc:"[4]byte"`
	Nsessions       uint32
	Nstaticsessions uint32
}

func (*Nat44UserDetails) GetMessageName() string {
	return "nat44_user_details"
}
func (*Nat44UserDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44UserDetails) GetCrcString() string {
	return "abb91460"
}
func NewNat44UserDetails() api.Message {
	return &Nat44UserDetails{}
}

// Nat44UserSessionDump represents the VPP binary API message 'nat44_user_session_dump'.
//
type Nat44UserSessionDump struct {
	IPAddress []byte `struc:"[4]byte"`
	VrfID     uint32
}

func (*Nat44UserSessionDump) GetMessageName() string {
	return "nat44_user_session_dump"
}
func (*Nat44UserSessionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44UserSessionDump) GetCrcString() string {
	return "bca4d2c2"
}
func NewNat44UserSessionDump() api.Message {
	return &Nat44UserSessionDump{}
}

// Nat44UserSessionDetails represents the VPP binary API message 'nat44_user_session_details'.
//
type Nat44UserSessionDetails struct {
	OutsideIPAddress []byte `struc:"[4]byte"`
	OutsidePort      uint16
	InsideIPAddress  []byte `struc:"[4]byte"`
	InsidePort       uint16
	Protocol         uint16
	IsStatic         uint8
	LastHeard        uint64
	TotalBytes       uint64
	TotalPkts        uint32
}

func (*Nat44UserSessionDetails) GetMessageName() string {
	return "nat44_user_session_details"
}
func (*Nat44UserSessionDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44UserSessionDetails) GetCrcString() string {
	return "2c9755fa"
}
func NewNat44UserSessionDetails() api.Message {
	return &Nat44UserSessionDetails{}
}

// Nat44AddDelLbStaticMapping represents the VPP binary API message 'nat44_add_del_lb_static_mapping'.
//
type Nat44AddDelLbStaticMapping struct {
	IsAdd        uint8
	ExternalAddr []byte `struc:"[4]byte"`
	ExternalPort uint16
	Protocol     uint8
	VrfID        uint32
	TwiceNat     uint8
	Out2inOnly   uint8
	LocalNum     uint8 `struc:"sizeof=Locals"`
	Locals       []Nat44LbAddrPort
}

func (*Nat44AddDelLbStaticMapping) GetMessageName() string {
	return "nat44_add_del_lb_static_mapping"
}
func (*Nat44AddDelLbStaticMapping) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44AddDelLbStaticMapping) GetCrcString() string {
	return "48886b32"
}
func NewNat44AddDelLbStaticMapping() api.Message {
	return &Nat44AddDelLbStaticMapping{}
}

// Nat44AddDelLbStaticMappingReply represents the VPP binary API message 'nat44_add_del_lb_static_mapping_reply'.
//
type Nat44AddDelLbStaticMappingReply struct {
	Retval int32
}

func (*Nat44AddDelLbStaticMappingReply) GetMessageName() string {
	return "nat44_add_del_lb_static_mapping_reply"
}
func (*Nat44AddDelLbStaticMappingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44AddDelLbStaticMappingReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44AddDelLbStaticMappingReply() api.Message {
	return &Nat44AddDelLbStaticMappingReply{}
}

// Nat44LbStaticMappingDump represents the VPP binary API message 'nat44_lb_static_mapping_dump'.
//
type Nat44LbStaticMappingDump struct {
}

func (*Nat44LbStaticMappingDump) GetMessageName() string {
	return "nat44_lb_static_mapping_dump"
}
func (*Nat44LbStaticMappingDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44LbStaticMappingDump) GetCrcString() string {
	return "51077d14"
}
func NewNat44LbStaticMappingDump() api.Message {
	return &Nat44LbStaticMappingDump{}
}

// Nat44LbStaticMappingDetails represents the VPP binary API message 'nat44_lb_static_mapping_details'.
//
type Nat44LbStaticMappingDetails struct {
	ExternalAddr []byte `struc:"[4]byte"`
	ExternalPort uint16
	Protocol     uint8
	VrfID        uint32
	TwiceNat     uint8
	Out2inOnly   uint8
	LocalNum     uint8 `struc:"sizeof=Locals"`
	Locals       []Nat44LbAddrPort
}

func (*Nat44LbStaticMappingDetails) GetMessageName() string {
	return "nat44_lb_static_mapping_details"
}
func (*Nat44LbStaticMappingDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44LbStaticMappingDetails) GetCrcString() string {
	return "df9f0583"
}
func NewNat44LbStaticMappingDetails() api.Message {
	return &Nat44LbStaticMappingDetails{}
}

// Nat44DelSession represents the VPP binary API message 'nat44_del_session'.
//
type Nat44DelSession struct {
	IsIn     uint8
	Address  []byte `struc:"[4]byte"`
	Protocol uint8
	Port     uint16
	VrfID    uint32
}

func (*Nat44DelSession) GetMessageName() string {
	return "nat44_del_session"
}
func (*Nat44DelSession) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44DelSession) GetCrcString() string {
	return "9d56a36e"
}
func NewNat44DelSession() api.Message {
	return &Nat44DelSession{}
}

// Nat44DelSessionReply represents the VPP binary API message 'nat44_del_session_reply'.
//
type Nat44DelSessionReply struct {
	Retval int32
}

func (*Nat44DelSessionReply) GetMessageName() string {
	return "nat44_del_session_reply"
}
func (*Nat44DelSessionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44DelSessionReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44DelSessionReply() api.Message {
	return &Nat44DelSessionReply{}
}

// Nat44ForwardingEnableDisable represents the VPP binary API message 'nat44_forwarding_enable_disable'.
//
type Nat44ForwardingEnableDisable struct {
	Enable uint8
}

func (*Nat44ForwardingEnableDisable) GetMessageName() string {
	return "nat44_forwarding_enable_disable"
}
func (*Nat44ForwardingEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44ForwardingEnableDisable) GetCrcString() string {
	return "8050327d"
}
func NewNat44ForwardingEnableDisable() api.Message {
	return &Nat44ForwardingEnableDisable{}
}

// Nat44ForwardingEnableDisableReply represents the VPP binary API message 'nat44_forwarding_enable_disable_reply'.
//
type Nat44ForwardingEnableDisableReply struct {
	Retval int32
}

func (*Nat44ForwardingEnableDisableReply) GetMessageName() string {
	return "nat44_forwarding_enable_disable_reply"
}
func (*Nat44ForwardingEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44ForwardingEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat44ForwardingEnableDisableReply() api.Message {
	return &Nat44ForwardingEnableDisableReply{}
}

// Nat44ForwardingIsEnabled represents the VPP binary API message 'nat44_forwarding_is_enabled'.
//
type Nat44ForwardingIsEnabled struct {
}

func (*Nat44ForwardingIsEnabled) GetMessageName() string {
	return "nat44_forwarding_is_enabled"
}
func (*Nat44ForwardingIsEnabled) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat44ForwardingIsEnabled) GetCrcString() string {
	return "51077d14"
}
func NewNat44ForwardingIsEnabled() api.Message {
	return &Nat44ForwardingIsEnabled{}
}

// Nat44ForwardingIsEnabledReply represents the VPP binary API message 'nat44_forwarding_is_enabled_reply'.
//
type Nat44ForwardingIsEnabledReply struct {
	Enabled uint8
}

func (*Nat44ForwardingIsEnabledReply) GetMessageName() string {
	return "nat44_forwarding_is_enabled_reply"
}
func (*Nat44ForwardingIsEnabledReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat44ForwardingIsEnabledReply) GetCrcString() string {
	return "9c4a7828"
}
func NewNat44ForwardingIsEnabledReply() api.Message {
	return &Nat44ForwardingIsEnabledReply{}
}

// NatDetAddDelMap represents the VPP binary API message 'nat_det_add_del_map'.
//
type NatDetAddDelMap struct {
	IsAdd    uint8
	IsNat44  uint8
	AddrOnly uint8
	InAddr   []byte `struc:"[16]byte"`
	InPlen   uint8
	OutAddr  []byte `struc:"[4]byte"`
	OutPlen  uint8
}

func (*NatDetAddDelMap) GetMessageName() string {
	return "nat_det_add_del_map"
}
func (*NatDetAddDelMap) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetAddDelMap) GetCrcString() string {
	return "5bd37d5b"
}
func NewNatDetAddDelMap() api.Message {
	return &NatDetAddDelMap{}
}

// NatDetAddDelMapReply represents the VPP binary API message 'nat_det_add_del_map_reply'.
//
type NatDetAddDelMapReply struct {
	Retval int32
}

func (*NatDetAddDelMapReply) GetMessageName() string {
	return "nat_det_add_del_map_reply"
}
func (*NatDetAddDelMapReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatDetAddDelMapReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNatDetAddDelMapReply() api.Message {
	return &NatDetAddDelMapReply{}
}

// NatDetForward represents the VPP binary API message 'nat_det_forward'.
//
type NatDetForward struct {
	IsNat44 uint8
	InAddr  []byte `struc:"[16]byte"`
}

func (*NatDetForward) GetMessageName() string {
	return "nat_det_forward"
}
func (*NatDetForward) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetForward) GetCrcString() string {
	return "037d399b"
}
func NewNatDetForward() api.Message {
	return &NatDetForward{}
}

// NatDetForwardReply represents the VPP binary API message 'nat_det_forward_reply'.
//
type NatDetForwardReply struct {
	Retval    int32
	OutPortLo uint16
	OutPortHi uint16
	OutAddr   []byte `struc:"[4]byte"`
}

func (*NatDetForwardReply) GetMessageName() string {
	return "nat_det_forward_reply"
}
func (*NatDetForwardReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatDetForwardReply) GetCrcString() string {
	return "bf9b96ea"
}
func NewNatDetForwardReply() api.Message {
	return &NatDetForwardReply{}
}

// NatDetReverse represents the VPP binary API message 'nat_det_reverse'.
//
type NatDetReverse struct {
	OutPort uint16
	OutAddr []byte `struc:"[4]byte"`
}

func (*NatDetReverse) GetMessageName() string {
	return "nat_det_reverse"
}
func (*NatDetReverse) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetReverse) GetCrcString() string {
	return "80ab12d2"
}
func NewNatDetReverse() api.Message {
	return &NatDetReverse{}
}

// NatDetReverseReply represents the VPP binary API message 'nat_det_reverse_reply'.
//
type NatDetReverseReply struct {
	Retval  int32
	IsNat44 uint8
	InAddr  []byte `struc:"[16]byte"`
}

func (*NatDetReverseReply) GetMessageName() string {
	return "nat_det_reverse_reply"
}
func (*NatDetReverseReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatDetReverseReply) GetCrcString() string {
	return "26139a2f"
}
func NewNatDetReverseReply() api.Message {
	return &NatDetReverseReply{}
}

// NatDetMapDump represents the VPP binary API message 'nat_det_map_dump'.
//
type NatDetMapDump struct {
}

func (*NatDetMapDump) GetMessageName() string {
	return "nat_det_map_dump"
}
func (*NatDetMapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetMapDump) GetCrcString() string {
	return "51077d14"
}
func NewNatDetMapDump() api.Message {
	return &NatDetMapDump{}
}

// NatDetMapDetails represents the VPP binary API message 'nat_det_map_details'.
//
type NatDetMapDetails struct {
	IsNat44      uint8
	InAddr       []byte `struc:"[16]byte"`
	InPlen       uint8
	OutAddr      []byte `struc:"[4]byte"`
	OutPlen      uint8
	SharingRatio uint32
	PortsPerHost uint16
	SesNum       uint32
}

func (*NatDetMapDetails) GetMessageName() string {
	return "nat_det_map_details"
}
func (*NatDetMapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatDetMapDetails) GetCrcString() string {
	return "886138a8"
}
func NewNatDetMapDetails() api.Message {
	return &NatDetMapDetails{}
}

// NatDetSetTimeouts represents the VPP binary API message 'nat_det_set_timeouts'.
//
type NatDetSetTimeouts struct {
	UDP            uint32
	TCPEstablished uint32
	TCPTransitory  uint32
	ICMP           uint32
}

func (*NatDetSetTimeouts) GetMessageName() string {
	return "nat_det_set_timeouts"
}
func (*NatDetSetTimeouts) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetSetTimeouts) GetCrcString() string {
	return "d4746b16"
}
func NewNatDetSetTimeouts() api.Message {
	return &NatDetSetTimeouts{}
}

// NatDetSetTimeoutsReply represents the VPP binary API message 'nat_det_set_timeouts_reply'.
//
type NatDetSetTimeoutsReply struct {
	Retval int32
}

func (*NatDetSetTimeoutsReply) GetMessageName() string {
	return "nat_det_set_timeouts_reply"
}
func (*NatDetSetTimeoutsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatDetSetTimeoutsReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNatDetSetTimeoutsReply() api.Message {
	return &NatDetSetTimeoutsReply{}
}

// NatDetGetTimeouts represents the VPP binary API message 'nat_det_get_timeouts'.
//
type NatDetGetTimeouts struct {
}

func (*NatDetGetTimeouts) GetMessageName() string {
	return "nat_det_get_timeouts"
}
func (*NatDetGetTimeouts) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetGetTimeouts) GetCrcString() string {
	return "51077d14"
}
func NewNatDetGetTimeouts() api.Message {
	return &NatDetGetTimeouts{}
}

// NatDetGetTimeoutsReply represents the VPP binary API message 'nat_det_get_timeouts_reply'.
//
type NatDetGetTimeoutsReply struct {
	Retval         int32
	UDP            uint32
	TCPEstablished uint32
	TCPTransitory  uint32
	ICMP           uint32
}

func (*NatDetGetTimeoutsReply) GetMessageName() string {
	return "nat_det_get_timeouts_reply"
}
func (*NatDetGetTimeoutsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatDetGetTimeoutsReply) GetCrcString() string {
	return "3c4df4e1"
}
func NewNatDetGetTimeoutsReply() api.Message {
	return &NatDetGetTimeoutsReply{}
}

// NatDetCloseSessionOut represents the VPP binary API message 'nat_det_close_session_out'.
//
type NatDetCloseSessionOut struct {
	OutAddr []byte `struc:"[4]byte"`
	OutPort uint16
	ExtAddr []byte `struc:"[4]byte"`
	ExtPort uint16
}

func (*NatDetCloseSessionOut) GetMessageName() string {
	return "nat_det_close_session_out"
}
func (*NatDetCloseSessionOut) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetCloseSessionOut) GetCrcString() string {
	return "2e165938"
}
func NewNatDetCloseSessionOut() api.Message {
	return &NatDetCloseSessionOut{}
}

// NatDetCloseSessionOutReply represents the VPP binary API message 'nat_det_close_session_out_reply'.
//
type NatDetCloseSessionOutReply struct {
	Retval int32
}

func (*NatDetCloseSessionOutReply) GetMessageName() string {
	return "nat_det_close_session_out_reply"
}
func (*NatDetCloseSessionOutReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatDetCloseSessionOutReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNatDetCloseSessionOutReply() api.Message {
	return &NatDetCloseSessionOutReply{}
}

// NatDetCloseSessionIn represents the VPP binary API message 'nat_det_close_session_in'.
//
type NatDetCloseSessionIn struct {
	IsNat44 uint8
	InAddr  []byte `struc:"[16]byte"`
	InPort  uint16
	ExtAddr []byte `struc:"[16]byte"`
	ExtPort uint16
}

func (*NatDetCloseSessionIn) GetMessageName() string {
	return "nat_det_close_session_in"
}
func (*NatDetCloseSessionIn) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetCloseSessionIn) GetCrcString() string {
	return "147e430c"
}
func NewNatDetCloseSessionIn() api.Message {
	return &NatDetCloseSessionIn{}
}

// NatDetCloseSessionInReply represents the VPP binary API message 'nat_det_close_session_in_reply'.
//
type NatDetCloseSessionInReply struct {
	Retval int32
}

func (*NatDetCloseSessionInReply) GetMessageName() string {
	return "nat_det_close_session_in_reply"
}
func (*NatDetCloseSessionInReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*NatDetCloseSessionInReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNatDetCloseSessionInReply() api.Message {
	return &NatDetCloseSessionInReply{}
}

// NatDetSessionDump represents the VPP binary API message 'nat_det_session_dump'.
//
type NatDetSessionDump struct {
	IsNat44  uint8
	UserAddr []byte `struc:"[16]byte"`
}

func (*NatDetSessionDump) GetMessageName() string {
	return "nat_det_session_dump"
}
func (*NatDetSessionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetSessionDump) GetCrcString() string {
	return "ddfb6b28"
}
func NewNatDetSessionDump() api.Message {
	return &NatDetSessionDump{}
}

// NatDetSessionDetails represents the VPP binary API message 'nat_det_session_details'.
//
type NatDetSessionDetails struct {
	InPort  uint16
	ExtAddr []byte `struc:"[4]byte"`
	ExtPort uint16
	OutPort uint16
	State   uint8
	Expire  uint32
}

func (*NatDetSessionDetails) GetMessageName() string {
	return "nat_det_session_details"
}
func (*NatDetSessionDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*NatDetSessionDetails) GetCrcString() string {
	return "699b5164"
}
func NewNatDetSessionDetails() api.Message {
	return &NatDetSessionDetails{}
}

// Nat64AddDelPoolAddrRange represents the VPP binary API message 'nat64_add_del_pool_addr_range'.
//
type Nat64AddDelPoolAddrRange struct {
	StartAddr []byte `struc:"[4]byte"`
	EndAddr   []byte `struc:"[4]byte"`
	VrfID     uint32
	IsAdd     uint8
}

func (*Nat64AddDelPoolAddrRange) GetMessageName() string {
	return "nat64_add_del_pool_addr_range"
}
func (*Nat64AddDelPoolAddrRange) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64AddDelPoolAddrRange) GetCrcString() string {
	return "5eb50214"
}
func NewNat64AddDelPoolAddrRange() api.Message {
	return &Nat64AddDelPoolAddrRange{}
}

// Nat64AddDelPoolAddrRangeReply represents the VPP binary API message 'nat64_add_del_pool_addr_range_reply'.
//
type Nat64AddDelPoolAddrRangeReply struct {
	Retval int32
}

func (*Nat64AddDelPoolAddrRangeReply) GetMessageName() string {
	return "nat64_add_del_pool_addr_range_reply"
}
func (*Nat64AddDelPoolAddrRangeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64AddDelPoolAddrRangeReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat64AddDelPoolAddrRangeReply() api.Message {
	return &Nat64AddDelPoolAddrRangeReply{}
}

// Nat64PoolAddrDump represents the VPP binary API message 'nat64_pool_addr_dump'.
//
type Nat64PoolAddrDump struct {
}

func (*Nat64PoolAddrDump) GetMessageName() string {
	return "nat64_pool_addr_dump"
}
func (*Nat64PoolAddrDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64PoolAddrDump) GetCrcString() string {
	return "51077d14"
}
func NewNat64PoolAddrDump() api.Message {
	return &Nat64PoolAddrDump{}
}

// Nat64PoolAddrDetails represents the VPP binary API message 'nat64_pool_addr_details'.
//
type Nat64PoolAddrDetails struct {
	Address []byte `struc:"[4]byte"`
	VrfID   uint32
}

func (*Nat64PoolAddrDetails) GetMessageName() string {
	return "nat64_pool_addr_details"
}
func (*Nat64PoolAddrDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64PoolAddrDetails) GetCrcString() string {
	return "8d10231c"
}
func NewNat64PoolAddrDetails() api.Message {
	return &Nat64PoolAddrDetails{}
}

// Nat64AddDelInterface represents the VPP binary API message 'nat64_add_del_interface'.
//
type Nat64AddDelInterface struct {
	SwIfIndex uint32
	IsInside  uint8
	IsAdd     uint8
}

func (*Nat64AddDelInterface) GetMessageName() string {
	return "nat64_add_del_interface"
}
func (*Nat64AddDelInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64AddDelInterface) GetCrcString() string {
	return "efbda9ce"
}
func NewNat64AddDelInterface() api.Message {
	return &Nat64AddDelInterface{}
}

// Nat64AddDelInterfaceReply represents the VPP binary API message 'nat64_add_del_interface_reply'.
//
type Nat64AddDelInterfaceReply struct {
	Retval int32
}

func (*Nat64AddDelInterfaceReply) GetMessageName() string {
	return "nat64_add_del_interface_reply"
}
func (*Nat64AddDelInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64AddDelInterfaceReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat64AddDelInterfaceReply() api.Message {
	return &Nat64AddDelInterfaceReply{}
}

// Nat64InterfaceDump represents the VPP binary API message 'nat64_interface_dump'.
//
type Nat64InterfaceDump struct {
}

func (*Nat64InterfaceDump) GetMessageName() string {
	return "nat64_interface_dump"
}
func (*Nat64InterfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64InterfaceDump) GetCrcString() string {
	return "51077d14"
}
func NewNat64InterfaceDump() api.Message {
	return &Nat64InterfaceDump{}
}

// Nat64InterfaceDetails represents the VPP binary API message 'nat64_interface_details'.
//
type Nat64InterfaceDetails struct {
	IsInside  uint8
	SwIfIndex uint32
}

func (*Nat64InterfaceDetails) GetMessageName() string {
	return "nat64_interface_details"
}
func (*Nat64InterfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64InterfaceDetails) GetCrcString() string {
	return "2b15e8e4"
}
func NewNat64InterfaceDetails() api.Message {
	return &Nat64InterfaceDetails{}
}

// Nat64AddDelStaticBib represents the VPP binary API message 'nat64_add_del_static_bib'.
//
type Nat64AddDelStaticBib struct {
	IAddr []byte `struc:"[16]byte"`
	OAddr []byte `struc:"[4]byte"`
	IPort uint16
	OPort uint16
	VrfID uint32
	Proto uint8
	IsAdd uint8
}

func (*Nat64AddDelStaticBib) GetMessageName() string {
	return "nat64_add_del_static_bib"
}
func (*Nat64AddDelStaticBib) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64AddDelStaticBib) GetCrcString() string {
	return "e36c7813"
}
func NewNat64AddDelStaticBib() api.Message {
	return &Nat64AddDelStaticBib{}
}

// Nat64AddDelStaticBibReply represents the VPP binary API message 'nat64_add_del_static_bib_reply'.
//
type Nat64AddDelStaticBibReply struct {
	Retval int32
}

func (*Nat64AddDelStaticBibReply) GetMessageName() string {
	return "nat64_add_del_static_bib_reply"
}
func (*Nat64AddDelStaticBibReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64AddDelStaticBibReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat64AddDelStaticBibReply() api.Message {
	return &Nat64AddDelStaticBibReply{}
}

// Nat64BibDump represents the VPP binary API message 'nat64_bib_dump'.
//
type Nat64BibDump struct {
	Proto uint8
}

func (*Nat64BibDump) GetMessageName() string {
	return "nat64_bib_dump"
}
func (*Nat64BibDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64BibDump) GetCrcString() string {
	return "cfcb6b75"
}
func NewNat64BibDump() api.Message {
	return &Nat64BibDump{}
}

// Nat64BibDetails represents the VPP binary API message 'nat64_bib_details'.
//
type Nat64BibDetails struct {
	IAddr    []byte `struc:"[16]byte"`
	OAddr    []byte `struc:"[4]byte"`
	IPort    uint16
	OPort    uint16
	VrfID    uint32
	Proto    uint8
	IsStatic uint8
	SesNum   uint32
}

func (*Nat64BibDetails) GetMessageName() string {
	return "nat64_bib_details"
}
func (*Nat64BibDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64BibDetails) GetCrcString() string {
	return "372e7a98"
}
func NewNat64BibDetails() api.Message {
	return &Nat64BibDetails{}
}

// Nat64SetTimeouts represents the VPP binary API message 'nat64_set_timeouts'.
//
type Nat64SetTimeouts struct {
	UDP            uint32
	ICMP           uint32
	TCPTrans       uint32
	TCPEst         uint32
	TCPIncomingSyn uint32
}

func (*Nat64SetTimeouts) GetMessageName() string {
	return "nat64_set_timeouts"
}
func (*Nat64SetTimeouts) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64SetTimeouts) GetCrcString() string {
	return "1cc51cf1"
}
func NewNat64SetTimeouts() api.Message {
	return &Nat64SetTimeouts{}
}

// Nat64SetTimeoutsReply represents the VPP binary API message 'nat64_set_timeouts_reply'.
//
type Nat64SetTimeoutsReply struct {
	Retval int32
}

func (*Nat64SetTimeoutsReply) GetMessageName() string {
	return "nat64_set_timeouts_reply"
}
func (*Nat64SetTimeoutsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64SetTimeoutsReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat64SetTimeoutsReply() api.Message {
	return &Nat64SetTimeoutsReply{}
}

// Nat64GetTimeouts represents the VPP binary API message 'nat64_get_timeouts'.
//
type Nat64GetTimeouts struct {
}

func (*Nat64GetTimeouts) GetMessageName() string {
	return "nat64_get_timeouts"
}
func (*Nat64GetTimeouts) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64GetTimeouts) GetCrcString() string {
	return "51077d14"
}
func NewNat64GetTimeouts() api.Message {
	return &Nat64GetTimeouts{}
}

// Nat64GetTimeoutsReply represents the VPP binary API message 'nat64_get_timeouts_reply'.
//
type Nat64GetTimeoutsReply struct {
	Retval         int32
	UDP            uint32
	ICMP           uint32
	TCPTrans       uint32
	TCPEst         uint32
	TCPIncomingSyn uint32
}

func (*Nat64GetTimeoutsReply) GetMessageName() string {
	return "nat64_get_timeouts_reply"
}
func (*Nat64GetTimeoutsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64GetTimeoutsReply) GetCrcString() string {
	return "cdd081d0"
}
func NewNat64GetTimeoutsReply() api.Message {
	return &Nat64GetTimeoutsReply{}
}

// Nat64StDump represents the VPP binary API message 'nat64_st_dump'.
//
type Nat64StDump struct {
	Proto uint8
}

func (*Nat64StDump) GetMessageName() string {
	return "nat64_st_dump"
}
func (*Nat64StDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64StDump) GetCrcString() string {
	return "cfcb6b75"
}
func NewNat64StDump() api.Message {
	return &Nat64StDump{}
}

// Nat64StDetails represents the VPP binary API message 'nat64_st_details'.
//
type Nat64StDetails struct {
	IlAddr []byte `struc:"[16]byte"`
	OlAddr []byte `struc:"[4]byte"`
	IlPort uint16
	OlPort uint16
	IrAddr []byte `struc:"[16]byte"`
	OrAddr []byte `struc:"[4]byte"`
	RPort  uint16
	VrfID  uint32
	Proto  uint8
}

func (*Nat64StDetails) GetMessageName() string {
	return "nat64_st_details"
}
func (*Nat64StDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64StDetails) GetCrcString() string {
	return "1aaf4631"
}
func NewNat64StDetails() api.Message {
	return &Nat64StDetails{}
}

// Nat64AddDelPrefix represents the VPP binary API message 'nat64_add_del_prefix'.
//
type Nat64AddDelPrefix struct {
	Prefix    []byte `struc:"[16]byte"`
	PrefixLen uint8
	VrfID     uint32
	IsAdd     uint8
}

func (*Nat64AddDelPrefix) GetMessageName() string {
	return "nat64_add_del_prefix"
}
func (*Nat64AddDelPrefix) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64AddDelPrefix) GetCrcString() string {
	return "f4ae7173"
}
func NewNat64AddDelPrefix() api.Message {
	return &Nat64AddDelPrefix{}
}

// Nat64AddDelPrefixReply represents the VPP binary API message 'nat64_add_del_prefix_reply'.
//
type Nat64AddDelPrefixReply struct {
	Retval int32
}

func (*Nat64AddDelPrefixReply) GetMessageName() string {
	return "nat64_add_del_prefix_reply"
}
func (*Nat64AddDelPrefixReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64AddDelPrefixReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat64AddDelPrefixReply() api.Message {
	return &Nat64AddDelPrefixReply{}
}

// Nat64PrefixDump represents the VPP binary API message 'nat64_prefix_dump'.
//
type Nat64PrefixDump struct {
}

func (*Nat64PrefixDump) GetMessageName() string {
	return "nat64_prefix_dump"
}
func (*Nat64PrefixDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64PrefixDump) GetCrcString() string {
	return "51077d14"
}
func NewNat64PrefixDump() api.Message {
	return &Nat64PrefixDump{}
}

// Nat64PrefixDetails represents the VPP binary API message 'nat64_prefix_details'.
//
type Nat64PrefixDetails struct {
	Prefix    []byte `struc:"[16]byte"`
	PrefixLen uint8
	VrfID     uint32
}

func (*Nat64PrefixDetails) GetMessageName() string {
	return "nat64_prefix_details"
}
func (*Nat64PrefixDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64PrefixDetails) GetCrcString() string {
	return "fb08875c"
}
func NewNat64PrefixDetails() api.Message {
	return &Nat64PrefixDetails{}
}

// Nat64AddDelInterfaceAddr represents the VPP binary API message 'nat64_add_del_interface_addr'.
//
type Nat64AddDelInterfaceAddr struct {
	IsAdd     uint8
	IsInside  uint8
	SwIfIndex uint32
}

func (*Nat64AddDelInterfaceAddr) GetMessageName() string {
	return "nat64_add_del_interface_addr"
}
func (*Nat64AddDelInterfaceAddr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Nat64AddDelInterfaceAddr) GetCrcString() string {
	return "9b1ac600"
}
func NewNat64AddDelInterfaceAddr() api.Message {
	return &Nat64AddDelInterfaceAddr{}
}

// Nat64AddDelInterfaceAddrReply represents the VPP binary API message 'nat64_add_del_interface_addr_reply'.
//
type Nat64AddDelInterfaceAddrReply struct {
	Retval int32
}

func (*Nat64AddDelInterfaceAddrReply) GetMessageName() string {
	return "nat64_add_del_interface_addr_reply"
}
func (*Nat64AddDelInterfaceAddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*Nat64AddDelInterfaceAddrReply) GetCrcString() string {
	return "e8d4e804"
}
func NewNat64AddDelInterfaceAddrReply() api.Message {
	return &Nat64AddDelInterfaceAddrReply{}
}

// DsliteAddDelPoolAddrRange represents the VPP binary API message 'dslite_add_del_pool_addr_range'.
//
type DsliteAddDelPoolAddrRange struct {
	StartAddr []byte `struc:"[4]byte"`
	EndAddr   []byte `struc:"[4]byte"`
	IsAdd     uint8
}

func (*DsliteAddDelPoolAddrRange) GetMessageName() string {
	return "dslite_add_del_pool_addr_range"
}
func (*DsliteAddDelPoolAddrRange) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DsliteAddDelPoolAddrRange) GetCrcString() string {
	return "258bff2a"
}
func NewDsliteAddDelPoolAddrRange() api.Message {
	return &DsliteAddDelPoolAddrRange{}
}

// DsliteAddDelPoolAddrRangeReply represents the VPP binary API message 'dslite_add_del_pool_addr_range_reply'.
//
type DsliteAddDelPoolAddrRangeReply struct {
	Retval int32
}

func (*DsliteAddDelPoolAddrRangeReply) GetMessageName() string {
	return "dslite_add_del_pool_addr_range_reply"
}
func (*DsliteAddDelPoolAddrRangeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DsliteAddDelPoolAddrRangeReply) GetCrcString() string {
	return "e8d4e804"
}
func NewDsliteAddDelPoolAddrRangeReply() api.Message {
	return &DsliteAddDelPoolAddrRangeReply{}
}

// DsliteSetAftrAddr represents the VPP binary API message 'dslite_set_aftr_addr'.
//
type DsliteSetAftrAddr struct {
	IP4Addr []byte `struc:"[4]byte"`
	IP6Addr []byte `struc:"[16]byte"`
}

func (*DsliteSetAftrAddr) GetMessageName() string {
	return "dslite_set_aftr_addr"
}
func (*DsliteSetAftrAddr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DsliteSetAftrAddr) GetCrcString() string {
	return "2e9c01ef"
}
func NewDsliteSetAftrAddr() api.Message {
	return &DsliteSetAftrAddr{}
}

// DsliteSetAftrAddrReply represents the VPP binary API message 'dslite_set_aftr_addr_reply'.
//
type DsliteSetAftrAddrReply struct {
	Retval int32
}

func (*DsliteSetAftrAddrReply) GetMessageName() string {
	return "dslite_set_aftr_addr_reply"
}
func (*DsliteSetAftrAddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DsliteSetAftrAddrReply) GetCrcString() string {
	return "e8d4e804"
}
func NewDsliteSetAftrAddrReply() api.Message {
	return &DsliteSetAftrAddrReply{}
}

// DsliteGetAftrAddr represents the VPP binary API message 'dslite_get_aftr_addr'.
//
type DsliteGetAftrAddr struct {
}

func (*DsliteGetAftrAddr) GetMessageName() string {
	return "dslite_get_aftr_addr"
}
func (*DsliteGetAftrAddr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DsliteGetAftrAddr) GetCrcString() string {
	return "51077d14"
}
func NewDsliteGetAftrAddr() api.Message {
	return &DsliteGetAftrAddr{}
}

// DsliteGetAftrAddrReply represents the VPP binary API message 'dslite_get_aftr_addr_reply'.
//
type DsliteGetAftrAddrReply struct {
	Retval  int32
	IP4Addr []byte `struc:"[4]byte"`
	IP6Addr []byte `struc:"[16]byte"`
}

func (*DsliteGetAftrAddrReply) GetMessageName() string {
	return "dslite_get_aftr_addr_reply"
}
func (*DsliteGetAftrAddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DsliteGetAftrAddrReply) GetCrcString() string {
	return "2c4c3037"
}
func NewDsliteGetAftrAddrReply() api.Message {
	return &DsliteGetAftrAddrReply{}
}

// DsliteSetB4Addr represents the VPP binary API message 'dslite_set_b4_addr'.
//
type DsliteSetB4Addr struct {
	IP4Addr []byte `struc:"[4]byte"`
	IP6Addr []byte `struc:"[16]byte"`
}

func (*DsliteSetB4Addr) GetMessageName() string {
	return "dslite_set_b4_addr"
}
func (*DsliteSetB4Addr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DsliteSetB4Addr) GetCrcString() string {
	return "2e9c01ef"
}
func NewDsliteSetB4Addr() api.Message {
	return &DsliteSetB4Addr{}
}

// DsliteSetB4AddrReply represents the VPP binary API message 'dslite_set_b4_addr_reply'.
//
type DsliteSetB4AddrReply struct {
	Retval int32
}

func (*DsliteSetB4AddrReply) GetMessageName() string {
	return "dslite_set_b4_addr_reply"
}
func (*DsliteSetB4AddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DsliteSetB4AddrReply) GetCrcString() string {
	return "e8d4e804"
}
func NewDsliteSetB4AddrReply() api.Message {
	return &DsliteSetB4AddrReply{}
}

// DsliteGetB4Addr represents the VPP binary API message 'dslite_get_b4_addr'.
//
type DsliteGetB4Addr struct {
}

func (*DsliteGetB4Addr) GetMessageName() string {
	return "dslite_get_b4_addr"
}
func (*DsliteGetB4Addr) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DsliteGetB4Addr) GetCrcString() string {
	return "51077d14"
}
func NewDsliteGetB4Addr() api.Message {
	return &DsliteGetB4Addr{}
}

// DsliteGetB4AddrReply represents the VPP binary API message 'dslite_get_b4_addr_reply'.
//
type DsliteGetB4AddrReply struct {
	Retval  int32
	IP4Addr []byte `struc:"[4]byte"`
	IP6Addr []byte `struc:"[16]byte"`
}

func (*DsliteGetB4AddrReply) GetMessageName() string {
	return "dslite_get_b4_addr_reply"
}
func (*DsliteGetB4AddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DsliteGetB4AddrReply) GetCrcString() string {
	return "2c4c3037"
}
func NewDsliteGetB4AddrReply() api.Message {
	return &DsliteGetB4AddrReply{}
}
