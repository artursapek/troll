(function () {
  var  w
    , brushHeight = 80
    , h = innerHeight
    , hRSI = 100
    , candles
    ;

  var padding = {
    x: 0
  , y: 20
  }

  var HOST = 'http://' + window.location.hostname + ':8001'

  var svg, x, y
    , green = '#ffffff'
    , red = '#c51c1c'
    , span = '#56595d'
    , crosshairYColor = '#1f2021'
    , crosshairXColor = '#343637'
    ;

  var settings = {
    'legend'      : true
  , 'prices'      : true
  , 'heikin-ashi' : false
  , 'ichimoku'    : false
  , 'ema-10'      : true
  , 'ema-21'      : true
  , 'sar'         : true
  , 'keltner'     : true
  , 'rsi'         : true
  }

  function storeSetting(key, setting) {
    localStorage.setItem('setting-' + key, setting);
  }

  function getSetting(key) {
    // lol:
    return { 'false': false, 'true': true }[localStorage.getItem('setting-' + key)];
  }

  function Candle(c) {
    if (settings['heikin-ashi']) {
      return c.HeikinAshi
    } else {
      return c.CandleStick;
    }
  }

  function roundToFive(n) {
    var off = n % 5
    if (off <= 2) {
      return n - off
    } else {
      return n + (5 - off)
    }
  }

  var cachedCandles = localStorage.getItem('candles')
    , lastTime      = 0
    ;

  // START DEFER
  setTimeout(function() {

  if (cachedCandles != null) {
    cachedCandles = JSON.parse(cachedCandles);
    lastTime = cachedCandles[0].Time.Close;
  }

  $.ajax({
    type: 'get',
    datatype: 'json',
    url: HOST + '/intervals.json?after=' + lastTime,
    success: drawCandles,
    error: error
  });

  function error() {
    $('#loading').text('server error. shit. sorry.');
  }

  function drawCandles(response) {
    var responseCandles = response.Intervals;
    if (responseCandles !== null) {
      responseCandles = responseCandles.filter(function (c) {
        return c.CandleStick.Close > 0 && c.CandleStick.Open > 0
      });
    }

    if (cachedCandles !== null) {
      candles = cachedCandles;
      if (responseCandles !== null) {
        candles = candles.concat(responseCandles);
      }
    } else {
      candles = responseCandles;
    }

    localStorage.setItem('candles', JSON.stringify(candles));
    console.log(candles[0].Time.Close, candles[candles.length - 1].Time.Close)


    w = candles.length * 5

    candles = candles.sort(function (a, b) {
      if (a.Time.Close < b.Time.Close) {
        return -1;
      } else {
        return 1;
      }
    })

    var minPrice = candles.reduce(function (a, c) {
      return Math.min(a, Candle(c).Low, Candle(c).Close, Candle(c).Open);
    }, 1/0);
    var maxPrice = candles.reduce(function (a, c) {
      return Math.max(a, Candle(c).High, Candle(c).Close, Candle(c).Open);
    }, 0);

    ww = w + 5

    console.log(ww);

    svg = d3.select('body')
      .append('svg')
      .attr('shape-rendering', 'crispEdges')
      .attr('width', Math.max(ww + innerWidth / 3, window.innerWidth))
      .attr('height', h)


    // Draw RSI

    x = d3.scale.linear()
      .domain(d3.extent(candles, function (c) { return c.Time.Close }))
      .range([0, ww])

    // This scale never changes
    var yrRSI = d3.scale
      .linear()
      .domain([0,100])
      .range([0 + padding.y, 0 + padding.y + hRSI])

    var RSIGenerator = d3.svg.line()
      .x(function (c) { return x(c.Time.Close) + 2 })
      .y(function (c) { return h - yrRSI(c.RSI) })

    var RSIThreshold = 25;

    function RSIGuide(n) {
      svg.append('line')
        .attr('x1', 0)
        .attr('x2', w*2)
        .attr('y1', h - yrRSI(n))
        .attr('y2', h - yrRSI(n))
        .attr('class', 'rsi-guide')
    }

    RSIGuide(0)
    RSIGuide(RSIThreshold)
    RSIGuide(100 - RSIThreshold)
    RSIGuide(100)

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

    svg.append('path')
      .attr('fill', 'none')
      .attr('class', 'rsi')
      .attr('d', RSIGenerator(candles))

    svg.append('path')
      .attr('id', 'rsi-dull')
      .attr('clip-path', 'url(#rsi-inside)')
      .attr('fill', 'none')
      .attr('class', 'rsi-dull')
      .attr('d', RSIGenerator(candles))

    svg.append('clipPath')
      .attr('id', 'rsi-inside')
    .append('rect')
      .attr('x', 0)
      .attr('y', h - 100 - yrRSI(RSIThreshold) + (100 - (RSIThreshold * 2)))
      .attr('width', w * 2)
      .attr('height', 100 - (RSIThreshold * 2))


    // draw the CANDLES

    function clearCandles() {
      svg.selectAll('rect.high-low').data([]).exit().remove()
      svg.selectAll('rect.open-close').data([]).exit().remove()
      svg.selectAll('rect.sar').data([]).exit().remove()
      svg.selectAll('path.kumo').data([]).exit().remove()
      svg.selectAll('path.keltner-inside').data([]).exit().remove()
    }

    var lastYAxis;

    function draw(cs, yaxis) {
      clearCandles()

      lastYAxis = yaxis;

      drawLine(function (c) { return c.EMA10 }, 'ema-10')
      drawLine(function (c) { return c.EMA21 }, 'ema-21')
      drawLine(function (c) { return c.EMA21 - c.ATR * 2 }, 'keltner-upper', 0, 2)
      drawLine(function (c) { return c.EMA21 + c.ATR * 2 }, 'keltner-lower', 0, 2)

      var keltnerGenerator = d3.svg.area()
        .x(function (c) { return x(c.Time.Close) + 2 })
        .y0(function (c) { return h - yaxis(c.EMA21 - c.ATR * 2) })
        .y1(function (c) { return h - yaxis(c.EMA21 + c.ATR * 2) })

      var keltnerInside = svg.append('path')
        .attr('d', keltnerGenerator(cs))
        .attr('class', 'keltner-inside')

      drawLine(function (c) { return c.Ichimoku.SenkouSpanA }, 'senkou-span-a', (60 * 60 * 2 * 11))
      drawLine(function (c) { return c.Ichimoku.SenkouSpanB }, 'senkou-span-b', (60 * 60 * 2 * 11))
      drawLine(function (c) { return c.Ichimoku.TenkanSen }, 'tenkan-sen')
      drawLine(function (c) { return c.Ichimoku.KijunSen }, 'kijun-sen')
      drawLine(function (c) { return c.CandleStick.Close }, 'chikou-span', -(60 * 60 * 2 * 11))

      var minMax = svg.selectAll('rect.high-low').data(cs).enter().append('rect')
      var minMaxAttrs = minMax
        .attr('class', 'high-low')
        .attr('width', '1px')
        .attr('height', function (c) {
          return (yaxis(Candle(c).High) - yaxis(Candle(c).Low))
        })
        .attr('x', function (c) { return roundToFive(x(c.Time.Close)) + 1; })
        .attr('y', function (c) { return h - yaxis(Candle(c).High)})
        .attr('fill', span)
        ;

      var openClose = svg.selectAll('rect.open-close').data(cs).enter().append('rect')

      var candleAttrs = openClose
        .attr('class', 'open-close')
        .attr('width', '2px')
        .attr('height', function (c) {
          return Math.abs(yaxis(Candle(c).Open) - yaxis(Candle(c).Close))
        })
        .attr('x', function (c) {
          return roundToFive(Math.round(x(c.Time.Close)));
        })
        //.attr('data-range-x', function (c) { return x(c.Time.Close); })
        .attr('data-timestamp', function (c) { return c.Time.Close })
        .attr('y', function (c) { return h - (yaxis(Math.max(Candle(c).Close,Candle(c).Open))) })
        .attr('fill', '#171718')
        .attr('strokeWidth', '1')
        .attr('stroke', function (c) { return Candle(c).Close > Candle(c).Open ? green : red })
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

      function drawLine(getter, className, xOffset, xOffsetPx) {
        svg.selectAll('path.' + className).data([]).exit().remove()
        xOffset = xOffset || 0;
        xOffsetPx = xOffsetPx || 0;
        var generator = d3.svg.line()
          .x(function (c) { return x(c.Time.Close + xOffset) + xOffsetPx })
          .y(function (c) { return h - yaxis(getter(c)) })

        svg.append('path')
          .attr('class', className)
          .attr('fill', 'none')
          .attr('d', generator(cs))
      }
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
      var start = Math.round(window.scrollX / 5) - 11
        , amt   = Math.round(window.innerWidth / 5) + 11
      return candles.slice(start, start + amt);
    }

    var yRangeCache = {};

    function yRange(cs) {
      var low = Math.min.apply(Math, cs.map(function (c) { return Candle(c).Low }));
      var high = Math.max.apply(Math, cs.map(function (c) { return Candle(c).High }));
      return d3.scale
        .linear()
        .domain([low, high])
        .range([0 + padding.y + (settings.rsi ? hRSI + padding.y : 0), h - padding.y])
    }


    var drawTimeout;

    $(window).scroll(refresh);

    function refresh () {
      clearTimeout(drawTimeout)
      drawTimeout = setTimeout(function () {
        var cs = visibleCandles()
        if (cs.length === 0) return;
        var cacheKey = '' + cs[0].Time.Close + cs[cs.length - 1].Time.Close + (settings.rsi ? 'RSI' : '');
        var yr;
        if (yRangeCache[cacheKey] !== undefined) {
          yr = yRangeCache[cacheKey];
        } else {
          var yr = yRange(cs);
          yRangeCache[cacheKey] = yr;
        }

        draw(cs, yr, yrRSI);
      }, 10);
    };

    window.refresh = refresh;

    Object.keys(settings).forEach(function (key) {
      memorizedSetting = getSetting(key);
      if (memorizedSetting !== undefined) {
        settings[key] = memorizedSetting;
      }
      toggleSetting(key, settings[key]);
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

        priceAtXCross = 1 - ( (e.pageY - padding.y) / (window.innerHeight - (padding.y * 2) - (settings.rsi ? hRSI + padding.y : 0)))
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
      $('#loading, #spinner').remove();
      window.scrollBy(ww * 2, 0);
      clearInterval(spinnerInterval);
    },0);

    setTimeout(getLatestInterval, response.PingIn * 1000);
  }

  // END DEFER
  },1);

  function getLatestInterval() {
    $.ajax({
      type: 'get',
      datatype: 'json',
      url: HOST + '/latest-interval.json',
      success: function (res) {
        var interval = res.Intervals[0];
        addCandleAndRefresh(interval);
        // DO IT AGAIN! FOREVER. LETS WATCH THIS BABY ALL THE WAY TO DA MOON
        // This is some of the worst JS I've ever written
        if (res.PingIn > 0) {
          setTimeout(getLatestInterval, res.PingIn * 1000)
        }
      }
    });
  }

  function addCandleAndRefresh(candle) {
    candles.push(candle);
    refresh();
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
    hotkeys = {
      49: 'ema-10'
    , 50: 'ema-21'
    , 69: 'keltner'
    , 72: 'heikin-ashi'
    , 73: 'ichimoku'
    , 76: 'legend'
    , 80: 'prices'
    , 82: 'rsi'
    , 83: 'sar'
    }

    var key = hotkeys[e.which];
    toggleSetting(key, !settings[key])
    settings[key] = !settings[key]
  });

  function toggleSetting(key, setting) {
    var $body = $('body')
    storeSetting(key, setting);
    switch (key) {
      case 'legend':
        $body.toggleClass('hide-legend', !setting);
        break;
      case 'prices':
        $body.toggleClass('hide-prices', !setting);
        $('#hotkey-p').toggleClass('inactive', !setting);
        break;
      case 'heikin-ashi':
        $('#hotkey-h').toggleClass('inactive', !setting);
        refresh();
        break;
      case 'ema-10':
        $body.toggleClass('hide-ema-10', !setting);
        $('#hotkey-1').toggleClass('inactive', !setting);
        break;
      case 'ema-21': // 2
        $body.toggleClass('hide-ema-21', !setting);
        $('#hotkey-2').toggleClass('inactive', !setting);
        break;
      case 'sar': // S
        $body.toggleClass('hide-sar', !setting);
        $('#hotkey-s').toggleClass('inactive', !setting);
        break;
      case 'ichimoku': // I
        $body.toggleClass('hide-ichimoku', !setting);
        $('#hotkey-i').toggleClass('inactive', !setting);
        break;
      case 'keltner': // L
        $body.toggleClass('hide-keltner', !setting);
        $('#hotkey-e').toggleClass('inactive', !setting);
        break;
      case 'rsi': // R
        $('#hotkey-r').toggleClass('inactive', !setting);
        refresh()
        $body.toggleClass('hide-rsi', !setting);
        break;
    }
  }

  var $spinner = $('#spinner')
    , spinnerFrame = 0
    , spinnerFrames = ['|', '/', '-', '\\', '|', '/', '-', '\\']
    ;

  var spinnerInterval = setInterval(function () {
    if (spinnerFrame === spinnerFrames.length - 1) {
      spinnerFrame = 0;
    } else {
      spinnerFrame ++;
    }
    $spinner.text(spinnerFrames[spinnerFrame]);
  }, 200);

}());
