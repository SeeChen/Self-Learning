# import gzip
# import base64
# import json
# import requests
# from playwright.sync_api import sync_playwright

# GO_SERVER_URL = "http://localhost:8899/forward"

# def decode_post_data(post_data, headers):
#     if not post_data:
#         return ""

#     if isinstance(post_data, str):
#         return post_data

#     try:
#         return post_data.decode("utf-8")

#     except UnicodeDecodeError:
#         if "Content-Encoding" in headers and "gzip" in headers["Content-Encoding"].lower():
#             try:
#                 return gzip.decompress(post_data).decode("utf-8")
#             except Exception:
#                 pass

#         try:
#             return post_data.decode("ISO-8859-1")
#         except UnicodeDecodeError:
#             pass

#         return f"BASE64: {base64.b64encode(post_data).decode()}"

# def intercept_request(route, request):
#     try:
#         post_data = decode_post_data(request.post_data, request.headers)

#         request_data = {
#             "url": request.url,
#             "method": request.method,
#             "headers": request.headers,
#             "body": post_data,
#         }

#         print(f"[{request.method}] {request.url}")

#         route.continue_()

#     except Exception as e:
#         print(f"Error: {e}")
#     # try:
#     #     response = requests.post(GO_SERVER_URL, json=request_data, timeout=5)

#     #     server_status = response.status_code
#     #     server_headers = response.headers
#     #     server_body = response.content

#     #     response_headers = {k: v for k, v in server_headers.items() if k.lower() not in ["content-encoding", "transfer-encoding"]}

#     #     route.fulfill(
#     #         status=server_status,
#     #         headers=response_headers,
#     #         body=server_body
#     #     )

#     # except requests.exceptions.RequestException as e:
#     #     print(f"Error forwarding request: {e}")
#     #     route.fulfill(
#     #         status=500,
#     #         headers={"Content-Type": "application/json"},
#     #         body=json.dumps({"error": "Failed to forward request"})
#     #     )

# def get_storage_data(page):
#     local_storage = page.evaluate("() => JSON.stringify(localStorage)")
#     session_storage = page.evaluate("() => JSON.stringify(sessionStorage)")
#     return {"localStorage": local_storage, "sessionStorage": session_storage}

# def get_browser_data(context):
#     cookies = context.cookies()
#     return {"cookies": cookies}

# with sync_playwright() as p:
#     browser = p.chromium.launch(headless=False)
#     context = browser.new_context()
#     page = context.new_page()

#     open_pages = [page]

#     def on_new_page(new_page):
#         print(f"New tab opened: {new_page.url}")
#         open_pages.append(new_page)

#         new_page.route("**/*", intercept_request)

#         new_page.wait_for_event("close")
#         open_pages.remove(new_page)
#         print(f"Tab closed. Remaining tabs: {len(open_pages)}")

#     context.on("page", on_new_page)

#     page.route("**/*", intercept_request)

#     page.goto("https://www.google.com")

#     cookies_data = get_browser_data(context)
#     print("Cookies:", cookies_data["cookies"])

#     storage_data = get_storage_data(page)
#     print("LocalStorage:", storage_data["localStorage"])
#     print("SessionStorage:", storage_data["sessionStorage"])

#     while len(open_pages) > 0:
#         open_pages[0].wait_for_event("close")
#         open_pages.pop(0)

#     print("All tabs closed. Exiting...")
#     browser.close()

import requests
from playwright.sync_api import sync_playwright

def intercept_request(route, request):
    print(f"[{request.method}] {request.url}")
    route.continue_()

with sync_playwright() as p:

    data = {
        "username": "",
        "password": ""
    }
    res = requests.post("", json=data).json()
    browser = p.chromium.launch(
        headless=False,
        proxy={"server": f'{res["ip"]}:{res["port"]}'}
    )
    context = browser.new_context(
        user_agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
    )
    page = context.new_page()

    open_pages = [page]

    def on_new_page(new_page):
        print(f"New tab opened: {new_page.url}")
        open_pages.append(new_page)

        new_page.route("**/*", intercept_request)

        new_page.wait_for_event("close")
        open_pages.remove(new_page)
        print(f"Tab closed. Remaining tabs: {len(open_pages)}")

    context.on("page", on_new_page)

    page.route("**/*", intercept_request)

    try:
        page.goto("https://www.bing.com", timeout=0)
        page.wait_for_selector("body")
        print("Page loaded successfully!")
    except Exception as e:
        print(f"Error: {e}")

    while len(open_pages) > 0:
        open_pages[0].wait_for_event("close", timeout=0)
        open_pages.pop(0)

    print("All tabs closed. Exiting...")
    browser.close()
