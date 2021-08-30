import glob
from collections import namedtuple

import cv2
import numpy as np
from fastapi import HTTPException

from convert import PAI2NUM

Matching = namedtuple("Matching", ("number", "score"))

# parametters
WIDTH = 910
HEIGHT = 512
THRESHOLD = 0.7
BLACK_THRESHOLD = 20
DORA_THRESHOLD = 0.6
MAX_YOHAI = 70


def autocrop(image):
    """Crops any edges below or equal to threshold
    Crops blank image to 1x1.
    Returns cropped image.
    """
    if len(image.shape) == 3:
        flatImage = np.max(image, 2)
    else:
        flatImage = image
    assert len(flatImage.shape) == 2
    rows = np.where(np.max(flatImage, 0) > BLACK_THRESHOLD)[0]
    if rows.size:
        cols = np.where(np.max(flatImage, 1) > BLACK_THRESHOLD)[0]
        image = image[cols[0]: cols[-1] + 1, rows[0]: rows[-1] + 1]
    else:
        image = image[:1, :1]
    return image


def match_hands(hand_templates, img_rgb):
    # この辞書で各座標の最大マッチング牌を管理
    coordinate_dict = {
        107: Matching(-1, 0),
        152: Matching(-1, 0),
        197: Matching(-1, 0),
        242: Matching(-1, 0),
        287: Matching(-1, 0),
        332: Matching(-1, 0),
        377: Matching(-1, 0),
        422: Matching(-1, 0),
        467: Matching(-1, 0), 
        512: Matching(-1, 0), 
        557: Matching(-1, 0), 
        602: Matching(-1, 0), 
        647: Matching(-1, 0), 
        706: Matching(-1, 0), 
    }
    for template in hand_templates:
        res = cv2.matchTemplate(img_rgb, template.img_array, cv2.TM_CCOEFF_NORMED)
        loc = np.where(res >= THRESHOLD)

        for pt in zip(*loc[::-1]):
            if pt[1] == 444 and pt[0] in coordinate_dict:
                score = res[pt[1]][pt[0]]
                if coordinate_dict[pt[0]].score < score:
                    coordinate_dict[pt[0]] = Matching(template.number, score)
    
    return coordinate_dict


def match_bakaze_jikaze(templates, img_rgb):
    matching = Matching(-1, THRESHOLD)
    for template in templates:
        res = cv2.matchTemplate(img_rgb, template.img_array, cv2.TM_CCOEFF_NORMED)
        score = np.max(res)
        if score > matching.score:
            matching = Matching(template.number, score)

    return matching


def match_dora(dora_templates, img_rgb):
    # この辞書で各座標の最大マッチング牌を管理
    coordinate_dict = {
        15: Matching(-1, 0),
        41: Matching(-1, 0),
        67: Matching(-1, 0),
        94: Matching(-1, 0),
        118: Matching(-1, 0)
    }
    for template in dora_templates:
        res = cv2.matchTemplate(img_rgb, template.img_array, cv2.TM_CCOEFF_NORMED)
        loc = np.where( res >= DORA_THRESHOLD)
        for pt in zip(*loc[::-1]):
            if pt[1] == 27 and pt[0] in coordinate_dict:
                score = res[pt[1]][pt[0]]
                if coordinate_dict[pt[0]].score < score:
                    coordinate_dict[pt[0]] = Matching(template.number, score)

    return coordinate_dict


def match_yohai(yohai_templates, img_rgb):
    # 座標系のずれが危ういので、2ドット分でカバー
    coordinate_dict = {
        "450 <= x <= 452": Matching(-1, 0),
        "462 <= x <= 464": Matching(-1, 0),
    }
    matching = Matching(-1, THRESHOLD)
    for template in yohai_templates:
        res = cv2.matchTemplate(img_rgb, template.img_array, cv2.TM_CCOEFF_NORMED)
        loc = np.where( res >= THRESHOLD)
        for pt in zip(*loc[::-1]):
            x = pt[0]
            for k, v in coordinate_dict.items():
                if pt[1] == 204 and eval(k):
                    score = res[pt[1]][pt[0]]
                    if coordinate_dict[k].score < score:
                        coordinate_dict[k] = Matching(template.number, score)
    yohai = int(''.join(v.number for _, v in coordinate_dict.items()))

    return yohai


def get_junme(yohai_templates, img_rgb, jikaze):
    yohai = match_yohai(yohai_templates, img_rgb)
    if jikaze == PAI2NUM['Ton']:
        margin = 3
    elif jikaze == PAI2NUM['Nan']:
        margin = 2
    elif jikaze == PAI2NUM['Sya']:
        margin = 1
    elif jikaze == PAI2NUM['Pe']:
        margin = 0
    else:
        raise HTTPException(status_code=500, detail="can not recoginize jikaze")
    junme = (MAX_YOHAI + margin - yohai) // 4
    # Analyzerの有効な巡目が1~17なので、調整（鳴きが発生するとどうしても正確には対応できなさそう）
    junme = min(junme, 17)
    junme = max(junme, 1)
    return junme


def extract(
    recognition_img, 
    hand_templates, 
    bakaze_templates, 
    jikaze_templates, 
    dora_templates, 
    yohai_templates):

    # crop and resize
    img_rgb = autocrop(recognition_img)
    if round(img_rgb.shape[1] / img_rgb.shape[0], 2) != 1.78:
        raise HTTPException(
            status_code=400, 
            detail=f"Expected image aspect ratio is 1.77... but the given image aspect ratio is {img_rgb.shape[1] / img_rgb.shape[0]}")
    img_rgb = cv2.resize(img_rgb, (WIDTH, HEIGHT))

    # 手牌の抽出
    hands = match_hands(hand_templates, img_rgb)
    # 場風の抽出
    bakaze = match_bakaze_jikaze(bakaze_templates, img_rgb[:255,:544,:])
    # 自風の抽出
    jikaze = match_bakaze_jikaze(jikaze_templates, img_rgb[:255,:544,:])
    # ドラの抽出
    dora = match_dora(dora_templates, img_rgb[:90,:154,:])
    # 巡目の計算
    junme = get_junme(yohai_templates, img_rgb[:255,:544,:], jikaze.number)

    resp = {
        'zikaze': jikaze.number,
        'bakaze': bakaze.number,
        'turn': junme,
        'dora_indicators': [v.number for _, v in dora.items() if v.number != -1],
        'hand_tiles': [v.number for _, v in hands.items() if v.number != -1],
    }
    return resp
