pub type Color {
  Blue
  Brown
  Cyan
  Green
  Grey
  Orange
  Pink
  Purple
  Red
  White
  Yellow
}

pub fn to_string(color: Color) -> String {
  case color {
    Blue -> "Blue"
    Brown -> "Brown"
    Cyan -> "Cyan"
    Green -> "Green"
    Grey -> "Grey"
    Orange -> "Orange"
    Pink -> "Pink"
    Purple -> "Purple"
    Red -> "Red"
    White -> "White"
    Yellow -> "Yellow"
  }
}
