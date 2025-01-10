import gleeunit/should

import poe/actions/label/border_color.{BorderColor}
import poe/actions/label/rgba.{RGB}

pub fn border_color_test() {
  BorderColor(RGB(66, 33, 99))
  |> border_color.to_string
  |> should.equal("SetBorderColor 66 33 99 255")
}
