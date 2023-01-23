package configer

import (
	"fmt"
	"github.com/gozelle/toml"
	"reflect"
)

func Load(data []byte, p interface{}) (err error) {
	err = toml.Unmarshal(data, p)
	if err != nil {
		err = fmt.Errorf("unmarshal config error: %s", err)
		return
	}
	err = handle(p)
	if err != nil {
		return
	}
	return
}

func LoadFile(file string, p interface{}) (err error) {
	
	err = toml.UnmarshalFile(file, p)
	if err != nil {
		err = fmt.Errorf("unmarshal config file error: %s", err)
		return
	}
	err = handle(p)
	if err != nil {
		return
	}
	return
}

func handle(p interface{}) (err error) {
	
	v := reflect.ValueOf(p)
	fmt.Println(v)
	
	fmt.Println(v.Elem())
	
	for i := 0; i < v.Elem().Type().NumField(); i++ {
		field := v.Elem().Field(i)
		fmt.Println(field.Type().Name(), field.String())
	}
	
	return
}

func isPointerOrSlice() {

}

func getDefaultValue() (val string, ok bool) {
	return
}

func assignDefaultValue() {

}
