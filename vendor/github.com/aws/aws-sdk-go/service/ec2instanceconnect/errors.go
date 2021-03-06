// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2instanceconnect

const (

	// ErrCodeAuthException for service response error code
	// "AuthException".
	//
	// Indicates that either your AWS credentials are invalid or you do not have
	// access to the EC2 instance.
	ErrCodeAuthException = "AuthException"

	// ErrCodeEC2InstanceNotFoundException for service response error code
	// "EC2InstanceNotFoundException".
	//
	// Indicates that the instance requested was not found in the given zone. Check
	// that you have provided a valid instance ID and the correct zone.
	ErrCodeEC2InstanceNotFoundException = "EC2InstanceNotFoundException"

	// ErrCodeInvalidArgsException for service response error code
	// "InvalidArgsException".
	//
	// Indicates that you provided bad input. Ensure you have a valid instance ID,
	// the correct zone, and a valid SSH public key.
	ErrCodeInvalidArgsException = "InvalidArgsException"

	// ErrCodeServiceException for service response error code
	// "ServiceException".
	//
	// Indicates that the service encountered an error. Follow the message's instructions
	// and try again.
	ErrCodeServiceException = "ServiceException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Indicates you have been making requests too frequently and have been throttled.
	// Wait for a while and try again. If higher call volume is warranted contact
	// AWS Support.
	ErrCodeThrottlingException = "ThrottlingException"
)
