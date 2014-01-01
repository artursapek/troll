db.test_intervals.find().forEach(function (x) {
  var i = x.ichimoku
  print([x.time.close.valueOf(), x.candlestick.low, x.candlestick.open, x.candlestick.close, x.candlestick.high, i.tenkensen, i.kijunsen, i.senkouspana, i.senkouspanb].join(','));
});
