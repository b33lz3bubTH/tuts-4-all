package database

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func NewTracer() trace.Tracer {
	return otel.Tracer("tuts-4-all-backend")
}
