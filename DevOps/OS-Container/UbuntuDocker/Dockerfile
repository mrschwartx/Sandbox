# Use an official Ubuntu base image
FROM ubuntu:jammy

# Update and install necessary packages
RUN apt-get update && apt-get install -y \
    sudo \
    openssh-server \
    vim \
    net-tools

# Create a new user and set password
RUN useradd -ms /bin/bash myuser && echo "myuser:password" | chpasswd && adduser myuser sudo

# Configure SSH
RUN mkdir /var/run/sshd

# Expose SSH port
EXPOSE 22

# Start SSH service
CMD ["/usr/sbin/sshd", "-D"]
