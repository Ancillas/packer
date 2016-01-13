# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

  config.vm.box = "puppetlabs/ubuntu-14.04-64-puppet"

  config.vm.provider "virtualbox" do |v|
    v.memory = 1024
  end

  config.vm.provision "shell", inline: <<-SHELL
    set -e
    sudo apt-get update
    set +e
    sudo /opt/puppetlabs/bin/puppet apply -e 'package { ['curl', 'git', 'make', 'binutils', 'bison', 'gcc', 'build-essential', 'unzip', 'bzr', 'mercurial']: ensure => installed, }'
    set -e
    sudo -i -u vagrant bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
    mkdir -p /home/vagrant/gocode
    mkdir -p /home/vagrant/gocode/bin
    mkdir -p /home/vagrant/gocode/pkg
    mkdir -p /home/vagrant/gocode/src
    echo "source /home/vagrant/.gvm/scripts/gvm" >> /home/vagrant/.profile
    echo -e "export GOPATH=\\"/home/vagrant/gocode\\"" >> /home/vagrant/.profile
    echo -e "export PATH=\\"/home/vagrant/gocode/bin:\\$PATH\\"" >> /home/vagrant/.profile
    source /home/vagrant/.profile
    gvm install go1.4.1
    gvm use go1.4.1 --default
    mkdir -p /home/vagrant/gocode
    sudo chown -R vagrant:vagrant /home/vagrant/gocode
    mkdir -p /home/vagrant/gocode/src/github.com/mitchellh
    ln -s /vagrant /home/vagrant/gocode/src/github.com/mitchellh/packer
    sudo chown -R vagrant:vagrant /home/vagrant/gocode/src
  SHELL

end
