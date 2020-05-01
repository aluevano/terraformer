// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

const (

	// ErrCodeInvalidArgsException for service response error code
	// "InvalidArgsException".
	//
	// One or more of of request parameters specified is not valid.
	ErrCodeInvalidArgsException = "InvalidArgsException"

	// ErrCodeInvalidArnException for service response error code
	// "InvalidArnException".
	//
	// The requested Amazon Resource Name (ARN) does not refer to an existing resource.
	ErrCodeInvalidArnException = "InvalidArnException"

	// ErrCodeInvalidDomainValidationOptionsException for service response error code
	// "InvalidDomainValidationOptionsException".
	//
	// One or more values in the DomainValidationOption structure is incorrect.
	ErrCodeInvalidDomainValidationOptionsException = "InvalidDomainValidationOptionsException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// An input parameter was invalid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// Processing has reached an invalid state.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeInvalidTagException for service response error code
	// "InvalidTagException".
	//
	// One or both of the values that make up the key-value pair is not valid. For
	// example, you cannot specify a tag value that begins with aws:.
	ErrCodeInvalidTagException = "InvalidTagException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// An ACM quota has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeRequestInProgressException for service response error code
	// "RequestInProgressException".
	//
	// The certificate request is in process and the certificate in your account
	// has not yet been issued.
	ErrCodeRequestInProgressException = "RequestInProgressException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The certificate is in use by another AWS service in the caller's account.
	// Remove the association and try again.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified certificate cannot be found in the caller's account or the
	// caller's account cannot be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTagPolicyException for service response error code
	// "TagPolicyException".
	//
	// A specified tag did not comply with an existing tag policy and was rejected.
	ErrCodeTagPolicyException = "TagPolicyException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The request contains too many tags. Try the request again with fewer tags.
	ErrCodeTooManyTagsException = "TooManyTagsException"
)
