<div align=center>

# REST Using Express JS
<!-- 
[Summary](#1-summary)</br>
[Environment Setup](#2-environment-setup)</br>
[Dataset Setup](#3-dataset-setup)</br>
[Model Training](#4-model-training)</br>
[Result](#5-result)</br>
[Running the Training Process](#6-running-the-training-process) -->

</div>

<!--
## 1. Summary
This task involved identifying and tracking a staff member wearing a nametag using [video](sample.mp4) footage from a 3D sensor. The goal was to determine which frames contained the staff member and locate their coordinates when present. I selected the YOLO (You Only Look Once) model for this task because of its fast object detection capabilities.

<!--
## 2. Environment Setup
|CATEGORY|DETAILS|
|---:|:---:|
|Platform|Windows 11|
|Base Model|YOLO (You Only Look Once)|
|Python|3.11|
|Environment|Conda environment|
|RAM|16 GB|
|Graphics Card|RTX 3050|
|Additional|CUDA 12.6|

### Environmental Documents
> 1. [environment.yml](./environment.yml)
> 2. [requirements.txt](./requirements.txt)

## 3. Dataset Setup
The dataset is split 80/20, with 80 percent of the images used for the training dataset and 20 percent for the validation dataset. All training in this project follows this split ratio.

The dataset is structured as follows for all training:
```
dataset
├── train
│   ├── images
│   └── labels
└── valid
    ├── images
    └── labels
```


All data is configured for training as specified in the [data.yml](./data.yaml) file.

## 4. Model Training
Regarding the training approach, I am using a self-iterative training method. This approach involves multiple iterations of model training and refinement based on progressively extracted data from previous outcomes.

Self-iteration refers to a process where the model continuously uses its own outputs as new training data in subsequent training rounds. After each round, the model’s predictions are used to adjust the training data, either by adding newly predicted labels or refining the data. The model then retrains with the updated data, iterating this process. Over time, the model improves by learning from its own results, essentially "bootstrapping" its knowledge. This method is a form of semi-supervised learning, starting with a small labeled dataset and gradually expanding the training set by using the model’s predictions on unlabeled data to enhance its performance.

### 4.1 Initial Training
First of all, I randomly selected 370 pictures from a total of 1370 frames, and then annotated the nametags, person, and staff. The initial training aimed to establish a baseline for object detection.

### 4.2 First Detection and Refinement
After the initial model training, the model was used to detect objects in the video frames. Based on the detection results, I extracted a set of images that were either correctly or incorrectly detected and re-annotated them to correct the labels.

After that, similar to the first step, I randomly sampled 100 images for training.

### 4.3 Subsequent Iterative Training
The second round of training used the refined dataset generated from the first detection. Each subsequent round of training used a refined dataset generated from the previous round's detection results. 

This iterative process of training, detecting, and refining continued through several rounds, with each iteration aimed at enhancing the model’s ability to accurately detect the objects of interest.

## 5. Result
The performance of the model was evaluated using a variety of metrics and visualizations, providing insight into both its accuracy and areas for improvement. Below is a breakdown of the key results:

### 5.1 Directory Structure
The results from the training process are organized as follows:
```
train_model
├── F1_curve.png
├── PR_curve.png
├── P_curve.png
├── R_curve.png
├── args.yaml
├── confusion_matrix.png
├── confusion_matrix_normalized.png
├── labels.jpg
├── labels_correlogram.jpg
├── results.csv
├── results.png
├── train_batch0.jpg
├── train_batch1.jpg
├── train_batch2.jpg
├── val_batch0_labels.jpg
├── val_batch0_pred.jpg
├── val_batch1_labels.jpg
├── val_batch1_pred.jpg
├── val_batch2_labels.jpg
├── val_batch2_pred.jpg
└── weights
    ├── best.onnx
    ├── best.pt
    └── last.pt
```

### 5.2 Precision, Recall, and F1 Scores
The [F1 score curve](./v3/runs/detect/train_model/F1_curve.png) shows the harmonic mean of precision and recall across different confidence thresholds. This curve highlights how the balance between precision and recall changes and helps identify the optimal operating threshold.

The [Precision-Recall curve](./v3/runs/detect/train_model/PR_curve.png) further illustrates the trade-off between precision and recall. As precision increases, recall tends to decrease, and vice versa.

The [P curve](./v3/runs/detect/train_model/P_curve.png) displays precision across various confidence levels, while the [R curve](./v3/runs/detect/train_model/R_curve.png) illustrates how recall changes with confidence.

### 5.3 Confusion Matrix
The [confusion matrix](./v3/runs/detect/train_model/confusion_matrix.png) provides a visualization of true positives, true negatives, false positives, and false negatives. 

The [normalized version](./v3/runs/detect/train_model/confusion_matrix_normalized.png) gives a clearer picture of classification performance, particularly when dealing with imbalanced data.

It shows the proportion of correct predictions relative to the total number of instances for each class.

### 5.4 Label Analysis
The [labels.jpg](./v3/runs/detect/train_model/labels.jpg) and [labels_correlogram.jpg](./v3/runs/detect/train_model/labels_correlogram.jpg) files provide insights into the distribution and correlation of the labels within the dataset. These visualizations help to understand if there are any patterns or correlations in the training data that could affect the model's performance.

### 5.5 Training Process Visualization
During training, batch-level performance is visualized in the [train_batch0.jpg](./v3/runs/detect/train_model/train_batch0.jpg), [train_batch1.jpg](./v3/runs/detect/train_model/train_batch1.jpg), and [train_batch2.jpg](./v3/runs/detect/train_model/train_batch2.jpg) files. These show a snapshot of the model's predictions on the training data. Similar visualizations for the validation data are provided in the ***val_batchX_labels.jpg*** and ***val_batchX_pred.jpg*** files, showing both the ground truth and predicted results for the validation batches.

### 5.6 Model Weights
The final model weights are saved in the [weights](./v3/runs/detect/train_model/weights/) directory, with different versions representing various stages of training. The file [best.pt](./v3/runs/detect/train_model/weights/best.pt) contains the best-performing model in PyTorch format, [last.pt](./v3/runs/detect/train_model/weights/last.pt) represents the final model from the last epoch, and [best.onnx](./v3/runs/detect/train_model/weights/best.onnx) provides an ONNX version of the best model for deployment or integration into other systems.

### 5.7 Overall Performance and Metrics
The [results.csv](./v3/runs/detect/train_model/results.csv) file provides detailed numerical metrics for each training epoch, including precision, recall, and mAP (mean Average Precision). 

The [results.png](./v3/runs/detect/train_model/results.png) image provides a visual summary of the overall training process, showcasing how the model's performance evolved over time, particularly with respect to key metrics like loss, precision, and recall.


## 6. Running the Training Process
### 6.1 Environment Setup and Dependencies
`conda create --name <env> --file requirements.txt`

### 6.2 Code Overview
The complete codebase for this training is available in [Code.py](./Code.py).

#### 6.2.1 Key Functions
`train()`
> Starts the training process for the model.
    
`detect_video()`
> Runs object detection on the [sample video](./sample.mp4).

`detect_all_frame()`
> Detects objects in every frame of the video.

`devide_data()`
> Splits the dataset to [train](./dataset/train/) and [validation](./dataset/valid/) subsets.

---
<div align="right">
-->

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 17-OCT-2024 14:12 UTC +08:00*
</div>