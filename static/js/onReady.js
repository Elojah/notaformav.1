/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   onReady.js                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: drabahi <drabahi@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/24 19:07:50 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/26 19:07:59 by drabahi          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

jQuery(document).ready(function () {
	initRating()
	// Set
	$("#post-comment").attr("action", window.location.pathname + "/comments")
	// Accordion initilization
	$(".ui.accordion").accordion()
});
