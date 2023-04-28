// Package fake implements a fake gantry.
package fake

import (
	"context"
	"fmt"

	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"

	"go.viam.com/rdk/components/gantry"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/testutils"
)

func init() {
	resource.RegisterComponent(
		gantry.API,
		resource.DefaultModelFamily.WithModel("fake"),
		resource.Registration[gantry.Gantry, resource.NoNativeConfig]{
			Constructor: func(
				ctx context.Context,
				_ resource.Dependencies,
				conf resource.Config,
				logger golog.Logger,
			) (gantry.Gantry, error) {
				return NewGantry(conf.ResourceName()), nil
			},
		})
}

// NewGantry returns a new fake gantry.
func NewGantry(name resource.Name) gantry.Gantry {
	return &Gantry{
		testutils.NewUnimplementedResource(name),
		resource.TriviallyReconfigurable{},
		resource.TriviallyCloseable{},
		[]float64{1.2},
		[]float64{5},
		r3.Vector{1, 0, 0},
		2,
	}
}

// Gantry is a fake gantry that can simply read and set properties.
type Gantry struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	positionsMm  []float64
	lengths      []float64
	axis         r3.Vector
	lengthMeters float64
}

// Position returns the position in meters.
func (g *Gantry) Position(ctx context.Context, extra map[string]interface{}) ([]float64, error) {
	return g.positionsMm, nil
}

// Lengths returns the position in meters.
func (g *Gantry) Lengths(ctx context.Context, extra map[string]interface{}) ([]float64, error) {
	return g.lengths, nil
}

// MoveToPosition is in meters.
func (g *Gantry) MoveToPosition(ctx context.Context, positionsMm []float64, extra map[string]interface{}) error {
	g.positionsMm = positionsMm
	return nil
}

// Stop doesn't do anything for a fake gantry.
func (g *Gantry) Stop(ctx context.Context, extra map[string]interface{}) error {
	return nil
}

// IsMoving is always false for a fake gantry.
func (g *Gantry) IsMoving(ctx context.Context) (bool, error) {
	return false, nil
}

// ModelFrame returns a Gantry frame.
func (g *Gantry) ModelFrame() referenceframe.Model {
	m := referenceframe.NewSimpleModel("")
	f, err := referenceframe.NewTranslationalFrame(g.Name().ShortName(), g.axis, referenceframe.Limit{0, g.lengthMeters})
	if err != nil {
		panic(fmt.Errorf("error creating frame: %w", err))
	}
	m.OrdTransforms = append(m.OrdTransforms, f)
	return m
}

// CurrentInputs returns positions in the Gantry frame model..
func (g *Gantry) CurrentInputs(ctx context.Context) ([]referenceframe.Input, error) {
	res, err := g.Position(ctx, nil)
	if err != nil {
		return nil, err
	}
	return referenceframe.FloatsToInputs(res), nil
}

// GoToInputs moves using the Gantry frames..
func (g *Gantry) GoToInputs(ctx context.Context, goal []referenceframe.Input) error {
	return g.MoveToPosition(ctx, referenceframe.InputsToFloats(goal), nil)
}