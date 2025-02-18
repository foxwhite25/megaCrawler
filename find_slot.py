import glob
import os
import os.path

for file in glob.glob("./plugins/storage/*.go"):
    if "keep" in file:
        continue
    for n in range(1000, 9999):
        if not os.path.isfile(f"./plugins/production/{n}.go"):
            os.rename(file, f"./plugins/production/{n}.go")
            print(f"./plugins/production/{n}.go")
            break