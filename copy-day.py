import os
import sys
import shutil

day_number = sys.argv[1]

dirname = "day-{}".format(day_number)
print("Creating {} directory...".format(dirname))
os.mkdir(dirname)

run_filename = "{}/day{}.py".format(dirname, day_number)
shutil.copy("templates/dayN.py", run_filename)
test_filename = "{}/day{}-test.py".format(dirname, day_number)
shutil.copy("templates/dayN-test.py", test_filename)

with open(test_filename) as file:
    filedata = file.read()

filedata = filedata.replace("dayN", "day{}".format(day_number))
filedata = filedata.replace("DayN", "Day{}".format(day_number))

with open(test_filename, "w") as file:
    file.write(filedata)
