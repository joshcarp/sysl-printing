package main

import (
	"testing"

	"github.com/joshcarp/sysl-printing/pkg/printer"
	"github.com/joshcarp/sysl-printing/pkg/syslutil"
	"github.com/sirupsen/logrus"
)

func TestPrinting(t *testing.T) {
	_, fs := syslutil.WriteToMemOverlayFs("../../tests")
	log := logrus.Logger{}

	module, _, _ := LoadSyslModule("/", "call.sysl", fs, &log)
	printer.PrintModule(module)
}
