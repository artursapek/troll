package analysis

func (slope Metrics) IsUpwards() bool {
  // Only last ten minutes really matter
  return slope["5"] > 0 && slope["10"] > 0
}

func (slope Metrics) Flat() bool {
  // Only last ten minutes really matter
  return (slope["10"] < 3 && slope["10"] > -3)
}

func (slope Metrics) Accelerating() bool {
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
  // Compare relatively. Not exactly the opposite of Accelerating
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


