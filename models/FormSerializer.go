/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   FormSerializer.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/18 20:47:08 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/25 17:38:09 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package models

import (
	"net/url"
	"reflect"
	"strconv"
)

func FormToStruct(model interface{}, data *url.Values) (result interface{}) {
	v := reflect.Indirect(reflect.ValueOf(model))
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Struct:
			v.Field(i).Set(reflect.ValueOf(FormToStruct(v.Field(i).Addr().Interface(), data)))
		case reflect.Int:
			tag := v.Type().Field(i).Tag.Get(`db`)
			if len(tag) == 0 {
				continue
			}
			intValue, err := strconv.ParseInt(data.Get(tag), 10, 64)
			if err != nil {
				// fmt.Println(`Can't convert this string to valid int`)
				continue
			}
			v.Field(i).SetInt(intValue)
		default:
			tag := v.Type().Field(i).Tag.Get(`db`)
			if len(tag) == 0 {
				continue
			}
			v.Field(i).Set(reflect.ValueOf(data.Get(tag)))
		}
	}
	return v.Interface()
}
