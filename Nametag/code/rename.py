
import os

file_path = './to_mark'
files = os.listdir(file_path)

for i, file_name in enumerate(files):
    
    old_path = os.path.join(file_path, file_name)
    
    new_file_nama = f'{(i + 1):>04}.png'
    new_path = os.path.join(file_path, new_file_nama)
    
    os.rename(old_path, new_path)
