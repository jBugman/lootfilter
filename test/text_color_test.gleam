import gleeunit/should

import poe/actions/label/rgba.{RGB}
import poe/actions/label/text_color.{TextColor}

pub fn text_color_test() {
  TextColor(RGB(66, 33, 99))
  |> text_color.to_string
  |> should.equal("SetTextColor 66 33 99 255")
}
