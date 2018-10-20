import cv2
import numpy as np

import matplotlib
matplotlib.use("Agg")
import matplotlib.pyplot as plt

def main():
    # 入力画像の読み込み
    img0 = cv2.imread("photo0.jpg")
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


if __name__ == "__main__":
    main()