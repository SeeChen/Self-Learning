
import os
import cv2

path_video = 'sample.mp4'
path_save_frame = 'frames'

cap = cv2.VideoCapture(path_video)

count_frame = 0
while cap.isOpened():
    ret, frame = cap.read()
    if not ret:
        break
    
    filename_frame = os.path.join(path_save_frame, f'frame_{count_frame:04d}.png')
    cv2.imwrite(filename_frame, frame)
    
    count_frame += 1
    
cap.release()