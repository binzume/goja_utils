const baseUrl = globalThis.testServerUrl || "http://binzume.net";
console.log("Test url:", baseUrl);

async function test_get_text() {
    let res = await fetch(baseUrl);
    if (res.status != 200) {
        throw "res.status != 200";
    }
    let text = await res.text();
    if (!text) {
        throw "!text";
    }
}

async function test_get_bytes() {
    let res = await fetch(baseUrl);
    if (res.status != 200) {
        throw "res.status != 200";
    }
    let data = await res.bytes();
    if (!data instanceof Uint8Array) {
        throw `!data instanceof Uint8Array`;
    }
}

async function test_get_json() {
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
    await test_get_text();
    await test_get_bytes();
    await test_get_json();
    await test_404();

    return "pass"
}

test();
