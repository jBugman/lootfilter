import gleam/int
import gleam/string

pub type RGBA {
  RGB(Int, Int, Int)
  RGBA(Int, Int, Int, Int)
  Greyscale(Int)
}

pub fn to_string(color: RGBA) -> String {
  case color {
    RGBA(r, g, b, a) -> {
      [int.to_string(r), int.to_string(g), int.to_string(b), int.to_string(a)]
      |> string.join(" ")
    }

    RGB(r, g, b) -> RGBA(r, g, b, 255) |> to_string

    Greyscale(v) -> RGBA(v, v, v, 255) |> to_string
  }
}
