// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-10 14:36:11.053646913 -0500 CDT m=+0.000159207
package aws_resources

import (
	"reflect"
)

func (i *KinesisAnalyticsV2Type) CallAWS() (map[string]reflect.Value, error) {

	outcome := map[string]reflect.Value{}

	if i.methodName == "" {
		return nil, NoServiceMethodName
	}

	instance, err := typeRegistry.Get("kinesisanalyticsv2", i.inputName)
	if err != nil {
		return nil, err
	}

	method := reflect.ValueOf(i.service).MethodByName(i.methodName)
	called := method.Call([]reflect.Value{reflect.ValueOf(instance)})

	send := reflect.Indirect(called[0]).MethodByName("Send")
	calledSend := send.Call([]reflect.Value{})
	outcome[i.resourceType] = calledSend[0]

	return outcome, nil
}
