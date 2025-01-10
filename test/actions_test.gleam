import gleeunit/should

import poe/actions
import poe/actions/beam.{Beam}
import poe/actions/color
import poe/actions/label/rgba.{Greyscale, RGB, RGBA}
import poe/actions/map_icon/map_icon.{MapIcon}
import poe/actions/map_icon/shape
import poe/actions/map_icon/size

pub fn action_map_icon_test() {
  MapIcon(size.Medium, color.White, shape.Square)
  |> actions.MapIcon
  |> actions.to_string
  |> should.equal("MinimapIcon 1 White Square")
}

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

pub fn rgba_test() {
  RGBA(255, 0, 0, 80)
  |> rgba.to_string
  |> should.equal("255 0 0 80")
}

pub fn rgb_test() {
  RGB(0, 235, 0)
  |> rgba.to_string
  |> should.equal("0 235 0 255")
}

pub fn greyscale_test() {
  Greyscale(100)
  |> rgba.to_string
  |> should.equal("100 100 100 255")
}
