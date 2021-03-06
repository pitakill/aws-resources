// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package aws_resources

import (
  "reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// Type Factory
type Factory interface {
	Configure(interface{}) error
	GetServices() (reflect.Value, error)
	GetResources() error
	GetResourcesDetail() ([]reflect.Value, error)
	SetService(aws.Config)
}

type Info func(aws.Config) Factory

type TypeConfig struct {
	resourceType string
}


var Relations = map[string]Info{
	"cloudformation": CloudFormationFactory,
	{{- range $index, $typ := .Resources }}
	"{{ ToLower $typ.Name }}": {{ $typ.Name }}Factory,
	{{- end }}
}
