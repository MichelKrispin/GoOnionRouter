// =======================
// Functions
// =======================

const createBox = (title, x, y, borderColor) => {
  let rect = two.makeRoundedRectangle(x, y, boxW, boxH, 5);
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

const makeText = (text, x, y) => {
  let textBox = two.makeText(text, x, y, { family: fontFamily });
  let bounding = textBox.getBoundingClientRect();
  let rect = two.makeRectangle(
    x,
    y,
    bounding.right - bounding.left,
    bounding.bottom - bounding.top
  );
  rect.fill = '#FFF';
  rect.stroke = '#FFF';
  textBox.remove();
  textBox = two.makeText(text, x, y, { family: fontFamily });
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
const boxW = 100;
const boxH = 50;
const boxWhalf = boxW / 2;
const boxHhalf = boxH / 2;
let pause = 1000;
const timeDiff = 1000;

// =======================
// Drawing
// =======================

let two = new Two({
  width: width,
  height: height,
}).appendTo(dashboard);

// Drawing
base();
two.update();

// Draw route request
setTimeout(function () {
  two.clear();
  two.makeArrow(
    clientPos.x + boxWhalf,
    clientPos.y - boxHhalf,
    directoryPos.x - boxWhalf,
    directoryPos.y + boxHhalf
  );
  makeText(
    'GET /route',
    clientPos.x + (directoryPos.x - clientPos.x) / 2,
    clientPos.y + (directoryPos.y - clientPos.y) / 2
  );
  base();
  two.update();
}, pause);
pause += timeDiff;

// Return route
setTimeout(function () {
  two.clear();
  two.makeArrow(
    directoryPos.x - boxWhalf,
    directoryPos.y + boxHhalf,
    clientPos.x + boxWhalf,
    clientPos.y - boxHhalf
  );
  makeText(
    '[8000] -> [8001] -> [8002]',
    clientPos.x + (directoryPos.x - clientPos.x) / 2,
    clientPos.y + (directoryPos.y - clientPos.y) / 2
  );
  base();
  two.update();
}, pause);
pause += timeDiff;

// Reset
setTimeout(function () {
  two.clear();
  base();
  two.update();
}, pause);
pause += timeDiff;
