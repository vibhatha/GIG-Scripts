import os

def get_downloaded_pdfs(directory):
    pdf_names = []
    for file_name in os.listdir(directory):
        if file_name.endswith('.pdf'):
            pdf_names.append(file_name)
    return pdf_names