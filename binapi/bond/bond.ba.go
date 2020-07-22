// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/core/bond.api.json

// Package bond contains generated bindings for API file bond.api.
//
// Contents:
//   2 enums
//  14 messages
//
package bond

import (
	api "git.fd.io/govpp.git/api"
	ethernet_types "git.fd.io/govpp.git/binapi/ethernet_types"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	codec "git.fd.io/govpp.git/codec"
	"strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "bond"
	APIVersion = "2.0.0"
	VersionCrc = 0xa65a4a1e
)

// BondLbAlgo defines enum 'bond_lb_algo'.
type BondLbAlgo uint32

const (
	BOND_API_LB_ALGO_L2  BondLbAlgo = 0
	BOND_API_LB_ALGO_L34 BondLbAlgo = 1
	BOND_API_LB_ALGO_L23 BondLbAlgo = 2
	BOND_API_LB_ALGO_RR  BondLbAlgo = 3
	BOND_API_LB_ALGO_BC  BondLbAlgo = 4
	BOND_API_LB_ALGO_AB  BondLbAlgo = 5
)

var (
	BondLbAlgo_name = map[uint32]string{
		0: "BOND_API_LB_ALGO_L2",
		1: "BOND_API_LB_ALGO_L34",
		2: "BOND_API_LB_ALGO_L23",
		3: "BOND_API_LB_ALGO_RR",
		4: "BOND_API_LB_ALGO_BC",
		5: "BOND_API_LB_ALGO_AB",
	}
	BondLbAlgo_value = map[string]uint32{
		"BOND_API_LB_ALGO_L2":  0,
		"BOND_API_LB_ALGO_L34": 1,
		"BOND_API_LB_ALGO_L23": 2,
		"BOND_API_LB_ALGO_RR":  3,
		"BOND_API_LB_ALGO_BC":  4,
		"BOND_API_LB_ALGO_AB":  5,
	}
)

func (x BondLbAlgo) String() string {
	s, ok := BondLbAlgo_name[uint32(x)]
	if ok {
		return s
	}
	return "BondLbAlgo(" + strconv.Itoa(int(x)) + ")"
}

// BondMode defines enum 'bond_mode'.
type BondMode uint32

const (
	BOND_API_MODE_ROUND_ROBIN   BondMode = 1
	BOND_API_MODE_ACTIVE_BACKUP BondMode = 2
	BOND_API_MODE_XOR           BondMode = 3
	BOND_API_MODE_BROADCAST     BondMode = 4
	BOND_API_MODE_LACP          BondMode = 5
)

var (
	BondMode_name = map[uint32]string{
		1: "BOND_API_MODE_ROUND_ROBIN",
		2: "BOND_API_MODE_ACTIVE_BACKUP",
		3: "BOND_API_MODE_XOR",
		4: "BOND_API_MODE_BROADCAST",
		5: "BOND_API_MODE_LACP",
	}
	BondMode_value = map[string]uint32{
		"BOND_API_MODE_ROUND_ROBIN":   1,
		"BOND_API_MODE_ACTIVE_BACKUP": 2,
		"BOND_API_MODE_XOR":           3,
		"BOND_API_MODE_BROADCAST":     4,
		"BOND_API_MODE_LACP":          5,
	}
)

func (x BondMode) String() string {
	s, ok := BondMode_name[uint32(x)]
	if ok {
		return s
	}
	return "BondMode(" + strconv.Itoa(int(x)) + ")"
}

// BondCreate defines message 'bond_create'.
type BondCreate struct {
	ID           uint32                    `binapi:"u32,name=id,default=4294967295" json:"id,omitempty"`
	UseCustomMac bool                      `binapi:"bool,name=use_custom_mac" json:"use_custom_mac,omitempty"`
	MacAddress   ethernet_types.MacAddress `binapi:"mac_address,name=mac_address" json:"mac_address,omitempty"`
	Mode         BondMode                  `binapi:"bond_mode,name=mode" json:"mode,omitempty"`
	Lb           BondLbAlgo                `binapi:"bond_lb_algo,name=lb" json:"lb,omitempty"`
	NumaOnly     bool                      `binapi:"bool,name=numa_only" json:"numa_only,omitempty"`
}

