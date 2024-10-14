import cv2
from ultralytics import YOLO

model = YOLO('yolov8n.pt')
model_target = YOLO('best.pt')

video_path = 'sample.mp4'
cap = cv2.VideoCapture(video_path)

fps = cap.get(cv2.CAP_PROP_FPS)
width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))

output_path = 'result.mp4'
fourcc = cv2.VideoWriter_fourcc(*'mp4v')
out = cv2.VideoWriter(output_path, fourcc, fps, (width, height))

while True:
    ret, frame = cap.read()
    if not ret:
        break

    result_personal = model(frame, classes=[0])
    personal_box = result_personal[0].boxes
    
    for box in personal_box:
        x1, y1, x2, y2 = box.xyxy[0]
        person_area = frame[int(y1):int(y2), int(x1):int(x2)]

        results_target = model_target(person_area)

        if results_target[0].boxes:
            
            cv2.rectangle(frame, (int(x1), int(y1)), (int(x2), int(y2)), (0, 255, 0), 2)

    out.write(frame)

    cv2.imshow('YOLOv8 Person Detection', frame)
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

cap.release()
out.release()
cv2.destroyAllWindows()
