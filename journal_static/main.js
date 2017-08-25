var journalNumber, journalText, num = 0, lastTitle, pageNum, byteCount, byteCountNum = 0;

window.onload = function() {
  var nextButton = document.getElementById("lastpage");
  nextButton.onclick = last;
  var lastButton = document.getElementById("nextpage");
  lastButton.onclick = next;

  journalNumber = document.getElementById("entrynum");
  journalText = document.getElementById("journal");

  journalNumber.innerHTML = reqNumber;
  loadContent(reqNumber);

  byteCount = document.getElementById("bytes");
  pageNum = document.getElementById("pagenum");

  pageNum.innerHTML = reqNumber;
  sendContent();
  countBytes();
};

function next() {
  if(reqNumber >= 10000){
    return;
  }
  reqNumber++;
  loadContent(reqNumber);
  pageNum.innerHTML = reqNumber;
}

function last() {
  if(reqNumber <= 1){
    return;
  }
  reqNumber--;
  loadContent(reqNumber);
  pageNum.innerHTML = reqNumber;
}

function loadContent(num) {
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    journalText.innerHTML = xhttp.responseText;
  };
  xhttp.open("GET", "entry/" + String(num), true);
  xhttp.send();

  var xhttp2 = new XMLHttpRequest();
  xhttp2.onreadystatechange = function() {
    journalNumber.innerHTML = xhttp2.responseText;
  };
  xhttp2.open("GET", "title/" + String(num), true);
  xhttp2.send();
}

var lastJournal = "";

function sendContent(){
  setTimeout(function(){
    if(lastJournal != journalText.innerHTML || num < 2){
      if(lastJournal != journalText.innerHTML){
        num = 0;
      }
      var xhttp = new XMLHttpRequest();
      xhttp.open("POST", "entryedit/" + String(reqNumber), true);
      xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhttp.send("text=" + journalText.innerHTML);
      lastJournal = journalText.innerHTML;
      num++;
    }
    if(lastTitle != journalNumber.innerHTML) {
      var xhttp = new XMLHttpRequest();
      xhttp.open("POST", "titleedit/" + String(reqNumber), true);
      xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhttp.send("text=" + journalNumber.innerHTML);
      lastTitle = journalNumber.innerHTML;
    }
    sendContent();
  }, 500);
}

function countBytes(){
  setTimeout(function(){
    if(byteCountNum != journalText.innerHTML.length){
      byteCount.innerHTML = journalText.innerHTML.length;
      byteCountNum = byteCount.innerHTML;
    }
    countBytes();
  }, 100);
}
