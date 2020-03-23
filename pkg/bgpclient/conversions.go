package bgpclient

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	api "github.com/osrg/gobgp/api"
	"github.com/osrg/gobgp/pkg/packet/bgp"
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
		glog.V(5).Infof("Admin: %+v Assigned: %+v", v.Admin, v.Assigned)
		r = &api.RouteDistinguisherTwoOctetAS{
			Admin:    uint32(v.Admin),
			Assigned: v.Assigned,
		}
	case *bgp.RouteDistinguisherIPAddressAS:
		glog.V(5).Infof("Admin: %+v Assigned: %+v", v.Admin, v.Assigned)
		r = &api.RouteDistinguisherIPAddress{
			Admin:    v.Admin.String(),
			Assigned: uint32(v.Assigned),
		}
	case *bgp.RouteDistinguisherFourOctetAS:
		glog.V(5).Infof("Admin: %+v Assigned: %+v", v.Admin, v.Assigned)
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
