FILES = [
    "./terraform-provider-mistapi/internal/resource_networktemplate/networktemplate_resource_gen.go"
]


for file in FILES:
    with open(file, "r") as f_in:
        with open(f"{file}.bak", "w") as f_bak:
            f_bak.writelines(f_in)
    with open(f"{file}.bak", "r") as f_in:
        with open(file, "w") as f_out:
            SKIP = False
            UNSKIP_NEXT = False
            done = []
            for line in f_in.readlines():
                if UNSKIP_NEXT:
                    UNSKIP_NEXT = False
                    SKIP = False
                if line.startswith("}"):
                    UNSKIP_NEXT = True
                if line.startswith("var") or line.startswith("func"):
                    if line not in done:
                        done.append(line)
                    else:
                        SKIP = True
                        print(f"skipping {line}", end="")
                elif line.startswith("type"):
                    if line not in done:
                        done.append(line)
                    else:
                        SKIP = True
                        UNSKIP_NEXT = False
                        print(f"skipping {line}", end="")
                if not SKIP:
                    f_out.write(f"{line}")

