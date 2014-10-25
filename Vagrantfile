#
# 2014+ Copyright (c) Anton Tiurin <noxiouz@yandex.ru>
# All rights reserved.
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU Lesser General Public License as published by
# the Free Software Foundation; either version 2 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Lesser General Public License for more details.
#


BOX_URL = 'https://github.com/cocaine/cocaine-vagrant/releases/download/v0.11/precise64-docker.box'
BOX_NAME = 'precise64-docker'

Vagrant.configure("2") do |config|

	(2..3).each do |i|
		config.vm.define "influx#{i}" do |influx|
			influx.vm.box = BOX_NAME
			influx.vm.box_url = BOX_URL
			influx.vm.hostname = "influx#{i}.local"

			influx.vm.network "private_network", ip: "192.168.50.#{i}"
			# port for influxdb web interface
			influx.vm.network :forwarded_port, guest: 8083, host: 8083 + 100*i
			# port for influxdb HTTP api. Used by client to send points
			influx.vm.network :forwarded_port, guest: 8086, host: 8086 + 100*i
			# for grafana. nginx port, for example
			influx.vm.network :forwarded_port, guest: 8080, host: 8080 + 100*i


			influx.vm.provision "shell", path: "script.sh"

			influx.vm.provision "shell" do |s|
				s.inline = "cp /vagrant/influx#{i}.config.toml /opt/influxdb/shared/config.toml"
			end

			influx.vm.provision "shell" do |s|
				s.inline = "cp /vagrant/config.js /home/vagrant/grafana-1.8.1/"
			end

			influx.vm.provision "shell" do |s|
				s.inline = "service influxdb restart"
			end

			influx.vm.provision "shell", path: "run_grafana.sh"
		end
	end
end
