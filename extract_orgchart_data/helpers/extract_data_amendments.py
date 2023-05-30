# from PyPDF2 import PdfReader
import pdfplumber

# def get_data_amendments(pdf_file):
#     reader = PdfReader(pdf_file)
#     parts = []

#     def visitor_body(text, cm, tm, fontDict, fontSize):
#         y = tm[5]
#         if y > 50 and y < 720:
#             parts.append(text)

#     for i in range(len(reader.pages)):
#         page = reader.pages[i]
#         page.extract_text(visitor_text=visitor_body)

#     text_body = "".join(parts)

#     return text_body

def get_data_amendments(pdf_path):
    all_text = ''
    with pdfplumber.open(pdf_path) as pdf:
        for pdf_page in pdf.pages:
            single_page_text = pdf_page.extract_text()
            all_text = all_text + '\n' + single_page_text
    return all_text