import csv
import re
import os

def write_to_csv(extracted_data,pdf_file,path):
    x = re.findall('[0-9]+', pdf_file)

    # Check if the directory exists
    if not os.path.exists(path):
        # If it doesn't exist, create it
        os.makedirs(path)

    csv_name = path+"gazette-"+x[0]+"-"+x[1]+"-"+x[2]+".csv"
    
    with open(csv_name, 'w', encoding='UTF8', newline='') as f:
        writer = csv.writer(f)

        for ministry in extracted_data:
            for department in extracted_data[ministry]:
                row = [ministry, department]
                writer.writerow(row)
       