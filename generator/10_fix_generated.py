import sys
import getopt
import yaml

def fix(fix_file:str):
    with open(fix_file, "r") as f:
        files = yaml.load(f, Loader=yaml.loader.SafeLoader)
    for file, actions in files.items():
        with open(file, "r") as f_in:
            with open(f"{file}.bak", "w") as f_bak:
                f_bak.writelines(f_in)
        with open(file, "w") as f_out:
            rename = actions.get("rename", {})
            dedup = actions.get("dedup", [])
            with open(f"{file}.bak", "r") as f_in:
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

    try:
        opts, args = getopt.getopt(
            sys.argv[1:],
            "f:",
            [
                "fix_file=",
            ],
        )
    except getopt.GetoptError as err:
        print(err)

    for o, a in opts:
        if o in ["-f", "--fix_file"]:
            FIX_FILE = a
        else:
            assert False, "unhandled option"

    fix(FIX_FILE)
