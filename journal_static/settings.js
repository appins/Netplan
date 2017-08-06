var theme_n, theme_d, theme_db, theme_r, theme_g;

window.onload = function(){
  theme_n = document.getElementById("theme_normal");
  theme_d = document.getElementById("theme_dark");
  theme_db = document.getElementById("theme_darkblue");
  theme_r = document.getElementById("theme_red");
  theme_g = document.getElementById("theme_grey");

  theme_n.onclick = function(){sendData("normal")};
  theme_d.onclick = function(){sendData("dark")};
  theme_db.onclick = function(){sendData("darkblue")};
  theme_r.onclick = function(){sendData("red")};
  theme_g.onclick = function(){sendData("grey")};

  var xhttptheme = new XMLHttpRequest();
  xhttptheme.open("GET", "theme.css", false)
  xhttptheme.send();
  document.getElementById("theme").innerHTML = xhttptheme.responseText;

  if(theme == "normal"){
    theme_n.checked = true;
  }
  else if(theme == "dark"){
    theme_d.checked = true;
  }
  else if(theme == "darkblue"){
    theme_db.checked = true;
  }
  else if(theme == "red"){
    theme_r.checked = true;
  }
  else if(theme == "grey"){
    theme_g.checked = true;
  }
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
