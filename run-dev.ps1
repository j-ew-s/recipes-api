
Write-Output "**************************************************"
Write-Output "* Initializing Receipts-API project using docker *"
Write-Output "**************************************************"

Write-Output "Step 1"
Write-Output "Creating image of Receipts-API"
Write-Output "docker build -t recipes-api ."

docker build -t recipes-api .


Write-Output "Step 2"
Write-Output "Pulling MongoDB on latest version"
Write-Output "docker pull mongo:latest"

docker pull mongo:latest


Write-Output "Step 3"
Write-Output "Creating Volume mongodata for storing data from mongodb container"
Write-Output "** any change on volume name should reflect on docker-compose.yml"
Write-Output "docker volume create --name=mongodata"

docker volume create --name=mongodata


Write-Output "Step 4"
Write-Output "Executing detached docker compose  building with force recreate"
Write-Output "docker-compose up -d --build --force-recreate"

docker-compose up -d --build --force-recreate


Write-Output "****************************************************"
Write-Output "* Its done. You should be able to run on port 8087 *"
Write-Output "****************************************************"
