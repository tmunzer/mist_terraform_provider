openapi-generator generate \
    -i ./mist_openapi/workdir/mist.openapi.yml \
    -g go \
    -o ./mist-sdk-go \
    --api-package mist-sdk-go \
    --additional-properties=packageName=mistsdkgo \
    --additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
    --additional-properties=enumClassPrefix=true \
    --git-repo-id mistsdkgo \
    --git-user-id tmunzer