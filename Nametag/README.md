<div align=center>

# AI Evaluation Test

[Summary](#1-summary)</br>
[Environment Setup](#2-environment-setup)</br>
[Dataset Setup](#3-dataset-setup)</br>
[Model Training](#4-model-training)</br>
[Result](#5-result)</br>
[Conclusion](#6-conclusion)

</div>

## 1. Summary
This task involved identifying and tracking a staff member wearing a nametag using [video](sample.mp4) footage from a 3D sensor. The goal was to determine which frames contained the staff member and locate their coordinates when present. I selected the YOLO (You Only Look Once) model for this task because of its fast object detection capabilities.


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

## 4. Model Training
Regarding the training approach, I am using a self-iterative training method. This approach involves multiple iterations of model training and refinement based on progressively extracted data from previous outcomes.

Self-iteration refers to a process where the model continuously uses its own outputs as new training data in subsequent training rounds. After each round, the model’s predictions are used to adjust the training data, either by adding newly predicted labels or refining the data. The model then retrains with the updated data, iterating this process. Over time, the model improves by learning from its own results, essentially "bootstrapping" its knowledge. This method is a form of semi-supervised learning, starting with a small labeled dataset and gradually expanding the training set by using the model’s predictions on unlabeled data to enhance its performance.

### 4.1 Initial Training
First of all, I randomly selected 370 pictures from a total of 1370 frames, and then annotated the nametags, person, and staff. The initial training aimed to establish a baseline for object detection.

### 4.2 First Detection and Refinement
After the initial model training, the model was used to detect objects in the video frames. Based on the detection results, I extracted a set of images that were either correctly or incorrectly detected and re-annotated them to correct the labels. After that, similar to the first step, I randomly sampled 100 images for training.

### 4.3 Subsequent Iterative Training


## 5. Result


## 6. Conclusion



---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 16-OCT-2024 23:14 UTC +08:00*
</div>