#/bin/bash

PNAME="terraform-provider-mistapi"
CUR_FOLDER=$(pwd)
GEN_FOLDER="/Users/tmunzer/4_dev/Mist/mist_terraform_provider/generator"
PVD_FOLDER="/Users/tmunzer/4_dev/Mist/mist_terraform_provider/terraform-provider-mist"
CODESPECS_FOLDER="code-spec-fix"

##############################
############### COMMON
code_specs(){
    code_specs_gen
    code_specs_fix $1
}

code_specs_gen() {
    echo "TF Specs: generating..."
    tfplugingen-openapi generate \
        --config ./generator_config.yml \
        --output ./provider-code-spec.json \
        $GEN_FOLDER/mist.provider.yml
}

code_specs_fix() {
    echo "TF Specs: applying fixes..."
    cp ./provider-code-spec.json ./provider-code-spec.json.bak
    python3 $GEN_FOLDER/10_fix_code_specs.py --spec_in=$1/provider-code-spec.json --fix_folder=$FOLDER/$CODESPECS_FOLDER --tf_type=resources
}

##############################
############### RESOURCE
resource() {
    RESOURCE_NAME=$1
    FOLDER="$GEN_FOLDER/resource/$RESOURCE_NAME"
    echo "----------------------------------------------"
    echo "RESOURCE GENERATION: $RESOURCE_NAME"
    echo "FOLDER: $FOLDER"
    if [ -d $FOLDER ]
    then
        cd $FOLDER
        resource_pre_fix
        code_specs $FOLDER
        resource_gen
        resource_post_fix $FOLDER
    else
        echo "ERROR: Resource not found"
    fi
    echo "----------------------------------------------"
    cd $CUR_FOLDER

}

resource_gen() {
    echo "TF Resource: generating..."
    tfplugingen-framework generate resources \
        --input ./provider-code-spec.json \
        --output $PVD_FOLDER/internal
}

resource_pre_fix(){
    if [ -f "pre_fix.py" ]
    then 
        echo "OpenAPI Specs: applying prefix ..."
        ./pre_fix.py
    else
        echo "OpenAPI Specs: no prefix to the apply..."
    fi
}

resource_post_fix() {
    if [ -f "post_fix.yml" ]
    then 
        echo "TF Resource: applying postfix..."
        python3 ./10_fix_generated.py -f $1/post_fix.yml
    elif [ -f "post_fix.yaml" ]
    then
        echo "TF Resource: applying postfix..."
        python3 ./10_fix_generated.py -f $1/post_fix.yaml
    else
        echo "TF Resource: no postfix to apply..."
    fi
}

##############################
############### DATASOURCE
datasource() {
    echo "DATASOURCE: $1"
}

datasource_gen() {
    tfplugingen-framework generate data-sources \
        --input ./provider-code-spec.json \
        --output ./terraform-provider-mist/internal
}

datasource_fix() {
    echo "fix provider..."
    python3 ./gen_provider_post.py
}

fmt() {
    cd terraform-provider-mist
    make fmt
    cd ..
}


###### ENTRYPOINT
ARG=$1
while echo "$ARG" | grep -qv "^$"
do
    case "$ARG" in
        -r )  shift
                resource $1
                shift
                ARG=$1;;
        -d )  shift
                datasource $1
                shift
                ARG=$1;;
    esac
done
