db.test_prices_analyzed.find().forEach(function (p) {
  print([p.servertime, p.price, p.analysis.volatility["6"], p.analysis.volatility["12"], p.analysis.volatility["24"]].join(','))
});
