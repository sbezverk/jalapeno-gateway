package types

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

// UnmarshalJSON unmarshals a byte slice into VRF structure
func (vrf *VRF) UnmarshalJSON(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}
	// vrf_name is a mandatory field, if missing, then return error
	if err := json.Unmarshal(objmap["vrf_name"], &vrf.VRFName); err != nil {
		return err
	}
	if sz, ok := objmap["security_zone"]; ok {
		if err := json.Unmarshal(sz, &vrf.SecurityZone); err != nil {
			return err
		}
	}
	if cp, ok := objmap["config_parameters"]; ok {
		vrf.ConfigParameters = &ConfigParameters{}
		if err := json.Unmarshal(cp, &vrf.ConfigParameters); err != nil {
			return err
		}
	}

	return nil
}

// Unmarshal unmarshals a map of JSON Raw messages into VRF structure
func (vrf *VRF) Unmarshal(objmap map[string]json.RawMessage) error {
	if err := json.Unmarshal(objmap["vrf_name"], &vrf.VRFName); err != nil {
		return err
	}
	if sz, ok := objmap["security_zone"]; ok {
		if err := json.Unmarshal(sz, &vrf.SecurityZone); err != nil {
			return err
		}
	}
	if cp, ok := objmap["config_parameters"]; ok {
		vrf.ConfigParameters = &ConfigParameters{}
		if err := json.Unmarshal(cp, &vrf.ConfigParameters); err != nil {
			return err
		}
	}

	return nil
}

// UnmarshalJSON unmarshal byte slice into BGP
func (r *StaticRoute) UnmarshalJSON(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}
	if v, ok := objmap["description"]; ok {
		if err := json.Unmarshal(v, &r.Description); err != nil {
			return err
		}
	}
	if v, ok := objmap["ite_id"]; ok {
		if err := json.Unmarshal(v, &r.SiteID); err != nil {
			return err
		}
	}
	if v, ok := objmap["prefix_length"]; ok {
		if err := json.Unmarshal(v, &r.PrefixLength); err != nil {
			return err
		}
	}
	var prefix string
	if v, ok := objmap["prefix"]; ok {
		if err := json.Unmarshal(v, &prefix); err != nil {
			return err
		}
		if p, err := prefixToBytes(prefix); err == nil {
			r.Prefix = make([]byte, len(p))
			copy(r.Prefix, p)
		}
	}
	var nexthop string
	if v, ok := objmap["next_hopn"]; ok {
		if err := json.Unmarshal(v, &nexthop); err != nil {
			return err
		}
		if p, err := prefixToBytes(nexthop); err == nil {
			r.NextHop = make([]byte, len(p))
			copy(r.NextHop, p)
		}
	}
	return nil
}

// UnmarshalJSON unmarshal byte slice into BGP
func (bgp *BGP) UnmarshalJSON(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}
	bgp.DefaultInfoOriginate = stringToBool(objmap, "default_info_originate")
	bgp.RDAuto = stringToBool(objmap, "rd_auto")

	return nil
}

// UnmarshalJSON unmarshal byte slice into ConfigParameters
func (cp *ConfigParameters) UnmarshalJSON(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}

	cp.BGP = &BGP{}
	if bgp, ok := objmap["bgp"]; ok {
		if err := json.Unmarshal(bgp, cp.BGP); err != nil {
			return err
		}
	}
	cp.AddressFamilies = make([]*AddressFamily, 0)
	if af, ok := objmap["address_families"]; ok {
		if err := json.Unmarshal(af, &cp.AddressFamilies); err != nil {
			return err
		}
	}

	return nil
}

func (rtt RouteTargetType) populate(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}
	for k, v := range objmap {
		var rts []json.RawMessage
		if err := json.Unmarshal(v, &rts); err != nil {
			return err
		}
		rtes := make([]string, 0)
		for _, e := range rts {
			var rte RouteTargetElement
			if err := json.Unmarshal(e, &rte); err != nil {
				return err
			}
			rtes = append(rtes, rte.String())
		}
		rtt[k] = rtes
	}
	return nil
}

func (rta RouteTargetAction) populate(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}
	for k, v := range objmap {
		rta[k] = make(RouteTargetType)
		rtt := make(RouteTargetType)
		if err := rtt.populate(v); err != nil {
			return err
		}
		rta[k] = rtt
	}
	return nil
}

// UnmarshalJSON unmarshals route targets structure associated with vrf
func (rt *RouteTargetLocation) UnmarshalJSON(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}
	*rt = make(RouteTargetLocation)

	for k, v := range objmap {
		// ignoring all route target locations with exception of core
		if strings.Compare(k, "core") != 0 {
			continue
		}
		rta := make(RouteTargetAction)
		if err := rta.populate(v); err != nil {
			return err
		}
		(*rt)[k] = rta
	}

	return nil
}

