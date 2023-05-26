from helpers.extract_data_amendments import get_data_amendments
from helpers.create_orgchart_data import create_orgchart_data,get_existing_data
from helpers.get_latest_full_gazette import get_latest_full_gazette
from helpers.llm_calls import update_existing_data, get_changes,config_openai
csv_directory = "./extracted"
pdf = 'functions_2022-07-22_E.pdf'
config_openai()

latest_full_gazette = get_latest_full_gazette(csv_directory)

json_path = create_orgchart_data(latest_full_gazette)
existing_data = get_existing_data(json_path)

# gets the text from the amendment pdf
amendment = get_data_amendments("./pdfs/functions_2022-09-16_E.pdf")

# find the changes done in the amendment, return a list of changes with corresponding ministry
changes_done = get_changes(amendment)

# todo: get the changes ministers from the existing data
# todo: create a promt with the changes and the ministers and get the ministers updated
# todo: save to the orgchart json
# todo: create a csv file with the updated data
# todo: fine-tune a openai model with the daata to increase accuracy

print(changes_done)
