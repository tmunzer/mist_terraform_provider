echo ">> Updating Source OpenAPI Spec files"
cd ./mist_openapi
git pull
cd ..
echo ">> Generating TF OpenAPI Spec file"
python3 ./00_update_specs_custom.py