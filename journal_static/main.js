var journalNumber, journalText;
var reqNumber = 1;

window.onload = function() {
  var nextButton = document.getElementById("lastpage");
  nextButton.onclick = last;
  var lastButton = document.getElementById("nextpage");
  lastButton.onclick = next;

  journalNumber = document.getElementById("entrynum");
  journalText = document.getElementById("journal");

  journalNumber.innerHTML = reqNumber;
  loadContent(1);

  sendContent();
};

function next() {
  if(reqNumber >= 10000){
    return;
  }
  reqNumber++;
  journalNumber.innerHTML = String(reqNumber);
  loadContent(reqNumber);
}

function last() {
  if(reqNumber <= 1){
    return;
  }
  reqNumber--;
  journalNumber.innerHTML = String(reqNumber);
  loadContent(reqNumber);
}

function loadContent(num) {
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    journalText.innerHTML = xhttp.responseText;
  };
  xhttp.open("GET", "entry/" + String(num), true);
  xhttp.send();
}

var lastJournal = "";

function sendContent(){
  setTimeout(function(){
    if(lastJournal != journalText.innerHTML){
      var xhttp = new XMLHttpRequest();
      xhttp.open("POST", "entryedit/" + String(reqNumber), true);
      xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhttp.send("text=" + journalText.innerHTML);
      lastJournal = journalText.innerHTML;
    }
    sendContent();
  }, 500);
}
