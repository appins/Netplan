var theme_n, theme_d, theme_db, theme_r, theme_g;

window.onload = function(){
  theme_n = document.getElementById("theme_normal");
  theme_d = document.getElementById("theme_dark");
  theme_db = document.getElementById("theme_darkblue");
  theme_r = document.getElementById("theme_red");
  theme_g = document.getElementById("theme_grey");
  theme_gr = document.getElementById("theme_green");

  theme_n.onclick = function(){sendData("normal")};
  theme_d.onclick = function(){sendData("dark")};
  theme_db.onclick = function(){sendData("darkblue")};
  theme_r.onclick = function(){sendData("red")};
  theme_g.onclick = function(){sendData("grey")};
  theme_gr.onclick = function(){sendData("green")};

  var xhttptheme = new XMLHttpRequest();
  xhttptheme.open("GET", "theme.css", false)
  xhttptheme.send();
  document.getElementById("theme").innerHTML = xhttptheme.responseText;
};

// NOTE: Currently only theme is implemented
function sendData(value){
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", "settingschange", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send("theme=" + value);

  xhttp.onreadystatechange = function() {
    var xhttptheme = new XMLHttpRequest();
    xhttptheme.open("GET", "theme.css", false)
    xhttptheme.send();
    document.getElementById("theme").innerHTML = xhttptheme.responseText;
  }
}
