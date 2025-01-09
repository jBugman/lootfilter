import gleam/string

import actions/rgba.{type RGBA}

pub type BorderColor {
  BorderColor(color: RGBA)
}

pub fn to_string(border_color: BorderColor) -> String {
  let BorderColor(color) = border_color

  ["SetBorderColor", rgba.to_string(color)]
  |> string.join(" ")
}
