cd ./mist_openapi
git pull
cd ..

python3 ./gen_sdk_pre.py

openapi-generator generate \
    -i ./mist.sdk.yml \
    -g go \
    -o ./mist-sdk-go \
    --api-package mist-sdk-go \
    --additional-properties=packageName=mistsdkgo \
    --additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
    --additional-properties=enumClassPrefix=true \
    --additional-properties=generateInterfaces=true \
    --git-repo-id mistsdkgo \
    --git-user-id tmunzer


cp ./fix_sdk/* ./mist-sdk-go/