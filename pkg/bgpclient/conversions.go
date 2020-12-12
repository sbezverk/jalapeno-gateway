package bgpclient

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	api "github.com/osrg/gobgp/api"
	"github.com/osrg/gobgp/pkg/packet/bgp"
	"github.com/sbezverk/gobmp/pkg/prefixsid"
	"github.com/sbezverk/gobmp/pkg/srv6"
)

// UnmarshalRD unmarshals Route Distinguisher from Proto message
func UnmarshalRD(rd *any.Any) (bgp.RouteDistinguisherInterface, error) {
	var rdValue ptypes.DynamicAny
	if err := ptypes.UnmarshalAny(rd, &rdValue); err != nil {
		return nil, fmt.Errorf("failed to unmarshal route distinguisher with error: %+v", err)
	}

	switch v := rdValue.Message.(type) {
	case *api.RouteDistinguisherTwoOctetAS:
		return bgp.NewRouteDistinguisherTwoOctetAS(uint16(v.Admin), v.Assigned), nil
	case *api.RouteDistinguisherIPAddress:
		return bgp.NewRouteDistinguisherIPAddressAS(v.Admin, uint16(v.Assigned)), nil
	case *api.RouteDistinguisherFourOctetAS:
		return bgp.NewRouteDistinguisherFourOctetAS(v.Admin, uint16(v.Assigned)), nil
	default:
		return nil, fmt.Errorf("Unknown route distinguisher type: %+v", v)
	}
}

// UnmarshalRT unmarshals Route Target extended community from Proto message
func UnmarshalRT(rts []*any.Any) ([]bgp.ExtendedCommunityInterface, error) {
	repl := make([]bgp.ExtendedCommunityInterface, 0)
	for i := 0; i < len(rts); i++ {
		var rtValue ptypes.DynamicAny
		if err := ptypes.UnmarshalAny(rts[i], &rtValue); err != nil {
			return nil, fmt.Errorf("failed to unmarshal route target with error: %+v", err)
		}
		switch v := rtValue.Message.(type) {
		case *api.TwoOctetAsSpecificExtended:
			repl = append(repl, bgp.NewTwoOctetAsSpecificExtended(bgp.ExtendedCommunityAttrSubType(v.SubType), uint16(v.As), v.LocalAdmin, v.IsTransitive))
		case *api.IPv4AddressSpecificExtended:
			repl = append(repl, bgp.NewIPv4AddressSpecificExtended(bgp.ExtendedCommunityAttrSubType(v.SubType), v.Address, uint16(v.LocalAdmin), v.IsTransitive))
		case *api.FourOctetAsSpecificExtended:
			repl = append(repl, bgp.NewFourOctetAsSpecificExtended(bgp.ExtendedCommunityAttrSubType(v.SubType), v.As, uint16(v.LocalAdmin), v.IsTransitive))
		default:
			return nil, fmt.Errorf("Unknown route target type: %+v", v)
		}
	}

	return repl, nil
}

// MarshalRD marshals Route Distinguisher into Proto Any format
func MarshalRD(rd bgp.RouteDistinguisherInterface) *any.Any {
	var r proto.Message
	switch v := rd.(type) {
	case *bgp.RouteDistinguisherTwoOctetAS:
		glog.V(6).Infof("Admin: %+v Assigned: %+v", v.Admin, v.Assigned)
		r = &api.RouteDistinguisherTwoOctetAS{
			Admin:    uint32(v.Admin),
			Assigned: v.Assigned,
		}
	case *bgp.RouteDistinguisherIPAddressAS:
		glog.V(6).Infof("Admin: %+v Assigned: %+v", v.Admin, v.Assigned)
		r = &api.RouteDistinguisherIPAddress{
			Admin:    v.Admin.String(),
			Assigned: uint32(v.Assigned),
		}
	case *bgp.RouteDistinguisherFourOctetAS:
		glog.V(6).Infof("Admin: %+v Assigned: %+v", v.Admin, v.Assigned)
		r = &api.RouteDistinguisherFourOctetAS{
			Admin:    v.Admin,
			Assigned: uint32(v.Assigned),
		}
	default:
		glog.V(5).Infof("Unknown type: %+v", v)
		return nil
	}
	a, _ := ptypes.MarshalAny(r)
	return a
}

