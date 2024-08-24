# Define colors for output
$Blue = [System.ConsoleColor]::Blue
$Yellow = [System.ConsoleColor]::Yellow
$Green = [System.ConsoleColor]::Green
$Default = [System.ConsoleColor]::Gray
$Cyan = [System.ConsoleColor]::Cyan
$Magenta = [System.ConsoleColor]::Magenta

function Write-Colored {
    param (
        [string]$Text,
        [System.ConsoleColor]$Color
    )
    $originalColor = [System.Console]::ForegroundColor
    [System.Console]::ForegroundColor = $Color
    Write-Output $Text
    [System.Console]::ForegroundColor = $originalColor
}

Write-Output ""
Write-Colored "--------------------------------------" $Default
Write-Output ""
Write-Colored "Starting Docker operations..." $Blue
Write-Output ""
Write-Colored "--------------------------------------" $Default

Write-Output ""
Write-Colored "Stopping and cleaning any existing containers..." $Yellow
cd docker
docker-compose down -v
docker-compose rm -f
docker rmi docker-login

Write-Output ""
Write-Colored "Downloading the most recent versions of containers..." $Yellow
docker-compose pull
docker-compose build --no-cache docker-login

Write-Output ""
Write-Colored "Starting containers..." $Yellow
docker-compose up -d
cd ..

# Retrieve IP and port
$containerIP = docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' myaac
$containerPort = docker port myaac
$containerPort = $containerPort.Split(":")[1]

# Print completion message and link
Write-Output ""
Write-Colored "Docker operations completed. You can now close this window." $Green
$link = "http://localhost:${containerPort}"
Write-Colored "Access it here: $link" $Cyan
Write-Output ""
