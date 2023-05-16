from helpers.extract_pdf_text import extract_pdf_text
import re

pdf_text = []
column_heading = "Column I"
extracted_data = dict()

def extract_ministers_departments(pdf_file):
    pdf_text = extract_pdf_text(pdf_file).body

    # iterate through the pdf_text lists
    for i, text1 in enumerate(pdf_text):
        for j, table_data in enumerate(text1):
            # getting headings list in pdf_text
            table_heading = table_data[0]

            # extract ministers if  table_heading list contains "Column I"
            if search_in_sublists(table_heading,column_heading):
                extract_ministers(pdf_text, i)

            extract_departments(table_data)

    print("No: of Ministers : " , len(extracted_data))
    for key, value in extracted_data.items():
        print(key, ' : ', value,'No. of Departments : ',len(value),'\n\n')

    return extracted_data



def is_department_cell(table_data):
    # returning false for unwanted cells in column II
    if "Column II" in table_data:
        return False
    if "Departments,  Statutory \nInstitutions & Public Corporations" in table_data:
        return False
    if "Departments, Statutory Institutions and Public Corporations" in table_data:
        return False
    if len(table_data) == 0:
        return False
    return True


def extract_ministers(pdf_text, i):
    # getting list containing ministers and merging
    minister_data = pdf_text[i-1][0][0][-1]
    minister_data = ''.join(minister_data)

    # check whether the minister_data is valid
    minister_len = len(minister_data)
    minimum_len = 10
    if minister_len < minimum_len: return

    # check whether the minister_data contains a number
    temp = re.findall(r'\d+', minister_data)
    no_lst_in_minister_str = list(map(int, temp))

    # search for the minister number in minister_data
    if len(no_lst_in_minister_str) > 0:
        minister_name = clean_minister_data(minister_data)

        if minister_name not in extracted_data:
            extracted_data[minister_name] = []
    return

    
def extract_departments(table_data):
    # find the list which containing 3 columns in the table
    if len(table_data) == 3:
        # getting the 2nd column data to extract "Column II"
        deparment_string = ''.join(table_data[1])

        # checking whether it is the department cell
        if is_department_cell(deparment_string):
            # clean department names and add the list to extracted_data
            department_lst = clean_department_data(deparment_string)

            try:
                minister_name = list(extracted_data.keys())[-1]
                extracted_data[minister_name] = extracted_data[minister_name] + department_lst
            except:
                print("No Ministry Found")



def clean_department_data(department_data):
    # Remove newlines and tabs
    data = department_data.replace('\n', '').replace('\t', '').replace('�', ' ')

    # Remove any non-printable characters
    data = ''.join(c for c in data if c.isprintable())

    # split the string by numbers and create a list
    lst = re.split('[0-9]+', data)
    for i,x in enumerate(lst):
        lst[i] = x.replace('. ', '')

    # remove empty strings and whitespace from list
    lst = [x.strip() for x in lst if x.strip()]

    return lst



def clean_minister_data(merged_str):
    # Remove "SCHEDULE" and "(Contd.)"
    remove_text = "SCHEDULE"
    compiled = re.compile(re.escape(remove_text), re.IGNORECASE)
    merged_str = compiled.sub('', merged_str)

    remove_text = "(Contd.)"
    compiled = re.compile(re.escape(remove_text), re.IGNORECASE)
    merged_str = compiled.sub('', merged_str)
    # remove unnessasary characters
    merged_str = merged_str.replace('.','').replace('•','').replace('/n','').replace('/t','')

    # Remove all digits
    merged_str = ''.join(c for c in merged_str if not c.isdigit())

    # remove trailing spaces
    return merged_str.strip()



def search_in_sublists(sublist, search_term):
    # searches an element in sublists
    for item in sublist:
        if search_term == item.strip():
            return True
    return False