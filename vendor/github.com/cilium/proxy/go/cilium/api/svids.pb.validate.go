// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cilium/api/svids.proto

package cilium

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on SVIDs with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SVIDs) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Identity

	for idx, item := range m.GetSvids() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SVIDsValidationError{
					field:  fmt.Sprintf("Svids[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SVIDsValidationError is the validation error returned by SVIDs.Validate if
// the designated constraints aren't met.
type SVIDsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SVIDsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SVIDsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SVIDsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SVIDsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SVIDsValidationError) ErrorName() string { return "SVIDsValidationError" }

// Error satisfies the builtin error interface
func (e SVIDsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSVIDs.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SVIDsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SVIDsValidationError{}

// Validate checks the field values on X509SVID with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *X509SVID) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SpiffeId

	// no validation rules for Cert

	// no validation rules for Key

	return nil
}

// X509SVIDValidationError is the validation error returned by
// X509SVID.Validate if the designated constraints aren't met.
type X509SVIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e X509SVIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e X509SVIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e X509SVIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e X509SVIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e X509SVIDValidationError) ErrorName() string { return "X509SVIDValidationError" }

// Error satisfies the builtin error interface
func (e X509SVIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sX509SVID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = X509SVIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = X509SVIDValidationError{}
