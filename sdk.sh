openapi-generator generate \
    -i ./openapi.yml \
    -g go \
    -o ./mist-sdk-go \
    --api-package mist-sdk-go \
    --additional-properties=packageName=mistsdkgo \
    --additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
    --additional-properties=enumClassPrefix=true \