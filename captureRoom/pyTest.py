#!/usr/bin/python

print ("Content-Type:text/html\n\n")

import time
import subprocess

#python3    /home/ubuntu/KB_1802/ImageProcessing/main.py 

cmd = "python3 /home/ubuntu/KB_1802/ImageProcessing/main.py"
print (time.time())
1#subprocess.call(cdCmd,shell = True)
ret = subprocess.call(cmd, shell = True)
print (cmd + " : " + str(ret))

