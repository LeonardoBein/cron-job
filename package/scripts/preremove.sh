systemctl stop cron-job
systemctl disable cron-job
rm /etc/systemd/system/cron-job.service
systemctl daemon-reload
systemctl reset-failed