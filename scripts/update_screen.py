#!/usr/bin/env python3

import sys

from PIL import Image
from inky.inky_ac073tc1a import Inky

inky = Inky()
saturation = 0.8

if len(sys.argv) == 1:
    print("""
Usage: {file} image-file
""".format(file=sys.argv[0]))
    sys.exit(1)

image = Image.open(sys.argv[1])

width, height = image.size
print(f'image size: {width}x{height}')

image = image.rotate(90, expand=1)
if width != 800 or height != 480:
    print("resizing to 800x480")
    # image = image.resize((800, 480), Image.BICUBIC)
    image = image.resize(inky.resolution)

if len(sys.argv) > 2:
    saturation = float(sys.argv[2])

inky.set_image(image, saturation=saturation)
inky.show()
