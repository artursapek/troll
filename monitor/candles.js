(function () {
  var  w = innerWidth
     , h = innerHeight
     ;

  var padding = {
    x: 0
  , y: 100
  }

  var svg, x, y;

  var green = '#ffffff', red = '#c51c1c', span = '#56595d', crosshair = '#1f2021';

  $.getJSON('http://localhost:8080/prices.json', drawCandles);

  function drawCandles(candles) {
    candles = candles.slice(1000)
    var minPrice = candles.reduce(function (a, c) {
      return Math.min(a, c.CandleStick.Low, c.CandleStick.Close, c.CandleStick.Open);
    }, 1/0);
    var maxPrice = candles.reduce(function (a, c) {
      return Math.max(a, c.CandleStick.High, c.CandleStick.Close, c.CandleStick.Open);
    }, 0);

    svg = d3.select('body')
      .append('svg')
      .attr('shape-rendering', 'crispEdges')
      .attr('width', w * 4)
      .attr('height', h)

    x = d3.scale.linear().domain(candles.map(function (c) { return c.Time.Close }))
    y = d3.scale.linear().domain([minPrice - padding.y, maxPrice + padding.y])

    var crosshairY = svg.append('line')
      .attr('x1', 100)
      .attr('x2', 100)
      .attr('y1', 0)
      .attr('y2', w * 10)
      .attr('stroke', crosshair)
      .attr('stroke-width', '3')

    var crosshairLabelPrice = svg.append('text')
      .attr('x', 0)
      .attr('y', 60)
      .attr('class', 'anchor-label')

    var crosshairLabelTime = svg.append('text')
      .attr('x', 0)
      .attr('y', 40)
      .attr('class', 'anchor-label')

    var minMax = svg.selectAll('rect.high-low').data(candles).enter().append('rect')
    var minMaxAttrs = minMax
      .attr('class', 'high-low')
      .attr('width', '1px')
      .attr('height', function (c) {
        return (y(c.CandleStick.High) - y(c.CandleStick.Low)) * h
      })
      .attr('x', function (c) { return x(c.Time.Close) * 5; })
      .attr('y', function (c) { return h - y(c.CandleStick.High) * h })
      .attr('fill', span)
      ;

    var openClose = svg.selectAll('rect.open-close').data(candles).enter().append('rect')

    var candleAttrs = openClose
      .attr('class', 'open-close')
      .attr('width', '2px')
      .attr('height', function (c) {
        return Math.abs(y(c.CandleStick.Open) - y(c.CandleStick.Close)) * h
      })
      .attr('x', function (c) { return x(c.Time.Close) * 5 - 1; })
      .attr('data-range-x', function (c) { return Math.round(x(c.Time.Close) * 5 - 1); })
      .attr('data-timestamp', function (c) { return c.Time.Close })
      .attr('y', function (c) { return h - (y(Math.max(c.CandleStick.Close,c.CandleStick.Open)) * h) })
      .attr('fill', '#171718')
      .attr('strokeWidth', '1')
      .attr('stroke', function (c) { return c.CandleStick.Close > c.CandleStick.Open ? green : red })
      ;

    $(document).mousemove(function (e) {
      var x = e.pageX;
      x -= (x % 5);
      crosshairY.attr('x1', x).attr('x2', x)

      var $selectedCandleElem = $('rect[data-range-x="'+ (x - 1) +'"]')
        , selectedTime = $selectedCandleElem.attr('data-timestamp')
        , selectedCandle = candles.filter(function (can) { return can.Time.Close == selectedTime })[0]

      crosshairLabelPrice
        .attr('x', x + 10)
        .text('$' + selectedCandle.CandleStick.Close.toFixed(3))
        ;
      var date = new Date(parseInt(selectedTime, 10) * 1000);

      crosshairLabelTime
        .attr('x', x + 10)
        .text(date.toDateString() + ' ' + date.getHours() + ':00')
    });

    $.getJSON('http://localhost:8080/trades.json', markTrades);
  }

  function markTrades(trades) {
    var tradeMarks = svg.selectAll('line.trade')
      .data(trades)
      .enter()
      .append('line');

    var tradeAttrs = tradeMarks
      .attr('class', function (t) { return t.Type })
      .attr('x1', function (t) { return x(t.Timestamp) * 5 })
      .attr('x2', function (t) { return x(t.Timestamp) * 5 })
      .attr('y1', 0)
      .attr('y1', h)
  }

}());
