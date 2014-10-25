#!/bin/sh

sudo apt-get update && sudo apt-get install nginx -y
cp /vagrant/default /etc/nginx/sites-enabled/default
service nginx restart
cp -R /home/vagrant/grafana-1.8.1 /usr/share/nginx/www/grafana
