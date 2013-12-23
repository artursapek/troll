package analysis

func (slope Metrics) IsUpwards() bool {
  // Only last ten minutes really matter
  return slope["5"] > 0 && slope["10"] > 0
}

func (status MarketStatus) SlopeIsSevere() bool {
  // Hard-code this for now. Dont know how else to do it.
  slope := status.Analysis.Slope
  threshold := status.Price / 20
  if slope.IsUpwards() {
    return (slope["5"] > threshold)
  } else {
    return (slope["5"] < -threshold)
  }
}

func (slope Metrics) IsAccelerating() bool {
  // In whichever direction it's going
  if slope.IsUpwards() {
    return slope["5"]  > slope["10"] &&
           slope["10"] > slope["30"] && 
           slope["30"] > slope["60"]
  } else {
    return slope["5"]  < slope["10"] &&
           slope["10"] < slope["30"] && 
           slope["30"] < slope["60"]
  }
  /*
   *                *
   *               *
   *              *
   *           **
   *        **
   *   ****
   */
}

func (slope Metrics) HasSettled() bool {
  // Compare relatively. Not exactly the opposite of IsAccelerating
  if slope.IsUpwards() {
    return slope["5"]  < (slope["10"] / 1.5) &&
           slope["10"] < (slope["30"] / 1.5)
  } else {
    return slope["5"]  > (slope["10"] / 1.5) &&
           slope["10"] > (slope["30"] / 1.5)
  }

  /*
   *         *****
   *      **
   *    *
   *   *
   *   *
   *  *
   *
   */
}
