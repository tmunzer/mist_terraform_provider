import os
import json
import yaml

FIX_PATH = "./provider-code-spec-fix"
FIX_FOLDERS = ["datasources", "resources"]


for folder in FIX_FOLDERS:
    FOLDER_PATH = os.path.join(FIX_PATH, folder)
    FIX_FILES = os.listdir(FOLDER_PATH)
    for file in FIX_FILES:
        fix_in_path = os.path.join(FOLDER_PATH, file)
        fix_out_path = os.path.join(FOLDER_PATH, file.replace(".json", ".yaml"))
        print(f" {fix_in_path} ".ljust(80, "."))
        with open(fix_in_path, "r") as fix_in:
            fix = json.load(fix_in)
            fix_out = fix_in.
            with open(fix_out_path, "w") as fix_out:
                yaml.dump(fix, fix_out)