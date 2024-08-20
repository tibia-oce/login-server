# Invoke-RestMethod -Uri http://localhost:80/login.php -Method Post -Body '{"email": "test@example.com", "password": "test", "stayloggedin": true, "type": "login"}' -ContentType "application/json"

$loginUrl = "http://localhost:80/login.php"
$body = '{"email": "test@example.com", "password": "test", "stayloggedin": true, "type": "login"}'

try {
    $response = Invoke-RestMethod -Uri $loginUrl -Method Post -Body $body -ContentType "application/json"
    Write-Output $response
}
catch {
    Write-Output $_.Exception.Message
}
