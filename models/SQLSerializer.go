/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   SQLSerializer.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/04 16:50:20 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/04 18:07:00 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/*
	Basic type settings
*/

var (
	GoToSQLType = map[reflect.Kind]string{
		reflect.Invalid: `invalid!`,
		reflect.Bool:    `boolean`,
		reflect.Int:     `integer`,
		reflect.String:  `text`,
		reflect.Struct:  `text`,
		reflect.Map:     `text`,
		reflect.Slice:   `text`,
	}

	ContainerTypes = map[string]reflect.Type{
		`timestamp`: reflect.TypeOf(time.Time{}),
	}
)

// Helper
func addPrefix(prefix string) string {
	if len(prefix) != 0 {
		return (prefix + `_`)
	} else {
		return (``)
	}
}

func IsContainerStructSQL(model interface{}) bool {
	modelType := reflect.TypeOf(model)
	for _, val := range ContainerTypes {
		if modelType == val {
			return true
		}
	}
	return false
}

func SerializeContainerStructSQL(model interface{}) string {
	modelType := reflect.TypeOf(model)
	if modelType == ContainerTypes[`timestamp`] {
		return `timestamp with time zone '` + model.(time.Time).Format(time.RFC3339) + `'`
	}
	return ``
}

func SerializeFieldsNameContainerSQLCreation(name string, model interface{}) string {
	modelType := reflect.TypeOf(model)
	if modelType == ContainerTypes[`timestamp`] {
		return `"` + name + `" timestamp with time zone not null default now()`
	}
	return ``
}

func SerializeFieldsNameContainerSQL(name string, model interface{}) string {
	return name
}

func RowsToStructContainer(rowMap map[string]interface{}, model reflect.Value, prefix string) (res interface{}, err error) {
	modelType := model.Type()
	if modelType == ContainerTypes[`timestamp`] {
		return rowMap[prefix], nil
	}
	return nil, errors.New(`Try to unserialize a container struct which is NOT a container struct`)
}

func SerializeStructSQL(model interface{}) string {
	var b bytes.Buffer

	if IsContainerStructSQL(model) {
		b.WriteString(SerializeContainerStructSQL(model))
		return b.String()
	}
	v := reflect.ValueOf(model)

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Struct:
			b.WriteString(SerializeStructSQL(v.Field(i).Interface()))
		case reflect.String:
			b.WriteString(SerializeSQL(v.Field(i).String()))
		case reflect.Int:
			b.WriteString(strconv.FormatInt(v.Field(i).Int(), 10))
		case reflect.Bool:
			b.WriteString(func(val bool) string {
				if val {
					return `TRUE`
				} else {
					return `FALSE`
				}
			}(v.Field(i).Bool()))
		case reflect.Slice:
			b.WriteString(`ARRAY[`)
			switch v.Field(i).Type().Elem().Kind() {
			case reflect.String:
				b.WriteString(SerializeStringArraySQL(v.Field(i).Interface().([]string)))
			case reflect.Struct:
				b.WriteString(SerializeStructArraySQL(v.Field(i).Interface()))
			default:
				fmt.Println(`SQLSerializer: Unrecognized type`)
			}
			b.WriteString(`]`)
		// Only one weird case of map FTM, DANGER!!!
		case reflect.Map:
			b.WriteString(`ARRAY[`)
			b.WriteString(SerializeMapSQL(v.Field(i).Interface().(map[string][]string)))
			b.WriteString(`]`)
		default:
			fmt.Println(`SQLSerializer: Unrecognized type`)
		}
		b.WriteString(`, `)
	}
	if b.Len() < 2 {
		return ``
	}
	return b.String()[:b.Len()-2]
}

/*
Intern only, use SerializeStructSQL instead except if you know what you do
*/
func SerializeStructArraySQL(models interface{}) string {
	var b bytes.Buffer
	v := reflect.ValueOf(models)
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			str, err := json.Marshal(v.Index(i).Interface())
			if err != nil {
				fmt.Println(`Error occured while serializing struct array: `, err.Error())
				continue
			}
			b.WriteString(SerializeByteArraySQL(str))
			b.WriteString(`, `)
		}
	}
	if b.Len() < 2 {
		return `''`
	}
	return b.String()[:b.Len()-2]
}

