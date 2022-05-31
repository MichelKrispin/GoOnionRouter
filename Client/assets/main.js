$("#loader").hide();

var params = {
  fullscreen: false,
};
var elem = document.getElementById("dashboard");
var two = new Two(params).appendTo(elem);

// Create the client box
let clientX = 110;
let clientY = 100;
let clientRect = two.makeRectangle(clientX, clientY, 100, 50);
clientRect.fill = "rgba(0, 200, 255, 0.75)";
clientRect.stroke = "#1C75BC";
let clientText = two.makeText("Client", clientX, clientY);

// Create the directory box
let directoryX = 500;
let directoryY = 50;
let directoryRect = two.makeRectangle(directoryX, directoryY, 100, 50);
directoryRect.fill = "rgba(0, 200, 255, 0.75)";
directoryRect.stroke = "#1C75BC";
let directoryText = two.makeText("Directory", directoryX, directoryY);

// Draw everything to the screen
two.update();
