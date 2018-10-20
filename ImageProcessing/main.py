import cv2
import numpy as np

import matplotlib
matplotlib.use("Agg")
import matplotlib.pyplot as plt
from pylab import *

def main():
    # 入力画像の読み込み
    img0 = cv2.imread("photo1.jpg")
<<<<<<< HEAD
    img1 = cv2.imread("photo2.jpg")

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

    overray(img1)





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
    """
    print(heatpic)
    for i in range(len(heatpic)):
        for j in range(len(heatpic[i])):
            if heatpic[i][j][0]>220 and heatpic[i][j][1]>220 and heatpic[i][j][2]>220:
                heatpic[i][j]=[0,0,0,0]
            else:
                heatpic[i][j][3]=180
    cv2.imwrite("heatmap2.png",heatpic)
    heatpic=cv2.imread("heatmap2.png")
    """
    result = cv2.addWeighted(img, 0.7, heatpic, 0.3, 1)
    cv2.imwrite("result.jpg", result)

def main2():
    # 入力画像の読み込み
    img0 = cv2.imread("photo0.jpg")
=======
>>>>>>> master
    img1 = cv2.imread("photo2.jpg")

    #エッジを出してみる
    edge=cv2.Canny(img1,100,200)
    print(np.sum(edge==255))
    cv2.imwrite("edge.jpg",edge)

    #画像の量子化
    Z = img1.reshape((-1, 3))

    # convert to np.float32
    Z = np.float32(Z)

    # define criteria, number of clusters(K) and apply kmeans()
    criteria = (cv2.TERM_CRITERIA_EPS + cv2.TERM_CRITERIA_MAX_ITER, 10, 1.0)
    K = 8
    ret, label, center = cv2.kmeans(Z, K, None, criteria, 10, cv2.KMEANS_RANDOM_CENTERS)

    # Now convert back into uint8, and make original image
    center = np.uint8(center)
    res = center[label.flatten()]
    res2 = res.reshape((img0.shape))
    cv2.imwrite("res8.jpg",res2)

    picSize=1280*720

    # グレースケール変換
    gray0 = cv2.cvtColor(img0, cv2.COLOR_BGR2GRAY)
    gray1 = cv2.cvtColor(img1, cv2.COLOR_BGR2GRAY)

    # フレームの絶対差分
    diff = cv2.absdiff(gray0, gray1)
    cv2.imwrite("diff.jpg", diff)
    print(diff)
    print(np.sum(diff))

    x=np.shape(diff)[0]
    y=np.shape(diff)[1]

    picx=np.shape(diff)[0]/x
    picy=np.shape(diff)[1]/y


    heatmap=np.zeros((x,y))
    pro=0
    for xi in range(x):
        for yi in range(y):
            if diff[xi][yi]<50:
                diff[xi][yi]=0
                #print(diff[xi][yi])
                pro+=1
            else:
                diff[xi][yi]=255
    print(pro)
    print((picSize-pro)/picSize*100)

    cv2.imwrite("diff2.jpg",diff)
    print(diff)

def draw_heatmap(x, y):
    heatmap, xedges, yedges = np.histogram2d(x, y, bins=50)
    extent = [xedges[0], xedges[-1], yedges[0], yedges[-1]]

    plt.figure()
    plt.imshow(heatmap, extent=extent)
    plt.show()
    plt.savefig('image.png')

if __name__ == "__main__":
    main()

