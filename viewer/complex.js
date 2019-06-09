function Complex(re, im) {
  this.re = re;
  this.im = im;
  this.add = function(that) {
    var re = this.re + that.re;
    var im = this.im + that.im;
    return new Complex(re, im);
  };
  this.mult = function(that) {
    var re = this.re * that.re - this.im * that.im;
    var im = this.im * that.re + this.re * that.im;
    return new Complex(re, im);
  }
  this.mod2 = function() {
    return this.re * this.re + this.im * this.im;
  }
}
