// =======================
// Functions
// =======================

const createBox = (title, x, y, borderColor) => {
  let rect = two.makeRoundedRectangle(x, y, 100, 50, 5);
  rect.fill = '#FFF';
  rect.stroke = borderColor;
  rect.linewidth = 5;
  two.makeText(title, x, y, { family: fontFamily, weight: 800 });
};

const base = () => {
  // Create the client box
  createBox('Client', clientPos.x, clientPos.y, '#1C75BC');

  // Create the directory box
  createBox('Directory', directoryPos.x, directoryPos.y, '#751CBC');

  // Create the nodes
  createBox('Node #1', width * 0.8, height * 0.3, '#751CBC');
  createBox('Node #2', width * 0.8, height * 0.5, '#751CBC');
  createBox('Node #3', width * 0.8, height * 0.7, '#751CBC');

  // Create the service box
  createBox('Service', width * 0.5, height * 0.9, '#751CBC');
};

const directoryClientRoute = (frameCount) => {
  base();
};

const clientDirectoryArrow = (frameCount) => {
  if (arrowToX < directoryPos.x) {
    arrowToX += 1;
  } else {
    two.bind('update', directoryClientRoute);
  }
  if (arrowToY > directoryPos.y) {
    arrowToY -= 0.9;
  }
  arrow.remove();
  arrow = two.makeArrow(arrowFromX, arrowFromY, arrowToX, arrowToY);
  base();
};

// =======================
// Initialization
// =======================

$('#loader').hide();

// =======================
// Global constants
// =======================

let dashboard = document.getElementById('dashboard');
const width = dashboard.offsetWidth;
const height = 600;

const fontFamily = 'Helvetica';
const clientPos = { x: width * 0.1, y: height * 0.4 };
const directoryPos = { x: width * 0.5, y: height * 0.1 };

// =======================
// Drawing
// =======================

let two = new Two({
  width: width,
  height: height,
}).appendTo(dashboard);

// Initialize all shapes first
let arrowFromX = clientPos.x;
let arrowFromY = clientPos.y;
let arrowToX = arrowFromX + 0.1;
let arrowToY = arrowFromY - 0.1;
let arrow = two.makeArrow(arrowFromX, arrowFromY, arrowToX, arrowToY);

let routeText = two.makeText(
  '[8000] -> [8001] -> [8002]',
  directoryPos.x,
  directoryPos.y,
  {
    family: fontFamily,
    weight: 800,
  }
);
routeText.remove();

// Then start with the first animation sequence
two.bind('update', clientDirectoryArrow);
two.play();
