package server

import (
	"testing"

	"github.com/Project-Prismatica/prismatica_integration_nessus"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestReportNoElements(t *testing.T) {
	ctx := context.Background()
	req := &prismatica_integration_nessus.SubmissionRequest{}

	req.XmlReport = "asdf"

	res, err := cli.SubmitReport(ctx, req)
	assert.Nil(t, res)

	assert.NotNil(t, err)

	assert.Contains(t, err.Error(), "no report elements")
}

func TestBadReportParse(t *testing.T) {
	ctx := context.Background()
	req := &prismatica_integration_nessus.SubmissionRequest{}

	req.XmlReport = "asdf"

	res, err := cli.SubmitReport(ctx, req)
	assert.Nil(t, res)

	assert.NotNil(t, err)

	assert.Contains(t, err.Error(), "could not parse xml report")
}

func TestEmptyReportXml(t *testing.T) {
	ctx := context.Background()
	req := &prismatica_integration_nessus.SubmissionRequest{}

	res, err := cli.SubmitReport(ctx, req)
	assert.Nil(t, res)

	assert.NotNil(t, err)

	assert.Contains(t, err.Error(), "empty report xml")
}
