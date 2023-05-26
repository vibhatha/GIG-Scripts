from PyPDF2 import PdfReader

def get_data_amendments(pdf_file):
    reader = PdfReader(pdf_file)
    parts = []

    def visitor_body(text, cm, tm, fontDict, fontSize):
        y = tm[5]
        if y > 50 and y < 720:
            parts.append(text)

    for i in range(len(reader.pages)):
        page = reader.pages[i]
        page.extract_text(visitor_text=visitor_body)

    text_body = "".join(parts)

    return text_body