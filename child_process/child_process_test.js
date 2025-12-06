
const child_process = require("child_process");

async function test_execSync() {

    let a = child_process.execSync('go version', { encoding: 'utf8' });
    if (!a.startsWith('go version ')) {
        throw "go version error: " + a;
    }

    let exception = null;
    try {
        child_process.execFileSync('not_exist_command');
    } catch (e) {
        exception = e;
    }
    if (exception == null) {
        throw "a exception should be thrown";
    }
}

async function test() {
    await test_execSync();

    return "pass";
}

test();
