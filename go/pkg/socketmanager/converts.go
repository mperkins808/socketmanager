package socketmanager

import "fmt"

// converts a value to an interface
func (r ArbResult) Interface() (interface{}, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	return r.Value, nil
}

// converts a value to an interface
func (r ArbResult) InterfaceArray() ([]interface{}, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	ifc, ok := r.Value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("value is not a string")
	}
	return ifc, nil
}

// converts a value to an interface
func (r ArbResult) InterfaceMap() (map[string]interface{}, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	ifc, ok := r.Value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("value is not a string")
	}
	return ifc, nil
}

// converts a value to a string
func (r ArbResult) String() (string, error) {
	if r.Err != nil {
		return "", r.Err
	}
	str, ok := r.Value.(string)
	if !ok {
		return "", fmt.Errorf("value is not a string")
	}
	return str, nil
}

// converts a value to an int
func (r ArbResult) Int() (int, error) {
	if r.Err != nil {
		return 0, r.Err
	}
	result, ok := r.Value.(int)
	if !ok {
		return 0, fmt.Errorf("value is not an int")
	}
	return result, nil
}

// converts a value to an int32
func (r ArbResult) Int32() (int32, error) {
	if r.Err != nil {
		return 0, r.Err
	}
	result, ok := r.Value.(int32)
	if !ok {
		return 0, fmt.Errorf("value is not an int32")
	}
	return result, nil
}

// converts a value to an int64
func (r ArbResult) Int64() (int64, error) {
	if r.Err != nil {
		return 0, r.Err
	}
	result, ok := r.Value.(int64)
	if !ok {
		return 0, fmt.Errorf("value is not an int64")
	}
	return result, nil
}

// converts a value to a float32
func (r ArbResult) Float32() (float32, error) {
	if r.Err != nil {
		return 0, r.Err
	}
	result, ok := r.Value.(float32)
	if !ok {
		return 0, fmt.Errorf("value is not a float32")
	}
	return result, nil
}

// converts a value to a float64
func (r ArbResult) Float64() (float64, error) {
	if r.Err != nil {
		return 0, r.Err
	}
	result, ok := r.Value.(float64)
	if !ok {
		return 0, fmt.Errorf("value is not a float64")
	}
	return result, nil
}

////////////////////////////
// ARRAYS
////////////////////////////

// converts a value to a string
func (r ArbResult) StringArray() ([]string, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	str, ok := r.Value.([]string)
	if !ok {
		return nil, fmt.Errorf("value is not a string array")
	}
	return str, nil
}

// converts a value to a int array
func (r ArbResult) IntArray() ([]int, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	result, ok := r.Value.([]int)
	if !ok {
		return nil, fmt.Errorf("value is not an int array")
	}
	return result, nil
}

// converts a value to an int32 array
func (r ArbResult) Int32Array() ([]int32, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	result, ok := r.Value.([]int32)
	if !ok {
		return nil, fmt.Errorf("value is not an int32 array")
	}
	return result, nil
}

// converts a value to an int64
func (r ArbResult) Int64Array() ([]int64, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	result, ok := r.Value.([]int64)
	if !ok {
		return nil, fmt.Errorf("value is not an int64 array")
	}
	return result, nil
}

// converts a value to a float32 array
func (r ArbResult) Float32Array() ([]float32, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	result, ok := r.Value.([]float32)
	if !ok {
		return nil, fmt.Errorf("value is not a float32 array")
	}
	return result, nil
}

// converts a value to a float64 array
func (r ArbResult) Float64Array() ([]float64, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	result, ok := r.Value.([]float64)
	if !ok {
		return nil, fmt.Errorf("value is not a float64 array")
	}
	return result, nil
}

// converts a value to a bool
func (r ArbResult) Bool() (bool, error) {
	if r.Err != nil {
		return false, r.Err
	}
	result, ok := r.Value.(bool)
	if !ok {
		return false, fmt.Errorf("value is not a float64 array")
	}
	return result, nil
}

// converts a value to a bool array
func (r ArbResult) BoolArray() ([]bool, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	result, ok := r.Value.([]bool)
	if !ok {
		return nil, fmt.Errorf("value is not a float64 array")
	}
	return result, nil
}
