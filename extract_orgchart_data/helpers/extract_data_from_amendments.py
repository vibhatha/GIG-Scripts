import openai
from PyPDF2 import PdfReader
import re


def extract_data_from_amendments(pdf_file):

    text_info = ""
    reader = PdfReader("functions_2022-09-16_E.pdf")
    number_of_pages = len(reader.pages)
    for i in range(number_of_pages):
        page = reader.pages[i]
        text = page.extract_text()
        text_info = text_info+text
        # print(text)
    
    cleaned_text = clean_text_information(text_info)
    
    extract_prompt = "Extract data from this gazette."

    response = openai.Completion.create(
    model="text-davinci-003",
    prompt= extract_prompt + cleaned_text,
    temperature=0,
    max_tokens=64,
    top_p=1.0,
    frequency_penalty=0.0,
    presence_penalty=0.0
    )

def clean_text_information(text_info):

    split_list = text_info.split("SChEDuLE", 1)
    if len(split_list) > 0:
        # print(split_list[1])
        remove_text = "PRINTED AT THE DEPARTMENT OF GOVERNMENT PRINTING,  SRI LANKA."
        compiled = re.compile(re.escape(remove_text), re.IGNORECASE)
        cleaned_text = compiled.sub('', split_list[1])
        # print(merged_str)
        return cleaned_text
     
    return text_info