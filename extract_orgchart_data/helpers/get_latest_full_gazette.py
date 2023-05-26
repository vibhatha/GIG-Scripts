import os
import csv

def get_latest_full_gazette(directory):
    """
    finds the latest full gazette in the given directory
    """
    csv_file = ''
    
    # Traverse through all files and directories in the given directory
    for root, dirs, files in os.walk(directory):
        
        for file in reversed(list(files)):
            # Check if the file is a CSV file
            if file.endswith('.csv'):
                file_path = os.path.join(root, file)
                with open(file_path, 'r', encoding='utf-8') as csv_file:
                    reader = csv.reader(csv_file)
                    
                    for row in reader:
                        # Check if "Terminate, All" is in the row
                        flag = False
                        for value in row:
                            if "Terminate" in value: flag = True
                            if flag and "All" in value:
                                csv_file = file_path
                                break
    
    return csv_file
