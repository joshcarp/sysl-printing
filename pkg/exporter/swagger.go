package exporter

import (
	yaml "github.com/ghodss/yaml"
	"github.com/go-openapi/spec"
	proto "github.com/joshcarp/sysl-printing/pkg/sysl"
	"github.com/sirupsen/logrus"
)

type SwaggerExporter struct {
	app          *proto.Application
	buildSwagger *spec.Swagger
	log          *logrus.Logger
}

func MakeSwaggerExporter(app *proto.Application, logger *logrus.Logger) *SwaggerExporter {
	return &SwaggerExporter{
		app:          app,
		buildSwagger: &spec.Swagger{},
		log:          logger,
	}
}

func (s *SwaggerExporter) GenerateSwagger() error {
	s.buildSwagger.Swagger = "2.0"
	s.buildSwagger.Host = s.app.GetAttrs()["host"].GetS()

	s.buildSwagger.SwaggerProps.Info = &spec.Info{}
	s.buildSwagger.Paths = &spec.Paths{}
	s.buildSwagger.Paths.Paths = map[string]spec.PathItem{}
	s.buildSwagger.Definitions = spec.Definitions{}

	s.buildSwagger.SwaggerProps.Info.Title = s.app.LongName
	s.buildSwagger.SwaggerProps.Info.Description = s.app.GetAttrs()["description"].GetS()
	s.buildSwagger.SwaggerProps.Info.Version = s.app.GetAttrs()["version"].GetS()
	if s.buildSwagger.SwaggerProps.Info.Version == "" {
		s.buildSwagger.SwaggerProps.Info.Version = "0.0.0"
	}

	// parse type defs
	typeExporter := makeTypeExporter(s.log)
	typeExportError := typeExporter.populateTypes(s.app.GetTypes(), s.buildSwagger.Definitions)
	if typeExportError != nil {
		return typeExportError
	}

	// parse endpoints
	endpointExporter := makeEndpointExporter(typeExporter, s.log)
	for endpointName, endpoint := range s.app.Endpoints {
		err := endpointExporter.populateEndpoint(endpointName, endpoint, s.buildSwagger.Paths.Paths)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SwaggerExporter) SerializeOutput(mode string) ([]byte, error) {
	jsonSpec, err := s.buildSwagger.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if mode == "json" {
		return jsonSpec, nil
	}
	yamlSpec, err := yaml.JSONToYAML(jsonSpec)
	if err != nil {
		return nil, err
	}
	return yamlSpec, nil
}
