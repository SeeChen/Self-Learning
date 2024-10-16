
import os
import cv2
import random
import shutil

from ultralytics import YOLO

def train():

    model = YOLO('best_v1.pt')
    
    model.train(
        data='data.yaml',
        epochs=1000,
        imgsz=640,
        batch=8,
        workers=16,
        name='train_model',
        save=True
    )
    
    model.export(format='onnx')
    

def check_model():
    model = YOLO('best_v1.pt')
    # print(model.model)
    # print(len(list(model.model.children())))
   
    
def detect_video():

    model = YOLO('best_v2.pt')

    video_path = 'sample.mp4'
    cap = cv2.VideoCapture(video_path)

    fps = cap.get(cv2.CAP_PROP_FPS)
    width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
    height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
    
    output_path = 'result.mp4'
    fourcc = cv2.VideoWriter_fourcc(*'mp4v')
    out = cv2.VideoWriter(output_path, fourcc, fps, (width, height))

    if not cap.isOpened():
        print("Error: Could not open video.")
        exit()

    while True:
        ret, frame = cap.read()
        if not ret:
            break

        results = model(frame)
        annotated_frame = results[0].plot()

        out.write(annotated_frame)

        cv2.imshow('Detection', annotated_frame)
        if cv2.waitKey(1) & 0xFF == ord('q'):
            break

    cap.release()
    out.release()
    cv2.destroyAllWindows()
    
    
def detect_all_frame():

    model = YOLO('best_v2.pt')
    
    frames_all_path = 'tt'
    frames = os.listdir(frames_all_path)
    
    path_output = 'output'
    
    for frame in frames:
        
        path_image = os.path.join(frames_all_path, frame)
        image = cv2.imread(path_image)
        
        results = model(image)
        
        path_output_txt = os.path.join(path_output, frame.replace('png', 'txt'))
        
        with open(path_output_txt, 'w') as f:
            
            for result in results:
                for box in result.boxes:
                    class_id = int(box.cls)
                    
                    confidence = box.conf.item()
                    x_center = box.xywh[0, 0].item()
                    y_center = box.xywh[0, 1].item()
                    width = box.xywh[0, 2].item()
                    height = box.xywh[0, 3].item()
                    
                    h, w, _ = image.shape
                    x_center /= w
                    y_center /= h
                    width /= w
                    height /= h
    
                    f.write(f"{class_id} {x_center} {y_center} {width} {height}\n")


def devide_data():

    all_images_dir = './mark/images'
    all_labels_dir = './mark/labels'

    train_images_dir = './dataset/train/images'
    train_labels_dir = './dataset/train/labels'

    valid_images_dir = './dataset/valid/images'
    valid_labels_dir = './dataset/valid/labels'

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


if __name__ == '__main__':
    
    # train()
    # detect_video()
    # detect_all_frame()
    devide_data()
    # check_model()
    pass
