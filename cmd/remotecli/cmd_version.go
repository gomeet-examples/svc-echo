// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/echo.proto
package remotecli

import (
	"fmt"
)

func (c remoteCli) cmdVersion(args []string) (string, error) {
	r, err := c.RemoteVersion()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Version: %v", r), nil
}
