import gleeunit/should

import poe/actions/label/background_color.{BackgroundColor}
import poe/actions/label/rgba.{RGB}

pub fn background_color_test() {
  BackgroundColor(RGB(66, 33, 99))
  |> background_color.to_string
  |> should.equal("SetBackgroundColor 66 33 99 255")
}
