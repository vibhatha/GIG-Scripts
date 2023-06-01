from pdf2docx import Converter
from docx2python import docx2python
import os

DOCX_DIRECTORY = "docx"

def extract_pdf_text(pdf_file):
    print("Converting PDF to text...","PDF file: ", os.path.basename(pdf_file))
    cv = Converter(pdf_file)
    docx_file = pdf_file.replace(".pdf",".docx")
    docx_file = os.path.join(os.getcwd(), DOCX_DIRECTORY, os.path.basename(docx_file))

    # Check if the directory exists
    if not os.path.exists(DOCX_DIRECTORY):
        # If it doesn't exist, create it
        os.makedirs(DOCX_DIRECTORY)

    if not os.path.exists(docx_file):
        cv.convert(docx_file)
    
    docx_content = docx2python(docx_file)

    return docx_content
    