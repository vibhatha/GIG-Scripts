from helpers.extract_ministers_departments import extract_ministers_departments
from helpers.write_to_csv import write_to_csv

# pdf name with location
pdf_file = "functions_2022-07-22_E.pdf"

# extract ministers and corresponding departments
extracted_data = extract_ministers_departments(pdf_file)

# writing to csv
write_to_csv(extracted_data,pdf_file)
