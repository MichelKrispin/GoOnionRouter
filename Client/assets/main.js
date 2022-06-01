// =======================
// Functions
// =======================

const createBox = (title, x, y, color, borderColor) => {
  let directoryRect = two.makeRectangle(x, y, 100, 50);
  directoryRect.fill = color;
  directoryRect.stroke = borderColor;
  two.makeText(title, x, y);
};

// =======================
// Initialization
// =======================

$("#loader").hide();

// =======================
// Drawing
// =======================

let dashboard = document.getElementById("dashboard");
const width = dashboard.offsetWidth;
const height = 600;

let two = new Two({
  width: width,
  height: height,
}).appendTo(dashboard);

// Create the client box
createBox(
  "Client",
  width * 0.1,
  height * 0.4,
  "rgba(0, 200, 255, 0.75)",
  "#1C75BC"
);

// Create the directory box
createBox(
  "Directory",
  width * 0.5,
  height * 0.1,
  "rgba(200, 0, 255, 0.75)",
  "#751CBC"
);

// Create the nodes
createBox(
  "Node 1",
  width * 0.8,
  height * 0.3,
  "rgba(200, 0, 255, 0.75)",
  "#751CBC"
);

createBox(
  "Node 2",
  width * 0.8,
  height * 0.5,
  "rgba(200, 0, 255, 0.75)",
  "#751CBC"
);

createBox(
  "Node 2",
  width * 0.8,
  height * 0.7,
  "rgba(200, 0, 255, 0.75)",
  "#751CBC"
);

// Create the service box
createBox(
  "Service",
  width * 0.5,
  height * 0.9,
  "rgba(200, 0, 255, 0.75)",
  "#751CBC"
);

// Draw everything to the screen
two.update();
