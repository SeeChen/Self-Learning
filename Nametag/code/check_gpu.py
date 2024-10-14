import os
import multiprocessing

# Get the total number of CPU cores
num_cores = multiprocessing.cpu_count()
print(f'Total number of CPU cores: {num_cores}')