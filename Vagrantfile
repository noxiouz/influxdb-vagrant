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

Vagrant.configure("2") do |config|
	config.vm.define "influx1" do |influx1|
		influx1.vm.box = 'precise64-docker'
		influx1.vm.box_url = 'https://github.com/cocaine/cocaine-vagrant/releases/download/v0.11/precise64-docker.box'
		influx1.vm.network "private_network", ip: "192.168.50.4"

		# provision
		influx1.vm.provision "shell",
			path: "script.sh"
	end

	config.vm.define "influx2" do |influx2|
		influx2.vm.box = 'precise64-docker'
		influx2.vm.box_url = 'https://github.com/cocaine/cocaine-vagrant/releases/download/v0.11/precise64-docker.box'
		influx2.vm.network "private_network", ip: "192.168.50.5"

		# provision
		influx2.vm.provision "shell",
			path: "script.sh"
	end
end
