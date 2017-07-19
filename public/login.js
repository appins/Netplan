var idinput, linkoutput;

window.onload = function() {
  idinput = document.getElementById("inputid");
  linkoutput = document.getElementById("gologin");
  setlink();
};

function setlink() {
  setTimeout(function() {
    linkoutput.href = "/journal/" + idinput.value;
    setlink();
  }, 50);
}
