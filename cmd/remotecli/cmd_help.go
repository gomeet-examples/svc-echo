// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/echo.proto
package remotecli

func (c *remoteCli) cmdHelp(args []string) (string, error) {
	h := `HELP :

	┌─ version
	└─ call version service

	┌─ services_status
	└─ call services_status service

	┌─ echo <uuid [string]> <content [string]>
	└─ call echo service

	┌─ service_address
	└─ return service address

	┌─ jwt [<token>]
	└─ display current jwt or save none if it's set

	┌─ console_version
	└─ return console version

	┌─ tls_config
	└─ display TLS client configuration

	┌─ help
	└─ display this help
`
	if c.ctxCall == ConsoleCall {
		h += `
	┌─ exit
	└─ exit the console
`
	}

	return h + "\n", nil
}
