import gleeunit/should

import actions/rgba.{RGB}
import block/label/text_color.{TextColor}

pub fn text_color_test() {
  TextColor(RGB(66, 33, 99))
  |> text_color.to_string
  |> should.equal("SetTextColor 66 33 99 255")
}
