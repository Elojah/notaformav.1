#!/bin/sh

a=0
while [ "$a" -lt 10 ]    # this is loop1
do
	curl --data "content=TEST_PAGINATION" http://localhost:4242/establishment/ANIq67cZ2N4/comments
done
