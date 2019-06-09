function Canvas(w, h) {
  var elem    = document.createElement('canvas');
  elem.width  = this.w = w;
  elem.height = this.h = h;
  this.elem   = elem
  this.ctx    = elem.getContext('2d');

  this.pixels = initPixels(this);

  // methods
  this.display = function(imgData) {
    this.ctx.putImageData(imgData, 0, 0);
  }
}

function initPixels(c) {
  var i = 0;
  var pixels = [];

  for (var k = 0; k < c.h; k++) {
    for (var l = 0; l < c.w; l++) {
      p = {
        i: i,
        k: l, l: k
      };

      pixels.push(p);

      i++;
    }
  }

  return pixels;
}
