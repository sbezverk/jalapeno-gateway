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
	requestLoop(ctx, gwclient)
}

func getVPN(ctx context.Context, gwclient pbapi.GatewayServiceClient, req *pbapi.L3VPNRequest) (uint32, error) {
	resp, err := gwclient.L3VPN(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("failed to request VPN label with error: %+v", err)
	}
	glog.Infof("Received label: %+v", resp.Label)

	if resp.VpnPrefix != nil {
		glog.Infof("Received prefixes: %+v", resp.VpnPrefix)
	}
	return resp.Label, nil
}

func requestLoop(ctx context.Context, gwclient pbapi.GatewayServiceClient) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter RD in a form of '2 Bytes ASN:4 Bytes Value', '4 Bytes ASN:2 Bytes Value' or '2 Bytes ASN:4 Bytes Value', 'q' to exit\n")
		rd, err := reader.ReadString('\n')
		if err != nil {
			glog.Errorf("failed to read input with error: %+v, try again...", err)
		}
		rd = strings.Replace(rd, "\n", "", -1)
		if strings.ToLower(rd) == "q" {
			glog.Infof("all done, exiting the loop..")
			return
		}
		if err := validateRD(rd); err != nil {
			glog.Errorf("failed to parse entered RD: %s with error: %+v, try again...", rd, err)
		}
		mrd, err := marshalRD(rd)
		if err != nil {
			glog.Errorf("failed to marshal RD: %s with error: %+v, try again...", rd, err)
		}
		req := &pbapi.L3VPNRequest{Rd: mrd}
		label, err := getVPN(ctx, gwclient, req)
		if err != nil {
			glog.Errorf("failed to receive label for  RD: %s with error: %+v, try again...", rd, err)
		}
		glog.Infof("RD: %s VPN Label: %d", rd, label)
	}
}

func marshalRD(rd string) (*any.Any, error) {
	if err := validateRD(rd); err != nil {
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

func validateRD(rd string) error {
	parts := strings.Split(rd, ":")
	if len(parts) != 2 {
		return fmt.Errorf("malformed RD, expected 2 fields separated by ':'")
	}
	part1 := strings.Trim(parts[0], " ")
	part2 := strings.Trim(parts[1], " ")

	if net.ParseIP(part1).To4() != nil {
		// Possible RD in format IP:Value, Value cannot exceed uint16 value.
		n, err := strconv.Atoi(part2)
		if err != nil {
			return fmt.Errorf("malformed RD, failed to parse Value field %s with error: %+v", part2, err)
		}
		if n > math.MaxUint16 {
			return fmt.Errorf("malformed RD, Value field %d exceeds maximum allowable %d", n, math.MaxUint16)
		}
		return nil
	}
	n1, err := strconv.Atoi(part1)
	if err != nil {
		return fmt.Errorf("malformed RD, failed to parse ASN field %s with error: %+v", part1, err)
	}
	n2, err := strconv.Atoi(part2)
	if err != nil {
		return fmt.Errorf("malformed RD, failed to parse Value field %s with error: %+v", part2, err)
	}
	// Check for ASN 4 bytes and Value 2 bytes
	if n1 > math.MaxUint16 && n1 <= math.MaxUint32 {
		if n2 > math.MaxUint16 {
			return fmt.Errorf("malformed RD, Value field %d exceeds maximum allowable %d", n2, math.MaxUint16)
		}
		return nil
	}
	// Check for ASN 2 bytes and Value 4 bytes
	if n1 <= math.MaxUint16 {
		if n2 > math.MaxUint32 {
			return fmt.Errorf("malformed RD, Value field %d exceeds maximum allowable %d", n2, math.MaxUint32)
		}
		return nil
	}
	return fmt.Errorf("malformed RD, ASN field %d exceeds maximum allowable %d", n1, math.MaxUint32)
}