func (m *BondCreate) Reset()               { *m = BondCreate{} }
func (*BondCreate) GetMessageName() string { return "bond_create" }
func (*BondCreate) GetCrcString() string   { return "48883c7e" }
func (*BondCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *BondCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.ID
	size += 1     // m.UseCustomMac
	size += 1 * 6 // m.MacAddress
	size += 4     // m.Mode
	size += 4     // m.Lb
	size += 1     // m.NumaOnly
	return size
}
func (m *BondCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ID)
	buf.EncodeBool(m.UseCustomMac)
	buf.EncodeBytes(m.MacAddress[:], 6)
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeUint32(uint32(m.Lb))
	buf.EncodeBool(m.NumaOnly)
	return buf.Bytes(), nil
}
func (m *BondCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint32()
	m.UseCustomMac = buf.DecodeBool()
	copy(m.MacAddress[:], buf.DecodeBytes(6))
	m.Mode = BondMode(buf.DecodeUint32())
	m.Lb = BondLbAlgo(buf.DecodeUint32())
	m.NumaOnly = buf.DecodeBool()
	return nil
}

// BondCreateReply defines message 'bond_create_reply'.
type BondCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *BondCreateReply) Reset()               { *m = BondCreateReply{} }
func (*BondCreateReply) GetMessageName() string { return "bond_create_reply" }
func (*BondCreateReply) GetCrcString() string   { return "5383d31f" }
func (*BondCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *BondCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *BondCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *BondCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// BondDelete defines message 'bond_delete'.
type BondDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *BondDelete) Reset()               { *m = BondDelete{} }
func (*BondDelete) GetMessageName() string { return "bond_delete" }
func (*BondDelete) GetCrcString() string   { return "f9e6675e" }
func (*BondDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *BondDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *BondDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *BondDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// BondDeleteReply defines message 'bond_delete_reply'.
type BondDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *BondDeleteReply) Reset()               { *m = BondDeleteReply{} }
func (*BondDeleteReply) GetMessageName() string { return "bond_delete_reply" }
func (*BondDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*BondDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *BondDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *BondDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *BondDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// BondDetachSlave defines message 'bond_detach_slave'.
type BondDetachSlave struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *BondDetachSlave) Reset()               { *m = BondDetachSlave{} }
func (*BondDetachSlave) GetMessageName() string { return "bond_detach_slave" }
func (*BondDetachSlave) GetCrcString() string   { return "f9e6675e" }
func (*BondDetachSlave) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *BondDetachSlave) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *BondDetachSlave) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *BondDetachSlave) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// BondDetachSlaveReply defines message 'bond_detach_slave_reply'.
type BondDetachSlaveReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *BondDetachSlaveReply) Reset()               { *m = BondDetachSlaveReply{} }
func (*BondDetachSlaveReply) GetMessageName() string { return "bond_detach_slave_reply" }
func (*BondDetachSlaveReply) GetCrcString() string   { return "e8d4e804" }
func (*BondDetachSlaveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *BondDetachSlaveReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *BondDetachSlaveReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *BondDetachSlaveReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// BondEnslave defines message 'bond_enslave'.
type BondEnslave struct {
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	BondSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=bond_sw_if_index" json:"bond_sw_if_index,omitempty"`
	IsPassive     bool                           `binapi:"bool,name=is_passive" json:"is_passive,omitempty"`
	IsLongTimeout bool                           `binapi:"bool,name=is_long_timeout" json:"is_long_timeout,omitempty"`
}

func (m *BondEnslave) Reset()               { *m = BondEnslave{} }
func (*BondEnslave) GetMessageName() string { return "bond_enslave" }
func (*BondEnslave) GetCrcString() string   { return "076ecfa7" }
func (*BondEnslave) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *BondEnslave) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 4 // m.BondSwIfIndex
	size += 1 // m.IsPassive
	size += 1 // m.IsLongTimeout
	return size
}
func (m *BondEnslave) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(uint32(m.BondSwIfIndex))
	buf.EncodeBool(m.IsPassive)
	buf.EncodeBool(m.IsLongTimeout)
	return buf.Bytes(), nil
}
func (m *BondEnslave) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.BondSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsPassive = buf.DecodeBool()
	m.IsLongTimeout = buf.DecodeBool()
	return nil
}

