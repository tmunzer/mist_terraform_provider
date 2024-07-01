# /!\ WORK IN PROGRESS /!\

# Terraform Provider for Juniper Mist

### How to use the pre-version
1. Make sure to have Go installed: https://go.dev/doc/install
2. create a `.terraformrc` file in your home folder with (replace `<home_folder_path>` with your actual home folder paht):
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/juniper/mist" = "<home_folder_path>/go/bin/",

  }
  direct {}
}
```
3. from the `terraform-provider-mist` folder, do `go install .` to install the provider
4. Create a configuration file (you can use the one from the `test` folder)

Documentation is currently available as Markdown in the `terraform-provider-mist/doc` folder
