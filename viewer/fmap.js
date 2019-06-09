function Julia(Cj, Cmap, R, iMax) {
  var cmap = Cmap()

  return function(z) {
    var i = 0;

    while (z.mod2() < R) {
      z = z.mult(z).add(Cj);
      if (i == iMax) {
        break;
      }
      i++;
    }

    var n = Math.floor( i*(cmap.length-1)/iMax );
    return cmap[n];
  }
}


function Mandelbrot(Cmap, R, iMax) {
  var cmap = Cmap();

  return function(c) {
    z = Complex(0, 0);
    var i = 0;

    while (z.mod2() < R) {
      z = z.mult(z).add(c);
      if (i == iMax) {
        break;
      }
      i++;
    }

    var n = i*(cmap.length-1)/iMax;

    return cmap[i];
  }
}
