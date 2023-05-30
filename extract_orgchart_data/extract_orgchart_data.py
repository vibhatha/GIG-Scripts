from helpers.extract_ministers_departments import extract_ministers_departments
# from helpers.extract_data_from_amendments import extract_data_from_amendments
from helpers.write_to_csv import write_to_csv
from helpers.crawl_pdfs import download_all_pdfs
from helpers.get_pdf_names import get_pdf_names

website_url = "http://www.cabinetoffice.gov.lk/cab/index.php?option=com_content&view=article&id=54&Itemid=93&lang=en"
pdf_directory = "./pdfs"
csv_directory = "extracted/"


download_all_pdfs(website_url, pdf_directory)
pdf_names = get_pdf_names(pdf_directory)

for pdf in pdf_names:
    # extract ministers and corresponding departments
    pdf_location = pdf_directory+"/"+pdf
    extracted_data = extract_ministers_departments(pdf_location)
    # writing to csv
    write_to_csv(extracted_data,pdf,csv_directory)
    extracted_data.clear()

