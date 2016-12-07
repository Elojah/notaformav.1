/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   json.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/19 20:14:46 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/06 18:12:43 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package psql

import (
// "github.com/go-server/models"
// "github.com/jmoiron/sqlx"
)

// func RowsToJSONEstablishment(rows *sqlx.Rows) (result models.JSONContent, err error) {
// 	for rows.Next() {
// 		var establishment models.EstablishmentModel
// 		err = rows.StructScan(&establishment)
// 		if err != nil {
// 			return result, err
// 		}
// 		var data models.JSONData
// 		data.Id = establishment.Id
// 		data.Attributes = establishment
// 		result.Data = append(result.Data, data)
// 	}
// 	return
// }

// func RowsToJSONCommentEstablishment(rows *sqlx.Rows) (result models.JSONContent, err error) {
// 	for rows.Next() {
// 		var comment models.CommentEstablishmentModel
// 		err = rows.StructScan(&comment)
// 		if err != nil {
// 			return result, err
// 		}
// 		var data models.JSONData
// 		data.Id = comment.Id
// 		data.Attributes = comment
// 		result.Data = append(result.Data, data)
// 	}
// 	return
// }
