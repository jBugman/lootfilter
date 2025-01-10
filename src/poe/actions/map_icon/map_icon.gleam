import gleam/string

import poe/actions/color.{type Color}
import poe/actions/map_icon/shape.{type Shape}
import poe/actions/map_icon/size.{type Size}

pub type MapIcon {
  MapIcon(size: Size, color: Color, shape: Shape)
}

pub fn to_string(icon: MapIcon) -> String {
  [
    "MinimapIcon",
    size.to_string(icon.size),
    color.to_string(icon.color),
    shape.to_string(icon.shape),
  ]
  |> string.join(" ")
}
