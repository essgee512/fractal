function CoolBlue() {
  var cmap = []

  for (var i = 0; i <= 255; i++) {
    cmap.push({
      r: 255-i,
      g: 255-i,
      b: 255-i,
      a: i
    })
  }

  for (var i = 0; i <= 255; i++) {
    cmap.push({
      r: 0,
      g: 255-i,
      b: 255,
      a: 255-i
    })
  }

  for (var i = 0; i <= 255; i++) {
    cmap.push({
      r: 255-i,
      g: 255-i,
      b: 255-i,
      a: i
    })
  }

  for (var i = 0; i <= 255; i++) {
    cmap.push({
      r: 0,
      g: 255-i,
      b: 255,
      a: 255
    })
  }

  return cmap
}
