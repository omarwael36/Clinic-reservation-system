#!/bin/bash
# EC2 Deployment Script
# This script pulls the latest images and restarts the containers

set -e

echo "========================================="
echo "Clinic Reservation System - EC2 Deploy"
echo "========================================="

# Navigate to app directory
cd /home/ec2-user/clinic-app || cd /home/ubuntu/clinic-app

# Pull latest images from DockerHub
echo ""
echo "[1/4] Pulling latest images from DockerHub..."
docker-compose pull

# Stop existing containers (if running)
echo ""
echo "[2/4] Stopping existing containers..."
docker-compose down || true

# Start containers with new images
echo ""
echo "[3/4] Starting containers with new images..."
docker-compose up -d

# Clean up old images
echo ""
echo "[4/4] Cleaning up old images..."
docker image prune -f

# Show status
echo ""
echo "========================================="
echo "Deployment Complete!"
echo "========================================="
echo ""
docker-compose ps
echo ""
echo "Application URLs:"
echo "  Frontend: http://$(curl -s http://169.254.169.254/latest/meta-data/public-ipv4):80"
echo "  Backend:  http://$(curl -s http://169.254.169.254/latest/meta-data/public-ipv4):8080"
echo ""

