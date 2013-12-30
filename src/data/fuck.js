db.test_prices.find().forEach(function (p) {
  db.test_prices.update({ _id: p._id }, { time: { server: p.time.server, local: p.time.server }});
  print(p._id);
});
