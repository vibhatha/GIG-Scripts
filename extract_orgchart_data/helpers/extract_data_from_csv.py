import csv
import json
import os
import re

def extract_data_from_csv(csv_path):
    """
    extracts data from the latest full gazette and store in a json with department numbers
    """
    data_dict = {}

    with open(csv_path, 'r', encoding='utf-8') as csv_file:
        reader = csv.reader(csv_file)
        for row in reader:
            key = row[0]
            value = row[1]

            if key.lower() == 'terminate':
                continue

            if key not in data_dict:
                data_dict[key] = []
            
            if data_dict[key] == []:
                last_department_no = 0            
            else:
                last_department_no = get_last_department_no(data_dict, key)

            if last_department_no is None:
                continue

            department_with_no = str(last_department_no + 1) + '. ' + value
            data_dict[key] = data_dict[key] + [department_with_no]
            
    return dict_to_json(data_dict)

def get_last_department_no(data_dict, key):
    department_arr = data_dict[key]
    last_department = department_arr[-1]
    match = re.search(r'\d+', last_department)
    
    if match:
        # Extract the matched number from the string
        number = int(match.group())
        return number
    
    # Return None if no number is found
    return None


def dict_to_json(dictionary):
    json_data = []
    for ministry in dictionary:
        json_obj = {
            "Ministry": ministry,
            "Organizations": dictionary[ministry]
        }
        json_data.append(json_obj)

    json_object = json.dumps(json_data)

    return json_object

def save_json(json_str,path):
    json_data = json.loads(json_str)
    # Check if the directory exists
    if not os.path.exists(path):
        # If it doesn't exist, create it
        os.makedirs(path)

    json_path = path+"orgchart_data.json"

    with open(json_path, 'w') as json_file:
        json.dump(json_data, json_file)

    return json_path

