
var save = document.getElementById('save')
var bind = document.getElementById('bindMail')


function httpPostAsync(theUrl, body)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", theUrl, true); // true for asynchronous
    xmlHttp.setRequestHeader("Content-Type", "application/json");
    xmlHttp.onreadystatechange = function() {//Call a function when the state changes.
      console.log("ok");
    if(xmlHttp.readyState == XMLHttpRequest.DONE && xmlHttp.status == 200) {

    }
}
    console.log(JSON.stringify(body));
    xmlHttp.send(body);
}


function saveServices(){

  var user = document.getElementById('users').value

  var all = document.querySelectorAll('input[class=form-check-input]:checked');


  var body = "{\"Name\": \""+user+"\", \"Services\":["

  for (var i = 0; i < all.length; i++) {
    body += "\""+ all[i].value + "\""
    if(i < all.length - 1){
      body +=","
    }

  }

  body += "]}"
  console.log(body);

  httpPostAsync("/saveServices/",body)
}

function bindEmail(){
  var user = document.getElementById('users').value


  var m = document.getElementById('email').value

  var p = document.getElementById('pwd').value



  var body = "{ \"User\": "+user+", \"Mail\": \""+m+"\", \"Pwd\": \"" + p + "\" }"


  console.log(body);

  httpPostAsync("/bindMail/",body)
}



save.addEventListener("click",saveServices)
bind.addEventListener("click",bindEmail)
