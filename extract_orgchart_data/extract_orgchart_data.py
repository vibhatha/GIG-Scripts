from helpers.extract_ministers_departments import extract_ministers_departments
from helpers.extract_data_from_amendments import extract_data_from_amendments
from helpers.write_to_csv import write_to_csv
from helpers.crawl_pdfs import download_all_pdfs
from helpers.get_downloaded_pdfs import get_downloaded_pdfs

# from docx2python.iterators import enum_at_depth
# from docx2python import docx2python

website_url = "http://www.cabinetoffice.gov.lk/cab/index.php?option=com_content&view=article&id=54&Itemid=93&lang=en"
pdf_directory = "./pdfs"
csv_directory = "extracted/"

download_all_pdfs(website_url, pdf_directory)
pdf_names = get_downloaded_pdfs(pdf_directory)


for pdf in pdf_names:
    # extract ministers and corresponding departments
    pdf_location = pdf_directory+"/"+pdf
    extracted_data = extract_ministers_departments(pdf_location)
    # writing to csv
    write_to_csv(extracted_data,pdf,csv_directory)
    extracted_data.clear()




# extract_data_from_amendments(pdf_file)

   
   
   
   
   
# # # # pdf name with location
# pdf_file = "functions_2022-09-16_E.pdf"   

    
# docx_content = docx2python("functions_2022-07-22_E.docx")


# def html_map(tables) -> str:
#     """Create an HTML map of document contents.

#     Render this in a browser to visually search for data.

#     :tables: value could come from, e.g.,
#         * docx_to_text_output.document
#         * docx_to_text_output.body
#     """

#     # prepend index tuple to each paragraph
#     for (i, j, k, l), paragraph in enum_at_depth(tables, 4):
#         tables[i][j][k][l] = " ".join([str((i, j, k, l)), paragraph])
        
#     # print(tables)
#     # # wrap each paragraph in <pre> tags
#     # for (i, j, k), cell in enum_at_depth(tables, 3):
#     #     tables[i][j][k] = "".join(["<pre>{x}</pre>".format(x) for x in cell])

#     # # wrap each cell in <td> tags
#     # for (i, j), row in enum_at_depth(tables, 2):
#     #     tables[i][j] = "".join(["<td>{x}</td>".format(x) for x in row])

#     # # wrap each row in <tr> tags
#     # for (i,), table in enum_at_depth(tables, 1):
#     #     tables[i] = "".join("<tr>{x}</tr>".format(x) for x in table)

#     # # wrap each table in <table> tags
#     # tables = "".join(['<table border="1">{x}</table>'.format(x) for x in tables])

#     # return ["<html><body>"] + tables + ["</body></html>"]

# # print(docx_content.text)
# # print(html_map(docx_content.body))
