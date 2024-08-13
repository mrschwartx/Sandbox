# Build Docker Images
docker build -t ubuntu-jammy .

# Run Docker Container
docker run -d -p 2222:22 --name my-ubuntu-jammy ubuntu-jammy

# SSH into Container
ssh myuser@localhost -p 2222
