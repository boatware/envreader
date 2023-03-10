package envreader

import (
	"os"
	"strconv"
	"strings"
)

// ReadString reads the environment variable with the given key.
// Example:
//
//	// .env
//	FOO="bar"
//
//	// main.go
//	var foo string
//	foo = env.ReadString("FOO")
//
//	// foo = "bar"
func ReadString(key string) string {
	return os.Getenv(key)
}

// ReadInt reads the environment variable with the given key and converts it to an int.
// Example:
//
//	// .env
//	FOO="1"
//
//	// main.go
//	var foo int
//	foo = env.ReadInt("FOO")
//
//	// foo = 1
func ReadInt(key string) int {
	value := ReadString(key)
	if value == "" {
		return 0
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return intVal
}

// ReadBool reads the environment variable with the given key and converts it to a bool.
// Example:
//
//	// .env
//	FOO="true"
//
//	// main.go
//	var foo bool
//	foo = env.ReadBool("FOO")
//
//	// foo = true
func ReadBool(key string) bool {
	value := ReadString(key)
	if value == "" {
		return false
	}

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return boolVal
}

// ReadFloat reads the environment variable with the given key and converts it to a float.
// Example:
//
//	// .env
//	FOO="1.1"
//
//	// main.go
//	var foo float64
//	foo = env.ReadFloat("FOO")
//
//	// foo = 1.1
func ReadFloat(key string) float64 {
	value := ReadString(key)
	if value == "" {
		return 0
	}

	floatVal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return floatVal
}

// ReadStringSlice reads the environment variable with the given key and converts it to a slice of strings.
// Example:
//
//	// .env
//	FOO="foo,bar,baz"
//
//	// main.go
//	var foo []string
//	foo = env.ReadStringSlice("FOO")
//
//	// foo = ["foo", "bar", "baz"]
func ReadStringSlice(key string) []string {
	value := ReadString(key)
	if value == "" {
		return []string{}
	}

	array := strings.Split(value, ",")

	return array
}

// ReadIntSlice reads the environment variable with the given key and converts it to a slice of ints.
// Example:
//
//	// .env
//	FOO="1,2,3"
//
//	// main.go
//	var foo []int
//	foo = env.ReadIntSlice("FOO")
//
//	// foo = [1, 2, 3]
func ReadIntSlice(key string) []int {
	value := ReadString(key)
	if value == "" {
		return []int{}
	}

	array := strings.Split(value, ",")
	intArray := []int{}

	for _, val := range array {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			continue
		}

		intArray = append(intArray, intVal)
	}

	return intArray
}

// ReadBoolSlice reads the environment variable with the given key and converts it to a slice of bools.
// Example:
//
//	// .env
//	FOO="true,false,true"
//
//	// main.go
//	var foo []bool
//	foo = env.ReadBoolSlice("FOO")
//
//	// foo = [true, false, true]
func ReadBoolSlice(key string) []bool {
	value := ReadString(key)
	if value == "" {
		return []bool{}
	}

	array := strings.Split(value, ",")
	var boolArray []bool

	for _, val := range array {
		boolVal, err := strconv.ParseBool(val)
		if err != nil {
			continue
		}

		boolArray = append(boolArray, boolVal)
	}

	return boolArray
}

// ReadFloatSlice reads the environment variable with the given key and converts it to a slice of floats.
// Example:
//
//	// .env
//	FOO="1.1,2.2,3.3"
//
//	// main.go
//	var foo []float64
//	foo = env.ReadFloatSlice("FOO")
//
//	// foo = [1.1, 2.2, 3.3]
func ReadFloatSlice(key string) []float64 {
	value := ReadString(key)
	if value == "" {
		return []float64{}
	}

	array := strings.Split(value, ",")
	var floatArray []float64

	for _, val := range array {
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			continue
		}

		floatArray = append(floatArray, floatVal)
	}

	return floatArray
}

// ReadStringMap reads the environment variable with the given key and converts it to a map of strings.
// Example:
//
//	// .env
//	FOO="foo=bar,baz=qux"
//
//	// main.go
//	var foo map[string]string
//	foo = env.ReadStringMap("FOO")
//
//	// foo = map[string]string{"foo": "bar", "baz": "qux"}
func ReadStringMap(key string) map[string]string {
	value := ReadString(key)
	if value == "" {
		return map[string]string{}
	}

	array := strings.Split(value, ",")
	mapArray := map[string]string{}

	for _, val := range array {
		keyVal := strings.Split(val, "=")
		if len(keyVal) != 2 {
			continue
		}

		mapArray[keyVal[0]] = keyVal[1]
	}

	return mapArray
}

// ReadIntMap reads the environment variable with the given key and converts it to a map of ints.
// Example:
//
//	// .env
//	FOO="foo=1,baz=2"
//
//	// main.go
//	var foo map[string]int
//	foo = env.ReadIntMap("FOO")
//
//	// foo = map[string]int{"foo": 1, "baz": 2}
func ReadIntMap(key string) map[string]int {
	value := ReadString(key)
	if value == "" {
		return map[string]int{}
	}

	array := strings.Split(value, ",")
	mapArray := map[string]int{}

	for _, val := range array {
		keyVal := strings.Split(val, "=")
		if len(keyVal) != 2 {
			continue
		}

		intVal, err := strconv.Atoi(keyVal[1])
		if err != nil {
			continue
		}

		mapArray[keyVal[0]] = intVal
	}

	return mapArray
}

// ReadBoolMap reads the environment variable with the given key and converts it to a map of bools.
// Example:
//
//	// .env
//	FOO="foo=true,baz=false"
//
//	// main.go
//	var foo map[string]bool
//	foo = env.ReadBoolMap("FOO")
//
//	// foo = map[string]bool{"foo": true, "baz": false}
func ReadBoolMap(key string) map[string]bool {
	value := ReadString(key)
	if value == "" {
		return map[string]bool{}
	}

	array := strings.Split(value, ",")
	mapArray := map[string]bool{}

	for _, val := range array {
		keyVal := strings.Split(val, "=")
		if len(keyVal) != 2 {
			continue
		}

		boolVal, err := strconv.ParseBool(keyVal[1])
		if err != nil {
			continue
		}

		mapArray[keyVal[0]] = boolVal
	}

	return mapArray
}

// ReadFloatMap reads the environment variable with the given key and converts it to a map of floats.
// Example:
//
//	// .env
//	FOO="foo=1.1,baz=2.2"
//
//	// main.go
//	var foo map[string]float64
//	foo = env.ReadFloatMap("FOO")
//
//	// foo = map[string]float64{"foo": 1.1, "baz": 2.2}
func ReadFloatMap(key string) map[string]float64 {
	value := ReadString(key)
	if value == "" {
		return map[string]float64{}
	}

	array := strings.Split(value, ",")
	mapArray := map[string]float64{}

	for _, val := range array {
		keyVal := strings.Split(val, "=")
		if len(keyVal) != 2 {
			continue
		}

		floatVal, err := strconv.ParseFloat(keyVal[1], 64)
		if err != nil {
			continue
		}

		mapArray[keyVal[0]] = floatVal
	}

	return mapArray
}
