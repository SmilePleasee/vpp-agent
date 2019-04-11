package utils

import (
	"fmt"
	"sort"

	"github.com/gogo/protobuf/proto"

	"github.com/ligato/cn-infra/db/keyval"
	acl "github.com/ligato/vpp-agent/api/models/vpp/acl"
	interfaces "github.com/ligato/vpp-agent/api/models/vpp/interfaces"
	ipsec "github.com/ligato/vpp-agent/api/models/vpp/ipsec"
	l2 "github.com/ligato/vpp-agent/api/models/vpp/l2"
	l3 "github.com/ligato/vpp-agent/api/models/vpp/l3"
	nat "github.com/ligato/vpp-agent/api/models/vpp/nat"

	linterface "github.com/ligato/vpp-agent/api/models/linux/interfaces"
	ll3 "github.com/ligato/vpp-agent/api/models/linux/l3"

	"errors"
)

// VppMetaData defines the etcd metadata.
type VppMetaData struct {
	Rev int64
	Key string
}

// ACLConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type ACLConfigWithMD struct {
	Metadata VppMetaData
	ACL      *acl.ACL
}

// InterfaceWithMD contains a data record for interface and its
// etcd metadata.
type ACLWithMD struct {
	Config *ACLConfigWithMD
}

// IfConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type IfConfigWithMD struct {
	Metadata  VppMetaData
	Interface *interfaces.Interface
}

// InterfaceWithMD contains a data record for interface and its
// etcd metadata.
type InterfaceWithMD struct {
	Config *IfConfigWithMD
	//	State  *IfStateWithMD
}

// BdConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type BdConfigWithMD struct {
	Metadata     VppMetaData
	BridgeDomain *l2.BridgeDomain
}

// BdWithMD contains a data record for interface and its
// etcd metadata.
type BdWithMD struct {
	Config *BdConfigWithMD
}

// FibTableConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type FibTableConfigWithMD struct {
	Metadata VppMetaData
	FIBEntry *l2.FIBEntry
}

// FibTableWithMD contains a data record for interface and its
// etcd metadata.
type FibTableWithMD struct {
	Config *FibTableConfigWithMD
}

// XconnectConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type XconnectConfigWithMD struct {
	Metadata VppMetaData
	Xconnect *l2.XConnectPair
}

// XconnectWithMD contains a data record for interface and its
// etcd metadata.
type XconnectWithMD struct {
	Config *XconnectConfigWithMD
}

// ARPConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type ARPConfigWithMD struct {
	Metadata VppMetaData
	ARPEntry *l3.ARPEntry
}

// ARPWithMD contains a data record for interface and its
// etcd metadata.
type ARPWithMD struct {
	Config *ARPConfigWithMD
}

// StaticRouterConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type StaticRoutesConfigWithMD struct {
	Metadata VppMetaData
	Route    *l3.Route
}

// StaticRouterWithMD contains a data record for interface and its
// etcd metadata.
type StaticRoutesWithMD struct {
	Config *StaticRoutesConfigWithMD
}

// ProxyARPConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type ProxyARPConfigWithMD struct {
	Metadata VppMetaData
	ProxyARP *l3.ProxyARP
}

// ProxyARPWithMD contains a data record for interface and its
// etcd metadata.
type ProxyARPWithMD struct {
	Config *ProxyARPConfigWithMD
}

// IPScanNeighConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type IPScanNeighConfigWithMD struct {
	Metadata       VppMetaData
	IPScanNeighbor *l3.IPScanNeighbor
}

// IPScanNeighWithMD contains a data record for interface and its
// etcd metadata.
type IPScanNeighWithMD struct {
	Config *IPScanNeighConfigWithMD
}

// NATConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type NATConfigWithMD struct {
	Metadata    VppMetaData
	Nat44Global *nat.Nat44Global
}

// NATWithMD contains a data record for interface and its
// etcd metadata.
type NATWithMD struct {
	Config *NATConfigWithMD
}

// DNATConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type DNATConfigWithMD struct {
	Metadata VppMetaData
	DNat44   *nat.DNat44
}

// DNATWithMD contains a data record for interface and its
// etcd metadata.
type DNATWithMD struct {
	Config *DNATConfigWithMD
}

// IPSecPolicyConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type IPSecPolicyConfigWithMD struct {
	Metadata               VppMetaData
	SecurityPolicyDatabase *ipsec.SecurityPolicyDatabase
}

// IPSecPolicyWithMD contains a data record for interface and its
// etcd metadata.
type IPSecPolicyWithMD struct {
	Config *IPSecPolicyConfigWithMD
}

// IPSecAssociationConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type IPSecAssosiciationConfigWithMD struct {
	Metadata            VppMetaData
	SecurityAssociation *ipsec.SecurityAssociation
}

// IPSecAssosiciationWithMD contains a data record for interface and its
// etcd metadata.
type IPSecAssosiciationWithMD struct {
	Config *IPSecAssosiciationConfigWithMD
}

// lInterfaceConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type lInterfaceConfigWithMD struct {
	Metadata  VppMetaData
	Interface *linterface.Interface
}

// lInterfaceWithMD contains a data record for interface and its
// etcd metadata.
type lInterfaceWithMD struct {
	Config *lInterfaceConfigWithMD
}

// lARPConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type lARPConfigWithMD struct {
	Metadata VppMetaData
	ARPEntry *ll3.ARPEntry
}

// lARPWithMD contains a data record for interface and its
// etcd metadata.
type lARPWithMD struct {
	Config *lARPConfigWithMD
}

// lRouteConfigWithMD contains a data record for interface configuration
// and its etcd metadata.
type lRouteConfigWithMD struct {
	Metadata VppMetaData
	Route    *ll3.Route
}

// lRouteWithMD contains a data record for interface and its
// etcd metadata.
type lRouteWithMD struct {
	Config *lRouteConfigWithMD
}

// VppData defines a structure to hold all etcd data records (of all
// types) for one VPP.
type VppData struct {
	ACL        map[string]ACLWithMD
	Interfaces map[string]InterfaceWithMD
	//	InterfaceErrors    map[string]InterfaceErrorWithMD
	BridgeDomains map[string]BdWithMD
	//	BridgeDomainErrors map[string]BridgeDomainErrorWithMD
	FibTableEntries FibTableWithMD
	XConnectPairs   map[string]XconnectWithMD
	ARP             ARPWithMD
	StaticRoutes    StaticRoutesWithMD
	ProxyARP        ProxyARPWithMD
	IPScanNeight    IPScanNeighWithMD
	NAT             NATWithMD
	DNAT            map[string]DNATWithMD
	IPSecPolicyDb   map[string]IPSecPolicyWithMD
	IPSecAssociate  map[string]IPSecAssosiciationWithMD
	//	Status             map[string]VppStatusWithMD
	LInterfaces map[string]lInterfaceWithMD
	LARP        map[string]lARPWithMD
	LRoute      map[string]lRouteWithMD
	ShowEtcd    bool
	ShowConf    bool
}

// EtcdDump is a map of VppData records. It constitutes a temporary
// storage for data retrieved from etcd. "Temporary" means during
// the execution of an agentctl command. Every command reads
// data from etcd first, then processes it, and finally either outputs
// the processed data to the user or updates one or more data records
// in etcd.
type EtcdDump map[string]*VppData

// NewEtcdDump returns a new instance of the temporary storage
// that will hold data retrieved from etcd.
func NewEtcdDump() EtcdDump {
	return make(EtcdDump)
}

// ReadDataFromDb reads a data record from etcd, parses it according to
// the expected record type and stores it in the EtcdDump temporary
// storage. A record is identified by a Key.
//
// The function returns an error if the etcd client encountered an
// error. The function returns true if the specified item has been
// found.
func (ed EtcdDump) ReadDataFromDb(db keyval.ProtoBroker, key string) (found bool, err error) {
	label, dataType, params := ParseKey(key)

	vd, ok := ed[label]
	if !ok {
		vd = newVppDataRecord()
	}

	switch dataType {
	case acl.ModelACL.KeyPrefix():
		ed[label], err = readACLConfigFromDb(db, vd, key, params)
	case interfaces.ModelInterface.KeyPrefix():
		ed[label], err = readInterfaceConfigFromDb(db, vd, key, params)
	case l2.ModelBridgeDomain.KeyPrefix():
		ed[label], err = readBridgeConfigFromDb(db, vd, key, params)
	case l2.ModelFIBEntry.KeyPrefix():
		ed[label], err = readFibTableConfigFromDb(db, vd, key, params)
	case l2.ModelXConnectPair.KeyPrefix():
		ed[label], err = readXConnectConfigFromDb(db, vd, key, params)
	case l3.ModelARPEntry.KeyPrefix():
		ed[label], err = readARPConfigFromDb(db, vd, key, params)
	case l3.ModelRoute.KeyPrefix():
		ed[label], err = readStatiRouteConfigFromDb(db, vd, key, params)
	case l3.ModelProxyARP.KeyPrefix():
		ed[label], err = readProxyARPConfigFromDb(db, vd, key)
	case l3.ModelIPScanNeighbor.KeyPrefix():
		ed[label], err = readIPScanNeightConfigFromDb(db, vd, key)
		//FIXME: Error in key
	//case NATPath:
	//	ed[label], err = readNATConfigFromDb(db, vd, key)
	//case DNATPath:
	//	ed[label], err = readDNATConfigFromDb(db, vd, key, params)
	case ipsec.ModelSecurityPolicyDatabase.KeyPrefix():
		ed[label], err = readIPSecPolicyConfigFromDb(db, vd, key, params)
	case ipsec.ModelSecurityAssociation.KeyPrefix():
		ed[label], err = readIPSecAssociateConfigFromDb(db, vd, key, params)
	case linterface.ModelInterface.KeyPrefix():
		ed[label], err = readLinuxInterfaceConfigFromDb(db, vd, key, params)
	case ll3.ModelARPEntry.KeyPrefix():
		ed[label], err = readLinuxARPConfigFromDb(db, vd, key, params)
	case ll3.ModelRoute.KeyPrefix():
		ed[label], err = readLinuxRouteConfigFromDb(db, vd, key, params)
	}

	return true, err
}

func readACLConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid ACL config Key '%s'\n", key)
		return vd, nil
	}

	acl := &acl.ACL{}

	found, rev, err := readDataFromDb(db, key, acl)
	if found && err == nil {
		vd.ACL[name] = ACLWithMD{
			Config: &ACLConfigWithMD{VppMetaData{rev, key}, acl},
		}
	}
	return vd, err
}

func readInterfaceConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid interface config Key '%s'\n", key)
		return vd, nil
	}

	int := &interfaces.Interface{}

	found, rev, err := readDataFromDb(db, key, int)
	if found && err == nil {
		vd.Interfaces[name] = InterfaceWithMD{
			Config: &IfConfigWithMD{VppMetaData{rev, key}, int},
		}
	}
	return vd, err
}

func readBridgeConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid Bridge domain config Key '%s'\n", key)
		return vd, nil
	}

	br := &l2.BridgeDomain{}

	found, rev, err := readDataFromDb(db, key, br)
	if found && err == nil {
		vd.BridgeDomains[name] = BdWithMD{
			Config: &BdConfigWithMD{VppMetaData{rev, key}, br},
		}
	}
	return vd, err
}

func readFibTableConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid Fib table config Key '%s'\n", key)
		return vd, nil
	}

	fib := &l2.FIBEntry{}

	found, rev, err := readDataFromDb(db, key, fib)
	if found && err == nil {
		vd.FibTableEntries = FibTableWithMD{
			Config: &FibTableConfigWithMD{VppMetaData{rev, key}, fib},
		}
	}
	return vd, err
}

func readXConnectConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid xconnect config Key '%s'\n", key)
		return vd, nil
	}

	xconnect := &l2.XConnectPair{}

	found, rev, err := readDataFromDb(db, key, xconnect)
	if found && err == nil {
		vd.XConnectPairs[name] = XconnectWithMD{
			Config: &XconnectConfigWithMD{VppMetaData{rev, key}, xconnect},
		}
	}
	return vd, err
}

func readARPConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid arp config Key '%s'\n", key)
		return vd, nil
	}

	arp := &l3.ARPEntry{}

	found, rev, err := readDataFromDb(db, key, arp)
	if found && err == nil {
		vd.ARP = ARPWithMD{
			Config: &ARPConfigWithMD{VppMetaData{rev, key}, arp},
		}
	}
	return vd, err
}

func readStatiRouteConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid static route config Key '%s'\n", key)
		return vd, nil
	}

	route := &l3.Route{}

	found, rev, err := readDataFromDb(db, key, route)
	if found && err == nil {
		vd.StaticRoutes = StaticRoutesWithMD{
			Config: &StaticRoutesConfigWithMD{VppMetaData{rev, key}, route},
		}
	}
	return vd, err
}

func readProxyARPConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string) (*VppData, error) {
	parp := &l3.ProxyARP{}

	found, rev, err := readDataFromDb(db, key, parp)
	if found && err == nil {
		vd.ProxyARP = ProxyARPWithMD{
			Config: &ProxyARPConfigWithMD{VppMetaData{rev, key}, parp},
		}
	}
	return vd, err
}

func readIPScanNeightConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string) (*VppData, error) {
	scan := &l3.IPScanNeighbor{}

	found, rev, err := readDataFromDb(db, key, scan)
	if found && err == nil {
		vd.IPScanNeight = IPScanNeighWithMD{
			Config: &IPScanNeighConfigWithMD{VppMetaData{rev, key}, scan},
		}
	}
	return vd, err
}

func readNATConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string) (*VppData, error) {
	nat := &nat.Nat44Global{}

	found, rev, err := readDataFromDb(db, key, nat)
	if found && err == nil {
		vd.NAT = NATWithMD{
			Config: &NATConfigWithMD{VppMetaData{rev, key}, nat},
		}
	}
	return vd, err
}

func readDNATConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid dnat config Key '%s'\n", key)
		return vd, nil
	}

	dnat := &nat.DNat44{}

	found, rev, err := readDataFromDb(db, key, dnat)
	if found && err == nil {
		vd.DNAT[name] = DNATWithMD{
			Config: &DNATConfigWithMD{VppMetaData{rev, key}, dnat},
		}
	}
	return vd, err
}

func readIPSecPolicyConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid ip sec policy database config Key '%s'\n", key)
		return vd, nil
	}

	policy := &ipsec.SecurityPolicyDatabase{}

	found, rev, err := readDataFromDb(db, key, policy)
	if found && err == nil {
		vd.IPSecPolicyDb[name] = IPSecPolicyWithMD{
			Config: &IPSecPolicyConfigWithMD{VppMetaData{rev, key}, policy},
		}
	}
	return vd, err
}

func readIPSecAssociateConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid ip sec associate config Key '%s'\n", key)
		return vd, nil
	}

	ipsec := &ipsec.SecurityAssociation{}

	found, rev, err := readDataFromDb(db, key, ipsec)
	if found && err == nil {
		vd.IPSecAssociate[name] = IPSecAssosiciationWithMD{
			Config: &IPSecAssosiciationConfigWithMD{VppMetaData{rev, key}, ipsec},
		}
	}
	return vd, err
}

func readLinuxInterfaceConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid linux interface config Key '%s'\n", key)
		return vd, nil
	}

	int := &linterface.Interface{}

	found, rev, err := readDataFromDb(db, key, int)
	if found && err == nil {
		vd.LInterfaces[name] = lInterfaceWithMD{
			Config: &lInterfaceConfigWithMD{VppMetaData{rev, key}, int},
		}
	}
	return vd, err
}

func readLinuxARPConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid linux arp config Key '%s'\n", key)
		return vd, nil
	}

	arp := &ll3.ARPEntry{}

	found, rev, err := readDataFromDb(db, key, arp)
	if found && err == nil {
		vd.LARP[name] = lARPWithMD{
			Config: &lARPConfigWithMD{VppMetaData{rev, key}, arp},
		}
	}
	return vd, err
}

func readLinuxRouteConfigFromDb(db keyval.ProtoBroker, vd *VppData, key string, name string) (*VppData, error) {
	if name == "" {
		fmt.Printf("WARNING: Invalid linux route config Key '%s'\n", key)
		return vd, nil
	}

	route := &ll3.Route{}

	found, rev, err := readDataFromDb(db, key, route)
	if found && err == nil {
		vd.LRoute[name] = lRouteWithMD{
			Config: &lRouteConfigWithMD{VppMetaData{rev, key}, route},
		}
	}
	return vd, err
}

func readDataFromDb(db keyval.ProtoBroker, key string, obj proto.Message) (bool, int64, error) {
	found, rev, err := db.GetValue(key, obj)
	if err != nil {
		return false, rev, errors.New("Could not read from database, Key:" + key + ", error" + err.Error())
	}
	if !found {
		fmt.Printf("WARNING: data for Key '%s' not found\n", key)
	}
	return found, rev, nil
}

func newVppDataRecord() *VppData {
	return &VppData{
		ACL:             make(map[string]ACLWithMD),
		Interfaces:      make(map[string]InterfaceWithMD),
		BridgeDomains:   make(map[string]BdWithMD),
		FibTableEntries: FibTableWithMD{},
		XConnectPairs:   make(map[string]XconnectWithMD),
		ARP:             ARPWithMD{},
		StaticRoutes:    StaticRoutesWithMD{},
		ProxyARP:        ProxyARPWithMD{},
		IPScanNeight:    IPScanNeighWithMD{},
		NAT:             NATWithMD{},
		DNAT:            make(map[string]DNATWithMD),
		IPSecPolicyDb:   make(map[string]IPSecPolicyWithMD),
		IPSecAssociate:  make(map[string]IPSecAssosiciationWithMD),
		LInterfaces:     make(map[string]lInterfaceWithMD),
		LARP:            make(map[string]lARPWithMD),
		LRoute:          make(map[string]lRouteWithMD),
		ShowEtcd:        false,
		ShowConf:        false,
	}
}

func (ed EtcdDump) getSortedKeys() []string {
	keys := []string{}
	for k := range ed {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
