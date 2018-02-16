// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/echo.proto
package functest

import (
	"fmt"

	pb "github.com/gomeet-examples/svc-echo/pb"
	"github.com/gomeet-examples/svc-echo/service"
)

func testVersionResponse(req *pb.EmptyMessage, res *pb.VersionResponse) (failures []TestFailure) {
	svc := service.NewService()

	if res.GetName() != svc.Name {
		failureMsg := fmt.Sprintf("expected name \"%s\" but got \"%s\"", svc.Name, res.GetName())
		failures = append(failures, TestFailure{Procedure: "Version", Message: failureMsg})
	}

	if res.GetVersion() != svc.Version {
		failureMsg := fmt.Sprintf("expected version number \"%s\" but got \"%s\"", svc.Version, res.GetVersion())
		failures = append(failures, TestFailure{Procedure: "Version", Message: failureMsg})
	}

	return failures
}
