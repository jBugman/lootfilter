import gleeunit/should

import block/label/font_size.{FontSize}

pub fn font_size_test() {
  FontSize(40)
  |> font_size.to_string
  |> should.equal("SetFontSize 40")
}
