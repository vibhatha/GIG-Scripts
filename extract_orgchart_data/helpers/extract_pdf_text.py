from pdf2docx import Converter
from docx2python import docx2python
import os


def extract_pdf_text(pdf_file):
    cv = Converter(pdf_file)
    docx_file = pdf_file.replace(".pdf",".docx")
    cv.convert(docx_file)
    
    docx_content = docx2python(docx_file)
    # docx_content.close()
    # os.remove(docx_file)    
    # print(docx_content.body)
    return docx_content.body
    