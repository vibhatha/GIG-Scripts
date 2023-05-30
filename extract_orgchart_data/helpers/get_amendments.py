import re
from helpers.get_pdf_names import get_pdf_names

def get_amendments(latest_full_gazette, pdf_directory):
    print("Getting amendments...")
    x = re.findall('[0-9]+', latest_full_gazette)
    
    pdf_name = "functions_"+x[0]+"-"+x[1]+"-"+x[2]+"_E.pdf"

    pdf_names = get_pdf_names(pdf_directory)
    
    return pdf_names[pdf_names.index(pdf_name)+1:]

