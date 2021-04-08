// Code generated by go-swagger; DO NOT EDIT.

package security_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeSecurityChecksReader is a Reader for the ChangeSecurityChecks structure.
type ChangeSecurityChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeSecurityChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeSecurityChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeSecurityChecksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeSecurityChecksOK creates a ChangeSecurityChecksOK with default headers values
func NewChangeSecurityChecksOK() *ChangeSecurityChecksOK {
	return &ChangeSecurityChecksOK{}
}

/*ChangeSecurityChecksOK handles this case with default header values.

A successful response.
*/
type ChangeSecurityChecksOK struct {
	Payload interface{}
}

func (o *ChangeSecurityChecksOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/Change][%d] changeSecurityChecksOk  %+v", 200, o.Payload)
}

func (o *ChangeSecurityChecksOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeSecurityChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeSecurityChecksDefault creates a ChangeSecurityChecksDefault with default headers values
func NewChangeSecurityChecksDefault(code int) *ChangeSecurityChecksDefault {
	return &ChangeSecurityChecksDefault{
		_statusCode: code,
	}
}

/*ChangeSecurityChecksDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeSecurityChecksDefault struct {
	_statusCode int

	Payload *ChangeSecurityChecksDefaultBody
}

// Code gets the status code for the change security checks default response
func (o *ChangeSecurityChecksDefault) Code() int {
	return o._statusCode
}

func (o *ChangeSecurityChecksDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/Change][%d] ChangeSecurityChecks default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeSecurityChecksDefault) GetPayload() *ChangeSecurityChecksDefaultBody {
	return o.Payload
}

func (o *ChangeSecurityChecksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeSecurityChecksDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeSecurityChecksBody change security checks body
swagger:model ChangeSecurityChecksBody
*/
type ChangeSecurityChecksBody struct {

	// params
	Params []*ParamsItems0 `json:"params"`
}

// Validate validates this change security checks body
func (o *ChangeSecurityChecksBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSecurityChecksBody) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	for i := 0; i < len(o.Params); i++ {
		if swag.IsZero(o.Params[i]) { // not required
			continue
		}

		if o.Params[i] != nil {
			if err := o.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSecurityChecksBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSecurityChecksBody) UnmarshalBinary(b []byte) error {
	var res ChangeSecurityChecksBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSecurityChecksDefaultBody change security checks default body
swagger:model ChangeSecurityChecksDefaultBody
*/
type ChangeSecurityChecksDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change security checks default body
func (o *ChangeSecurityChecksDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSecurityChecksDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSecurityChecksDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSecurityChecksDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeSecurityChecksDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeURL string `json:"type_url,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ParamsItems0 ChangeSecurityCheckParams specifies a single check parameters.
swagger:model ParamsItems0
*/
type ParamsItems0 struct {

	// The name of the check to change.
	Name string `json:"name,omitempty"`

	// enable
	Enable bool `json:"enable,omitempty"`

	// disable
	Disable bool `json:"disable,omitempty"`

	// SecurityCheckInterval represents possible execution interval values for checks.
	// Enum: [SECURITY_CHECK_INTERVAL_INVALID STANDARD FREQUENT RARE]
	Interval *string `json:"interval,omitempty"`
}

// Validate validates this params items0
func (o *ParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var paramsItems0TypeIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SECURITY_CHECK_INTERVAL_INVALID","STANDARD","FREQUENT","RARE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paramsItems0TypeIntervalPropEnum = append(paramsItems0TypeIntervalPropEnum, v)
	}
}

const (

	// ParamsItems0IntervalSECURITYCHECKINTERVALINVALID captures enum value "SECURITY_CHECK_INTERVAL_INVALID"
	ParamsItems0IntervalSECURITYCHECKINTERVALINVALID string = "SECURITY_CHECK_INTERVAL_INVALID"

	// ParamsItems0IntervalSTANDARD captures enum value "STANDARD"
	ParamsItems0IntervalSTANDARD string = "STANDARD"

	// ParamsItems0IntervalFREQUENT captures enum value "FREQUENT"
	ParamsItems0IntervalFREQUENT string = "FREQUENT"

	// ParamsItems0IntervalRARE captures enum value "RARE"
	ParamsItems0IntervalRARE string = "RARE"
)

// prop value enum
func (o *ParamsItems0) validateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paramsItems0TypeIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ParamsItems0) validateInterval(formats strfmt.Registry) error {

	if swag.IsZero(o.Interval) { // not required
		return nil
	}

	// value enum
	if err := o.validateIntervalEnum("interval", "body", *o.Interval); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ParamsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ParamsItems0) UnmarshalBinary(b []byte) error {
	var res ParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