func SerializeFieldsNameSQLCreation(model interface{}) (result string) {
	return SerializeFieldsNameSQLCreationImpl(model, ``)
}

func SerializeFieldsNameSQLCreationImpl(model interface{}, prefix string) (result string) {
	var b bytes.Buffer

	if IsContainerStructSQL(model) {
		b.WriteString(SerializeFieldsNameContainerSQLCreation(prefix, model))
		return b.String()
	}
	v := reflect.ValueOf(model)

	for i := 0; i < v.NumField(); i++ {
		vType := v.Field(i).Kind()
		switch vType {
		case reflect.Struct:
			b.WriteString(SerializeFieldsNameSQLCreationImpl(
				v.Field(i).Interface(),
				addPrefix(prefix)+v.Type().Field(i).Tag.Get(`db`),
			))
		default:
			b.WriteString(`"`)
			b.WriteString(addPrefix(prefix))
			b.WriteString(v.Type().Field(i).Tag.Get(`db`))
			b.WriteString(`" `)
			switch vType {
			case reflect.Slice, reflect.Map:
				b.WriteString(GoToSQLType[v.Field(i).Type().Elem().Kind()])
				b.WriteString(`[]`)
			default:
				b.WriteString(GoToSQLType[vType])
				b.WriteString(` `)
				b.WriteString(v.Type().Field(i).Tag.Get(`modifiers`))
			}
		}
		b.WriteString(`, `)
	}
	if b.Len() < 2 {
		return ``
	}
	result = b.String()[:b.Len()-2]
	return
}

func SerializeFieldsNameSQL(model interface{}) string {
	return SerializeFieldsNameSQLImpl(model, ``)
}
func SerializeFieldsNameSQLImpl(model interface{}, prefix string) string {
	var b bytes.Buffer
	if IsContainerStructSQL(model) {
		b.WriteString(SerializeFieldsNameContainerSQL(prefix, model))
		return b.String()
	}

	v := reflect.ValueOf(model)

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Struct:
			b.WriteString(SerializeFieldsNameSQLImpl(
				v.Field(i).Interface(),
				addPrefix(prefix)+v.Type().Field(i).Tag.Get(`db`),
			))
		default:
			b.WriteString(addPrefix(prefix))
			b.WriteString(v.Type().Field(i).Tag.Get(`db`))
		}
		b.WriteString(`, `)
	}
	if b.Len() < 2 {
		return ``
	}
	return b.String()[:b.Len()-2]
}

func SerializeStructSQLIndex(model interface{}) string {
	return SerializeStructSQLIndexImpl(model, ``)
}

func addSuffix(s string, suffix string) string {
	if len(s) != 0 {
		return s + suffix
	}
	return ``
}
func SerializeStructSQLIndexImpl(model interface{}, prefix string) string {
	var b bytes.Buffer
	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Struct:
			b.WriteString(addSuffix(
				SerializeStructSQLIndexImpl(
					v.Field(i).Interface(),
					addPrefix(prefix)+v.Type().Field(i).Tag.Get(`db`),
				), ` || `),
			)
		default:
			if len(v.Type().Field(i).Tag.Get(`search`)) == 0 {
				continue
			}
			b.WriteString(`setweight(to_tsvector(coalesce(unaccent(`)
			switch v.Field(i).Kind() {
			case reflect.Slice, reflect.Map:
				b.WriteString(`array_to_string(`)
				b.WriteString(addPrefix(prefix) + v.Type().Field(i).Tag.Get(`db`))
				b.WriteString(`, ',')`)
			default:
				b.WriteString(addPrefix(prefix) + v.Type().Field(i).Tag.Get(`db`))
			}
			b.WriteString(`),`)
			b.WriteString(`'')), '`)
			b.WriteString(v.Type().Field(i).Tag.Get(`search`))
			b.WriteString(`') || `)
		}
	}
	if b.Len() < 3 {
		return ``
	}
	return b.String()[:b.Len()-3]
}

