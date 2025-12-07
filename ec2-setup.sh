#!/bin/bash
# EC2 Initial Setup Script
# Run this ONCE on your EC2 instance to set everything up

set -e

echo "========================================="
echo "EC2 Initial Setup for Clinic App"
echo "========================================="

# Update system
echo ""
echo "[1/6] Updating system packages..."
sudo yum update -y || sudo apt-get update -y

# Install Docker
echo ""
echo "[2/6] Installing Docker..."
if command -v amazon-linux-extras &> /dev/null; then
    # Amazon Linux 2
    sudo amazon-linux-extras install docker -y
else
    # Ubuntu/Debian
    sudo apt-get install -y docker.io
fi

# Start Docker service
echo ""
echo "[3/6] Starting Docker service..."
sudo systemctl start docker
sudo systemctl enable docker

# Add current user to docker group
echo ""
echo "[4/6] Adding user to docker group..."
sudo usermod -aG docker $USER

# Install Docker Compose
echo ""
echo "[5/6] Installing Docker Compose..."
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Create app directory
echo ""
echo "[6/6] Creating application directory..."
mkdir -p ~/clinic-app
cd ~/clinic-app

# Download docker-compose.yml from GitHub
echo "Downloading docker-compose.yml..."
curl -O https://raw.githubusercontent.com/omarwael36/Clinic-reservation-system/main/docker-compose.yml

# Download deploy script
echo "Downloading deploy script..."
curl -O https://raw.githubusercontent.com/omarwael36/Clinic-reservation-system/main/deploy.sh
chmod +x deploy.sh

echo ""
echo "========================================="
echo "Setup Complete!"
echo "========================================="
echo ""
echo "IMPORTANT: Log out and log back in for docker group to take effect!"
echo ""
echo "Then run: cd ~/clinic-app && ./deploy.sh"
echo ""
echo "Or to start manually:"
echo "  cd ~/clinic-app"
echo "  docker-compose pull"
echo "  docker-compose up -d"
echo ""

