// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package transform

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/cosi-project/runtime/pkg/controller"
	"github.com/cosi-project/runtime/pkg/state"
)

// ControllerOptions configures TransformController.
type ControllerOptions struct {
	extraEventCh            <-chan struct{}
	onShutdownCallback      OnShutdownCallback
	inputListOptions        []state.ListOption
	extraInputs             []controller.Input
	extraOutputs            []controller.Output
	primaryOutputKind       controller.OutputKind
	requeueInterval         time.Duration
	inputFinalizers         bool
	ignoreTearingDownInputs bool
}

// ControllerOption is an option for TransformController.
type ControllerOption func(*ControllerOptions)

// WithInputListOptions adds a filter on input resource list.
//
// E.g., query only resources with specific labels.
func WithInputListOptions(opts ...state.ListOption) ControllerOption {
	return func(o *ControllerOptions) {
		o.inputListOptions = append(o.inputListOptions, opts...)
	}
}

// WithExtraInputs adds extra inputs to the controller.
func WithExtraInputs(inputs ...controller.Input) ControllerOption {
	return func(o *ControllerOptions) {
		o.extraInputs = append(o.extraInputs, inputs...)
	}
}

// WithExtraOutputs adds extra outputs to the controller.
func WithExtraOutputs(outputs ...controller.Output) ControllerOption {
	return func(o *ControllerOptions) {
		o.extraOutputs = append(o.extraOutputs, outputs...)
	}
}

// WithInputFinalizers enables setting finalizers on controller inputs.
//
// The finalizer on input will be removed only when matching output is destroyed.
func WithInputFinalizers() ControllerOption {
	return func(o *ControllerOptions) {
		if o.ignoreTearingDownInputs {
			panic("WithIgnoreTearingDownInputs is mutually exclusive with WithInputFinalizers")
		}

		o.inputFinalizers = true
	}
}

// WithIgnoreTearingDownInputs makes controller treat tearing down inputs as 'normal' inputs.
//
// With this setting enabled outputs will still exist until the input is destroyed.
// This setting is mutually exclusive with WithInputFinalizers.
func WithIgnoreTearingDownInputs() ControllerOption {
	return func(o *ControllerOptions) {
		if o.inputFinalizers {
			panic("WithIgnoreTearingDownInputs is mutually exclusive with WithInputFinalizers")
		}

		o.ignoreTearingDownInputs = true
	}
}

// WithExtraEventChannel adds an extra event channel to the controller.
//
// When this channel receives data, the controller will run the transform function.
// This is useful to wake up the controller from a goroutine.
func WithExtraEventChannel(extraEventCh <-chan struct{}) ControllerOption {
	return func(o *ControllerOptions) {
		o.extraEventCh = extraEventCh
	}
}

// OnShutdownCallback is a function called when the controller is shutting down, either gracefully or due to an error.
type OnShutdownCallback func(ctx context.Context, rw controller.ReaderWriter, logger *zap.Logger)

// WithOnShutdownCallback adds a callback to be called when the controller is shutting down, either gracefully or due to an error.
func WithOnShutdownCallback(onShutdownCallback OnShutdownCallback) ControllerOption {
	return func(o *ControllerOptions) {
		o.onShutdownCallback = onShutdownCallback
	}
}

// WithOutputKind sets main output resource kind.
func WithOutputKind(kind controller.OutputKind) ControllerOption {
	return func(o *ControllerOptions) {
		o.primaryOutputKind = kind
	}
}

// WithRequeueInterval sets the requeue interval of the transform controller.
// Requeue is triggered by returning an error tagged SkipReconcileAndRequeue from the input processor function.
func WithRequeueInterval(value time.Duration) ControllerOption {
	return func(o *ControllerOptions) {
		o.requeueInterval = value
	}
}
