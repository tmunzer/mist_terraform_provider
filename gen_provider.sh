#~/go/bin/tfplugingen-openapi generate --config ./generator_config.yml --output ./provider_code_spec.json ./openapi.yml
#~/go/bin/tfplugingen-framework generate all --input provider_code_spec.json --output internal/provider
#~/go/bin/tfplugingen-framework scaffold data-source --name services --output-dir ./terraform-provider-mist/internal/provider

# alias tfplugingen-framework="~/go/bin/tfplugingen-framework"
# alias tfplugingen-openapi="~/go/bin/tfplugingen-openapi"

PNAME="terraform-provider-mistapi"


# mkdir $PNAME
# cp openapi.yml $PNAME
# cd $PNAME
# go mod init $PNAME
# touch main.go

# cat <<EOF >>main.go
# package main
# import (
#     "context"
#     "log"
#     "terraform-provider-mistapi/internal/provider"
#     "github.com/hashicorp/terraform-plugin-framework/providerserver"
# )
# func main() {
#     opts := providerserver.ServeOpts{
# 		Address: "registry.terraform.io/juniper/mistapi",
#     }
#     err := providerserver.Serve(context.Background(), provider.New(), opts)
#     if err != nil {
#         log.Fatal(err.Error())
#     }
# }
# EOF


# mkdir -p internal/provider


# tfplugingen-framework scaffold provider \
#   --name mist \
#   --output-dir ./internal/provider

# ts=`date "+%s"`
# echo "backuping src"
# cp ./provider-code-spec.json ./provider-code-spec.json.$ts.bak

echo "fix openapi specs..."
python3 ./gen_provider_pre.py

echo "generate provider specs..."
tfplugingen-openapi generate \
    --config ./generator_config.yml \
    --output ./provider-code-spec.json \
    ./mist.provider.yml 


echo "fix provider specs..."
cp ./provider-code-spec.json ./provider-code-spec.json.bak
python3 ./gen_provider_mid.py

echo "generate provider..."
tfplugingen-framework generate resources \
    --input ./provider-code-spec.json \
    --output ./terraform-provider-mist/internal

echo "fix provider..."
python3 ./gen_provider_post.py



# tfplugingen-framework scaffold resource \
#     --name services \
#     --output-dir ./terraform-provider-mist/internal/provider