
const delayMs = (t) => new Promise(r => setTimeout(r, t));

(async () => {
    console.log('started.');

    for (let i = 0; i < 10; i++) {
        await delayMs(1000);
        console.log(i);
    }
})();
