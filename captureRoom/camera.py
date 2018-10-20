import time
import picamera
import subprocess

picDir = "./"
timeStr = time.strftime("%Y%m%d-%H%M%S")
filename = picDir + "header" + timeStr + ".jpg"

# take photo and save temp
with picamera.PiCamera() as camera:
    camera.resolution = (128,128)
    camera.start_preview()
    time.sleep(2)
    camera.capture(filename)

cmd = "ls -l"
result = subprocess.call(cmd.split())
if result == 0 :
    print("send succeeded!")
else
    print("send failed.")
    quit()


