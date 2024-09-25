# Gin Lambda Example

Gin Lambda example. This example spins up a toy API that allows callers to
create and list Dashes via CreateDash and ListDashes API operations
respectively.

NOTE: This sample currently persists Dashes in RAM for the sake of brevity, so
Dashes are lost once the Lambda function shuts down. This would realistically be
done with a database instead.

## Getting started

### Local

Build and run:

```
cd api
go build cmd/local/main.go
./main
```

Query with curl:

```
# CreateDash
curl -X POST http://localhost:8080/api/v1/dashes -d '{
  "name": "My Dash",
  "speed": 123
}"

# ListDashes
curl -X GET http://localhost:8080/api/v1/dashes
```

### Lambda

Deploy Lambda function

```
cd cdk
npm install
npx cdk bootstrap
npx cdk deploy
```

Go to Lambda function in AWS Console and test with payloads:

```
# CreateDash
{
  "path": "/api/v1/dashes",
  "httpMethod": "POST",
  "body": "{\"name\": \"My Dash\", \"speed\": 123}"
}

# ListDashes
{
  "path": "/api/v1/dashes",
  "httpMethod": "GET"
}
```