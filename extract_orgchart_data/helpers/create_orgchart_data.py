from  helpers.extract_data_from_csv import extract_data_from_csv,save_json 
import json

def create_orgchart_data(csv_path):
    json_str = extract_data_from_csv(csv_path)
    json_path = './json/'
    return save_json(json_str,json_path)


def get_existing_data(json_path):
    with open(json_path, 'r') as json_file:
        data = json.load(json_file)
    return data