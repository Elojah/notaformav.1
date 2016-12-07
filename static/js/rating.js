/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   rating.js                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/24 18:17:34 by hdezier           #+#    #+#             */
/*   Updated: 2016/09/24 19:18:54 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

function initRating() {
	$(".star-rating").rating({
		min: 0,
		max: 5,
		step: 1,
		animate: false,
		showCaption: false,
		hoverEnabled: false
	});
}
