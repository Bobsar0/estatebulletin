//Function to toggle cards
function myFunction() {
      var x = document.getElementById("Demo");
      var y = document.getElementById("cardbutton");
      
      if (x.className.indexOf("w3-show") == -1) {
          x.className += " w3-show";
          y.innerHTML = " X "
      } else {
          x.className = x.className.replace(" w3-show", "");
          y.innerHTML = "Chat"
          
      }
  }