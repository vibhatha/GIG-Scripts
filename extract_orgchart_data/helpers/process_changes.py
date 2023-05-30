import json

def get_data_to_be_updated(changes_done,existing_data):
    changed_ministries = get_changed_ministries(changes_done)
    filterd_data = filter_data(changed_ministries,existing_data)
    data_to_be_updated = merge_data(filterd_data,changes_done)
    return data_to_be_updated


def get_changed_ministries(changes_done):
    changed_ministries = []
    for item in changes_done:
        changed_ministries.append(item['ministry'])
    return changed_ministries

def filter_data(changed_ministries,existing_data):
    filtered_data = []
    
    key_list_lower = [key.lower() for key in changed_ministries]
    
    for item in existing_data:
        if item['Ministry'].lower() in key_list_lower:
            filtered_data.append(item)
    
    return filtered_data

def merge_data(filterd_data,changes_done):
    data_to_be_updated = []
    
    for item in changes_done:
        for data in filterd_data:
            if item['ministry'].lower() == data['Ministry'].lower():
                json_obj = {
                    'Ministry': data['Ministry'],
                    'Organizations': data['Organizations'],
                    'changes_to_be_done': item['Column II']
                }
                data_to_be_updated.append(json_obj)
    
    return data_to_be_updated
