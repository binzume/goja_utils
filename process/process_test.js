
const process = require("process");

async function test_env() {
    if (process.env.TEST != "Hello") {
        throw process.env.TEST
    }
}

async function test_stdin() {
    let txt = process.stdin.read();
    if (txt != "testtest") {
        throw txt;
    }
}

async function test() {
    await test_env();
    await test_stdin();

    return "pass";
}

test();