// MarshalRT marshals Route Target into Proto Any format
func MarshalRT(rt bgp.ExtendedCommunityInterface) *any.Any {
	var r proto.Message
	switch v := rt.(type) {
	case *bgp.TwoOctetAsSpecificExtended:
		r = &api.TwoOctetAsSpecificExtended{
			IsTransitive: true,
			SubType:      uint32(bgp.EC_SUBTYPE_ROUTE_TARGET),
			As:           uint32(v.AS),
			LocalAdmin:   uint32(v.LocalAdmin),
		}
	case *bgp.IPv4AddressSpecificExtended:
		r = &api.IPv4AddressSpecificExtended{
			IsTransitive: true,
			SubType:      uint32(bgp.EC_SUBTYPE_ROUTE_TARGET),
			Address:      v.IPv4.String(),
			LocalAdmin:   uint32(v.LocalAdmin),
		}
	case *bgp.FourOctetAsSpecificExtended:
		r = &api.FourOctetAsSpecificExtended{
			IsTransitive: true,
			SubType:      uint32(bgp.EC_SUBTYPE_ROUTE_TARGET),
			As:           uint32(v.AS),
			LocalAdmin:   uint32(v.LocalAdmin),
		}

	default:
		glog.V(5).Infof("Marshal RT Unknown type: %+v", v)
		return nil
	}
	a, _ := ptypes.MarshalAny(r)
	return a
}

// MarshalRTs marshals slice of Route Targets into Proto Any format
func MarshalRTs(rts []bgp.ExtendedCommunityInterface) []*any.Any {
	a := make([]*any.Any, len(rts))
	for i := 0; i < len(rts); i++ {
		a[i] = MarshalRT(rts[i])
	}

	return a
}

// MarshalPrefixSID marshals Prefix SID object into a slice of protobuf's anytype
func MarshalPrefixSID(psid *prefixsid.PSid) []*any.Any {
	mtlvs := make([]*any.Any, 0)
	if psid == nil {
		return mtlvs
	}
	var r proto.Message
	switch {
	case psid.SRv6L3Service != nil:
		o := &api.SRv6L3ServiceTLV{}
		o.SubTlvs = MarshalSRv6SubTLVs(psid.SRv6L3Service.SubTLVs)
		r = o
	default:
		return nil
	}
	a, _ := ptypes.MarshalAny(r)
	mtlvs = append(mtlvs, a)

	return mtlvs
}

// MarshalSRv6SubTLVs marshals SRv6 SubTLV map into a native protobuf map of TLVs
func MarshalSRv6SubTLVs(tlvs map[uint8][]srv6.SubTLV) map[uint32]*api.SRv6TLV {
	mtlvs := make(map[uint32]*api.SRv6TLV, len(tlvs))
	for t, tlv := range tlvs {
		var r proto.Message
		switch t {
		case 1:
			for _, stlv := range tlv {
				infoS, ok := stlv.(*srv6.InformationSubTLV)
				if !ok {
					continue
				}
				o := &api.SRv6InformationSubTLV{
					Flags: &api.SRv6SIDFlags{},
				}
				o.EndpointBehavior = uint32(infoS.EndpointBehavior)
				o.Sid = make([]byte, 16)
				copy(o.Sid, []byte(net.ParseIP(infoS.SID).To16()))
				o.SubSubTlvs = MarshalSRv6SubSubTLVs(infoS.SubSubTLVs)
				r = o
				a, _ := ptypes.MarshalAny(r)
				tlvs, ok := mtlvs[uint32(t)]
				if !ok {
					tlvs = &api.SRv6TLV{
						Tlv: make([]*any.Any, 0),
					}
					mtlvs[uint32(t)] = tlvs
				}
				tlvs.Tlv = append(tlvs.Tlv, a)
			}
		default:
			continue
		}
	}

	return mtlvs
}

