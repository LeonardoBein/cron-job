cat <<EOF
cron-job has been installed as a systemd service.

To start/stop cron-job:

sudo systemctl start cron-job
sudo systemctl stop cron-job

To enable/disable cron-job starting automatically on boot:

sudo systemctl enable cron-job
sudo systemctl disable cron-job

To reload cron-job:

sudo systemctl restart cron-job

To view cron-job logs:

journalctl -f -u cron-job

EOF