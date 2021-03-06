# Protocol Documentation
<a name="top"/>

## Table of Contents

- [echo.proto](#echo.proto)
    - [EchoRequest](#grpc.gomeetexamples.echo.EchoRequest)
    - [EchoResponse](#grpc.gomeetexamples.echo.EchoResponse)
    - [EmptyMessage](#grpc.gomeetexamples.echo.EmptyMessage)
    - [ServiceStatus](#grpc.gomeetexamples.echo.ServiceStatus)
    - [ServicesStatusList](#grpc.gomeetexamples.echo.ServicesStatusList)
    - [VersionResponse](#grpc.gomeetexamples.echo.VersionResponse)
  
    - [ServiceStatus.Status](#grpc.gomeetexamples.echo.ServiceStatus.Status)
  
  
    - [Echo](#grpc.gomeetexamples.echo.Echo)
  

- [Scalar Value Types](#scalar-value-types)



<a name="echo.proto"/>
<p align="right"><a href="#top">Top</a></p>

## echo.proto



<a name="grpc.gomeetexamples.echo.EchoRequest"/>

### EchoRequest
EchoRequest represents a simple message sent to the Echo service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | Uuid represents the message identifier. |
| content | [string](#string) |  | some content // note : lorem.string =&gt; generate a random sentence 3 to 11 words in length |






<a name="grpc.gomeetexamples.echo.EchoResponse"/>

### EchoResponse
EchoResponse represents a simple message that the Echo service return.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | Uuid represents the message identifier. |
| content | [string](#string) |  | some content |






<a name="grpc.gomeetexamples.echo.EmptyMessage"/>

### EmptyMessage







<a name="grpc.gomeetexamples.echo.ServiceStatus"/>

### ServiceStatus
SeviceStatus represents a sub services status message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name of service |
| version | [string](#string) |  | version of service |
| status | [ServiceStatus.Status](#grpc.gomeetexamples.echo.ServiceStatus.Status) |  | status of service see enum Status |
| e_msg | [string](#string) |  |  |






<a name="grpc.gomeetexamples.echo.ServicesStatusList"/>

### ServicesStatusList
ServicesStatusList is the sub services status list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| services | [ServiceStatus](#grpc.gomeetexamples.echo.ServiceStatus) | repeated |  |






<a name="grpc.gomeetexamples.echo.VersionResponse"/>

### VersionResponse
VersionMessage represents a version message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Id represents the message identifier. |
| version | [string](#string) |  |  |





 


<a name="grpc.gomeetexamples.echo.ServiceStatus.Status"/>

### ServiceStatus.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| UNAVAILABLE | 1 |  |


 

 


<a name="grpc.gomeetexamples.echo.Echo"/>

### Echo


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [EmptyMessage](#grpc.gomeetexamples.echo.EmptyMessage) | [VersionResponse](#grpc.gomeetexamples.echo.EmptyMessage) | Version method receives no paramaters and returns a version message. |
| ServicesStatus | [EmptyMessage](#grpc.gomeetexamples.echo.EmptyMessage) | [ServicesStatusList](#grpc.gomeetexamples.echo.EmptyMessage) | ServicesStatus method receives no paramaters and returns all services status message |
| Echo | [EchoRequest](#grpc.gomeetexamples.echo.EchoRequest) | [EchoResponse](#grpc.gomeetexamples.echo.EchoRequest) | Echo method receives a simple message and returns it. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

