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
    --git-repo-id mistsdkgo \
    --git-user-id tmunzer


cp ./fix_sdk/model_const_device_ap.go ./mist-sdk-go/
cp ./fix_sdk/model_dot11_band.go ./mist-sdk-go/
cp ./fix_sdk/model_dot11_bandwidth.go ./mist-sdk-go/
cp ./fix_sdk/model_rrm_event_pre_bandwidth.go ./mist-sdk-go/