
import os
import random
import shutil

folder_source = './to_mark'
folder_destination = './random'

num_image_to_sample = 100

images_all = [f for f in os.listdir(folder_source) if os.path.isfile(os.path.join(folder_source, f))]

sampled_images = random.sample(images_all, num_image_to_sample)

for image in sampled_images:
    path_src = os.path.join(folder_source, image)
    path_dst = os.path.join(folder_destination, image)
    shutil.copy(path_src, path_dst)


