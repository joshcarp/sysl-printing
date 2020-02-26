package main

import (
	"bytes"
	"fmt"
	"github.com/joshcarp/sysl-printing/pkg/printer"
	"github.com/joshcarp/sysl-printing/pkg/syslutil"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestPrinting(t *testing.T) {
	_, fs := syslutil.WriteToMemOverlayFs("../../tests")
	log := logrus.Logger{}

	module, _, _ := LoadSyslModule("/", "call.sysl", fs, &log)
	var buf bytes.Buffer
	pr := printer.NewPrinter(&buf)
	pr.PrintModule(module)
	fmt.Print(buf.String())
}
