from helpers.extract_data_amendments import get_data_amendments
from helpers.create_orgchart_data import create_orgchart_data,get_existing_data
from helpers.get_latest_full_gazette import get_latest_full_gazette
from helpers.llm_calls import get_changes,config_openai
from helpers.process_changes import get_data_to_be_updated
from helpers.update_existing_data import update_existing_data
from helpers.write_to_csv import write_to_csv
from helpers.get_amendments import get_amendments
from helpers.process_data import process_data
from dotenv.main import load_dotenv
import os

load_dotenv()
OPENAI_API_KEY = os.environ['OPENAI_API_KEY']

csv_directory = "./extracted/"
pdf_directory = "./pdfs/"
config_openai(OPENAI_API_KEY)

latest_full_gazette = get_latest_full_gazette(csv_directory)
amendment_pdfs_lst = get_amendments(latest_full_gazette,pdf_directory)
print(amendment_pdfs_lst)


for pdf in amendment_pdfs_lst:
    pdf_path = pdf_directory+pdf

    json_path = create_orgchart_data(latest_full_gazette)
    existing_data = get_existing_data(json_path)

    # gets the text from the amendment pdf
    amendment = get_data_amendments(pdf_path)
    # print(amendment)

    # find the changes done in the amendment, return a list of changes with corresponding ministry
    changes_done = get_changes(amendment)
    print(changes_done)

    data_to_be_updated = get_data_to_be_updated(changes_done,existing_data)

    updated_data = update_existing_data(data_to_be_updated, json_path)
    print(updated_data)

    processed_data_dict = process_data(updated_data)
    write_to_csv(processed_data_dict,pdf,csv_directory)
    processed_data_dict.clear()


# todo: fine-tune a openai model with the daata to increase accuracy and reduce the no. of tokens
