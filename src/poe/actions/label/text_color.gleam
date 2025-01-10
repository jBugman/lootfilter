import gleam/string

import poe/actions/label/rgba.{type RGBA}

pub type TextColor {
  TextColor(color: RGBA)
}

pub fn to_string(text_color: TextColor) -> String {
  let TextColor(color) = text_color

  ["SetTextColor", rgba.to_string(color)]
  |> string.join(" ")
}
