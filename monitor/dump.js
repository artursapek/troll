all = []
db.test_2hrs.find().forEach(function (d) {
  var c= d.candlestick;
  all.push({
    time: d.time.close.valueOf()
  , high: c.high
  , low: c.low
  , open: c.open
  , close: c.close
  });
});

printjson(all)
