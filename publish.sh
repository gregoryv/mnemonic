#!/bin/bash

file=$1

[ ! -e $file ] && echo "$file does not exist" && exit 1

rsync -av $file  warmachine:/var/www/www.7de.se/dl/$file
