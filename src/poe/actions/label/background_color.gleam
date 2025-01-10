import gleam/string

import poe/actions/label/rgba.{type RGBA}

pub type BackgroundColor {
  BackgroundColor(color: RGBA)
}

pub fn to_string(background_color: BackgroundColor) -> String {
  let BackgroundColor(color) = background_color

  ["SetBackgroundColor", rgba.to_string(color)]
  |> string.join(" ")
}
