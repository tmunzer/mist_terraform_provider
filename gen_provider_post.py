FILES = [
    "./terraform-provider-mist/internal/resource_networktemplate/networktemplate_resource_gen.go",
    "./terraform-provider-mist/internal/resource_network/network_resource_gen.go",
    "./terraform-provider-mist/internal/resource_gatewaytemplate/gatewaytemplate_resource_gen.go",
    "./terraform-provider-mist/internal/resource_rftemplate/rftemplate_resource_gen.go",
]

ADD = {
    # "./terraform-provider-mist/internal/resource_networktemplate/networktemplate_resource_gen.go": [
    #     "./tacacsAcctServers.fix",
    #     #      "./switchMatchingRules.fix",
    # ]
}

REPLACE_ITEMS = {
    # "": {
    #     # "Tacacs": [{"source": "AcctServers", "dest": "TacacsAcctServers"}],
    #     "*": [{"source": "matching_rules", "dest": "rules"}]
    # }
}

for file in FILES:
    with open(file, "r") as f_in:
        with open(f"{file}.bak", "w") as f_bak:
            f_bak.writelines(f_in)
    with open(file, "w") as f_out:
        with open(f"{file}.bak", "r") as f_in:
            SKIP = False
            UNSKIP_NEXT = False
            TO_REPLACE = []
            done = []
            for line in f_in.readlines():
                if UNSKIP_NEXT:
                    UNSKIP_NEXT = False
                    SKIP = False

                for item in TO_REPLACE:
                    line = line.replace(item["source"], item["dest"])

                if line.startswith("}"):
                    UNSKIP_NEXT = True
                    TO_REPLACE = []
                elif line.startswith("type") or line.startswith("func"):
                    for item, value in REPLACE_ITEMS.items():
                        if f"{item}Type" in line or f"{item}Value" in line:
                            TO_REPLACE = value
                    if line not in done:
                        done.append(line)
                    else:
                        SKIP = True
                        print(f"skipping {line}", end="")
                elif line.startswith("var"):
                    if line not in done:
                        done.append(line)
                    else:
                        SKIP = True
                        UNSKIP_NEXT = False
                        print(f"skipping {line}", end="")
                if not SKIP:
                    f_out.write(f"{line}")
        for fix in ADD.get(file, []):
            with open(fix, "r") as f_fix:
                f_out.writelines(f_fix.readlines())
