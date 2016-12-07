/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   data.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/14 15:27:47 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 17:30:25 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package psql

import (
	"bytes"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/go-server/models"
	"github.com/jmoiron/sqlx"
	"time"
)

func SerializePage(pages map[string]string) string {
	if len(pages) == 0 {
		return (``)
	}
	var b bytes.Buffer
	if val, ok := pages["order"]; ok && val != "" {
		b.WriteString(` ORDER BY `)
		b.WriteString(val)
	}
	if val, ok := pages["offset"]; ok && val != "" {
		b.WriteString(` OFFSET `)
		b.WriteString(val)
	}
	if val, ok := pages["limit"]; ok && val != "" {
		b.WriteString(` LIMIT `)
		b.WriteString(val)
	}
	return b.String()
}

func SerializeFilter(filters map[string]string) string {
	if len(filters) == 0 {
		return (``)
	}
	var b bytes.Buffer
	b.WriteString(` WHERE `)
	for key, val := range filters {
		b.WriteString(key)
		b.WriteString(`=`)
		b.WriteString(models.SerializeSQL(val))
		b.WriteString(` AND `)
	}
	return b.String()[:b.Len()-5]
}

func SerializeFilterNoQuote(filters map[string]string) string {
	if len(filters) == 0 {
		return (``)
	}
	var b bytes.Buffer
	b.WriteString(` WHERE `)
	for key, val := range filters {
		b.WriteString(key)
		b.WriteString(`=`)
		b.WriteString(val)
		b.WriteString(` AND `)
	}
	return b.String()[:b.Len()-5]
}

func SerializeFilterArray(filters map[string]string) string {
	if len(filters) == 0 {
		return (``)
	}
	var b bytes.Buffer
	b.WriteString(` WHERE `)
	for key, val := range filters {
		b.WriteString(` ARRAY[`)
		b.WriteString(val)
		b.WriteString(`] && `)
		b.WriteString(key)
		b.WriteString(` AND `)
	}
	return b.String()[:b.Len()-5]
}

func GetRow(table string, model interface{}, queryOptions models.QueryModel) (row *sqlx.Row, err error) {
	query := `SELECT ` + models.SerializeFieldsNameSQL(model) + ` FROM ` + table + SerializeFilter(queryOptions.Filter)
	fmt.Println(query)
	row = db.QueryRowx(query)
	err = row.Err()
	return
}

func GetRowsSpec(table string, model interface{}, queryOptions models.QueryModel, spec models.Converter) (rows *sqlx.Rows, err error) {
	query := `SELECT ` + spec(models.SerializeFieldsNameSQL(model)) +
		`, COUNT(*) OVER() AS totalcount FROM ` +
		table + SerializeFilter(queryOptions.Filter) +
		SerializePage(queryOptions.Page)
	fmt.Println(query)
	start := time.Now()
	rows, err = db.Queryx(query)
	elapsed := time.Since(start)
	fmt.Println("Request took ", elapsed)
	return
}

func RowExist(table string, model interface{}, queryOptions models.QueryModel) bool {
	query := `SELECT 1 FROM ` + table + SerializeFilter(queryOptions.Filter)
	fmt.Println(query)
	nResult := 0
	err := db.Get(&nResult, query)
	if err == nil || nResult == 1 {
		return true
	}
	return false
}

func GetRows(table string, model interface{}, queryOptions models.QueryModel) (rows *sqlx.Rows, err error) {
	query := `SELECT ` + models.SerializeFieldsNameSQL(model) +
		`, COUNT(*) OVER() AS totalcount FROM ` +
		table + SerializeFilter(queryOptions.Filter) +
		SerializePage(queryOptions.Page)
	fmt.Println(query)
	start := time.Now()
	rows, err = db.Queryx(query)
	elapsed := time.Since(start)
	fmt.Println("Request took ", elapsed)
	return
}

func GetRowsWithSubRequest(table string, model interface{}, queryOptions models.QueryModel, subreq string) (rows *sqlx.Rows, err error) {
	query := `SELECT ` + models.SerializeFieldsNameSQL(model) + ` ,` + subreq +
		`, COUNT(*) OVER() AS totalcount FROM ` +
		table + SerializeFilter(queryOptions.Filter) +
		SerializePage(queryOptions.Page)
	fmt.Println(query)
	start := time.Now()
	rows, err = db.Queryx(query)
	elapsed := time.Since(start)
	fmt.Println("Request took ", elapsed)
	return
}

func Search(table string, model interface{}, queryOptions models.QueryModel) (rows *sqlx.Rows, err error) {
	if db == nil {
		return
	}
	query := `SELECT ` + models.SerializeFieldsNameSQL(model) +
		`, COUNT(*) OVER() AS totalcount FROM ` +
		table +
		` WHERE textsearchable_index_col @@ to_tsquery('` + queryOptions.Search + `')` +
		SerializePage(queryOptions.Page)

	fmt.Println(query)
	start := time.Now()
	rows, err = db.Queryx(query)
	elapsed := time.Since(start)
	fmt.Println("Request took ", elapsed)
	if err != nil {
		fmt.Println("Error occured in search")
		return
	}
	return
}

func SearchSpec(table string, model interface{}, queryOptions models.QueryModel, spec models.Converter) (rows *sqlx.Rows, err error) {
	if db == nil {
		return
	}
	query := `SELECT ` + spec(models.SerializeFieldsNameSQL(model)) +
		`, COUNT(*) OVER() AS totalcount FROM ` +
		table +
		` WHERE textsearchable_index_col @@ to_tsquery('` + queryOptions.Search + `')` +
		SerializePage(queryOptions.Page)

	fmt.Println(query)
	start := time.Now()
	rows, err = db.Queryx(query)
	elapsed := time.Since(start)
	fmt.Println("Request took ", elapsed)
	if err != nil {
		fmt.Println("Error occured in search")
		return
	}
	return
}

func SearchMatchArray(table string, model interface{}, queryOptions models.QueryModel, spec models.Converter) (rows *sqlx.Rows, err error) {
	if db == nil {
		return
	}
	query := `SELECT ` + spec(models.SerializeFieldsNameSQL(model)) +
		`, COUNT(*) OVER() AS totalcount FROM ` +
		table + SerializeFilterArray(queryOptions.Filter) +
		SerializePage(queryOptions.Page)

	fmt.Println(query)
	start := time.Now()
	rows, err = db.Queryx(query)
	elapsed := time.Since(start)
	fmt.Println("Request took ", elapsed)
	if err != nil {
		fmt.Println("Error occured in search")
		return
	}
	return
}
