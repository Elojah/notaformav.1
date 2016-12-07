/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   insert.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/12 13:22:30 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/04 18:55:42 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */
package psql

import (
	"bytes"
	"fmt"
	"github.com/go-server/models"
	"reflect"
	"strings"
)

func InsertRow(table string, model interface{}) (err error) {
	query := `
		INSERT INTO ` + table + `
		(` + models.SerializeFieldsNameSQL(model) + `)
		VALUES
		(` + models.SerializeStructSQL(model) + `)
	`
	fmt.Println(query)
	_, err = db.Exec(query)
	return
}

func ConditionRowEquals(model interface{}, uniqueRows []string) string {
	if len(uniqueRows) == 0 {
		return (``)
	}
	var b bytes.Buffer
	b.WriteString(`WHERE `)
	v := reflect.ValueOf(model)
	for _, val := range uniqueRows {
		modelVal, ok := v.Type().FieldByName(val)
		if !ok {
			fmt.Println(val, ` is not a valid field for this model. Can't condition on this.`)
			continue
		}
		b.WriteString(modelVal.Tag.Get(`db`))
		b.WriteString(`=`)
		b.WriteString(models.SerializeSQL(v.FieldByName(val).String()))
		b.WriteString(` AND `)
	}
	return b.String()[:b.Len()-5]
}

func InsertRowUnique(table string, model interface{}, uniqueRows []string) bool {
	// Get first key for the return
	var firstKey string
	for _, firstKey = range uniqueRows {
		break
	}
	query := `
				INSERT INTO ` + table + `
				(` + models.SerializeFieldsNameSQL(model) + `)
					SELECT ` + models.SerializeStructSQL(model) + `
					WHERE
					NOT EXISTS (
						SELECT 1 FROM ` + table + `
						` + ConditionRowEquals(model, uniqueRows) + `
					)
				RETURNING ` + firstKey + `;`
	var id string
	fmt.Println(query)
	err := db.Get(&id, query)
	if err != nil || len(id) == 0 {
		fmt.Println(err)
		return false
	}
	return true
}

func UpdateAverage(table string, field string, id string) bool {
	query := `UPDATE ` + table + ` SET ` + field + `=(` +
		` SELECT AVG(` + field + `) from comment_` + table + ` WHERE parentid='` + id + `')` +
		`WHERE id='` + id + `'`
	fmt.Println(query)
	row := db.QueryRowx(query)
	err := row.Err()
	if err != nil || row == nil {
		return false
	}
	return true
}

func replaceWhereToSet(input string) string {
	return strings.NewReplacer(`WHERE`, `SET`).Replace(input)
}

/*
TO_SEE: Maybe change newValues in a struct ?
*/
func UpdateRow(table string, queryOptions models.QueryModel, newValues map[string]string) bool {
	query := `UPDATE ` + table + replaceWhereToSet(SerializeFilter(newValues)) + SerializeFilter(queryOptions.Filter)
	fmt.Println(query)
	row := db.QueryRowx(query)
	err := row.Err()
	if err != nil || row == nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func UpdateRowNoQuote(table string, queryOptions models.QueryModel, newValues map[string]string) bool {
	query := `UPDATE ` + table + replaceWhereToSet(SerializeFilterNoQuote(newValues)) + SerializeFilter(queryOptions.Filter)
	fmt.Println(query)
	row := db.QueryRowx(query)
	err := row.Err()
	if err != nil || row == nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
