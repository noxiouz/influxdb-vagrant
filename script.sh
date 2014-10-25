#!/usr/bin/env bash

echo "Fetching package"
wget -q http://s3.amazonaws.com/influxdb/influxdb_latest_amd64.deb
echo "Installing package"
sudo dpkg -i influxdb_latest_amd64.deb

echo "Copy configuration file"
cp /vagrant/config.toml /opt/influxdb/shared/config.toml

service influxdb restart
