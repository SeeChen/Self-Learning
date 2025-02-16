from flask import Flask, request, jsonify
import requests

app = Flask(__name__)

@app.route('/forward', methods=['POST'])
def forward_request():
    """接收 Playwright 发送的拦截请求，并转发到真实服务器"""
    data = request.json
    real_url = data.get("url")
    headers = data.get("headers", {})

    if not real_url:
        return jsonify({"error": "Invalid request"}), 400

    try:
        response = requests.get(real_url, headers=headers)
        print(response.headers.get("Content-Type", ""))
        return jsonify(response.json())
    except requests.exceptions.RequestException as e:
        print("Error")
        return jsonify({"error": f"Failed to reach {real_url}"}), 500

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8899, debug=True)
    