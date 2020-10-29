import requests

url = 'http://localhost:8080/send'
headers = {
  'Content-Type': 'application/json'
}
data = "{\n  \"page\": \"page1\",\n  \"type\": \"page_update\",\n  \"data\": {\n    \"object_type\": \"Person\",\n    \"id\": 24772,\n    \"value\": null\n  }\n}"


response = requests.request(
  'POST',
  url,
  data=data,
  headers=headers,
)

print(response)