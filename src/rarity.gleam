import gleam/option.{type Option, None, Some}

import actions/beam.{type Beam, Beam}
import actions/color.{type Color}
import actions/map_icon/size.{type Size}
import domain

pub type Rarity {
  Normal
  Uncommon
  Rare
  Epic
  Legendary
}

pub fn render(rarity: Rarity) -> List(String) {
  [map_icon(rarity), beam(rarity)]
  |> option.values
}

fn map_icon(rarity: Rarity) -> Option(String) {
  case rarity {
    Legendary -> circle(size.Large, color.Orange)
    Epic -> circle(size.Large, color.Purple)
    Rare -> circle(size.Medium, color.Blue)
    Uncommon -> circle(size.Small, color.Green)
    Normal -> None
  }
}

fn circle(size: Size, color: Color) -> Option(String) {
  domain.circle(size, color)
  |> Some
}

fn beam(rarity: Rarity) -> Option(String) {
  case rarity {
    Legendary -> beam_(color.Orange)
    Epic -> beam_(color.Purple)
    Rare -> beam_(color.Blue)
    Uncommon -> None
    Normal -> None
  }
}

fn beam_(c: Color) -> Option(String) {
  Beam(c)
  |> beam.to_string
  |> Some
}
