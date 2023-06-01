import os

def get_pdf_names(directory):
    pdf_names = []

    directory = os.path.join(os.getcwd(), directory)
    # Create the directory if it doesn't exist
    if not os.path.exists(directory):
        os.makedirs(directory)

    for file_name in os.listdir(directory):
        if file_name.endswith('.pdf'):
            pdf_names.append(file_name)
    return pdf_names