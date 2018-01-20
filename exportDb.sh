#! /bin/bash
now=$(date +"%d-%m-%Y")
docker exec -i serieswatchermysql mysqldump -proot series | tee ../../Backup/$now.sql > ../../Backup/current.sql
