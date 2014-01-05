(function () {
  var  w
    , brushHeight = 80
    , h = innerHeight
    ;

  var padding = {
    x: 0
  , y: 100
  }

  var svg, x, y;

  var green = '#ffffff', red = '#c51c1c', span = '#56595d', crosshair = '#1f2021';

  var HOST = 'http://' + window.location.hostname + ':8001'

  $.getJSON(HOST + '/prices.json', drawCandles);

  function drawCandles(candles) {
    candles = candles.filter(function (c) {
      return c.CandleStick.Close > 0
    });

    w = candles.length * 5

    candles = candles.sort(function (a, b) {
      if (a.Time.Close < b.Time.Close) {
        return -1;
      } else {
        return 1;
      }
    })

    var minPrice = candles.reduce(function (a, c) {
      return Math.min(a, c.CandleStick.Low, c.CandleStick.Close, c.CandleStick.Open);
    }, 1/0);
    var maxPrice = candles.reduce(function (a, c) {
      return Math.max(a, c.CandleStick.High, c.CandleStick.Close, c.CandleStick.Open);
    }, 0);

    ww = w + 5

    svg = d3.select('body')
      .append('svg')
      .attr('shape-rendering', 'crispEdges')
      .attr('width', Math.max(ww + innerWidth / 3, window.innerWidth))
      .attr('height', h)

    x = d3.scale.linear()
      .domain(d3.extent(candles, function (c) { return c.Time.Close }))
      .range([0, ww])

    y = d3.scale
      .linear()
      .domain([minPrice - padding.y, maxPrice + padding.y])
      .range([0, h])

    var crosshairY = svg.append('line')
      .attr('x1', 100)
      .attr('x2', 100)
      .attr('y1', 0)
      .attr('y2', h)
      .attr('stroke', crosshair)
      .attr('stroke-width', '3')

    var crosshairLabelPrice = svg.append('text')
      .attr('x', 0)
      .attr('y', 45)
      .attr('class', 'anchor-label strong')

    var crosshairLabelTime = svg.append('text')
      .attr('x', 0)
      .attr('y', 30)
      .attr('class', 'anchor-label')

    // Draw SAR
    var SAR = svg.selectAll('rect.sar').data(candles).enter().append('rect')

    SARAttrs = SAR
      .attr('class', 'sar')
      .attr('x', function (c) {
        return roundToFive(x(c.Time.Close)) + 1
      })
      .attr('y', function (c) {
        return h - Math.round(y(c.SAR.Value))
      })
      .attr('width', 1)
      .attr('height', 1)

    var kumoGenerator = d3.svg.area()
      .x(function (c) { return x(c.Time.Close) })
      .y0(function (c) { return h - y(c.Ichimoku.SenkouSpanA) })
      .y1(function (c) { return h - y(c.Ichimoku.SenkouSpanB) })

    var kumo = svg.append('path')
      .attr('d', kumoGenerator(candles))
      .attr('class', 'kumo')

    function drawLine(getter, className) {
      var generator = d3.svg.line()
        .x(function (c) { return x(c.Time.Close) })
        .y(function (c) { return h - y(getter(c)) })

      svg.append('path')
        .attr('class', className)
        .attr('fill', 'none')
        .attr('d', generator(candles))
    }

    function roundToFive(n) {
      var off = n % 5
      if (off <= 2) {
        return n - off
      } else {
        return n + (5 - off)
      }
    }

    drawLine(function (c) { return c.Ichimoku.SenkouSpanA }, 'senkou-span-a')
    drawLine(function (c) { return c.Ichimoku.SenkouSpanB }, 'senkou-span-b')
    drawLine(function (c) { return c.Ichimoku.TenkanSen }, 'tenkan-sen')
    drawLine(function (c) { return c.Ichimoku.KijunSen }, 'kijun-sen')

    // draw the CANDLES

    var minMax = svg.selectAll('rect.high-low').data(candles).enter().append('rect')
    var minMaxAttrs = minMax
      .attr('class', 'high-low')
      .attr('width', '1px')
      .attr('height', function (c) {
        return (y(c.CandleStick.High) - y(c.CandleStick.Low))
      })
      .attr('x', function (c) { return roundToFive(x(c.Time.Close)) + 1; })
      .attr('y', function (c) { return h - y(c.CandleStick.High)})
      .attr('fill', span)
      ;

    var openClose = svg.selectAll('rect.open-close').data(candles).enter().append('rect')

    var candleAttrs = openClose
      .attr('class', 'open-close')
      .attr('width', '2px')
      .attr('height', function (c) {
        return Math.abs(y(c.CandleStick.Open) - y(c.CandleStick.Close))
      })
      .attr('x', function (c) {
        return roundToFive(Math.round(x(c.Time.Close)));
      })
      //.attr('data-range-x', function (c) { return x(c.Time.Close); })
      .attr('data-timestamp', function (c) { return c.Time.Close })
      .attr('y', function (c) { return h - (y(Math.max(c.CandleStick.Close,c.CandleStick.Open))) })
      .attr('fill', '#171718')
      .attr('strokeWidth', '1')
      .attr('stroke', function (c) { return c.CandleStick.Close > c.CandleStick.Open ? green : red })
      ;


    function orientCrosshairText(dir, x) {
      var offset = 10;
      switch (dir) {
      case "left":
        crosshairLabelPrice
          .attr('x', x - offset)
          .attr('text-anchor', 'end')
        crosshairLabelTime
          .attr('x', x - offset)
          .attr('text-anchor', 'end')
        break;

      case "right":
        crosshairLabelPrice
          .attr('x', x + offset)
          .attr('text-anchor', 'start')
        crosshairLabelTime
          .attr('x', x + offset)
          .attr('text-anchor', 'start')
        break;
      }
    }

    $(document).mousemove(function (e) {
      var x = roundToFive(e.pageX);
      crosshairY.attr('x1', x + 1).attr('x2', x + 1)

      var $selectedCandleElem = $('rect.open-close[x="'+ (x) +'"]')
        , selectedTime = $selectedCandleElem.attr('data-timestamp')
        , selectedCandle = candles.filter(function (can) { return can.Time.Close == selectedTime })[0]

      if (selectedCandle) {

        dir = (e.clientX > (innerWidth / 2)) ? 'left' : 'right'
        orientCrosshairText(dir, x)

        crosshairLabelPrice
          .attr('class', 'anchor-label strong')
          .text('$' + selectedCandle.CandleStick.Close.toFixed(3))
          ;
        var date = new Date(parseInt(selectedTime, 10) * 1000);

        crosshairLabelTime
          .attr('class', 'anchor-label')
          .text(date.toDateString() + ' ' + date.getHours() + ':00')

      } else {
        crosshairLabelPrice.attr('class', 'hidden')
        crosshairLabelTime.attr('class', 'hidden')
      }
    });

    //$.getJSON(HOST + '/trades.json', markTrades);

    setTimeout(function () {
      $('#loading').remove();
      window.scrollTo(ww * 2, 0)
    },0);
  }

  function markTrades(trades) {
    var tradeMarks = svg.selectAll('line.trade')
      .data(trades)
      .enter()
      .append('line');

    var tradeAttrs = tradeMarks
      .attr('class', function (t) { return t.Type })
      .attr('x1', function (t) { return x(t.Timestamp) + 1 })
      .attr('x2', function (t) { return x(t.Timestamp) + 1 })
      .attr('y1', function (t) { return h - y(t.Rate) - 150})
      .attr('y2', function (t) { return h - y(t.Rate) + 150})
  }

  $(document).keyup(function (e) {
    var $body = $('body')
    switch (e.which) {
      case 75: // K
        $body.toggleClass('hide-kumo');
        break;
      case 83: // S
        $body.toggleClass('hide-sar');
        break;
      case 73: // I
        $body.toggleClass('hide-tenkan-kijun');
        break;
    }
  });

}());