// BondEnslaveReply defines message 'bond_enslave_reply'.
type BondEnslaveReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *BondEnslaveReply) Reset()               { *m = BondEnslaveReply{} }
func (*BondEnslaveReply) GetMessageName() string { return "bond_enslave_reply" }
func (*BondEnslaveReply) GetCrcString() string   { return "e8d4e804" }
func (*BondEnslaveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *BondEnslaveReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *BondEnslaveReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *BondEnslaveReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SwInterfaceBondDetails defines message 'sw_interface_bond_details'.
type SwInterfaceBondDetails struct {
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	ID            uint32                         `binapi:"u32,name=id" json:"id,omitempty"`
	Mode          BondMode                       `binapi:"bond_mode,name=mode" json:"mode,omitempty"`
	Lb            BondLbAlgo                     `binapi:"bond_lb_algo,name=lb" json:"lb,omitempty"`
	NumaOnly      bool                           `binapi:"bool,name=numa_only" json:"numa_only,omitempty"`
	ActiveSlaves  uint32                         `binapi:"u32,name=active_slaves" json:"active_slaves,omitempty"`
	Slaves        uint32                         `binapi:"u32,name=slaves" json:"slaves,omitempty"`
	InterfaceName string                         `binapi:"string[64],name=interface_name" json:"interface_name,omitempty"`
}

func (m *SwInterfaceBondDetails) Reset()               { *m = SwInterfaceBondDetails{} }
func (*SwInterfaceBondDetails) GetMessageName() string { return "sw_interface_bond_details" }
func (*SwInterfaceBondDetails) GetCrcString() string   { return "f5ef2106" }
func (*SwInterfaceBondDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceBondDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.SwIfIndex
	size += 4  // m.ID
	size += 4  // m.Mode
	size += 4  // m.Lb
	size += 1  // m.NumaOnly
	size += 4  // m.ActiveSlaves
	size += 4  // m.Slaves
	size += 64 // m.InterfaceName
	return size
}
func (m *SwInterfaceBondDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(m.ID)
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeUint32(uint32(m.Lb))
	buf.EncodeBool(m.NumaOnly)
	buf.EncodeUint32(m.ActiveSlaves)
	buf.EncodeUint32(m.Slaves)
	buf.EncodeString(m.InterfaceName, 64)
	return buf.Bytes(), nil
}
func (m *SwInterfaceBondDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.ID = buf.DecodeUint32()
	m.Mode = BondMode(buf.DecodeUint32())
	m.Lb = BondLbAlgo(buf.DecodeUint32())
	m.NumaOnly = buf.DecodeBool()
	m.ActiveSlaves = buf.DecodeUint32()
	m.Slaves = buf.DecodeUint32()
	m.InterfaceName = buf.DecodeString(64)
	return nil
}

// SwInterfaceBondDump defines message 'sw_interface_bond_dump'.
type SwInterfaceBondDump struct{}

func (m *SwInterfaceBondDump) Reset()               { *m = SwInterfaceBondDump{} }
func (*SwInterfaceBondDump) GetMessageName() string { return "sw_interface_bond_dump" }
func (*SwInterfaceBondDump) GetCrcString() string   { return "51077d14" }
func (*SwInterfaceBondDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceBondDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SwInterfaceBondDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SwInterfaceBondDump) Unmarshal(b []byte) error {
	return nil
}

// SwInterfaceSetBondWeight defines message 'sw_interface_set_bond_weight'.
type SwInterfaceSetBondWeight struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Weight    uint32                         `binapi:"u32,name=weight" json:"weight,omitempty"`
}

func (m *SwInterfaceSetBondWeight) Reset()               { *m = SwInterfaceSetBondWeight{} }
func (*SwInterfaceSetBondWeight) GetMessageName() string { return "sw_interface_set_bond_weight" }
func (*SwInterfaceSetBondWeight) GetCrcString() string   { return "deb510a0" }
func (*SwInterfaceSetBondWeight) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSetBondWeight) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 4 // m.Weight
	return size
}
func (m *SwInterfaceSetBondWeight) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(m.Weight)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetBondWeight) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Weight = buf.DecodeUint32()
	return nil
}

