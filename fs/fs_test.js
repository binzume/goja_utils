
const fs = require("fs");
const testdataDir = '../testdata/';

async function test_readFileSync() {
    let content = fs.readFileSync(testdataDir + "test.txt", { encoding: 'utf8' });
    if (content.trim() != "Hello") {
        throw `${content} != Hello`
    }

    let buf = fs.readFileSync(testdataDir + "test.txt", { encoding: 'buffer' });
    let bytes = new Int8Array(buf);
    if (bytes.at(0) != 72) {
        throw `${bytes.at(0)} != 72`
    }

    let exception = null;
    try {
        fs.readFileSync(testdataDir + "not_found.txt");
    } catch (e) {
        exception = e;
    }
    if (exception == null) {
        throw "a exception should be thrown";
    }
}

async function test_statSync() {
    let stat = fs.statSync(testdataDir + "test.txt")
    if (stat.isDirectory()) {
        throw "stat.isDirectory()";
    }
    if (stat.size <= 0) {
        throw "stat.size <= 0";
    }
    if (stat.mtimeMs <= 0) {
        throw "stat.mtimeMs";
    }

    stat = fs.statSync(testdataDir)
    if (!stat.isDirectory()) {
        throw "!stat.isDirectory()";
    }

    stat = fs.statSync(testdataDir + "not_found.txt")
    if (stat != null) {
        throw "stat != null";
    }


    let exception = null;
    try {
        fs.statSync(testdataDir + "not_found.txt", { throwIfNoEntry: true })
    } catch (e) {
        exception = e;
    }
    if (exception == null) {
        throw "a exception should be thrown";
    }
}

async function test_writeTest() {
    let d = testdataDir + "testtmp";
    let text = "Hello, world!";
    try {
        fs.rmSync(d, { recursive: true });
    } catch { }

    fs.mkdirSync(d)
    fs.writeFileSync(d + "/test.txt", text)
    let read = fs.readFileSync(d + "/test.txt");
    fs.rmSync(d, { recursive: true });
    if (read != text) {
        throw "read:" + read;
    }
}

async function test() {
    await test_readFileSync();
    await test_statSync();
    await test_writeTest();

    return "pass"
}

test();
