import gleam/option.{type Option, None, Some}

import actions/rgba.{type RGBA}
import actions/sound.{type Sound}
import rarity.{type Rarity}

pub type Decorations {
  Decorations(
    font_size: Option(Int),
    text: Option(RGBA),
    border: Option(RGBA),
    background: Option(RGBA),
    rarity: Rarity,
    sound: Sound,
  )
}

pub const hidden = Decorations(
  font_size: Some(10),
  rarity: rarity.Normal,
  text: None,
  border: None,
  background: None,
  sound: sound.Mute,
)
