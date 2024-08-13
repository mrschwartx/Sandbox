# Pull Base Image : on local
docker pull ubuntu:jammy

# Run Docker Container : on local
docker run -it --name my-ubuntu-jammy ubuntu:jammy

# Exect Container : on local
docker exec -it my-ubuntu-jammy bash

# Install Necessary Packages : on container
apt-get update
apt-get install -y sudo vim net-tools

# Create a User : on container
adduser myuser
usermod -aG sudo myuser

# SSH : on container
apt-get install -y openssh-server
service ssh start

# Map SSH : on local
docker run -d -p 2222:22 --name my-ubuntu-jammy ubuntu:jammy

# Use SSH : on client
ssh myuser@localhost -p 2222