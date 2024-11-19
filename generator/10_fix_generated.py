import sys
import getopt
import yaml

def fix(fix_file:str, provider_path:str):
    with open(fix_file, "r") as f:
        files = yaml.load(f, Loader=yaml.loader.SafeLoader)
    for tf_file, actions in files.items():
        tf_file_path = f"{provider_path}/{tf_file}"
        with open(tf_file_path, "r") as f_in:
            with open(f"{tf_file_path}.bak", "w") as f_bak:
                f_bak.writelines(f_in)
        with open(tf_file_path, "w") as f_out:
            rename = actions.get("rename", {})
            dedup = actions.get("dedup", [])
            with open(f"{tf_file_path}.bak", "r") as f_in:
                SKIP = False
                UNSKIP_NEXT = False
                DONE = []
                for line in f_in.readlines():
                    for src_string, dst_string in rename.items():
                        line = line.replace(src_string, dst_string)
                    if UNSKIP_NEXT:
                        UNSKIP_NEXT = False
                        SKIP = False
                    if line.startswith("}"):
                        UNSKIP_NEXT = True
                    if line.startswith("type") or line.startswith("func"):
                        for dedup_string in dedup:
                            if (
                                f"{dedup_string}Type" in line
                                or f"{dedup_string}Value" in line
                            ):
                                if not line in DONE:
                                    DONE.append(line)
                                else:
                                    SKIP = True
                                    UNSKIP_NEXT = False
                    elif line.startswith("var"):
                        if line not in DONE:
                            DONE.append(line)
                        else:
                            SKIP = True
                            UNSKIP_NEXT = True
                    if not SKIP:
                        f_out.write(f"{line}")


###############################################################################
###############################################################################
#### SCRIPT ENTRYPOINT ####
if __name__ == "__main__":

    FIX_FILE=""
    PROVIDER_PATH=""
    try:
        opts, args = getopt.getopt(
            sys.argv[1:],
            "f:p:",
            [
                "fix_file=",
                "provider_path="
            ],
        )
    except getopt.GetoptError as err:
        print(err)

    for o, a in opts:
        if o in ["-f", "--fix_file"]:
            FIX_FILE = a
        elif o in ["-p", "--provider_path"]:
            PROVIDER_PATH = a
        else:
            assert False, "unhandled option"

    fix(FIX_FILE, PROVIDER_PATH)
