echo "Backing up..."
tar --exclude-vcs -zcvf /backups/backup_$(date +%Y%m%d).tar.gz /data