import gleeunit/should

import actions/beam.{Beam}
import actions/color
import actions/map_icon/map_icon.{MapIcon}
import actions/map_icon/shape
import actions/map_icon/size
import actions/rgba.{Greyscale, RGB, RGBA}

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
