import requests
from uuid import uuid4


url = "http://localhost:4747/person"


for x in range(100):
    data = {
    "name" : str(uuid4()),
    "country": "us"
}
    response = requests.post(url, json=data)
    print(response.status_code)
