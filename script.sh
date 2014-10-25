#!/usr/bin/env bash

if [ $(dpkg -l | grep -c influx) -eq 0 ];
then
	echo "Fetching package"
	wget -q http://s3.amazonaws.com/influxdb/influxdb_latest_amd64.deb
	echo "Installing package"
	sudo dpkg -i influxdb_latest_amd64.deb
else
	echo "influxdb has been already installed"
fi
