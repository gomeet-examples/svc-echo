package functest

import (
	"fmt"

	"github.com/google/uuid"

	pb "github.com/gomeet-examples/svc-echo/pb"
)

func testGetEchoRequest(
	config FunctionalTestConfig,
) (reqs []*pb.EchoRequest, extras map[string]interface{}, err error) {
	// return an array of pb.EchoRequest struct pointers,
	// each of them will be passed as an argument to the grpc Echo method

	reqs = append(reqs, &pb.EchoRequest{Uuid: "", Content: "Hello message"})        // empty Uuid
	reqs = append(reqs, &pb.EchoRequest{Uuid: "invalid", Content: "Hello message"}) // invalid Uuid
	reqs = append(reqs, &pb.EchoRequest{Uuid: uuid.New().String(), Content: ""})    // empty Content
	reqs = append(reqs, &pb.EchoRequest{Uuid: uuid.New().String(), Content: "E"})   // invalid Content

	// valid index 4+
	reqs = append(reqs, &pb.EchoRequest{Uuid: uuid.New().String(), Content: "Hello message"})
	return reqs, extras, err
}

func testEchoResponse(
	config FunctionalTestConfig,
	testsType string,
	testCaseResults []*TestCaseResult,
	extras map[string]interface{},
) (failures []TestFailure) {
	// Do something useful functional test with
	// testCaseResults[n].Request, testCaseResults[n].Response and testCaseResults[n].Error
	// then return a array of TestFailure struct
	// testsType value is value of FUNCTEST_HTTP (HTTP) and FUNCTEST_GRPC (GRPC) constants cf. types.go
	for i, tr := range testCaseResults {
		var (
			req *pb.EchoRequest
			res *pb.EchoResponse
			err error
			ok  bool
		)
		if tr.Request == nil {
			failures = append(failures, TestFailure{Procedure: "Echo", Message: "expected request message type pb.EchoRequest - nil given"})
			continue
		}
		req, ok = tr.Request.(*pb.EchoRequest)
		if !ok {
			failures = append(failures, TestFailure{Procedure: "Echo", Message: "expected request message type pb.EchoRequest - cast fail"})
			continue
		}

		err = tr.Error
		if i < 4 {
			if err == nil {
				// if no error are expected do something like this
				failures = append(failures, TestFailure{Procedure: "Echo", Message: "an error is expected"})
			}
			continue
		}

		if tr.Response != nil {
			res, ok = tr.Response.(*pb.EchoResponse)
			if !ok {
				failures = append(failures, TestFailure{Procedure: "Echo", Message: "expected response message type pb.EchoRequest - cast fail"})
				continue
			}
		}

		if err != nil {
			// if no error are expected do something like this
			failures = append(failures, TestFailure{Procedure: "Echo", Message: fmt.Sprintf("no error expected - %s %d", err, i)})
			continue
		}

		if req != nil && res != nil {
			if res.GetUuid() != req.GetUuid() {
				failureMsg := fmt.Sprintf("expected Uuid \"%s\" but got \"%s\" for request: %v", req.GetUuid(), res.GetUuid(), req)
				failures = append(failures, TestFailure{Procedure: "Echo", Message: failureMsg})
			}
			if res.GetContent() != req.GetContent() {
				failureMsg := fmt.Sprintf("expected Content \"%s\" but got \"%s\" for request: %v", req.GetContent(), res.GetContent(), req)
				failures = append(failures, TestFailure{Procedure: "Echo", Message: failureMsg})
			}
		}
	}

	return failures
}