// SwInterfaceSetBondWeightReply defines message 'sw_interface_set_bond_weight_reply'.
type SwInterfaceSetBondWeightReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceSetBondWeightReply) Reset() { *m = SwInterfaceSetBondWeightReply{} }
func (*SwInterfaceSetBondWeightReply) GetMessageName() string {
	return "sw_interface_set_bond_weight_reply"
}
func (*SwInterfaceSetBondWeightReply) GetCrcString() string { return "e8d4e804" }
func (*SwInterfaceSetBondWeightReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSetBondWeightReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceSetBondWeightReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetBondWeightReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SwInterfaceSlaveDetails defines message 'sw_interface_slave_details'.
type SwInterfaceSlaveDetails struct {
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	InterfaceName string                         `binapi:"string[64],name=interface_name" json:"interface_name,omitempty"`
	IsPassive     bool                           `binapi:"bool,name=is_passive" json:"is_passive,omitempty"`
	IsLongTimeout bool                           `binapi:"bool,name=is_long_timeout" json:"is_long_timeout,omitempty"`
	IsLocalNuma   bool                           `binapi:"bool,name=is_local_numa" json:"is_local_numa,omitempty"`
	Weight        uint32                         `binapi:"u32,name=weight" json:"weight,omitempty"`
}

func (m *SwInterfaceSlaveDetails) Reset()               { *m = SwInterfaceSlaveDetails{} }
func (*SwInterfaceSlaveDetails) GetMessageName() string { return "sw_interface_slave_details" }
func (*SwInterfaceSlaveDetails) GetCrcString() string   { return "3c4a0e23" }
func (*SwInterfaceSlaveDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSlaveDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.SwIfIndex
	size += 64 // m.InterfaceName
	size += 1  // m.IsPassive
	size += 1  // m.IsLongTimeout
	size += 1  // m.IsLocalNuma
	size += 4  // m.Weight
	return size
}
func (m *SwInterfaceSlaveDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.InterfaceName, 64)
	buf.EncodeBool(m.IsPassive)
	buf.EncodeBool(m.IsLongTimeout)
	buf.EncodeBool(m.IsLocalNuma)
	buf.EncodeUint32(m.Weight)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSlaveDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.InterfaceName = buf.DecodeString(64)
	m.IsPassive = buf.DecodeBool()
	m.IsLongTimeout = buf.DecodeBool()
	m.IsLocalNuma = buf.DecodeBool()
	m.Weight = buf.DecodeUint32()
	return nil
}

// SwInterfaceSlaveDump defines message 'sw_interface_slave_dump'.
type SwInterfaceSlaveDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *SwInterfaceSlaveDump) Reset()               { *m = SwInterfaceSlaveDump{} }
func (*SwInterfaceSlaveDump) GetMessageName() string { return "sw_interface_slave_dump" }
func (*SwInterfaceSlaveDump) GetCrcString() string   { return "f9e6675e" }
func (*SwInterfaceSlaveDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSlaveDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *SwInterfaceSlaveDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *SwInterfaceSlaveDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

func init() { file_bond_binapi_init() }
func file_bond_binapi_init() {
	api.RegisterMessage((*BondCreate)(nil), "bond_create_48883c7e")
	api.RegisterMessage((*BondCreateReply)(nil), "bond_create_reply_5383d31f")
	api.RegisterMessage((*BondDelete)(nil), "bond_delete_f9e6675e")
	api.RegisterMessage((*BondDeleteReply)(nil), "bond_delete_reply_e8d4e804")
	api.RegisterMessage((*BondDetachSlave)(nil), "bond_detach_slave_f9e6675e")
	api.RegisterMessage((*BondDetachSlaveReply)(nil), "bond_detach_slave_reply_e8d4e804")
	api.RegisterMessage((*BondEnslave)(nil), "bond_enslave_076ecfa7")
	api.RegisterMessage((*BondEnslaveReply)(nil), "bond_enslave_reply_e8d4e804")
	api.RegisterMessage((*SwInterfaceBondDetails)(nil), "sw_interface_bond_details_f5ef2106")
	api.RegisterMessage((*SwInterfaceBondDump)(nil), "sw_interface_bond_dump_51077d14")
	api.RegisterMessage((*SwInterfaceSetBondWeight)(nil), "sw_interface_set_bond_weight_deb510a0")
	api.RegisterMessage((*SwInterfaceSetBondWeightReply)(nil), "sw_interface_set_bond_weight_reply_e8d4e804")
	api.RegisterMessage((*SwInterfaceSlaveDetails)(nil), "sw_interface_slave_details_3c4a0e23")
	api.RegisterMessage((*SwInterfaceSlaveDump)(nil), "sw_interface_slave_dump_f9e6675e")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*BondCreate)(nil),
		(*BondCreateReply)(nil),
		(*BondDelete)(nil),
		(*BondDeleteReply)(nil),
		(*BondDetachSlave)(nil),
		(*BondDetachSlaveReply)(nil),
		(*BondEnslave)(nil),
		(*BondEnslaveReply)(nil),
		(*SwInterfaceBondDetails)(nil),
		(*SwInterfaceBondDump)(nil),
		(*SwInterfaceSetBondWeight)(nil),
		(*SwInterfaceSetBondWeightReply)(nil),
		(*SwInterfaceSlaveDetails)(nil),
		(*SwInterfaceSlaveDump)(nil),
	}
}
