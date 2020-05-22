package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/osrg/gobgp/pkg/packet/bgp"
	"github.com/sbezverk/gobmp/pkg/srv6"
	pbapi "github.com/sbezverk/jalapeno-gateway/pkg/apis"
	"github.com/sbezverk/jalapeno-gateway/pkg/bgpclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	jalapenoGateway string
)

func init() {
	flag.StringVar(&jalapenoGateway, "gateway", "192.168.80.104:40040", "Address to access jalapeno gateway")
}

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	conn, err := grpc.DialContext(context.TODO(), jalapenoGateway, grpc.WithInsecure())
	if err != nil {
		glog.Errorf("failed to connect to Jalapeno Gateway at the address: %s with error: %+v", jalapenoGateway, err)
		os.Exit(1)
	}
	defer conn.Close()
	gwclient := pbapi.NewGatewayServiceClient(conn)

	mainLoop(gwclient)
}

func mainLoop(gwclient pbapi.GatewayServiceClient) {
	parameters := []parameter{
		// {
		// 	prompt:    "IPv4 address used by the application ",
		// 	validator: ipValidator,
		// },
		// {
		// 	prompt:    "IPv4 address for the VPNv4 next hop ",
		// 	validator: ipValidator,
		// },
		// {
		// 	prompt:    "Autonomous System Number ",
		// 	validator: asnValidator,
		// },
		{
			prompt:    "RD for the application VRF ",
			validator: rdValidator,
		},
		// {
		// 	prompt:    "RT for the application address ",
		// 	validator: rtValidator,
		// },
		// {
		// 	prompt:    "VPN Label ",
		// 	validator: labelValidator,
		// },
		// {
		// 	prompt:    "Unicast Label ",
		// 	validator: labelValidator,
		// },
	}
	for {
		getInput(parameters, 0)
		if err := processRequest(gwclient, parameters); err != nil {
			fmt.Printf("\nrequest failed with error: %+v\n\n\n", err)
			continue
		}
	}
}

func processRequest(gwclient pbapi.GatewayServiceClient, p []parameter) error {
	ctx := metadata.NewOutgoingContext(context.TODO(), metadata.New(map[string]string{
		"CLIENT_IP": net.ParseIP("57.57.57.7").String(),
	}))
	// Prepare the application IP
	//	prefix, maskLength, _ := getPrefixAndMask(p[0].input)
	// Prepare the next hop address
	//	nhPrefix, nhMask, _ := getPrefixAndMask(p[1].input)
	// Get ASN
	//	asn, _ := strconv.Atoi(p[2].input)
	// Get and marshal RD
	rd, _ := marshalRD(p[0].input)
	// Get and marshal a slice of RTs
	//	rt, _ := marshalRT(p[4].input)
	// Get VPN label
	//	vpnLabel, _ := getLabel(p[5].input)
	// Get Unicast label
	// ucastLabel, _ := getLabel(p[6].input)
	prefixes, err := getVpnPrefixByRD(ctx, gwclient, rd)
	if err != nil {
		return fmt.Errorf("failed to get vpn prefixes for RD: %s with error: %+v", p[0].input, err)
	}
	fmt.Printf("\nSRv6 L3 Prefixes for RD: %s\n", p[0].input)
	for _, p := range prefixes {
		if p.PrefixSid != nil {
			if psid, err := bgpclient.UnmarshalPrefixSID(p.PrefixSid.Tlvs); err == nil {
				if psid.SRv6L3Service != nil {
					for _, t := range psid.SRv6L3Service.SubTLVs[1] {
						s := t.(*srv6.InformationSubTLV).SubSubTLVs[1][0]
						fmt.Printf("SubTLV: %+v SubSubTLVs: %+v\n", t, s.(*srv6.SIDStructureSubSubTLV))
					}
				}
			}
		}
	}
	return nil
}

