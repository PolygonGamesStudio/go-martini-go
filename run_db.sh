#!/bin/sh
echo "enter your postgres-user password:"
su postgres
pg_ctl start
#сейчас не работает - прерываетс яна вводе пароля