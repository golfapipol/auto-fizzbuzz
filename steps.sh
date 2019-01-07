#!/bin/sh
registry="golfapipol"
backing_service="docker-compose-backing-service.yml"

# Unit Test
cd $(pwd)/src
go test ./... && echo "Unit Test Passed"

# Integration Test
docker-compose -f $backing_service up 
## tearup
docker-compose -f $backing_service run -v $(pwd)/seed/tearup-integrate.js:/tmp/tearup.js mongodb mongo mongodb://mongodb:27017/fizzbuzzdb /tmp/tearup.js
## test
go test -tags=integration ./... && echo "Integration Test Passed"
## teardown
docker-compose -f $backing_service run -v $(pwd)/seed/teardown-integrate.js:/tmp/teardown.js mongodb mongo mongodb://mongodb:27017/insurance_gateway /tmp/teardown.js
# Build
cd $(pwd)
docker-compose -f docker-compose.yml build && echo "Build Passed"

# Push Image
docker tag fizzbuzz-service:latest $registry/fizzbuzz-service:latest
docker push $registry/fizzbuzz-service:latest && echo "Push to registry Passed"

# Deploy
## dev environment
# ssh
docker-compose -f docker-compose.deploy.yml -f $backing_service up -d

## tearup data
docker-compose -f $backing_service run -v $(pwd)/seed/tearup-acceptance.js:/tmp/tearup.js mongodb mongo mongodb://mongodb:27017/fizzbuzzdb /tmp/tearup.js
## Acceptance Test
newman run atdd/fizzbuzz.postman_collection.json -d atdd/test-data/data.json
## teardown data
docker-compose -f $backing_service run -v $(pwd)/seed/teardown-acceptance.js:/tmp/teardown.js mongodb mongo mongodb://mongodb:27017/insurance_gateway /tmp/teardown.js
