import gleam/int
import gleam/string

pub type FontSize {
  FontSize(size: Int)
}

pub fn to_string(font_size: FontSize) -> String {
  let FontSize(size) = font_size

  ["SetFontSize", int.to_string(size)]
  |> string.join(" ")
}