func escapeQuote(input []byte) (output []byte) {
	return bytes.Replace(input, []byte("'"), []byte("''"), -1)
}
func SerializeByteArraySQL(input []byte) (output string) {
	var b bytes.Buffer
	b.WriteByte('\'')
	b.WriteString(string(escapeQuote(input)))
	b.WriteByte('\'')
	return b.String()
}
func SerializeSQL(input string) (output string) {
	return SerializeByteArraySQL([]byte(input))
}

func SerializeStringArraySQL(input []string) string {
	var b bytes.Buffer
	for _, val := range input {
		b.WriteString(SerializeSQL(val))
		b.WriteString(`, `)
	}
	if b.Len() < 2 {
		return `''`
	}
	return b.String()[:b.Len()-2]
}

func SerializeMapSQL(input map[string][]string) string {
	var b bytes.Buffer
	for key, val := range input {
		var bCurrent bytes.Buffer
		bCurrent.WriteString(`{"`)
		bCurrent.WriteString(key)
		bCurrent.WriteString(`":`)
		jsonVal, err := json.Marshal(val)
		if err != nil {
			continue
		}
		bCurrent.WriteString(string(jsonVal))
		bCurrent.WriteByte('}')
		b.WriteString(SerializeSQL(bCurrent.String()))
		b.WriteString(`, `)
	}
	if b.Len() < 2 {
		return `''`
	}
	return b.String()[:b.Len()-2]
}

/*
Ok so PSQL... if there is no space or comma in any elements, no double quotes to surround. Take this JSON !
*/
func MaybeNoQuoteJSONToSliceString(input string) (output []string) {
	for i := 0; i < len(input); {
		if input[i] == '"' {
			next := strings.Index(input[i+1:], `"`)
			if next != -1 {
				output = append(output, input[i+1:next+i])
				i = next + i + 2
			} else {
				break
			}
		} else if input[i] == ',' {
			i++
		} else {
			next := strings.Index(input[i+1:], `,`)
			if next != -1 {
				output = append(output, input[i:next+i+1])
				i = next + i + 2
			} else {
				output = append(output, input[i:])
				break
			}
		}
	}
	return
}

func SQLArrayToSliceString(input []byte) (output []string) {
	input[0] = '['
	input[len(input)-1] = ']'
	err := json.Unmarshal(input, &output)
	if err != nil {
		/*
			Ok so PSQL... if there is no space or comma in any elements, no double quotes to surround. Take this JSON !
		*/
		return MaybeNoQuoteJSONToSliceString(string(input[1 : len(input)-2]))
	}
	return
}

func SQLArrayToMapString(input []byte) (output map[string][]string) {
	if len(input) < 4 {
		return map[string][]string{}
	}
	input[0] = '['
	input[len(input)-1] = ']'
	var jsonKeyVal []string
	err := json.Unmarshal(input, &jsonKeyVal)
	if err != nil {
		fmt.Println(`Error occured while converting SQL array to map string (first slice): `, err.Error())
		return map[string][]string{}
	}
	for _, val := range jsonKeyVal {
		err = json.Unmarshal([]byte(val), &output)
		if err != nil {
			fmt.Println(`Error occured while converting SQL array to map string: `, err.Error())
			return map[string][]string{}
		}
	}
	return
}

func SQLArrayToMapStruct(input []byte) (output []interface{}) {
	jsonStrings := SQLArrayToSliceString(input)
	output = make([]interface{}, len(jsonStrings))
	for i, val := range jsonStrings {
		err := json.Unmarshal([]byte(val), &output[i])
		if err != nil {
			fmt.Println(`Error occured while converting SQL array to slice struct: `, err.Error())
		}
	}
	return
}

