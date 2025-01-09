import gleeunit/should

import actions/beam.{Beam}
import actions/color
import actions/map_icon/map_icon.{MapIcon}
import actions/map_icon/shape
import actions/map_icon/size

pub fn map_icon_test() {
  MapIcon(size.Small, color.Red, shape.Circle)
  |> map_icon.to_string
  |> should.equal("MinimapIcon 2 Red Circle")
}

pub fn beam_test() {
  Beam(color.Orange)
  |> beam.to_string
  |> should.equal("PlayEffect Orange")
}
