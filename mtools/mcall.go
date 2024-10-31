package mtools

import (
	"errors"
	"reflect"
	"sync"
)

func MCall(inputs interface{}, callFunc func(interface{})) (err error) {
	var (
		wg          sync.WaitGroup
		workerLimit chan struct{}
	)

	value := reflect.ValueOf(inputs)
	if value.Kind() != reflect.Slice {
		return errors.New("inputs must be slice")
	}

	count := value.Len()
	if count == 0 {
		return
	}

	workerLimit = make(chan struct{}, 10)
	for i := 0; i < count; i++ {
		wg.Add(1)
		workerLimit <- struct{}{} // 获取令牌
		go func(input interface{}) {
			defer func() {
				wg.Done()
				<-workerLimit // 释放令牌
			}()
			callFunc(input)
		}(value.Index(i).Interface())
	}
	wg.Wait()
	close(workerLimit)
	return
}

var (
	InputsMustBeSliceErr         = errors.New("inputs must be slice")
	OutputsMustBeSlicePointerErr = errors.New("outputs must be slice pointer")
	OutputsLenErr                = errors.New("outputs len error")
)

func MGet(inputs, outputs interface{}, callFunc func(interface{}) interface{}) (err error) {
	var (
		wg          sync.WaitGroup
		workerLimit chan struct{}
	)

	inputValue := reflect.ValueOf(inputs)
	if inputValue.Kind() != reflect.Slice {
		return InputsMustBeSliceErr
	}
	outputValuePtr := reflect.ValueOf(outputs)
	outputValue := outputValuePtr.Elem()
	if outputValuePtr.Kind() != reflect.Ptr || outputValue.Kind() != reflect.Slice {
		return OutputsMustBeSlicePointerErr
	}

	count := inputValue.Len()
	if count == 0 {
		return
	}
	outputCount := outputValue.Len()
	if count != outputCount {
		return OutputsLenErr
	}
	outputType := outputValue.Type().Elem()
	workerLimit = make(chan struct{}, 10)
	for i := 0; i < count; i++ {
		wg.Add(1)
		workerLimit <- struct{}{} // 获取令牌
		go func(index int, input interface{}) {
			defer func() {
				wg.Done()
				<-workerLimit // 释放令牌
			}()
			item := callFunc(input)
			if outputType != reflect.TypeOf(item) {
				outputValue.Index(index).Set(reflect.Zero(outputType))
			} else {
				outputValue.Index(index).Set(reflect.ValueOf(item))
			}
		}(i, inputValue.Index(i).Interface())
	}
	wg.Wait()
	close(workerLimit)
	return
}