// MarshalSRv6SubSubTLVs marshals SRv6 SubSubTLV map into a native protobuf map of TLVs
func MarshalSRv6SubSubTLVs(stlvs map[uint8][]srv6.SubSubTLV) map[uint32]*api.SRv6TLV {
	mtlvs := make(map[uint32]*api.SRv6TLV, len(stlvs))
	for t, tlv := range stlvs {
		var r proto.Message
		switch t {
		case 1:
			for _, stlv := range tlv {
				sstlv, ok := stlv.(*srv6.SIDStructureSubSubTLV)
				if !ok {
					continue
				}
				o := &api.SRv6StructureSubSubTLV{
					LocalBlockLength:    uint32(sstlv.LocalBlockLength),
					LocalNodeLength:     uint32(sstlv.LocalNodeLength),
					FunctionLength:      uint32(sstlv.FunctionLength),
					ArgumentLength:      uint32(sstlv.ArgumentLength),
					TranspositionLength: uint32(sstlv.TranspositionLength),
					TranspositionOffset: uint32(sstlv.TranspositionOffset),
				}
				r = o
				a, _ := ptypes.MarshalAny(r)
				tlvs, ok := mtlvs[uint32(t)]
				if !ok {
					tlvs = &api.SRv6TLV{
						Tlv: make([]*any.Any, 0),
					}
					mtlvs[uint32(t)] = tlvs
				}
				tlvs.Tlv = append(tlvs.Tlv, a)
			}
		default:
			continue
		}
	}

	return mtlvs
}

// UnmarshalPrefixSID unmarshals a slice of protobuf's anytype in Prefix SID object
func UnmarshalPrefixSID(svcs []*any.Any) (*prefixsid.PSid, error) {
	psid := &prefixsid.PSid{}
	for _, svc := range svcs {
		var svcValue ptypes.DynamicAny
		if err := ptypes.UnmarshalAny(svc, &svcValue); err != nil {
			return nil, fmt.Errorf("failed to unmarshal Prefix SID with error: %+v", err)
		}
		switch v := svcValue.Message.(type) {
		case *api.SRv6L3ServiceTLV:
			if l3, err := UnmarshalSRv6SubTLVs(v.SubTlvs); err == nil {
				psid.SRv6L3Service = &srv6.L3Service{
					SubTLVs: l3,
				}
			}
		default:
			continue
		}
	}
	return psid, nil
}

// UnmarshalSRv6SubTLVs unmarshals a native protobuf map of TLVs into SRv6 SubTLV map
func UnmarshalSRv6SubTLVs(stlvs map[uint32]*api.SRv6TLV) (map[uint8][]srv6.SubTLV, error) {
	var err error
	mtlvs := make(map[uint8][]srv6.SubTLV, len(stlvs))
	for t, tlv := range stlvs {
		switch t {
		case 1:
			for _, stlv := range tlv.Tlv {
				var stlvValue ptypes.DynamicAny
				o := &srv6.InformationSubTLV{}
				if err := ptypes.UnmarshalAny(stlv, &stlvValue); err != nil {
					return nil, fmt.Errorf("failed to unmarshal SRv6 SID Structure Sub Sub TLV with error: %+v", err)
				}
				v, ok := stlvValue.Message.(*api.SRv6InformationSubTLV)
				if !ok {
					continue
				}
				o.EndpointBehavior = uint16(v.EndpointBehavior)
				// TODO Once Flags are finalized in RFC, populate its value
				o.Flags = 0
				o.SID = net.IP(v.Sid).To16().String()
				o.SubSubTLVs, err = UnmarshalSRv6SubSubTLVs(v.SubSubTlvs)
				if err != nil {
					continue
				}
				var e srv6.SubSubTLV
				e = o
				stlvs, ok := mtlvs[uint8(t)]
				if !ok {
					stlvs = make([]srv6.SubTLV, 0)
				}
				stlvs = append(stlvs, e)
				mtlvs[uint8(t)] = stlvs
			}
		default:
			continue
		}
	}

	return mtlvs, nil
}

