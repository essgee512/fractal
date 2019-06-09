function Fractal(fmap) {
  this.fmap = fmap

  // methods
  this.render = function(view) {
    var img  = new ImageData(view.w, view.h)
    var data = img.data

    view.pixels.forEach(function(p) {
      renderPixel(fmap, data, p, view)
    });

    view.ctx.putImageData(img, 0, 0)
  };

  // this.renderPixels = function() {
  // };
}

function renderPixel(fmap, data, p, view) {
    var s  = view.s
    var cx = view.cx; var cy = view.cy;
    var w  = view.w;  var h  = view.h;
    var k  = p.k;     var l  = p.l

    var x = s * ((3.0/w)*k + cx - 1.5) + cx;
    var y = s * ((2.0/h)*l + cy - 1.0) + cy;

    var z = new Complex(x, y);

    var c = fmap(z); // c = {r: 128, g: 128, b: 128, a: 255};

    var s = 4 * p.i;
    data[s+0] = c.r;
    data[s+1] = c.g;
    data[s+2] = c.b;
    data[s+3] = c.a;
}
