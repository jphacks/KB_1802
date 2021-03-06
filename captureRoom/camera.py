import time
import picamera
import pexpect
import subprocess
import pymysql.cursors
import urllib.request

#for JPHACKS 2018


timeStr = time.strftime("%Y%m%d-%H%M%S")
filename = "header" + timeStr + ".jpg"

# take photo and save temp
with picamera.PiCamera() as camera:
    camera.resolution = (1280,720)
    camera.start_preview()
    time.sleep(2)
    camera.capture(filename)

#file send
# very danger code
"""
serverPath = "/home/ubuntu/camData"
cmd = "scp " + "./" + filename + " ubuntu@52.197.145.249:" + serverPath
scp = pexpect.spawn(cmd)
scp.expect("Enter passphrase")
scp.sendline("hogehogehugagahuga")
scp.interact()
scp.close()
"""

#register photo to SQL
captureTime = time.strftime("%Y-%m-%d %H:%M:%S")
conn = pymysql.connect(host = '52.197.145.249',
                        user = 'soisy',
                        password = 'boaboa',
                        db = 'soisy',
                        charset = 'utf8mb4',
                        cursorclass = pymysql.cursors.DictCursor)

try :
    with conn.cursor() as cursor :
        sql = "INSERT INTO camData (filepath,capturedate,isClean,dirtiness) VALUES(%s,%s,%s,%s)"
        cursor.execute(sql,(serverPath+"/"+filename,captureTime,0,0.2))
    conn.commit()
finally:
    conn.close()

#cleaning
cmd = "rm " + filename
subprocess.call(cmd.split())

#send HTTP Request
url = 'http://52.197.145.249/pyTest.py'
req = urllib.request.Request(url)
with urllib.request.urlopen(req) as res :
    body = res.read()
print (body)
