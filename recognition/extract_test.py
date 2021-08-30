import glob
from collections import namedtuple

import cv2
import numpy as np
import requests
from fastapi import HTTPException

Matching = namedtuple("Matching", ("name", "score"))

# parametters
input_dir = './inputs/'
output_dir = './outputs/'
# テンプレートディレクトリ
hand_dir = 'template/hand/'
bakaze_dir = 'template/bakaze/'
jikaze_dir = 'template/jikaze/'
dora_dir = 'template/dora/'
yohai_dir = 'template/yohai/'

WIDTH = 910
HEIGHT = 512
THRESHOLD = 0.7
BLACK_THRESHOLD = 20
DORA_THRESHOLD = 0.6
MAX_YOHAI = 70


def img2url(url):
    try:
        resp = requests.get(url, stream=True).raw
        img = np.asarray(bytearray(resp.read()), dtype="uint8")
        img_rgb = cv2.imdecode(img, cv2.IMREAD_COLOR)
        return img_rgb
    except:
        raise HTTPException(status_code=400, detail="Cannot load img url")


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
        107: Matching('', 0),
        152: Matching('', 0),
        197: Matching('', 0),
        242: Matching('', 0),
        287: Matching('', 0),
        332: Matching('', 0),
        377: Matching('', 0),
        422: Matching('', 0),
        467: Matching('', 0), 
        512: Matching('', 0), 
        557: Matching('', 0), 
        602: Matching('', 0), 
        647: Matching('', 0), 
        706: Matching('', 0), 
    }
    for template_name in hand_templates:
        template = cv2.imread(template_name)
        res = cv2.matchTemplate(img_rgb, template, cv2.TM_CCOEFF_NORMED)
        loc = np.where(res >= THRESHOLD)

        for pt in zip(*loc[::-1]):
            if pt[1] == 444 and pt[0] in coordinate_dict:
                name = template_name.replace(hand_dir, "").replace(".jpg", "")
                score = res[pt[1]][pt[0]]
                if coordinate_dict[pt[0]].score < score:
                    coordinate_dict[pt[0]] = Matching(name, score)
    
    return coordinate_dict


def match_bakaze(bakaze_templates, img_rgb):
    matching = Matching('', THRESHOLD)
    for template_name in bakaze_templates:
        template = cv2.imread(template_name)
        res = cv2.matchTemplate(img_rgb, template, cv2.TM_CCOEFF_NORMED)
        name = template_name.replace(bakaze_dir, "").replace(".jpg", "")
        score = np.max(res)
        if score > matching.score:
            matching = Matching(name, score)

    return matching


def match_jikaze(jikaze_templates, img_rgb):
    matching = Matching('', THRESHOLD)
    for template_name in jikaze_templates:
        template = cv2.imread(template_name)
        res = cv2.matchTemplate(img_rgb, template, cv2.TM_CCOEFF_NORMED)
        name = template_name.replace(jikaze_dir, "").replace(".jpg", "")
        score = np.max(res)
        if score > matching.score:
            matching = Matching(name, score)

    return matching


def match_dora(dora_templates, img_rgb):
    # この辞書で各座標の最大マッチング牌を管理
    coordinate_dict = {
        15: Matching('', 0),
        41: Matching('', 0),
        67: Matching('', 0),
        94: Matching('', 0),
        118: Matching('', 0)
    }
    for template_name in dora_templates:
        template = cv2.imread(template_name)
        res = cv2.matchTemplate(img_rgb, template, cv2.TM_CCOEFF_NORMED)
        loc = np.where( res >= DORA_THRESHOLD)
        for pt in zip(*loc[::-1]):
            if pt[1] == 27 and pt[0] in coordinate_dict:
                name = template_name.replace(dora_dir, "").replace(".jpg", "")
                score = res[pt[1]][pt[0]]
                if coordinate_dict[pt[0]].score < score:
                    coordinate_dict[pt[0]] = Matching(name, score)

    dora = []
    dora_rotate_dict = {
        'Ton': 'Nan',
        'Nan': 'Sya',
        'Sya': 'Pe',
        'Pe': 'Ton',
        'Haku': 'Hatu',
        'Hatu': 'Tyun',
        'Tyun': 'Haku'
    }
    for _, v in coordinate_dict.items():
        if v.name == '':
            continue
        if v.name in dora_rotate_dict:
            dora.append(dora_rotate_dict[v.name])
        elif "Aka" in v.name:
            dora.append(v.name.replace("Aka", "").replace("5", "6"))
        else:
            dora.append(v.name[:-1] + str(int(v.name[-1]) + 1)[0])

    return dora


def match_yohai(yohai_templates, img_rgb):
    # 座標系のずれが危ういので、2ドット分でカバー
    coordinate_dict = {
        "450 <= x <= 452": Matching('', 0),
        "462 <= x <= 464": Matching('', 0),
    }
    matching = Matching('', THRESHOLD)
    for template_name in yohai_templates:
        template = cv2.imread(template_name)
        res = cv2.matchTemplate(img_rgb, template, cv2.TM_CCOEFF_NORMED)
        loc = np.where( res >= THRESHOLD)
        name = template_name.replace(yohai_dir, "").replace(".jpg", "")
        for pt in zip(*loc[::-1]):
            x = pt[0]
            for k, v in coordinate_dict.items():
                if pt[1] == 204 and eval(k):
                    score = res[pt[1]][pt[0]]
                    if coordinate_dict[k].score < score:
                        coordinate_dict[k] = Matching(name, score)
    yohai = int(''.join(v.name for k, v in coordinate_dict.items()))

    return yohai


def get_junme(yohai_templates, img_rgb, jikaze):
    yohai = match_yohai(yohai_templates, img_rgb)
    if jikaze == 'Ton':
        margin = 3
    elif jikaze == 'Nan':
        margin = 2
    elif jikaze == 'Sya':
        margin = 1
    elif jikaze == 'Pe':
        margin = 0
    else:
        return -1
    return (MAX_YOHAI + margin - yohai) // 4


def extract_test(img_url):
    img_rgb = img2url(img_url)
    if type(img_rgb) != np.ndarray:
        return "Cannot load img url"

    # crop and resize
    img_rgb = autocrop(img_rgb)
    if round(img_rgb.shape[1] / img_rgb.shape[0], 2) != 1.78:
        return f"Expected image aspect ratio is 1.77... but the given image aspect ratio is {img_rgb.shape[1] / img_rgb.shape[0]}"
    img_rgb = cv2.resize(img_rgb, (WIDTH, HEIGHT))

    # 手牌の抽出
    hand_templates = glob.glob(f'{hand_dir}*.jpg')
    hands = match_hands(hand_templates, img_rgb)
    # 場風の抽出
    bakaze_templates = glob.glob(f'{bakaze_dir}*.jpg')
    bakaze = match_bakaze(bakaze_templates, img_rgb[:255,:544,:])
    # 自風の抽出
    jikaze_templates = glob.glob(f'{jikaze_dir}*.jpg')
    jikaze = match_jikaze(jikaze_templates, img_rgb[:255,:544,:])
    # ドラの抽出
    dora_templates = glob.glob(f'{dora_dir}*.jpg')
    dora = match_dora(dora_templates, img_rgb[:90,:154,:])
    # 巡目の計算
    yohai_templates = glob.glob(f'{yohai_dir}*.jpg')
    junme = get_junme(yohai_templates, img_rgb[:255,:544,:], jikaze.name)

    resp = {
        'hands': [v.name for _, v in hands.items()],
        'bakaze': bakaze.name,
        'jikaze': jikaze.name,
        'dora': dora,
        'junme': junme
    }
    return resp
