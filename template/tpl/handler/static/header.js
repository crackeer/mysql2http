
var styleFiles = [
    "/bootstrap.css",
    "/jsoneditor.css"
]
var jsFile1 = [
    "/jquery.js",
    "/vue.global.js",
    "/axios.js",
    "/jsoneditor.js"
]
var jsFile2 = [
    "/bootstrap.js",
]
document.addEventListener("DOMContentLoaded", async () => {
    loadStyles(styleFiles)
    await loadJs(jsFile1)
    await loadJs(jsFile2)
    await sleep(200)
    startWork()

}, false);


async function loadStyles(urls) {
    var head = document.getElementsByTagName("head")[0];
    for (var i = 0; i < urls.length; i++) {
        head.appendChild(createStyleNode(urls[i]));
    }
}

function createStyleNode(url) {
    var link = document.createElement("link");
    link.type = "text/css";
    link.rel = "stylesheet";
    link.href = url;
    return link
}

async function loadJs(urls) {
    for (var i = 0; i < urls.length; i++) {
        await loadJsUrl(urls[i])
    }
}

function loadJsUrl(url) {
    return new Promise((resolve) => {
        let domScript = createJsNode(url)
        domScript.onload = domScript.onreadystatechange = function () {
            if (!this.readyState || 'loaded' === this.readyState || 'complete' === this.readyState) {
                resolve()
            }
        }
        document.getElementsByTagName('head')[0].appendChild(domScript);
    });
}

function createJsNode(url) {
    var scriptNode = document.createElement("script");
    scriptNode.src = url;
    return scriptNode
}

function sleep(time) {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve();
        }, time);
    });
}

function initJSONEditor(target, value) {
    let jsonEditor = 'jsonEditor-' + target
    if (window[jsonEditor] == undefined) {
        window[jsonEditor] = new JSONEditor(document.getElementById(target), {
            "mode": "code",
            "search": true,
            "indentation": 4
        })
    }
    try {
        let jsonValue = JSON.parse(value)
        window[jsonEditor].set(jsonValue)
    } catch (e) {
    }
}