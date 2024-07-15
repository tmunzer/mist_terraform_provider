**This is a dev repo. Please go to [terraform-provider-mist](https://github.com/juniper/terraform-provider-mist) to get the provider**


# Terraform Provider for Juniper Mist


### How to use the pre-version
1. Make sure to have Go installed: https://go.dev/doc/install
2. Install Terraform: https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli
3. create a `.terraformrc` file in your home folder with (replace `<home_folder_path>` with your actual home folder paht):
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/juniper/mist" = "<home_folder_path>/go/bin/",

  }
  direct {}
}
```
3. Clone the repository
4. From the `terraform-provider-mist` folder, do 
  * `go mod tidy` to install the depencies
  * `go install .` to install the provider
5. Create a configuration file (you can use the one from the `test` folder)

**note**
It is recommended to store the provider configuration (api host and api token) in an environment file.
To do so, create a `.env` file in the `test` folder (or in the folder storing the configuration to apply to Mist) and define the settings like this:
```
HOST=api.mist.com
APITOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```



Documentation is currently available as Markdown in the `terraform-provider-mist/docs` folder
