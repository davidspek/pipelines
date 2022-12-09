// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V2beta1RuntimeState Describes the runtime state of an entity.
//
//  - RUNTIMESTATE_UNSPECIFIED: Default value. This value is not used.
//  - PENDING: Service is preparing to execute an entity.
//  - RUNNING: Entity execution is in progress.
//  - SUCCEEDED: Entity completed successfully.
//  - SKIPPED: Entity has been skipped. For example, due to caching.
//  - FAILED: Entity execution has failed.
//  - CANCELING: Entity is being canceled. From this state, an entity may only
// change its state to SUCCEEDED, FAILED or CANCELED.
//  - CANCELED: Entity has been canceled.
//  - PAUSED: Entity has been paused. It can be resumed.
// swagger:model v2beta1RuntimeState
type V2beta1RuntimeState string

const (

	// V2beta1RuntimeStateRUNTIMESTATEUNSPECIFIED captures enum value "RUNTIMESTATE_UNSPECIFIED"
	V2beta1RuntimeStateRUNTIMESTATEUNSPECIFIED V2beta1RuntimeState = "RUNTIMESTATE_UNSPECIFIED"

	// V2beta1RuntimeStatePENDING captures enum value "PENDING"
	V2beta1RuntimeStatePENDING V2beta1RuntimeState = "PENDING"

	// V2beta1RuntimeStateRUNNING captures enum value "RUNNING"
	V2beta1RuntimeStateRUNNING V2beta1RuntimeState = "RUNNING"

	// V2beta1RuntimeStateSUCCEEDED captures enum value "SUCCEEDED"
	V2beta1RuntimeStateSUCCEEDED V2beta1RuntimeState = "SUCCEEDED"

	// V2beta1RuntimeStateSKIPPED captures enum value "SKIPPED"
	V2beta1RuntimeStateSKIPPED V2beta1RuntimeState = "SKIPPED"

	// V2beta1RuntimeStateFAILED captures enum value "FAILED"
	V2beta1RuntimeStateFAILED V2beta1RuntimeState = "FAILED"

	// V2beta1RuntimeStateCANCELING captures enum value "CANCELING"
	V2beta1RuntimeStateCANCELING V2beta1RuntimeState = "CANCELING"

	// V2beta1RuntimeStateCANCELED captures enum value "CANCELED"
	V2beta1RuntimeStateCANCELED V2beta1RuntimeState = "CANCELED"

	// V2beta1RuntimeStatePAUSED captures enum value "PAUSED"
	V2beta1RuntimeStatePAUSED V2beta1RuntimeState = "PAUSED"
)

// for schema
var v2beta1RuntimeStateEnum []interface{}

func init() {
	var res []V2beta1RuntimeState
	if err := json.Unmarshal([]byte(`["RUNTIMESTATE_UNSPECIFIED","PENDING","RUNNING","SUCCEEDED","SKIPPED","FAILED","CANCELING","CANCELED","PAUSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2beta1RuntimeStateEnum = append(v2beta1RuntimeStateEnum, v)
	}
}

func (m V2beta1RuntimeState) validateV2beta1RuntimeStateEnum(path, location string, value V2beta1RuntimeState) error {
	if err := validate.Enum(path, location, value, v2beta1RuntimeStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v2beta1 runtime state
func (m V2beta1RuntimeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV2beta1RuntimeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
