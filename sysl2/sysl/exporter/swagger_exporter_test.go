package exporter

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/anz-bank/sysl/sysl2/sysl/parse"
	"github.com/anz-bank/sysl/sysl2/sysl/syslutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestExportAll(t *testing.T) {
	t.Parallel()
	modelParser := parse.NewParser()
	const syslTestDir = "test-data"
	files, err := ioutil.ReadDir(syslTestDir)
	require.NoError(t, err)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		parts := strings.Split(file.Name(), ".")
		if strings.EqualFold(parts[1], "sysl") {
			t.Run(parts[0], func(t *testing.T) {
				t.Parallel()
				mod, _, err1 := parse.LoadAndGetDefaultApp("exporter/test-data/"+parts[0]+`.sysl`,
					syslutil.NewChrootFs(afero.NewOsFs(), ".."), modelParser)
				require.NoError(t, err1)
				if err1 != nil {
					t.Errorf("Error reading sysl %s", parts[0]+`.sysl`)
				}
				swaggerExporter := MakeSwaggerExporter(mod.GetApps()["testapp"], logrus.StandardLogger())
				err2 := swaggerExporter.GenerateSwagger()
				require.NoError(t, err2)
				out, err := swaggerExporter.SerializeToYaml()
				require.NoError(t, err)
				yamlFileBytes, err := ioutil.ReadFile("../exporter/test-data/" + parts[0] + `.yaml`)
				require.NoError(t, err)
				if string(yamlFileBytes) != string(out) {
					t.Errorf("Content mismatched\n%s\n*******\n%s for Filename %s", string(yamlFileBytes),
						string(out), parts[0]+`.sysl`)
				}
			})
		}
	}
}