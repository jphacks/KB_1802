import cv2
import numpy as np

import matplotlib
import requests

matplotlib.use("Agg")
import matplotlib.pyplot as plt
from pylab import *

import mysql.connector

from urllib.parse import urlparse

PATH="home/ubuntu/result"

def main():
    # 元画像
    img0 = sqlreadold()
    #最新画像
    img1 = sqlreadnew()

    res0=Quantization(img0)
    res1=Quantization(img1)

    cv2.imwrite("res0.jpg",res0)
    cv2.imwrite("res1.jpg",res1)


    edge0=Edgepic(res0)
    edge1=Edgepic(res1)

    cv2.imwrite("res0.jpg",edge0)
    cv2.imwrite("res1.jpg",edge1)

    edgenum0=Edgenum(res0)
    edgenum1=Edgenum(res1)

    diffpic=diff(edge0,edge1)

    cv2.imwrite("diff2.jpg",diffpic)

    heatdata=mkheatdata(diffpic)

    heatmap(heatdata)

    result = overray(img1)

    nowtime=str(datetime.now().strftime("%Y/%m/%d-%H:%M:%S"))
    cv2.imwrite(PATH+nowtime,result)

    sqlregist(PATH,nowtime)

def Quantization(img):
    # 画像の量子化
    Z = img.reshape((-1, 3))

    # convert to np.float32
    Z = np.float32(Z)

    # define criteria, number of clusters(K) and apply kmeans()
    criteria = (cv2.TERM_CRITERIA_EPS + cv2.TERM_CRITERIA_MAX_ITER, 10, 1.0)
    K = 8
    ret, label, center = cv2.kmeans(Z, K, None, criteria, 10, cv2.KMEANS_RANDOM_CENTERS)

    # Now convert back into uint8, and make original image
    center = np.uint8(center)
    res = center[label.flatten()]
    res2 = res.reshape((img.shape))
    return res2

def Edgepic(img):
    #エッジを出してみる
    edge=cv2.Canny(img,100,600)
    return edge

def Edgenum(img):
    return np.sum(img==255)

def diff(img0,img1):
    # グレースケール変換
    #gray0 = cv2.cvtColor(img0, cv2.COLOR_BGR2GRAY)
    #gray1 = cv2.cvtColor(img1, cv2.COLOR_BGR2GRAY)

    # フレームの絶対差分
    diff = cv2.absdiff(img0, img1)
    return diff

def heatmap(heatdata):
    fig,ax=plt.subplots()
    fig=plt.figure(figsize=(128,72),dpi=10)
    ax.tick_params(labelbottom="off", bottom="off")  # x軸の削除
    ax.tick_params(labelleft="off", left="off")  # y軸の削除
    ax.set_xticklabels([])
    box("off")  # 枠線の削除
    fig.subplots_adjust(left=0,bottom=0,right= 1,top=1)

    plt.pcolor(heatdata,cmap=plt.cm.Reds)
    plt.savefig('heatmap.png')

def mkheatdata(img):
    N=10
    M=10
    y=int(np.shape(img)[1]/N)
    x=int(np.shape(img)[0]/M)
    array=np.zeros((N,M))
    cv2.imwrite("img.jpg",img)
    for n in range(N):
        for m in range(M):
            array[n][m]=np.sum(img[x*n:x*(n+1)-1,y*m:y*(m+1)-1])
    array=array[::-1][::1]
    return array

def overray(img):
    heatpic = cv2.imread("heatmap.png")

    result = cv2.addWeighted(img, 0.7, heatpic, 0.3, 1)
    #cv2.imwrite("result.jpg", result)
    return result


def draw_heatmap(x, y):
    heatmap, xedges, yedges = np.histogram2d(x, y, bins=50)
    extent = [xedges[0], xedges[-1], yedges[0], yedges[-1]]

    plt.figure()
    plt.imshow(heatmap, extent=extent)
    plt.show()
    plt.savefig('image.png')

def post():
    x=str(2)
    r=requests.post("http://52.197.145.249:9000/dirtinessCheck?level="+x)
    print(r)


def sqlreadnew():
    url=urlparse("mysql://soisy:boaboa@52.197.145.249")
    conn = mysql.connector.connect(
        host=url.hostname or 'localhost',
        port=url.port or 3306,
        user=url.username or 'root',
        password=url.password or '',
        database=url.path[1:],
        charset='utf8',
    )
    print(conn.is_connected())

    cur=conn.cursor()
    cur.execute("USE soisy")
#    cur.execute("UPDATE camData SET isClean=1 ")

    cur.execute("SELECT * FROM camData as m WHERE NOT EXISTS "
                    "(SELECT 1 FROM camData as s WHERE m.CAPTUREDATE>s.CAPTUREDATE)")
    add=str(cur.fetchall()[0][0])

    print(add)
    return cv2.imread(add)

def sqlreadold():
    url=urlparse("mysql://soisy:boaboa@52.197.145.249")
    conn = mysql.connector.connect(
        host=url.hostname or 'localhost',
        port=url.port or 3306,
        user=url.username or 'root',
        password=url.password or '',
        database=url.path[1:],
        charset='utf8',
    )
    print(conn.is_connected())

    cur=conn.cursor()
    cur.execute("USE soisy")
#    cur.execute("UPDATE camData SET isClean=1 ")

    cur.execute("SELECT * FROM camData as m WHERE NOT EXISTS "
                    "(SELECT 1 FROM camData as s WHERE isClean=1)")
    add=str(cur.fetchall()[0][0])

    print(add)
    return cv2.imread(add)


def sqlregist(path,nowtime):
    url=urlparse("mysql://soisy:boaboa@52.197.145.249")
    conn = mysql.connector.connect(
        host=url.hostname or 'localhost',
        port=url.port or 3306,
        user=url.username or 'root',
        password=url.password or '',
        database=url.path[1:],
        charset='utf8',
    )
    print(conn.is_connected())

    cur=conn.cursor()
    cur.execute("USE soisy")
#    cur.execute("UPDATE camData SET isClean=1 ")

    cur.execute("SELECT * FROM resultData as m WHERE NOT EXISTS "
                    "(SELECT 1 FROM resultData as s WHERE m.resultDate>s.resultDate)")
    i=int(cur.fetchall()[0][0])

    if i>5:
        post()
        i=0
    else:
        i+=1

    cur.execute("INSERT INTO resultData (CleanSeq,ResultDate,ResultFilePath) VALUES ("+i+","+nowtime+","+path+nowtime+")")

    print(add)

if __name__ == "__main__":
    main()
    post()


