var journalNumber, journalText;
var reqNumber = 1;

window.onload = function() {
  var nextButton = document.getElementById("lastpage");
  nextButton.onclick = last;
  var lastButton = document.getElementById("nextpage");
  lastButton.onclick = next;

  journalNumber = document.getElementById("entrynum");
  journalText = document.getElementById("jorunalText");

  journalNumber.innerHTML = reqNumber;
};

function next() {
  console.log("next");
}

function last() {
  console.log("last");
}
