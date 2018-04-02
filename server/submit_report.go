package server

import (
	"errors"

	context "golang.org/x/net/context"

	"github.com/lair-framework/go-nessus"
	log "github.com/sirupsen/logrus"

	"github.com/Project-Prismatica/prismatica_integration_nessus"

)

func (s PrismaticaIntegrationNessusServer) SubmitReport(ctx context.Context,
		r *prismatica_integration_nessus.SubmissionRequest) (
		*prismatica_integration_nessus.SubmissionResult, error) {

	if len(r.XmlReport) == 0 {
		return nil, errors.New("empty report xml")
	}

	reportData, err := nessus.Parse([]byte(r.XmlReport))
	if err != nil {
		log.WithFields(log.Fields{"nessus_error": err, "report": r.XmlReport}).
			Warn("could not parse requested report")
		return nil, errors.New("could not parse xml report")
	}

	log.WithFields(log.Fields{"reportData": reportData}).Info("loaded report")

	return nil, errors.New("not yet implemented")
}
