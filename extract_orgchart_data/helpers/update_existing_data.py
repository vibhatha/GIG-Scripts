from helpers.llm_calls import update_entry
import json

def update_existing_data(data_to_be_updated, json_path):
    updated_data = []
    for item in data_to_be_updated:
        updated_obj = update_entry(json.dumps(item))
        arranged_data = arrange_departments(json.loads(updated_obj))
        updated_data.append(arranged_data)
        update_json_file(json_path, arranged_data)

    return updated_data

def arrange_departments(json_data):
    i = 1
    for departmernt in json_data['Organizations']:
        departmernt_no = int(departmernt.split('.')[0])
        if i != departmernt_no:
            # print(i,departmernt_no)
            json_data['Organizations'][i-1] = f"{i}. {departmernt.split('. ')[1]}"
        i += 1

    return json_data

def update_json_file(filepath, new_data):
    # Load the existing JSON data from the file
    with open(filepath, 'r') as file:
        json_data = json.load(file)
    
    # Update the JSON data with the new data
    for item in json_data:
        if item['Ministry'].lower() == new_data['Ministry'].lower():
            item['Organizations'] = new_data['Organizations']

    # Save the updated JSON data back to the file
    with open(filepath, 'w') as file:
        json.dump(json_data, file, indent=4)
    
    # Return the updated data
    return json_data