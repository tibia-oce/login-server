# Define colors for output
$Blue = [System.ConsoleColor]::Blue
$Yellow = [System.ConsoleColor]::Yellow
$Green = [System.ConsoleColor]::Green
$Default = [System.ConsoleColor]::Gray

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
Write-Colored "Stopping and removing containers..." $Yellow
# Stop and remove containers
cd docker
docker-compose down -v
docker-compose rm -f

Write-Output ""
Write-Colored "Removing network and volume..." $Yellow
docker-compose rm -f
cd ..

Write-Output ""
Write-Colored "Docker operations completed. You can now close this window and exit Docker." $Green
Write-Output ""