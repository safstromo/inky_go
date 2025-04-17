#!/usr/bin/env python3

import sys

from PIL import Image
from inky.inky_ac073tc1a import Inky

inky = Inky()
saturation = 0.8
target_width = 800
target_height = 480

if len(sys.argv) == 1:
    print("""
Usage: {file} image-file [saturation]
""".format(file=sys.argv[0]))
    sys.exit(1)

image_path = sys.argv[1]
try:
    image = Image.open(image_path)
except FileNotFoundError:
    print(f"Error: Image file not found at {image_path}")
    sys.exit(1)
except Exception as e:
    print(f"Error opening image file: {e}")
    sys.exit(1)

original_width, original_height = image.size
print(f'Original image size: {original_width}x{original_height}')

# Rotate the original image by 90 degrees
rotated_image = image.rotate(90, expand=True)
rotated_width, rotated_height = rotated_image.size
print(f'Rotated image size: {rotated_width}x{rotated_height}')

target_width = 800
target_height = 480

rotated_aspect = rotated_width / rotated_height
target_aspect = target_width / target_height

if rotated_aspect > target_aspect:
    # Rotated image is wider, fit to target height
    final_height = target_height
    final_width = int(target_height * rotated_aspect)
else:
    # Rotated image is taller or same aspect, fit to target width
    final_width = target_width
    final_height = int(target_width / rotated_aspect)

final_image = rotated_image.resize((final_width, final_height), Image.Resampling.LANCZOS)
print(f'Final image size: {final_image.width}x{final_image.height}')

# Create background and paste (centering)
background = Image.new("RGB", (target_width, target_height), "white")
paste_x = (target_width - final_width) // 2
paste_y = (target_height - final_height) // 2
background.paste(final_image, (paste_x, paste_y))

final_display_image = background

if len(sys.argv) > 2:
    try:
        saturation = float(sys.argv[2])
    except ValueError:
        print("Error: Saturation value must be a float.")
        sys.exit(1)

inky.set_image(final_display_image, saturation=saturation)
inky.show()
