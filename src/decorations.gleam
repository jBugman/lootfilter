import gleam/int
import gleam/list
import gleam/option.{type Option, None, Some}

import actions/rgba.{type RGBA}
import actions/sound.{type Sound}
import rarity.{type Rarity}
import styles.{type Styles}

type Props =
  List(String)

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

pub fn render(decor: Decorations, global_style: Styles) -> Props {
  [
    decor.font_size
      |> font_size,
    decor.text
      |> text_color,
    decor.border
      |> border_color,
    decor.background
      |> background(global_style.background),
    decor.rarity
      |> rarity.render,
    decor.sound
      |> sound.render,
  ]
  |> list.flatten
}

fn apply(opt: Option(a), f: fn(a) -> List(String)) -> Props {
  opt
  |> option.map(f)
  |> option.unwrap([])
}

fn font_size(val: Option(Int)) -> Props {
  apply(val, fn(n) { ["SetFontSize " <> int.to_string(n)] })
}

fn text_color(val: Option(RGBA)) -> Props {
  apply(val, fn(c) { ["SetTextColor " <> rgba.to_string(c)] })
}

fn border_color(val: Option(RGBA)) -> Props {
  apply(val, fn(c) { ["SetBorderColor " <> rgba.to_string(c)] })
}

fn background(local: Option(RGBA), global: Option(RGBA)) -> Props {
  option.or(local, global)
  |> apply(fn(c) { ["SetBackgroundColor " <> rgba.to_string(c)] })
}
