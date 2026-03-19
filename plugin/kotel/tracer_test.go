package kotel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func TestNewTracer(t *testing.T) {
	prop := propagation.NewCompositeTextMapPropagator()

	testCases := []struct {
		name               string
		opts               []TracerOpt
		expectedPropagator propagation.TextMapPropagator
	}{
		{
			name:               "Empty (Use globals)",
			opts:               []TracerOpt{},
			expectedPropagator: otel.GetTextMapPropagator(),
		},
		{
			name:               "With TracerPropagator",
			opts:               []TracerOpt{TracerPropagator(prop)},
			expectedPropagator: prop,
		},
		{
			name:               "Nil TracerPropagator",
			opts:               []TracerOpt{TracerPropagator(nil)},
			expectedPropagator: otel.GetTextMapPropagator(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewTracer(tc.opts...)

			assert.NotNil(t, result.tracer)
			assert.Equal(t, otel.GetTracerProvider(), result.tracerProvider)
			assert.Equal(t, tc.expectedPropagator, result.propagators)
			assert.Empty(t, result.clientID)
			assert.Empty(t, result.consumerGroup)
			assert.Nil(t, result.keyFormatter)
			assert.False(t, result.linkSpans)
		})
	}
}
