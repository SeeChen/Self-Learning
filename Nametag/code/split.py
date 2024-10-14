
import os
import shutil
import random

all_images_dir = './all/images'
all_labels_dir = './all/labels'

train_images_dir = './dataset/images/train'
train_labels_dir = './dataset/labels/train'

valid_images_dir = './dataset/images/val'
valid_labels_dir = './dataset/labels/val'

images_all = os.listdir(all_images_dir)

ratio_train = 0.8
ratio_valid = 0.2

random.shuffle(images_all)

train_size = int(len(images_all) * ratio_train)

images_train = images_all[:train_size]
images_valid = images_all[train_size:]

def move_files(images_list, dest_images_dir, dest_labels_dir):
    for img in images_list:
        
        path_image = os.path.join(all_images_dir, img)
        path_label = os.path.join(all_labels_dir, img.replace('.png', '.txt'))
        
        shutil.copy(path_image, dest_images_dir)
        shutil.copy(path_label, dest_labels_dir)
        

move_files(images_train, train_images_dir, train_labels_dir)
move_files(images_valid, valid_images_dir, valid_labels_dir)
