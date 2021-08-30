import glob
from collections import namedtuple

import cv2
import numpy as np
import requests
from fastapi import HTTPException

from convert import PAI2NUM

Template = namedtuple('Template', ('number', 'img_array'))

def load_templates(template_dir):
    ret = []
    templates = glob.glob(f'{template_dir}*.jpg')
    for template_name in templates:
        img_array = cv2.imread(template_name)
        name = template_name.replace(template_dir, "").replace(".jpg", "")
        if template_dir == 'template/yohai/':
            # あえて文字列
            ret.append(Template(name, img_array))
        else:
            ret.append(Template(PAI2NUM[name], img_array))

    return ret


def img2url(url):
    try:
        resp = requests.get(url, stream=True).raw
        img = np.asarray(bytearray(resp.read()), dtype="uint8")
        img_rgb = cv2.imdecode(img, cv2.IMREAD_COLOR)
        return img_rgb
    except:
        raise HTTPException(status_code=400, detail="Cannot load img url")
