

(async () => {
    const url = "https://www.binzume.net/";
    console.log("fetch:", url);
    let res = await fetch(url);
    console.log(res.status);
    console.log(res.statusText);
    console.log(res.headers);

    let body = await res.text();
    console.log(body.substring(0, 200) + "...");
})();
