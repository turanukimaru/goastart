package calcapi

import (
	"context"
	"github.com/turanukimaru/goastart/gen/calc"
	"github.com/turanukimaru/gormstart/pkg/dummydb"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	//	s.logger.Print("calc.add")
	err = dummydb.DbAccess()
	return p.A + p.B, err
}
