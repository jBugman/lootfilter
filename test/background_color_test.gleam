import gleeunit/should

import actions/rgba.{RGB}
import block/label/background_color.{BackgroundColor}

pub fn background_color_test() {
  BackgroundColor(RGB(66, 33, 99))
  |> background_color.to_string
  |> should.equal("SetBackgroundColor 66 33 99 255")
}
