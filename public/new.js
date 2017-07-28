var h4, ae;

window.onload = function() {
  h1 = document.getElementById("h1");
  a = document.getElementById("a");

  h1.innerHTML = "Your id is <br>" + userid[0].toUpperCase() + userid.substring(1, userid.length);
  a.href = "/journal/" + userid + "/";
};
