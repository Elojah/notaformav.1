/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   vote.js                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/27 16:18:48 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/27 17:16:48 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

function	sendVoteResponse(id, data, status) {
	console.log(id, data, status)
	var evt = $.Event('end-voting')
	// TODO: Parse data and change state if err == none
	evt.state = null
	$(window).trigger(evt)
}

function	sendVote(id, upOrDown) {
	$.post(
		window.location.pathname + '/comment/' + id,
		{vote: upOrDown},
		function (data, status) {
			sendVoteResponse(id, data, status)
		}
	);
}