// UnmarshalSRv6SubSubTLVs unmarshals a native protobuf map of TLVs into SRv6 SubSubTLV map
func UnmarshalSRv6SubSubTLVs(stlvs map[uint32]*api.SRv6TLV) (map[uint8][]srv6.SubSubTLV, error) {
	mtlvs := make(map[uint8][]srv6.SubSubTLV, len(stlvs))
	for t, tlv := range stlvs {
		switch t {
		case 1:
			for _, sstlv := range tlv.Tlv {
				var sstlvValue ptypes.DynamicAny
				o := &srv6.SIDStructureSubSubTLV{}
				if err := ptypes.UnmarshalAny(sstlv, &sstlvValue); err != nil {
					return nil, fmt.Errorf("failed to unmarshal SRv6 SID Structure Sub Sub TLV with error: %+v", err)
				}
				v, ok := sstlvValue.Message.(*api.SRv6StructureSubSubTLV)
				if !ok {
					continue
				}
				o.LocalBlockLength = uint8(v.LocalBlockLength)
				o.LocalNodeLength = uint8(v.LocalNodeLength)
				o.FunctionLength = uint8(v.FunctionLength)
				o.ArgumentLength = uint8(v.ArgumentLength)
				o.TranspositionLength = uint8(v.TranspositionLength)
				o.TranspositionOffset = uint8(v.TranspositionOffset)
				var e srv6.SubSubTLV
				e = o
				sstlvs, ok := mtlvs[uint8(t)]
				if !ok {
					sstlvs = make([]srv6.SubSubTLV, 0)
				}
				sstlvs = append(sstlvs, e)
				mtlvs[uint8(t)] = sstlvs
			}
		default:
			continue
		}
	}

	return mtlvs, nil
}

// MarshalRDFromString marshals Route Distinguisher into Protobuf Any.any format
func MarshalRDFromString(rd string) (*any.Any, error) {
	if err := RDValidator(rd); err != nil {
		return nil, err
	}
	// Sice passed RD got already validated, it is safe to ignore any error processing
	parts := strings.Split(rd, ":")
	if net.ParseIP(parts[0]).To4() != nil {
		// If parts[0] is a valid IP, then it is IP:Value
		n, _ := strconv.Atoi(parts[1])
		return MarshalRD(bgp.NewRouteDistinguisherIPAddressAS(parts[0], uint16(n))), nil
	}
	n1, _ := strconv.Atoi(parts[0])
	n2, _ := strconv.Atoi(parts[1])
	if n1 < math.MaxUint16 {
		// If parts[0] is less than MaxUint16, then it is 2 Bytes ASN: 4 Bytes Value
		return MarshalRD(bgp.NewRouteDistinguisherTwoOctetAS(uint16(n1), uint32(n2))), nil
	}
	// Since no match before then, it is 4 Bytes ASN: 2 Bytes Value
	return MarshalRD(bgp.NewRouteDistinguisherFourOctetAS(uint32(n1), uint16(n2))), nil
}

// MarshalRTFromString marshals Route Target into Protobuf Any.any format
func MarshalRTFromString(rt string) (*any.Any, error) {
	if err := RTValidator(rt); err != nil {
		return nil, err
	}
	// Sice passed RD got already validated, it is safe to ignore any error processing
	parts := strings.Split(rt, ":")
	if net.ParseIP(parts[0]).To4() != nil {
		// If parts[0] is a valid IP, then it is IP:Value
		n, _ := strconv.Atoi(parts[1])
		return MarshalRT(bgp.NewIPv4AddressSpecificExtended(bgp.EC_SUBTYPE_ROUTE_TARGET, parts[0], uint16(n), true)), nil
	}
	n1, _ := strconv.Atoi(parts[0])
	n2, _ := strconv.Atoi(parts[1])
	if n1 < math.MaxUint16 {
		// If parts[0] is less than MaxUint16, then it is 2 Bytes ASN: 4 Bytes Value
		return MarshalRT(bgp.NewTwoOctetAsSpecificExtended(bgp.EC_SUBTYPE_ROUTE_TARGET, uint16(n1), uint32(n2), true)), nil
	}

	// Since no match before then, it is 4 Bytes ASN: 2 Bytes Value
	return MarshalRT(bgp.NewFourOctetAsSpecificExtended(bgp.EC_SUBTYPE_ROUTE_TARGET, uint32(n1), uint16(n2), true)), nil
}
