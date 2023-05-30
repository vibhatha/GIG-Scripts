
def process_data(json_data):
    data_dict = {}
    for item in json_data:
        item['Organizations'] = remove_departmernt_no(item['Organizations'])
        data_dict[item['Ministry']] = item['Organizations']
    return data_dict

def remove_departmernt_no(departments):
    new_departments = []
    for department in departments:
        new_departments.append(department.split('. ')[1])
    return new_departments