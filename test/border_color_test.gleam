import gleeunit/should

import actions/rgba.{RGB}
import block/label/border_color.{BorderColor}

pub fn border_color_test() {
  BorderColor(RGB(66, 33, 99))
  |> border_color.to_string
  |> should.equal("SetBorderColor 66 33 99 255")
}
