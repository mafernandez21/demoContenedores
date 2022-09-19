#!/bin/bash
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg lsb-release vim net-tools
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin
sudo docker run hello-world

echo 'PS1="\[\033[1;32m\]<?>_\[\033[00m\]"' >> ~/.bashrc