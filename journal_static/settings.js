window.onload = function(){
  document.getElementById("theme_normal").onclick = function(){
    sendData("normal");
  };
  document.getElementById("theme_dark").onclick = function(){
    sendData("dark");
  };
  document.getElementById("theme_darkblue").onclick = function(){
    sendData("darkblue");
  };
  document.getElementById("theme_red").onclick = function(){
    sendData("red");
  };
  document.getElementById("theme_grey").onclick = function(){
    sendData("grey");
  };
};

// NOTE: Currently only theme is implemented
function sendData(value){
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", "settingschange", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send("theme=" + value);
}