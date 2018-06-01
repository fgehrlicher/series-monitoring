#! /bin/bash
now=$(date +"%d-%m-%Y")
docker exec -i serieswatchermysql mysqldump -proot series | tee ../Ressources/Backup/$now.sql > ../Ressources/Backup/current.sql
