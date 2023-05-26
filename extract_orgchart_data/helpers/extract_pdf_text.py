from pdf2docx import Converter
from docx2python import docx2python
import os


def extract_pdf_text(pdf_file):
    cv = Converter(pdf_file)
    docx_directory = "./docx"
    docx_file = pdf_file.replace(".pdf",".docx")
    docx_file = docx_directory+"/"+os.path.basename(docx_file)
    
    # Check if the directory exists
    if not os.path.exists(docx_directory):
        # If it doesn't exist, create it
        os.makedirs(docx_directory)

    if not os.path.exists(docx_file):
        cv.convert(docx_file)
    
    docx_content = docx2python(docx_file)

    return docx_content
    