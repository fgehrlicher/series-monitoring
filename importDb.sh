#!/usr/bin/env bash

cat ../../Backup/current.sql | docker exec -i serieswatchermysql mysql -proot
