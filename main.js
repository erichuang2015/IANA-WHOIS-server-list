var xmlHttp = new XMLHttpRequest();

var latestTLD = "",
    latestflag = false,
    httpserver = "http://localhost:8081";

xmlHttp.open("GET", httpserver + "/whois/latestTLD", false);
xmlHttp.onload = function (e) {
    latestTLD = this.responseText;
};
xmlHttp.send(null);

var tldTags = document.getElementsByClassName("tld");
for (var i = 0; i < tldTags.length; i++) {

    var tld = tldTags[i].innerText,
        whiospage = tldTags[i].firstChild.getAttribute("href");

    if (latestTLD && !latestflag) {
        console.info("for TLD: " + tld)
        if (tld != latestTLD) {
            continue
        }
        console.debug("found flag: " + tld)
        latestflag = true
        continue
    }

    if (tld.indexOf(".") != 0) {//不是以(.)开头，说明是特殊符号域名，不采集
        continue
    }

    console.info(whiospage)

    xmlHttp.open("GET", whiospage, false);
    xmlHttp.onload = function (e) {
        var detailpage = document.createElement("div");
        detailpage.innerHTML = this.responseText;

        try {
            var serverP = detailpage.getElementsByTagName("h2")[4].nextElementSibling;
            if (serverP.innerHTML.trim() === "") {//whois所在的p元素内没有可用的内容，不采集
                return
            }

            var server = serverP.lastChild.textContent.trim();
            if (server.indexOf("whois.") === -1) {//p元素最后一个节点的内容不是以 whois 开头的，不采集
                return
            }

            var content = JSON.stringify({ "TLD": tld, "server": server });
            console.info(content)

            xmlHttp.open("POST", httpserver + "/whois", true);//异步提交采集结果
            xmlHttp.send(content);
        } catch (error) {
            console.error(error)
            return
        }
    }
    xmlHttp.send(null);
}