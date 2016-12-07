/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   create.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/06 18:03:33 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/18 16:12:15 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package psql

import (
	"bytes"
	"fmt"
	"github.com/go-server/models"
)

func CreateTable(model interface{}, name string) (err error) {
	var b bytes.Buffer
	b.WriteString(`CREATE TABLE IF NOT EXISTS ` + name + ` (`)
	b.WriteString(models.SerializeFieldsNameSQLCreation(model))
	b.WriteString(`);`)

	fmt.Println(b.String())
	_, err = db.Exec(b.String())
	return
}
