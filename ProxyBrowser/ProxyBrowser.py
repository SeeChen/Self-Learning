import json
import requests
from playwright.sync_api import sync_playwright

GO_SERVER_URL = "http://localhost:8899/forward"

def intercept_request(route, request):
    print(f"Intercepted request: {request.url}")
    print()

    try:
        response = requests.post(GO_SERVER_URL, json={"url": request.url, "headers": request.headers})
        server_response = response.json()
    except requests.exceptions.RequestException as e:
        print(f"Error forwarding request: {e}")
        server_response = {"error": "Failed to forward request"}

    if "error" in server_response:
        server_response = {"message": "Intercepted by Playwright"}
    
    route.fulfill(
        status=200,
        headers={"Content-Type": "application/json"},
        body=json.dumps(server_response)
    )

def get_storage_data(page):
    local_storage = page.evaluate("() => JSON.stringify(localStorage)")
    session_storage = page.evaluate("() => JSON.stringify(sessionStorage)")
    return {"localStorage": local_storage, "sessionStorage": session_storage}

def get_browser_data(context):
    cookies = context.cookies()
    return {"cookies": cookies}

with sync_playwright() as p:
    browser = p.chromium.launch(headless=False)
    context = browser.new_context()
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

    page.goto("https://www.google.com")

    cookies_data = get_browser_data(context)
    print("Cookies:", cookies_data["cookies"])

    storage_data = get_storage_data(page)
    print("LocalStorage:", storage_data["localStorage"])
    print("SessionStorage:", storage_data["sessionStorage"])

    while len(open_pages) > 0:
        open_pages[0].wait_for_event("close")
        open_pages.pop(0)

    print("All tabs closed. Exiting...")
    browser.close()