// UnmarshalJSON unmarshal byte slice into AddressFamily
func (af *AddressFamily) UnmarshalJSON(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}
	af.Policies = &Policy{}
	if v, ok := objmap["policies"]; ok {
		if err := json.Unmarshal(v, &af.Policies); err != nil {
			return err
		}
	}
	af.StaticRoutes = make([]*StaticRoute, 0)
	if sr, ok := objmap["static_routes"]; ok {
		if err := json.Unmarshal(sr, &af.StaticRoutes); err != nil {
			return err
		}
	}

	af.ConfigNeed = stringToBool(objmap, "bgp_address_family_config_needed")
	af.Enabled = stringToBool(objmap, "enabled")
	if afn, ok := objmap["af_name"]; ok {
		if err := json.Unmarshal(afn, &af.AFName); err != nil {
			return err
		}
	}
	if rts, ok := objmap["route_targets"]; ok {
		if err := json.Unmarshal(rts, &af.RouteTargets); err != nil {
			return err
		}
	}
	if sfn, ok := objmap["saf_name"]; ok {
		if err := json.Unmarshal(sfn, &af.SAFIName); err != nil {
			return err
		}
	}
	af.BGPAddressFamily = &BGPAddressFamily{}
	if baf, ok := objmap["bgp_address_family"]; ok {
		if err := json.Unmarshal(baf, &af.BGPAddressFamily); err != nil {
			return err
		}
	}

	return nil
}

// UnmarshalJSON unmarshal byte slice into BGPAddressFamily
func (bgpaf *BGPAddressFamily) UnmarshalJSON(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err != nil {
		return err
	}
	bgpaf.LabelAllocationPerVRF = stringToBool(objmap, "label_allocation_mode")
	bgpaf.RedistConnected = stringToBool(objmap, "redist_connected")
	bgpaf.RedistStatic = stringToBool(objmap, "redist_static")
	if v, ok := objmap["max_paths_ebgp"]; ok {
		if err := json.Unmarshal(v, &bgpaf.MaxPathEBGP); err != nil {
			return err
		}
	}
	if v, ok := objmap["max_paths_ibgp"]; ok {
		if err := json.Unmarshal(v, &bgpaf.MaxPathIBGP); err != nil {
			return err
		}
	}

	return nil
}

// UnmarshalJSON unmarshal byte slice into RouteTargetElement
func (rte *RouteTargetElement) UnmarshalJSON(b []byte) error {
	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(b, &objmap); err == nil {
		// If Unmarshal succeeded, then process object map
		as := objmap["as"]
		rte.AS = make([]byte, len(as))
		copy(rte.AS, []byte(as))
		index := objmap["index"]
		rte.Index = make([]byte, len(index))
		copy(rte.Index, []byte(index))
	} else {
		// Check if RT is just a string
		rt := strings.Split(strings.Replace(string(b), "\"", "", -1), ":")
		if len(rt) != 2 {
			return fmt.Errorf("invalid route target \"%s\"", string(b))
		}
		rte.AS = make([]byte, len(rt[0]))
		copy(rte.AS, []byte(rt[0]))
		rte.Index = make([]byte, len(rt[1]))
		copy(rte.Index, []byte(rt[1]))
	}
	return nil
}

func stringToBool(objmap map[string]json.RawMessage, key string) bool {
	var v string
	if err := json.Unmarshal(objmap[key], &v); err != nil {
		return false
	}
	if strings.Compare(strings.ToLower(v), "yes") == 0 {
		return true
	}

	return false
}

func prefixToBytes(p string) ([]byte, error) {
	b := net.ParseIP(p)
	if b == nil {
		return nil, fmt.Errorf("invalid prefix %s", p)
	}
	if b.To4() != nil {
		return []byte(b.To4()), nil
	}
	return []byte(b.To16()), nil
}

// MakeMessage build a message for sending over a channel between a puller process and the processor
func MakeMessage(b []byte) (*PullerMessage, error) {
	objmap := make(map[string]json.RawMessage)
	if err := json.Unmarshal(b, &objmap); err != nil {
		return nil, fmt.Errorf("failed to decode response body into object with error: %+v", err)
	}
	hit := make(map[string]json.RawMessage)
	if err := json.Unmarshal(objmap["hits"], &hit); err != nil {
		return nil, fmt.Errorf("failed to unmarshal hit with error: %+v", err)
	}
	hits := make([]json.RawMessage, 0)
	if err := json.Unmarshal(hit["hits"], &hits); err != nil {
		return nil, fmt.Errorf("failed to unmarshal hits with error: %+v", err)
	}
	msg := PullerMessage{
		Entries: make([]*VRF, len(hits)),
	}
	for i, h := range hits {
		hit := Hit{}
		if err := json.Unmarshal(h, &hit); err != nil {
			return nil, fmt.Errorf("failed to unmarshal hit with error: %+v", err)
		}
		vrf := &VRF{}
		if err := json.Unmarshal(hit.Source, &vrf); err != nil {
			return nil, fmt.Errorf("failed to unmarshal hits with error: %+v", err)
		}
		if hit.Version != nil {
			vrf.Version = hit.Version
		}
		msg.Entries[i] = vrf
	}

	return &msg, nil
}
