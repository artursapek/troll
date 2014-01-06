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

  var green = '#ffffff', red = '#c51c1c', span = '#56595d', crosshairYColor = '#1f2021', crosshairXColor = '#343637';


  var HOST = 'http://' + window.location.hostname + ':8001'

  function roundToFive(n) {
    var off = n % 5
    if (off <= 2) {
      return n - off
    } else {
      return n + (5 - off)
    }
  }


  $.ajax({
    type: 'get',
    datatype: 'json',
    url: HOST + '/prices.json',
    success: drawCandles,
    error: error
  });

  function error() {
    $('#loading').text('server error. shit. sorry.');
  }

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
      .range([0 - padding.y, h + padding.y])

    var crosshairY = svg.append('line')
      .attr('x1', 100)
      .attr('x2', 100)
      .attr('y1', 0)
      .attr('y2', h)
      .attr('stroke', crosshairYColor)
      .attr('stroke-width', '3')

    var crosshairX = svg.append('line')
      .attr('x1', 0)
      .attr('x2', w * 2)
      .attr('y1', 0)
      .attr('y2', 0)
      .attr('stroke', crosshairXColor)
      .attr('stroke-width', '1')

    var crosshairLabelPrice = svg.append('text')
      .attr('x', 0)
      .attr('y', 45)
      .attr('class', 'anchor-label strong')

    var crosshairXLabelPrice = svg.append('text')
      .attr('class', 'anchor-label strong')
      .attr('text-anchor', 'end')

    var crosshairLabelTime = svg.append('text')
      .attr('x', 0)
      .attr('y', 30)
      .attr('class', 'anchor-label')
    // draw the CANDLES

    function clearCandles() {
      svg.selectAll('rect.high-low').data([]).exit().remove()
      svg.selectAll('rect.open-close').data([]).exit().remove()
      svg.selectAll('rect.sar').data([]).exit().remove()
      svg.selectAll('path.kumo').data([]).exit().remove()
    }

    var lastYAxis;

    function draw(cs, yaxis) {
      clearCandles()

      lastYAxis = yaxis;

      var minMax = svg.selectAll('rect.high-low').data(cs).enter().append('rect')
      var minMaxAttrs = minMax
        .attr('class', 'high-low')
        .attr('width', '1px')
        .attr('height', function (c) {
          return (yaxis(c.CandleStick.High) - yaxis(c.CandleStick.Low))
        })
        .attr('x', function (c) { return roundToFive(x(c.Time.Close)) + 1; })
        .attr('y', function (c) { return h - yaxis(c.CandleStick.High)})
        .attr('fill', span)
        ;

      var openClose = svg.selectAll('rect.open-close').data(cs).enter().append('rect')

      var candleAttrs = openClose
        .attr('class', 'open-close')
        .attr('width', '2px')
        .attr('height', function (c) {
          return Math.abs(yaxis(c.CandleStick.Open) - yaxis(c.CandleStick.Close))
        })
        .attr('x', function (c) {
          return roundToFive(Math.round(x(c.Time.Close)));
        })
        //.attr('data-range-x', function (c) { return x(c.Time.Close); })
        .attr('data-timestamp', function (c) { return c.Time.Close })
        .attr('y', function (c) { return h - (yaxis(Math.max(c.CandleStick.Close,c.CandleStick.Open))) })
        .attr('fill', '#171718')
        .attr('strokeWidth', '1')
        .attr('stroke', function (c) { return c.CandleStick.Close > c.CandleStick.Open ? green : red })
        ;


      // Draw SAR
      var SAR = svg.selectAll('rect.sar').data(cs).enter().append('rect')

      SARAttrs = SAR
        .attr('class', 'sar')
        .attr('x', function (c) {
          return roundToFive(x(c.Time.Close)) + 1
        })
        .attr('y', function (c) {
          return h - Math.round(yaxis(c.SAR.Value))
        })
        .attr('width', 1)
        .attr('height', 1)

      var kumoGenerator = d3.svg.area()
        .x(function (c) { return x(c.Time.Close + (60 * 60 * 2 * 11)) })
        .y0(function (c) { return h - yaxis(c.Ichimoku.SenkouSpanA) })
        .y1(function (c) { return h - yaxis(c.Ichimoku.SenkouSpanB) })

      var kumo = svg.append('path')
        .attr('d', kumoGenerator(cs))
        .attr('class', 'kumo')

      function drawLine(getter, className, xOffset) {
        svg.selectAll('path.' + className).data([]).exit().remove()
        xOffset = xOffset || 0;
        var generator = d3.svg.line()
          .x(function (c) { return x(c.Time.Close + xOffset) })
          .y(function (c) { return h - yaxis(getter(c)) })

        svg.append('path')
          .attr('class', className)
          .attr('fill', 'none')
          .attr('d', generator(cs))
      }


      drawLine(function (c) { return c.Ichimoku.SenkouSpanA }, 'senkou-span-a', (60 * 60 * 2 * 11))
      drawLine(function (c) { return c.Ichimoku.SenkouSpanB }, 'senkou-span-b', (60 * 60 * 2 * 11))
      drawLine(function (c) { return c.Ichimoku.TenkanSen }, 'tenkan-sen')
      drawLine(function (c) { return c.Ichimoku.KijunSen }, 'kijun-sen')
      drawLine(function (c) { return c.CandleStick.Close }, 'chikou-span', -(60 * 60 * 2 * 11))
    }


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

    function visibleCandles() {
      var start = Math.round(window.scrollX / 5) - 22
        , amt   = Math.round(window.innerWidth / 5) + 22
      return candles.slice(start, start + amt);
    }

    var yRangeCache = {};

    function yRange(cs) {
      var low = Math.min.apply(Math, cs.map(function (c) { return c.CandleStick.Low }));
      var high = Math.max.apply(Math, cs.map(function (c) { return c.CandleStick.High }));
      return d3.scale
        .linear()
        .domain([low, high])
        .range([0 + padding.y, h - padding.y])
    }

    var drawTimeout;

    $(window).scroll(function () {
      clearTimeout(drawTimeout)
      drawTimeout = setTimeout(function () {
        var cs = visibleCandles()
        if (cs.length === 0) return;
        var cacheKey = '' + cs[0].Time.Close + cs[cs.length - 1].Time.Close;
        var yr;
        if (yRangeCache[cacheKey] !== undefined) {
          yr = yRangeCache[cacheKey];
        } else {
          var yr = yRange(cs);
          yRangeCache[cacheKey] = yr;
        }

        draw(cs, yr);
      }, 10);
    });

    $(document).mousemove(function (e) {
      var x = roundToFive(e.pageX);
      crosshairY.attr('x1', x + 1).attr('x2', x + 1)

      crosshairX
        .attr('y1', e.pageY)
        .attr('y2', e.pageY)

      if (lastYAxis !== undefined) {
        lyr = lastYAxis.domain()
        wir = [padding.y, window.innerHeight - padding.y]

        priceAtXCross = 1 - ( (e.pageY - padding.y) / (window.innerHeight - (padding.y * 2)))
        priceAtXCross *= (lyr[1] - lyr[0])
        priceAtXCross += lyr[0]

        crosshairXLabelPrice 
          .text('$' + priceAtXCross.toFixed(3))
          .attr('x', window.scrollX + window.innerWidth - 30)
          .attr('y', e.pageY - 10)
      }

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
      case 67: // C
        $body.toggleClass('hide-chikou');
        break;
      case 76: // C
        $body.toggleClass('hide-legend');
        break;
    }
  });

}());