func getVpnPrefixByRD(ctx context.Context, gwclient pbapi.GatewayServiceClient, rd *any.Any) ([]*pbapi.SRv6L3Prefix, error) {
	req := &pbapi.L3VpnRequest{Rd: rd, Ipv4: true}
	resp, err := gwclient.SRv6L3VPN(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Srv6Prefix, nil
}

// func advertiseVpnPrefix(ctx context.Context,
// 	gwclient pbapi.GatewayServiceClient,
// 	// VPN prefix
// 	prefix []byte,
// 	// VPN Prefix mask
// 	mask int,
// 	// VPN Prefix's Next hop address
// 	nhPrefix []byte,
// 	// Autonomous Systen Number
// 	asn int,
// 	// VPN Prefix's Route Distinguisher
// 	rd *any.Any,
// 	// VPN Prefix's Route Targets
// 	rt *any.Any,
// 	// VPN Prefix's VPN label
// 	vpnLabel int) error {
// 	req := &pbapi.VPNv4Prefix{
// 		Prefix: []*pbapi.VPNPrefix{
// 			{
// 				Prefix: &pbapi.Prefix{
// 					Address:    prefix,
// 					MaskLength: uint32(mask),
// 				},
// 				VpnLabel:  uint32(vpnLabel),
// 				NhAddress: nhPrefix,
// 				Asn:       uint32(asn),
// 				Rd:        rd,
// 				Rt: []*any.Any{
// 					rt,
// 				},
// 			},
// 		},
// 	}
// 	// Sending request to program VPNv4 Prefix
// 	//	_, err := gwclient.AdvBGPVPNv4(ctx, req)

// 	return err
// }

func getPrefixAndMask(addr string) ([]byte, int, error) {
	if _, pr, err := net.ParseCIDR(addr); err == nil {
		l, _ := pr.Mask.Size()
		return pr.IP, l, nil
	}
	if pr := net.ParseIP(addr); pr != nil {
		return pr, 32, nil
	}

	return nil, 0, fmt.Errorf("invalid address %s", addr)
}

func getInput(p []parameter, index int) int {
	reader := bufio.NewReader(os.Stdin)
	i := index
	for {
		if i >= len(p) {
			return i
		}
		fmt.Printf("Enter %s, 'b' to return to the previous parameter or 'q' to exit\n", p[i].prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("failed to read input with error: %+v, try again...", err)
			continue
		}
		input = strings.Replace(input, "\n", "", -1)
		switch strings.ToLower(input) {
		case "q":
			os.Exit(0)
		case "b":
			if i == 0 {
				continue
			}
			return i - 1
		default:
			p[i].input = input
			if err := p[i].validator(p[i]); err != nil {
				fmt.Printf("Validation failed with error: %+v, try again...\n", err)
				continue
			}
			if i+1 < len(p) {
				i = getInput(p, i+1)
			} else {
				return i + 1
			}
		}
	}
}

func marshalRD(rd string) (*any.Any, error) {
	if err := bgpclient.RDValidator(rd); err != nil {
		return nil, err
	}
	// Sice passed RD got already validated, it is safe to ignore any error processing
	parts := strings.Split(rd, ":")
	if net.ParseIP(parts[0]).To4() != nil {
		// If parts[0] is a valid IP, then it is IP:Value
		n, _ := strconv.Atoi(parts[1])
		return bgpclient.MarshalRD(bgp.NewRouteDistinguisherIPAddressAS(parts[0], uint16(n))), nil
	}
	n1, _ := strconv.Atoi(parts[0])
	n2, _ := strconv.Atoi(parts[1])
	if n1 < math.MaxUint16 {
		// If parts[0] is less than MaxUint16, then it is 2 Bytes ASN: 4 Bytes Value
		return bgpclient.MarshalRD(bgp.NewRouteDistinguisherTwoOctetAS(uint16(n1), uint32(n2))), nil
	}
	// Since no match before then, it is 4 Bytes ASN: 2 Bytes Value
	return bgpclient.MarshalRD(bgp.NewRouteDistinguisherFourOctetAS(uint32(n1), uint16(n2))), nil
}

func marshalRT(rt string) (*any.Any, error) {
	if err := bgpclient.RTValidator(rt); err != nil {
		return nil, err
	}
	// Sice passed RD got already validated, it is safe to ignore any error processing
	parts := strings.Split(rt, ":")
	if net.ParseIP(parts[0]).To4() != nil {
		// If parts[0] is a valid IP, then it is IP:Value
		n, _ := strconv.Atoi(parts[1])
		return bgpclient.MarshalRT(bgp.NewIPv4AddressSpecificExtended(bgp.EC_SUBTYPE_ROUTE_TARGET, parts[0], uint16(n), true)), nil
	}
	n1, _ := strconv.Atoi(parts[0])
	n2, _ := strconv.Atoi(parts[1])
	if n1 < math.MaxUint16 {
		// If parts[0] is less than MaxUint16, then it is 2 Bytes ASN: 4 Bytes Value
		return bgpclient.MarshalRT(bgp.NewTwoOctetAsSpecificExtended(bgp.EC_SUBTYPE_ROUTE_TARGET, uint16(n1), uint32(n2), true)), nil
	}

	// Since no match before then, it is 4 Bytes ASN: 2 Bytes Value
	return bgpclient.MarshalRT(bgp.NewFourOctetAsSpecificExtended(bgp.EC_SUBTYPE_ROUTE_TARGET, uint32(n1), uint16(n2), true)), nil
}

type parameter struct {
	prompt    string
	input     string
	validator func(parameter) error
}

func ipValidator(p parameter) error {
	if _, _, err := net.ParseCIDR(p.input); err == nil {
		return nil
	}
	if net.ParseIP(p.input) != nil {
		return nil
	}

	return fmt.Errorf("invalid ip address %s", p.input)
}

func asnValidator(p parameter) error {
	asn, err := strconv.Atoi(p.input)
	if err != nil {
		return err
	}
	if asn <= 0 || asn >= math.MaxUint32 {
		return fmt.Errorf("invalid ASN %d", asn)
	}

	return nil
}

func rdValidator(p parameter) error {
	if err := bgpclient.RDValidator(p.input); err != nil {
		return err
	}

	return nil
}

func rtValidator(p parameter) error {
	if err := bgpclient.RTValidator(p.input); err != nil {
		return err
	}

	return nil
}

func labelValidator(p parameter) error {
	label, err := strconv.Atoi(p.input)
	if err != nil {
		return err
	}
	// Validating vpn Label that it is not excedding 2^20
	if label <= 0 || label > 1048576 {
		return fmt.Errorf("invalid vpn label %d", label)
	}

	return nil
}

func getLabel(s string) (int, error) {
	label, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	// Validating vpn Label that it is not excedding 2^20
	if label <= 0 || label > 1048576 {
		return 0, fmt.Errorf("invalid vpn label %d", label)
	}

	return label, nil
}
