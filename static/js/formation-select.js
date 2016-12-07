/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   formation-select.js                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/12/07 16:36:07 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 20:56:27 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

var			currentFormationID = '';
function	showFormationEdition() {
	$('.dropdown#formation-selection')
	.dropdown({
		onChange: function(value, text) {
			if (value == currentFormationID) {
				return
			}
			currentFormationID = value
			refreshOne('formation-edit', {id:value} ,"/api/formation", renderFormationEdition);
		}
	});
}
