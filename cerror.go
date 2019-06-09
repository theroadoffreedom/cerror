package cerror

import (

	log "github.com/sirupsen/logrus"
	perrors "github.com/pingcap/errors"
)

// MustNil cleans up and fatals if err is not nil.
func MustNil(err error, closeFuns ...func()) {
	if err != nil {
		for _, f := range closeFuns {
			f()
		}
		log.Fatalf(perrors.ErrorStack(err))
	}
}
