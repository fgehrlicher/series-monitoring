#! /bin/bash
now=$(date +"%d-%m-%Y")
docker exec -i serieswatchermysql mysqldump -proot serieswatcher > Backup/$now.sql
