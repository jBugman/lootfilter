import gleam/int
import gleam/list
import gleam/string

pub type Color {
  Custom(String)
  Greyscale(Int)
}

pub fn to_string(color: Color) -> String {
  case color {
    Custom(color) -> color

    Greyscale(value) -> {
      int.to_string(value)
      |> list.repeat(3)
      |> string.join(" ")
    }
  }
}
