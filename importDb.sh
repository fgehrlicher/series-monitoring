#!/usr/bin/env bash

cat ./Ressources/Backup/current.sql | docker exec -i serieswatchermysql mysql -proot
