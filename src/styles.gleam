import gleam/dict.{type Dict}
import gleam/option.{type Option}
import tom.{type Toml}

import actions/rgba.{type RGBA, RGBA}

pub type Styles {
  Styles(background: Option(RGBA))
}

pub fn parse(src: Dict(String, Toml)) -> Styles {
  let bg_color =
    tom.get_string(src, ["global_styles", "background_color"])
    |> option.from_result
    |> option.map(RGBA)

  Styles(background: bg_color)
}
