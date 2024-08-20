# health.ps1

# Define the URL of the health check endpoint
$healthCheckUrl = "http://localhost:80/health"

try {
    # Perform the GET request to the health check endpoint
    $response = Invoke-RestMethod -Uri $healthCheckUrl -Method Get

    # Output the response to the console
    Write-Output "Health Check Response:"
    Write-Output $response
}
catch {
    # Handle any errors that occur during the request
    Write-Output "Error occurred during health check:"
    Write-Output $_.Exception.Message
}
