import os

path_root = 'mark'

path_images = os.path.join(path_root, 'images')
path_labels = os.path.join(path_root, 'labels')

files_images = sorted([f for f in os.listdir(path_images)])
files_labels = sorted([f for f in os.listdir(path_labels)])

for idx, file_image in enumerate(files_images):
    
    path_image = os.path.join(path_images, file_image)
    path_label = os.path.join(path_labels, file_image.replace('.png', '.txt'))
    
    if os.path.exists(path_label):
        
        new_name = f'{idx:>04}'
        
        new_image = f'{new_name}.png'
        new_label = f'{new_name}.txt'
        
        new_path_image = os.path.join(path_images, new_image)
        new_path_label = os.path.join(path_labels, new_label)
        
        os.rename(path_image, new_path_image)
        os.rename(path_label, new_path_label)
        
        print(f'Rename {path_image} to {new_path_image}')
        print(f'Rename {path_label} to {new_path_label}')
