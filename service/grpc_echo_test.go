package service

import (
	"testing"

	"github.com/gomeet-examples/svc-echo/pb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestEcho(t *testing.T) {
	server := newEchoServer()
	ctx := context.Background()

	req := pb.NewEchoRequestGomeetFaker()
	res, err := server.Echo(ctx, req)
	assert.Nil(t, err, "Echo: error on call")
	assert.NotNil(t, res, "Echo: error on call")

	assert.Equal(t, req.GetUuid(), res.GetUuid(), "Echo: Uuid field in response must be the same as that of the request")
	assert.Equal(t, req.GetContent(), res.GetContent(), "Echo: Content field in response must be the same as that of the request")

	// empty Uuid
	req = &pb.EchoRequest{
		Content: "Valid content",
	}
	assert.NotNil(t, req.Validate(), "Echo: validate true on empty Uuid")
	res, err = server.Echo(ctx, req)
	assert.NotNil(t, err, "Echo: no error on call empty Uuid")
	assert.Nil(t, res, "Echo: no error on call empty Uuid")

	// invalid Uuid
	req = &pb.EchoRequest{
		Uuid:    "invalid",
		Content: "Valid content",
	}
	assert.NotNil(t, req.Validate(), "Echo: validate true on invalid Uuid")
	res, err = server.Echo(ctx, req)
	assert.NotNil(t, err, "Echo: no error on call invalid Uuid")
	assert.Nil(t, res, "Echo: no error on call invalid Uuid")

	// empty content
	req = &pb.EchoRequest{
		Uuid: uuid.New().String(),
	}
	assert.NotNil(t, req.Validate(), "Echo: validate true on empty content")
	res, err = server.Echo(ctx, req)
	assert.NotNil(t, err, "Echo: no error on call on empty content")
	assert.Nil(t, res, "Echo: no error on call on empty content")

	// invalid content
	req = &pb.EchoRequest{
		Uuid:    uuid.New().String(),
		Content: "E",
	}
	assert.NotNil(t, req.Validate(), "Echo: validate true on invalid content")
	res, err = server.Echo(ctx, req)
	assert.NotNil(t, err, "Echo: no error on call on invalid content")
	assert.Nil(t, res, "Echo: no error on call on invalid content")
}
