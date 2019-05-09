// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-09 13:01:20.611969172 -0500 CDT m=+0.000147082
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
)

type LexRuntimeServiceType struct {
	service      *lexruntimeservice.LexRuntimeService
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type LexRuntimeServiceTypeConfig struct {
	resourceType string
}

func LexRuntimeServiceFactory(cfg aws.Config) Factory {
	i := new(LexRuntimeServiceType)

	i.SetService(cfg)

	return i
}

func (i *LexRuntimeServiceType) SetPartialName() {
	// "AWS::LexRuntimeService::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::LexRuntimeService::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *LexRuntimeServiceType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *LexRuntimeServiceType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *LexRuntimeServiceType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *LexRuntimeServiceType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *LexRuntimeServiceType) Configure(param interface{}) error {
	config, ok := param.(LexRuntimeServiceTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (LexRuntimeServiceTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *LexRuntimeServiceType) SetService(cfg aws.Config) {
	srv := lexruntimeservice.New(cfg)

	i.service = srv
}

func (i *LexRuntimeServiceType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("lexruntimeservice", i.inputName)
	if err != nil {
		// We can ignore this kind of errors because there is not resources by the
		// i.inputName
		//log.Println(err)
		return
	}

	method := reflect.ValueOf(i.service).MethodByName(i.methodName)
	called := method.Call([]reflect.Value{reflect.ValueOf(instance)})

	send := reflect.Indirect(called[0]).MethodByName("Send")
	calledSend := send.Call([]reflect.Value{})

	res := calledSend[0]

	fmt.Printf("%v\n", res)
}

func (i *LexRuntimeServiceType) GetResources() {}

func (i *LexRuntimeServiceType) GetResourcesDetail() {}