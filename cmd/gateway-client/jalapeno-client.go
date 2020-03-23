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

	ctx := metadata.NewOutgoingContext(context.TODO(), metadata.New(map[string]string{
		"CLIENT_IP": net.ParseIP("57.57.57.7").String(),
	}))

	mainLoop(ctx, gwclient)
}

func requestLoop(ctx context.Context, gwclient pbapi.GatewayServiceClient) {
	rd := "111:111"
	for {
		mrd, err := marshalRD(rd)
		if err != nil {
			glog.Errorf("failed to marshal RD: %s with error: %+v, try again...", rd, err)
		}
		req := &pbapi.L3VPNRequest{Rd: mrd, Ipv4: true}
		resp, err := gwclient.L3VPN(ctx, req)
		if err != nil {
			glog.Errorf("failed to request VPN label with error: %+v", err)
			continue
		}
		glog.Infof("Prefixes:")
		if resp.VpnPrefix != nil {
			for _, p := range resp.VpnPrefix {
				glog.Infof("- %s/%d VPN Label: %d Prefix SID label: %d", net.IP(p.Address).String(), p.MaskLength, p.VpnLabel, p.SidLabel)
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
	if asn <= 0 || asn >= math.MaxUint16 {
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

func mainLoop(ctx context.Context, gwclient pbapi.GatewayServiceClient) {
	parameters := []parameter{
		{
			prompt:    "IPv4 address used by the application ",
			validator: ipValidator,
		},
		{
			prompt:    "Autonomous System Number ",
			validator: asnValidator,
		},
		{
			prompt:    "RD for the application VRF ",
			validator: rdValidator,
		},
		{
			prompt:    "RT for the application address ",
			validator: rtValidator,
		},
	}
	for {
		getInput(parameters, 0)
		fmt.Printf("Acting on parameters: %+v\n", parameters)
	}
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
