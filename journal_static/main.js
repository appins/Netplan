var journalNumber, journalText, num = 0;

window.onload = function() {
  var nextButton = document.getElementById("lastpage");
  nextButton.onclick = last;
  var lastButton = document.getElementById("nextpage");
  lastButton.onclick = next;

  journalNumber = document.getElementById("entrynum");
  journalText = document.getElementById("journal");

  journalNumber.innerHTML = reqNumber;
  loadContent(reqNumber);

  setTimeout(sendContent, 1000);
};

function next() {
  var checkSave = "";
  while(checkSave != journalText.innerHTML){
    var xhttp0 = new XMLHttpRequest();
    xhttp0.open("POST", "entryedit/" + String(num), false);
    xhttp0.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp0.send("text=" + journalText.innerHTML)

    var xhttp1 = new XMLHttpRequest();
    xhttp1.open("GET", "entry/" + String(num), false);
    xhttp1.send();
    checkSave = xhttp1.responseText;
    console.log("makin sure");
  }
  if(reqNumber >= 2500){
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
    if(xhttp.readyState == 4)
      journalText.innerHTML = xhttp.responseText;
  };
  xhttp.open("GET", "entry/" + String(num), true);
  xhttp.send();
}

var lastJournal = "";

function sendContent(){
  setTimeout(function(){
    if(lastJournal != journalText.innerHTML || num < 5){
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
    sendContent();
  }, 150);
}
