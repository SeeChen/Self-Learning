from ultralytics import YOLO

def training():

    model = YOLO('./yolov8n.pt')
    
    model.train(
        data='data.yaml',
        epochs=1000,
        imgsz=640,
        batch=32,
        workers=16,
        name='train_model'
    )
    
    model.export(format='onnx')

if __name__ == '__main__':
    
    training()
