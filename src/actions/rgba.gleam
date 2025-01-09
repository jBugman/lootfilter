import gleam/int
import gleam/list
import gleam/string

pub type RGBA {
  RGBA(String)
  Greyscale(Int)
}

pub fn to_string(color: RGBA) -> String {
  case color {
    RGBA(color) -> color

    Greyscale(value) -> {
      int.to_string(value)
      |> list.repeat(3)
      |> string.join(" ")
    }
  }
}
