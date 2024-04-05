import gleam/option.{type Option}
import poe/color.{type Color}
import gleam/dict.{type Dict}
import tom.{type Toml}

pub type Styles {
  Styles(background: Option(Color))
}

pub fn parse(src: Dict(String, Toml)) -> Styles {
  let bg_color =
    tom.get_string(src, ["global_styles", "background_color"])
    |> option.from_result
    |> option.map(color.Custom)

  Styles(background: bg_color)
}
