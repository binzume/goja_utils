const baseUrl = globalThis.testServerUrl || "http://binzume.net";
console.log("Test url:", baseUrl);

async function test_get() {
    let res = await fetch(baseUrl);
    if (res.status != 200) {
        throw "res.status != 200";
    }
    let text = await res.text();
    if (!text) {
        throw "!text";
    }
}

async function test_json() {
    let res = await fetch(baseUrl);
    if (res.status != 200) {
        throw "res.status != 200";
    }
    let json = await res.json();
    if (json?.status != "ok") {
        throw "json?.status != ok";
    }
}

async function test_404() {
    let res = await fetch(baseUrl + "/_test_not_found");
    if (res.status != 404) {
        throw "res.status != 404";
    }
    let text = await res.text();
    if (!text) {
        throw "!text";
    }
}

async function test() {
    await test_get();
    await test_json();
    await test_404();

    return "pass"
}

test();
