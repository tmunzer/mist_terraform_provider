import re
SPEC_IN = "./mist_openapi/workdir/mist.openapi.yml"
SPEC_OUT = "./mist.sdk.yml"
REG = r"^( *)enum:$"

with open(SPEC_IN, "r") as f_in:
    with open(SPEC_OUT, "w") as f_out:
        f_lines = f_in.readlines()
        for line in f_lines:
            f_out.write(line)
            res = re.match(REG, line)
            if res:
                f_out.write(f"{res.groups()[0]}- \"\"\n")

