package service

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

/*
	fill cfg with env variables starting with prefix

	example, given a struct like:

	type LogConfig struct {
		level string `default:"info"`
		type  string `default:"fmt"`
	}

	lc:= LogConfig{}

	LoadIntoWithPrefix(&lc, "LOG")

	loads LOG_LEVEL from env into lc.level
	loads LOG_TYPE from env into lc.type

	If a variable is not in env, the default tag is parsed and assigned.
	types are handled correctly for string, int, float, bool

	Return false is not all struct fields could be assigned (missing defaults + missing variables)


*/
func LoadIntoWithPrefix(cfg interface{}, prefix string) bool {
	viper.AutomaticEnv()
	return fill(cfg, prefix)
}

func LoadInto(cfg interface{}) bool {
	return LoadIntoWithPrefix(cfg, "")
}

//given a struct pointer with annotated fields, tries to fill it from viper
//annotations can have a default key for the default value (if var not set).
//return true if everything is ok, else return false (missing required vars, conversions error etc)
//can panic if f is not settable or if the tagged field is not public
func fill(f interface{}, prefix string) bool {

	ok := true
	if prefix != "" {
		prefix = prefix + "_"
	}

	//get type of *F
	tpf := reflect.TypeOf(f)
	//dereference to get type of F
	tf := tpf.Elem()
	vf := reflect.ValueOf(f).Elem()
	//check that it is a struct
	if tf.Kind() != reflect.Struct {
		return false
	}
	//loop over fields
	for i := 0; i < tf.NumField(); i++ {
		fl := tf.Field(i)
		vname := strings.ToUpper(prefix + fl.Name)
		def, hasdef := fl.Tag.Lookup("default")
		if hasdef {
			viper.SetDefault(vname, def)
		}
		if !viper.IsSet(vname) {
			//no default and not in env!
			ok = false
			break
		}
		switch fl.Type.Kind() {
		case reflect.String:
			vf.Field(i).SetString(viper.GetString(vname))
		case reflect.Bool:
			vf.Field(i).SetBool(viper.GetBool(vname))
		case reflect.Int:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			vf.Field(i).SetInt(viper.GetInt64(vname))
		case reflect.Uint:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			vf.Field(i).SetUint(viper.GetUint64(vname))
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			vf.Field(i).SetFloat(viper.GetFloat64(vname))
		}
	}
	return ok
}
