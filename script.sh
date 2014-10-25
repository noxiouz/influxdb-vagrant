#!/usr/bin/env bash


# install influxdb
if [ $(dpkg -l | grep -c influx) -eq 0 ];
then
	echo "Fetching package"
	wget -q http://s3.amazonaws.com/influxdb/influxdb_latest_amd64.deb
	echo "Installing package"
	sudo dpkg -i influxdb_latest_amd64.deb
else
	echo "influxdb has been already installed"
fi


# install grafana
echo "Installing grafana"
wget -q "http://grafanarel.s3.amazonaws.com/grafana-1.8.1.tar.gz" && tar -xf grafana-1.8.1.tar.gz
