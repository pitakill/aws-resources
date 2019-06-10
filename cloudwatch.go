// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-10 14:36:11.053646913 -0500 CDT m=+0.000159207
package aws_resources

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type CloudWatchType struct {
	service      *cloudwatch.CloudWatch
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func CloudWatchFactory(cfg aws.Config) Factory {
	i := new(CloudWatchType)

	i.SetService(cfg)

	return i
}

func (i *CloudWatchType) SetPartialName() {
	// "AWS::CloudWatch::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::CloudWatch::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *CloudWatchType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "cloudwatch", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *CloudWatchType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "cloudwatch", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *CloudWatchType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "cloudwatch", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *CloudWatchType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *CloudWatchType) Configure(param interface{}) error {
	config, ok := param.(TypeConfig)
	if !ok {
		return errors.New("config is not a valid param (TypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *CloudWatchType) SetService(cfg aws.Config) {
	srv := cloudwatch.New(cfg)

	i.service = srv
}

func (i *CloudWatchType) GetServices() (reflect.Value, error) {
	if i.methodName == "" {
		return reflect.Value{}, errors.New("method can't be an empty string")
	}

	instance, err := typeRegistry.Get("cloudwatch", i.inputName)
	if err != nil {
		// We can ignore this kind of errors because there is not resources by the
		// i.inputName
		log.Println(err)
		// This seems odd, we can do better
		// We need to think about this again
		return reflect.Value{}, nil
	}

	method := reflect.ValueOf(i.service).MethodByName(i.methodName)
	called := method.Call([]reflect.Value{reflect.ValueOf(instance)})

	send := reflect.Indirect(called[0]).MethodByName("Send")
	calledSend := send.Call([]reflect.Value{})

	return calledSend[0], nil
}

func (i *CloudWatchType) GetResources() error { return nil }

func (i *CloudWatchType) GetResourcesDetail() ([]reflect.Value, error) {
	return []reflect.Value{}, nil
}
