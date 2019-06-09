function Viewer(params) {
  this.size     = params.size;     // 400
  this.center   = params.center;   // {x: 0, y: 0}
  this.scale    = params.scale;    // 1.0
  // this.fmap     = params.fmap;     // new Julia(c, color, R, iMax)
  viewerHTML    = document.getElementById('viewer');

  this.fractal  = new Fractal(params.fmap);

  // derived
  var w = 3 * params.size;
  var h = 2 * params.size;

  this.canvas = new Canvas(w, h);
  viewerHTML.appendChild(this.canvas.elem);

  // methods
  this.display = function() {
    var view = {
      cx:     this.center.x,
      cy:     this.center.y,
      s:      this.scale,
      w:      this.canvas.w,
      h:      this.canvas.h,
      pixels: this.canvas.pixels,
      ctx:    this.canvas.ctx
    }

    this.fractal.render(view);
  }

  registerHandlerFor(this);
}

function registerHandlerFor(v) {
  document.addEventListener('keydown',
    function(e) {
      switch(e.keyCode) {
        case 37: v.center.x -= 0.25*v.scale; // left
          break;
        case 38: v.center.y -= 0.25*v.scale; // up
          break;
        case 39: v.center.x += 0.25*v.scale; // right
          break;
        case 40: v.center.y += 0.25*v.scale; // down
          break;
        case 90: v.scale    *= 0.75; // zoom in  (z)
          break;
        case 65: v.scale    /= 0.75; // zoom out (a)
          break;
        case 67: v.center.x = 0; v.center.y = 0; // c
          break;
        default:
          console.log('not a valid key')
      }
      v.display();
    }, false);
}


window.onload = function() {
  Cj = new Complex(-0.8, 0.156);
  new Viewer({
    size:    250,
    center:  {x: 0, y: 0},
    scale:   1.0,
    fmap:    Julia(Cj, CoolBlue, 8, 100)
    // fmap:    Mandelbrot(cmap, R, iMax);
  }).display();
}
