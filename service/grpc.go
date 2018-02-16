// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/echo.proto
package service

import (
	pb "github.com/gomeet-examples/svc-echo/pb"
	// SUB-SERVICES DEFINITION : import-pb
	// svc{{SubServiceNamePascalCase}}Pb "github.com/gomeet-examples/svc-{{SubServiceNameKebabCase}}/pb"
	// END SUB-SERVICES DEFINITION : import-pb
)

// Implements of echoServer
type echoServer struct {
	jwtSecret     string
	caCertificate string
	certificate   string
	privateKey    string

	// EXTRA : var
	// END EXTRA : var

	// SUB-SERVICES DEFINITION : var-address
	// svc{{SubServiceNamePascalCase}}Address string
	// END SUB-SERVICES DEFINITION : var-address

	// SUB-SERVICES DEFINITION : var-client
	// svc{{SubServiceNamePascalCase}}GrpcClient svc{{SubServiceNamePascalCase}}Pb.{{SubServiceNamePascalCase}}Client
	// END SUB-SERVICES DEFINITION : var-client
}

func newEchoServer() pb.EchoServer {
	return new(echoServer)
}
