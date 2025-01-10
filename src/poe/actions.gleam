import gleam/string_tree.{type StringTree}

import poe/actions/beam
import poe/actions/label/border_color
import poe/actions/map_icon/map_icon

pub type Action {
  MapIcon(map_icon: map_icon.MapIcon)
  Beam(beam: beam.Beam)
  BorderColor(color: border_color.BorderColor)
}

pub fn to_string(action: Action) -> String {
  case action {
    MapIcon(x) -> map_icon.to_string(x)
    Beam(x) -> beam.to_string(x)
    BorderColor(x) -> border_color.to_string(x)
  }
}

pub fn to_string_tree(action: Action) -> StringTree {
  action |> to_string |> string_tree.from_string
}
