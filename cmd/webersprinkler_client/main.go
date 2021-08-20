package main

import (
	"context"
	"flag"
	"log"

	"github.com/ejweber/webersprinkler2/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			// TODO: Get default CA root cert file.
			// *caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewSprinklerToServerClient(conn)

	client.SendStatus(context.Background(), &api.SprinklerStatusRequest{})
}

// func main() {
// 	flag.Parse()
// 	var opts []grpc.DialOption
// 	if *tls {
// 		if *caFile == "" {
// 			*caFile = data.Path("x509/ca_cert.pem")
// 		}
// 		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
// 		if err != nil {
// 			log.Fatalf("Failed to create TLS credentials %v", err)
// 		}
// 		opts = append(opts, grpc.WithTransportCredentials(creds))
// 	} else {
// 		opts = append(opts, grpc.WithInsecure())
// 	}

// 	opts = append(opts, grpc.WithBlock())
// 	conn, err := grpc.Dial(*serverAddr, opts...)
// 	if err != nil {
// 		log.Fatalf("fail to dial: %v", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewRouteGuideClient(conn)

// 	// Looking for a valid feature
// 	printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})

// 	// Feature missing.
// 	printFeature(client, &pb.Point{Latitude: 0, Longitude: 0})

// 	// Looking for features between 40, -75 and 42, -73.
// 	printFeatures(client, &pb.Rectangle{
// 		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
// 		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
// 	})

// 	// RecordRoute
// 	runRecordRoute(client)

// 	// RouteChat
// 	runRouteChat(client)
// }
