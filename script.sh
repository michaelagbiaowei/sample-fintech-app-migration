# USING BASH SCRIPT TO TAG AND UPLOAD AN IMAGE TO DOCKER HUB

#!/usr/bin/env bash

# Assumes that an image is built via `run_docker-compose.sh`

# Step 1:
# Create dockerpath
dockerpath=go-sample-app
user=maiempire

# Step 2:  
# Authenticate & tag
echo "Docker ID and Image: $dockerpath"
docker login -u $user

#Step 3:
# Tag the images with your Docker ID
docker tag $dockerpath:v2 $user/$dockerpath:v2

# Step 4:
# Push image to a docker repository
docker push $user/$dockerpath:v2