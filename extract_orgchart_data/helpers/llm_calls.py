import openai
import json

def config_openai():
    openai.api_key = OPENAI_API_KEY
    return

def update_existing_data(existing_data, amendment):
    # print(existing_data)

    full_prompt = f'There is a table which consist of "Ministries" and corresponding "Departments, Statutory Institutions and Public Corporations". "Departments, Statutory Institutions and Public Corporations" is included in Column II of the table. That data in table is converted to a json file(every json object has "Ministry" key and "Organizations" key. "Ministry" key consist of the ministry and "Organizations" consist of the relavent "Departments, Statutory Institutions and Public Corporations"). Find the updates done to the table by going through this text : "{amendment}" and update this json file: "{existing_data}". Return result in JSON format without any explanation.'
    # full_prompt = f'find entities in this text: {amendment}'

    response = openai.Completion.create(
    model="text-davinci-003",
    prompt=full_prompt,
    temperature=0,
    max_tokens=64,
    top_p=1.0,
    frequency_penalty=0.0,
    presence_penalty=0.0
    )

    extracted_data = response.choices[0].text.strip()
    print(response)

    return extracted_data

def get_changes(amendment):
    full_prompt = f"""You are a professional entitiy exrtactor. The task is to identify the "Ministry" entities and the extract the instructions which is relavent to ONLY "Column II" under that ministry in the given text. Return the names of the ministries and changes done in json format.
     Below are some examples.

     Input: "(1) With reference to the Heading “No.01. Minister of Defence” of the said notification, as follows :-
        (a) In Column I thereof, by omitting item 16 ;
        (b) In Column I thereof, re-numbering of items 17 to 28 respectively, as numbers 16 to 27 ;
        (c) In Column II thereof, by omitting item 19 ;
        (d) In Column II thereof, re-numbering of items 20 to 28 respectively, as numbers 19 to 27 ;
        (e) In Column III thereof, by omitting the following items ;·
        • Convention against Illicit Traffic in Narcotic Drugs and Psychotropic Substances Act, No. 1 of 
        2008
        • Drug Dependent Persons (Treatment and Rehabilitation) Act, No. 54 of 2007
        • National Dangerous Drugs Control Board Act, No. 11 of 1984
        
        (2) With reference to the Heading, “No. 03. Minister of Technology” of the said notification, as follows:-
        (a) In Column I thereof, by omitting item 11 ;
        (b) In Column I thereof, re-numbering of items 12 and 13 respectively, as numbers 11 and 12;
        (c) In Column II thereof, by omitting items 1 and 5 ;
        (d) In Column II thereof, re-numbering of items 2 to 4 and item 6 respectively, as numbers 1 to 4;
        (e) In Column III thereof, by omitting the following items;·
        • Registration of Persons Act, No. 32 of 1968·
        • Science and Technology Development Act, No. 11 of 1994"
     Output: [
        {{"ministry": "Minister of Defence", "Column II": ["remove item 19","re-number items 20 to 28 respectively, as numbers 19 to 27 "]}},
        {{"ministry": "Ministry of Technology", "Column II": ["remove items 1 and 5","re-number items 2 to 4 and item 6 respectively, as numbers 1 to 4"]}}
        ]

     Input: "With reference to the Heading, “No. 19. Minister of Industries” of the said notification, as follows:-
        (a) In Column II thereof, by insertion of the following item immediately after item 27; 
        28. Industrial Technology Institute 
        (b) In Column III thereof, by insertion the following item after the item “National Gem and Jewellery Authority 
        Act, No. 50 of 1993”;·
        • Science and Technology Development Act, No. 11 of 1994"
     Output: [
        {{"ministry": "Ministry of Industries", "Column II": ["insert the following item after item 27; 28. Industrial Technology Institute"]}}
        ]


     Input: "(2) With reference to the Heading, No. “06. Minister of Investment Promotion” of the said notification, as follows:-
        (a) In Column I thereof, by insertion of the following items immediately after item 10;
        11. Development of Colombo Port City Special Economic Zone as an International Business and Services Hub 
        with specialized infrastructure and other facilities aimed at national interest and economic advancement.
        (b) In Column I thereof, re-numbering of items 11 and 12 respectively, as numbers 12 and 13;
        (c) In Column II thereof, by insertion of the following items immediately after item 4;
        5. Sri Lanka Export Development Board
        6. Board of Investment of Sri Lanka
        7. Colombo Port City Economic Commission
        (d) In Column III thereof, by insertion the following items after the item “Grant of Citizenship to Stateless Persons 
        (Special Provisions) Act, No. 39 of 1988”;
        •	 Sri Lanka Export Development Act No. 40 of 1979
        •	 Greater Colombo Economic Commission Law No. 4 of 1978 (Board of Investment of Sri Lanka Law)
        •	 Colombo Port City Economic Commission Act, No. 11 of 2021"
     Output: [
        {{"ministry": "Minister of Investment Promotion", "Column II": [insert the following items after item 4; 5. Sri Lanka Export Development Board, 6. Board of Investment of Sri Lanka, 7. Colombo Port City Economic Commission]}}
        ]
        

     Input: {amendment}
     Output: 
     """

    response = openai.Completion.create(
    model="text-davinci-003",
    prompt=full_prompt,
    temperature=0,
    max_tokens=2000,
    top_p=1.0,
    frequency_penalty=0.0,
    presence_penalty=0.0
    )

    changes_json_str = response.choices[0].text.strip()

    return json.loads(changes_json_str)

#      Output: ""ministry": "Minister of Defence", "Column II": ["remove item 19","re-number items 20 to 28 respectively, as numbers 19 to 27 "]"
