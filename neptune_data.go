// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-02 13:56:03.549666 -0500 CDT m=+0.005104325
package aws_resources

import (
	"reflect"
)

func (i *NeptuneType) CallAWS() (map[string]reflect.Value, error) {

	outcome := map[string]reflect.Value{}

	if i.methodName == "" {
		return nil, NoServiceMethodName
	}

	instance, err := typeRegistry.Get("neptune", i.inputName)
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
