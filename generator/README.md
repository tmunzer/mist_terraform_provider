This folder contains all the required scripts and files to generate the provider Resources and Datasources from the Mist OpenAPI specs.


## Scripts
### 00_update_specs.sh
This script is used to update the Mist OpenAPI specification from the remote Github repository. It will:
1. perform a `git pull` from the repo folder
2. run the `00_update_specs_custom.py` python script to apply custom changes to the OpenAPI Spec to adapt it to the Terraform constraints:
    * customize the org_inventory object
    * customize the switch matching rules
    * customize the DHCPD object in the switch and gateway configurations
    * remove webhook sample data
    * and others...

#### Script usage
This script can be run without any parameter: 
```bash
00_update_specs.sh
```

### 10_gen_provider.sh
This script is used to automate the Terraform Resource or Datasource from the OpenAPI Specification. 

> [!IMPORTANT]
> The script is using the [tfplugingen-openapi](https://developer.hashicorp.com/terraform/plugin/code-generation/openapi-generator) and [tfplugingen-framework](https://developer.hashicorp.com/terraform/plugin/code-generation/framework-generator) tools from Hashicorp to generate the final resource or datasource. Make sure to have them installed first

#### Create a new Resource or Datasource
1. create a new folder in the `./resource` or `./datasource` folder. The new folder name doesn't matter, but the best practice would be to name it with the resource/datasource name
2. in the new folder:
    * create the `generator_config.yml` file. This file is describing the resources and datasources than must be generated and the associated API Requests. An example can be found in the [tfplugingen-openapi documentation](https://developer.hashicorp.com/terraform/plugin/code-generation/openapi-generator#example)
    * create a `code-spec-fix` folder. This folder can be empty of have one or multiple yaml files. These files are used to manipulate the terraform specs generated by the `tfplugingen-openapi` tool (customize the required/computed/optional fields, add validators/planmodifiers, rename fields, ...)
    * if needed, create a `post_fix.yml` file that will be used to manipulate the resource / datasource generated by the `tfplugingen-framework` tool

#### Folder structure:
* generator (folder)
  * resource (folder)
    * my_resource_folder_name (folder)
      * code-spec-fix (folder)
        * my_first_fix.yml (file)
        * my_second_fix.yml (file)
      * generator_config.yml (file)
      * port_fix.yml (file)

#### Script usage
This script must be run with parameter: 
* `-r <resource_folder_name>` to generate a resource
* `-d <datasource_folder_name>` to generate a datasource

It is possible to generate multiple resources / datasources at the same time, e.g.
```bash
10_gen_provider.sh -r resource_one -r resource_two -d datasource_one -d datasource_two
```

