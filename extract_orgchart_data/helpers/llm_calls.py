import openai
import json

def config_openai(OPENAI_API_KEY):
    openai.api_key = OPENAI_API_KEY
    return

def update_entry (existing_data):

    full_prompt = f"""The task is to change the data in "Organizations" according to the instructions given in ""changes_to_be_done"" in given data. Return the updated json object with the changes done.
     Below are some examples.

     Input: {{"Ministry": "Minister of finance, Economic Stabilization and National Policies", "Organizations": ["1. National Planning Department", "2. Department of Project Management and Monitoring", "3. Department of External resources", "4. State resources Management Corporation", "5. General Treasury", "6. Department of fiscal Policy"], "changes_to_be_done": ["remove items 2, 4 and 6"]}}
     Output: {{"Ministry": "Minister of finance, Economic Stabilization and National Policies", "Organizations": ["1. National Planning Department", "3. Department of External resources", "5. General Treasury"]}}

     Input: {{"Ministry": "Minister of Investment Promotion", "Organizations": ["1. Department of immigration and Emigration", "2. Colombo lotus Tower Management Company (Pvt.) limited", "3. Techno Park Development Company (Pvt.) limited", "4. information Technology Parks"], "changes_to_be_done": ["insert the following items after item 4; 5. Sri Lanka Export Development Board, 6. Board of Investment of Sri Lanka"]}}
     Output: {{"Ministry": "Minister of Investment Promotion", "Organizations": ["1. Department of immigration and Emigration", "2. Colombo lotus Tower Management Company (Pvt.) limited", "3. Techno Park Development Company (Pvt.) limited", "4. information Technology Parks", "5. Sri Lanka Export Development Board", "6. Board of Investment of Sri Lanka"]}}

     Input: {{"Ministry": "Minister of Technology", "Organizations": ["1. Department of registration of Persons", "2. information and Communication Technology agency and allied    institutions", "3. Sri lanka Computer Emergency readiness Team", "4. Sri lanka Telecom and its Subsidiaries and allied institutions", "5. industrial Technology institute", "6. Sri lanka Standards institute"], "changes_to_be_done": ["remove items 1 and 5","re-number items 2 to 4 and item 6 respectively, as numbers 1 to 4"]}}
     Output: {{"Ministry": "Minister of Technology", "Organizations": ["1. information and Communication Technology agency and allied    institutions", "2. Sri lanka Computer Emergency readiness Team", ""3. Sri lanka Telecom and its Subsidiaries and allied institutions", "4. Sri lanka Standards institute"]}}

     Input: {existing_data}
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

    updated_data = response.choices[0].text.strip()
   #  updated_data = """{"Ministry": "Minister of Investment Promotion", "Organizations": ["1. Department of immigration and Emigration", "2. Colombo lotus Tower Management Company (Pvt.) limited", "3. Techno Park Development Company (Pvt.) limited", "4. information Technology Parks", "5. Sri Lanka Export Development Board", "6. Board of Investment of Sri Lanka", "7. Colombo Port City Economic Commission"]}"""
   #  updated_data = """{'Ministry': 'Minister of finance, Economic Stabilization and National Policies', 'Organizations': ['1. National Planning Department', '2. Department of Project Management and Monitoring', '3. Department of External resources', '4. State resources Management Corporation', '5. General Treasury', '6. Department of fiscal Policy', '7. Department of National Budget', '8. Department of Management Services', '9. Department of Public Finance', '10. Department of Treasury Operations', '11. Department of State Accounts', '12. Department of Trade and Investment Policies', '13. Department of Information Technology Management', '14. Department of Legal Affairs', '15. Department of Management Audit', '16. Department of Development Finance', '17. Department of Public Enterprises', '18. Office of Comptroller General', '19. Department of Inland Revenue', '20. Sri Lanka Customs', '21. Department of Excise', '22. National Lotteries Board', '23. Development Lotteries Board', '24. Department of Valuation', '25. Import and Export Control Department', '27. Central Bank of Sri Lanka', '28. All State Banks, Financial Institutions, Insurance Companies and their subsidiaries and related institutions', '29. insurance regulatory Commission of Sri lanka', '30. Sri lanka insurance Corporation and its subsidiaries and affiliated companies', '31. Credit information Bureau', '32. Department of the registrar of Companies', '33. Securities and Exchange Commission of Sri lanka', '34. Sri lanka accounting and auditing Standard Monitoring Board', '35. Public utilities Commission of Sri lanka', '36. Sri lanka Export Credit insurance Corporation', '37. Housing Development finance Corporation', '38. State Mortgage and investment Bank', '39. regional Development Bank', '40. Tax appeals Commission', '41. Department of Census and Statistics', '42. institute of Policy Studies', '43. Sustainable Development Council', '44. Welfare Benefits Board', '45. Public Service Mutual Provident fund', '46. Strike, riot, Civil Commotion and Terrorism fund', '47. National insurance Trust fund', "48. Employees' Trust fund", '49. lady lochore fund', '52. Department of Telecommunications', '53. Wildlife Trust', '54. Sri lanka Media Training institute', '55. Department of internal Trade', '56. Pulse Crops, Grain research and Production authority', '57. janatha fertilizer Enterprises ltd', '58. Protection of Children National Trust fund', '59. institutions coming under the revival (removal) of underperforming Enterprises or underutilized assets act vested to the Secretary to the Treasury']}"""
    return updated_data

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
    
   #  changes_json_str = """[
   #      {"ministry": "Minister of Finance, Economic Stabilization and National Policies", "Column II": ["remove items 26, 50 and 51"]},
   #      {"ministry": "Minister of Investment Promotion", "Column II": ["insert the following items after item 4; 5. Sri Lanka Export Development Board, 6. Board of Investment of Sri Lanka, 7. Colombo Port City Economic Commission"]}
   #      ]"""
        
    return json.loads(changes_json_str)

#      Output: ""ministry": "Minister of Defence", "Column II": ["remove item 19","re-number items 20 to 28 respectively, as numbers 19 to 27 "]"
