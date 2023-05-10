import os
import requests
from bs4 import BeautifulSoup
from urllib.parse import urlparse


def download_pdf(url, save_directory):
    response = requests.get(url)
    file_name = os.path.join(save_directory, url.split("/")[-1])
    with open(file_name, 'wb') as file:
        file.write(response.content)

def get_pdf_links(url):
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')
    pdf_links = []
    language = "E"
    for link in soup.find_all('a'):
        href = link.get('href')
        if href and href.endswith(str(language + '.pdf')):
            pdf_links.append(href)
    return pdf_links

def download_all_pdfs(url, save_directory):
    if not os.path.exists(save_directory):
        os.makedirs(save_directory)

    domain_name = urlparse(url).scheme + "://" + urlparse(url).netloc
    
    pdf_links = get_pdf_links(url)
    print(f"Found {len(pdf_links)} PDFs to download.")

    for link in pdf_links:
        pdf_url = link if link.startswith('http') else domain_name + link
        print(f"Downloading {pdf_url}...")
        download_pdf(pdf_url, save_directory)
        print("Download complete!")


