import time
import picamera

picDir = "./"
timeStr = time.strftime("%Y%m%d-%H%M%S")
filename = picDir + "header" + timeStr + ".jpg"

with picamera.PiCamera() as camera:
    camera.resolution = (128,128)
    camera.start_preview()
    time.sleep(2)
    camera.capture(filename)

