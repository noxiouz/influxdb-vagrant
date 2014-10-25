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

	(1..2).each do |i|
		config.vm.define "influx#{i}" do |influx|
			influx.vm.box = BOX_NAME
			influx.vm.box_url = BOX_URL
			influx.vm.network "private_network", ip: "192.168.50.#{i}"
			influx.vm.hostname = "influx#{i}.local"

			# provision
			influx.vm.provision "shell",
				path: "script.sh"
		end
	end
end
