import openai
import marvin as ai_fn

def update_existing_data(existing_data, amendment):
    openai.api_key = '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'
    # print(existing_data)

    full_prompt = f'There is a table which consist of "Ministries" and corresponding "Departments, Statutory Institutions and Public Corporations". "Departments, Statutory Institutions and Public Corporations" is included in Column II of the table. That data in table is converted to a json file(every json object has "Ministry" key and "Organizations" key. "Ministry" key consist of the ministry and "Organizations" consist of the relavent "Departments, Statutory Institutions and Public Corporations"). Find the updates done to the table by going through this text : "{amendment}" and update this json file: "{existing_data}". Return result in JSON format without any explanation.'
    full_prompt = f'find entities in this text: {amendment}'

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

# @ai_fn
# def update_existing_data(existing_data: str, amendment: str) -> str:
#     """There is a table which consist of "Ministries" and corresponding "Departments, Statutory Institutions and Public Corporations". "Departments, Statutory Institutions and Public Corporations" is included in Column II of the table. That data in table is converted to a json file(every json object has "Ministry" key and "Organizations" key. "Ministry" key consist of the ministry and "Organizations" consist of the relavent "Departments, Statutory Institutions and Public Corporations"). Find the updates done to the table by going through this text : "{amendment}" and update this json file: "{existing_data}". Return result in JSON format without any explanation."""