func JSONToSliceStructValue(input []byte, outParam reflect.Value) (output reflect.Value, err error) {
	jsonStrings := SQLArrayToSliceString(input)
	elemType := outParam.Type().Elem()
	output = reflect.ValueOf(outParam.Interface())
	for _, val := range jsonStrings {
		if len(val) == 0 {
			continue
		}
		f := reflect.New(elemType).Interface()
		err := json.Unmarshal([]byte(val), f)
		if err != nil {
			fmt.Println(`Error occured while converting SQL array to slice struct: `, err.Error())
			fmt.Println(val)
			return outParam, err
		}
		output = reflect.Append(output, reflect.Indirect(reflect.ValueOf(f)))
	}
	return
}

func RowToStruct(row *sqlx.Row, model reflect.Value) (res interface{}, err error) {
	rowMap := make(map[string]interface{})
	err = row.MapScan(rowMap)
	if err != nil {
		return
	}
	return RowToStructImpl(rowMap, model, ``)
}

func RowsToStruct(rows *sqlx.Rows, model reflect.Value) (res interface{}, err error) {
	rowsMap := make(map[string]interface{})
	err = rows.MapScan(rowsMap)
	if err != nil {
		return
	}
	return RowToStructImpl(rowsMap, model, ``)
}

func RowToStructImpl(rowMap map[string]interface{}, model reflect.Value, prefix string) (res interface{}, err error) {
	v := reflect.Indirect(model)
	if IsContainerStructSQL(v.Interface()) {
		return RowsToStructContainer(rowMap, v, prefix)
	}
	for i := 0; i < v.NumField(); i++ {
		val, ok := rowMap[addPrefix(prefix)+v.Type().Field(i).Tag.Get(`db`)] // Ok can be false, if a struct this is NOT an error
		if !ok && v.Field(i).Kind() != reflect.Struct {
			continue
		}
		switch v.Field(i).Kind() {
		case reflect.Struct:
			subModel, err := RowToStructImpl(rowMap, v.Field(i).Addr(), addPrefix(prefix)+v.Type().Field(i).Tag.Get(`db`))
			if err != nil {
				fmt.Println(`Error setting sub struct: `, err.Error())
				continue
			}
			v.Field(i).Set(reflect.ValueOf(subModel))
		case reflect.String:
			v.Field(i).SetString(val.(string)) // BruteForce
		case reflect.Int, reflect.Int64:
			v.Field(i).SetInt(valToInt64(val))
		case reflect.Float32, reflect.Float64:
			v.Field(i).SetFloat(valToFloat64(val))
		case reflect.Map:
			v.Field(i).Set(reflect.ValueOf(SQLArrayToMapString(reflect.ValueOf(val).Interface().([]byte))))
		case reflect.Slice:
			switch v.Field(i).Type().Elem().Kind() {
			case reflect.String:
				v.Field(i).Set(reflect.ValueOf(SQLArrayToSliceString(reflect.ValueOf(val).Interface().([]byte))))
			case reflect.Struct:
				structVal, err := JSONToSliceStructValue(reflect.ValueOf(val).Interface().([]byte), v.Field(i))
				if err != nil {
					fmt.Println(`Error setting slice struct: `, err.Error())
				}
				v.Field(i).Set(structVal)
			default:
				fmt.Println(`SQLSerializer: Unrecognized type`)
				continue
			}
		}
	}
	return v.Interface(), nil
}

func valToInt64(val interface{}) (res int64) {
	if val == nil {
		return 0
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Int64, reflect.Int:
		return v.Interface().(int64)
	case reflect.Float32, reflect.Float64:
		return int64(valToFloat64(val))
	case reflect.String:
		res, _ = strconv.ParseInt(v.Interface().(string), 10, 64)
	case reflect.Bool:
		if val.(bool) == true {
			return 1
		} else {
			return 0
		}
	default:
		fmt.Println(`Can't convert this val to int:`, val)
	}
	return 0
}

func valToFloat64(val interface{}) (res float64) {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Interface().(float64)
	case reflect.String:
		res, _ = strconv.ParseFloat(v.Interface().(string), 64)
	default:
		fmt.Println(`Can't convert this val to float:`, val)
	}
	return 0
}
