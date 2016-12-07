/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   common.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: drabahi <drabahi@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/12 12:00:46 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/23 14:01:12 by drabahi          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package psql

import (
	"fmt"
	"github.com/go-server/conf"
	"github.com/go-server/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"io/ioutil"
	"strings"
)

var db *sqlx.DB

func Open() (err error) {
	url := conf.DATABASE_URL
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err = sqlx.Open("postgres", connection)
	if err != nil {
		return
	}
	return
}

func Drop(table string) (err error) {
	query := `DROP TABLE IF EXISTS ` + table
	_, err = db.Exec(query)
	return
}

func Index(table string, model interface{}) (err error) {
	query := `ALTER TABLE ` + table + ` ADD COLUMN textsearchable_index_col tsvector`
	db.MustExec(query) // How get error here ?
	query = `UPDATE ` + table + ` SET textsearchable_index_col = `
	query += models.SerializeStructSQLIndex(model)
	fmt.Println(query)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	// db.MustExec(query) // How get error here ?
	query = `DROP INDEX IF EXISTS establishments_index`
	db.MustExec(query) // How get error here ?
	query = `CREATE INDEX search_index_` + table + ` ON ` + table + `
			USING GIN (textsearchable_index_col)`
	db.MustExec(query) // How get error here ?
	query = `
		CREATE OR REPLACE FUNCTION size() RETURNS int AS $func$
		DECLARE
			retval int;
		BEGIN
			GET DIAGNOSTICS retval = ROW_COUNT;
			RETURN retval;
		END;
		$func$ LANGUAGE plpgsql`
	db.MustExec(query)
	return nil
}

func Exec(filename string) (data *interface{}, err error) {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		if request == "" {
			continue
		}
		fmt.Println(request)
		_, err := db.Exec(request)
		if err != nil {
			fmt.Println("Error executing sql query...")
			return nil, err
		}
	}
	return
}
