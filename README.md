# volect
`volect` is a Volunteer reflection writing service made with Go.

## How to set up
```shell
# How to run in a local development environment
$ VOLUNTEER_ADMIN_ID=[ID] VOLUNTEER_ADMIN_PW=[PW] go run .

# How to run in GCE
$ chmod +x myserver
$ sudo VOLUNTEER_ADMIN_ID=[ID] VOLUNTEER_ADMIN_PW=[PW] ./myserver

# How to create build files for GCE
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myserver .

# How to view SQLite data
$ sqlite3 reviews.db
sqlite> .mode column
sqlite> .headers on
sqlite> SELECT * FROM reviews;

# How to extract sqlite data to csv
$ sqlite3 -header -csv reviews.db "SELECT * FROM reviews;" > report.csv
$ printf "\xEF\xBB\xBF" > final_report.csv
$ cat report.csv >> final_report.csv

# How to get GCE's NAME, ZONE
$ gcloud compute instances list

# How to upload files to GCE
$ gcloud compute scp [FILES] [NAME]:~/ --zone=[ZONE]

# How to import files from GCE
$ gcloud compute scp [NAME]:~/[FILES] ./ --zone=[ZONE]

# How to connect to GCE via ssh
$ gcloud compute ssh [NAME] --zone=[ZONE]

# How to set Ubuntu GCE
sudo apt-get update
sudo apt install build-essential -y
sudo apt install golang-go -y

# systemctl
sudo vim /etc/systemd/system/volunteer.service

sudo systemctl daemon-reload
sudo systemctl enable volunteer.service
sudo systemctl start volunteer.service
sudo systemctl status volunteer.service
sudo systemctl restart volunteer.service
sudo systemctl stop volunteer.service

sudo setcap CAP_NET_BIND_SERVICE=+eip /home/[USER_ID]/myserver
sudo journalctl -u volunteer.service -e --no-pager

# How to SWAP
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

## License
Copyright &copy; 2025-2026 Cha Haneum

This project is licensed under the MIT License.
